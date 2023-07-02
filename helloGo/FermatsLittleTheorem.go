package main

import (
	"fmt"
	"math/big"
	"math/rand"
	"time"
)

// 计算 a^b mod m 的结果
// 对指数值e大于1的情况，对任意的x,y和m，我们总可以分别计算x取模m和y取模m，而后将它们乘起来之后取模m，得到x乘y取模m
func powerMod(a, b, m *big.Int) *big.Int {
	res := big.NewInt(1)

	for b.Cmp(big.NewInt(0)) > 0 {
		// 如果 b 是奇数，则将结果乘以 a
		if b.Bit(0) == 1 {
			res.Mul(res, a)
			res.Mod(res, m)
		}

		// 将 a 自乘
		a.Mul(a, a)
		a.Mod(a, m)

		// 将 b 右移一位（相当于除以 2）
		b.Rsh(b, 1)
	}

	return res
}

// 使用不会被欺骗的变形 Miller-Rabin检查
func isPrimeMillerRabin(n *big.Int, iterations int) bool {
	if n.Cmp(big.NewInt(2)) == 0 {
		return true
	}

	// 偶数
	if n.Bit(0) == 0 {
		return false
	}

	one := big.NewInt(1)
	if n.Cmp(one) == 0 {
		return false
	}
	nMinusOne := big.NewInt(0).Sub(n, one) // Create a separate copy of n - 1

	// 随机数生成器
	//rand.Seed(time.Now().UnixNano())

	// 检查1取模n的非平凡平方根
	// In the context of primality testing, if nonTrivialSquareRoot is not nil and is different from 1 and n-1, it means that the number n is composite (non-prime).
	//This is because a prime number should not have a non-trivial square root modulo n.
	nonTrivialSquareRoot := big.NewInt(1).ModSqrt(one, n)
	if nonTrivialSquareRoot != nil && nonTrivialSquareRoot.Cmp(one) != 0 && nonTrivialSquareRoot.Cmp(nMinusOne) != 0 {
		return false
	}

	for i := 0; i < iterations; i++ {
		// 随机选择一个 a，1 < a < n
		a := big.NewInt(0).Rand(rand.New(rand.NewSource(time.Now().UnixNano())), nMinusOne)
		a.Add(a, one)
		// 检查 a^(n-1) ≡ 1 (mod n)
		if powerMod(a, new(big.Int).Set(nMinusOne), new(big.Int).Set(n)).Cmp(one) != 0 { // Use the copy of n for powerMod
			return false
		}
	}

	return true
}

// 使用费马小定理检查是否为素数: 如果n是一个素数，a是小于n的任意正整数，那么a的n次方与a模n同余，存在能骗过check的Carmichael数
func isPrimeFermat(n *big.Int, iterations int) bool {
	if n.Cmp(big.NewInt(2)) == 0 {
		return true
	}

	// 偶数
	if n.Bit(0) == 0 {
		return false
	}

	one := big.NewInt(1)
	if n.Cmp(one) == 0 {
		return false
	}
	nMinusOne := big.NewInt(0).Sub(n, one) // Create a separate copy of n - 1

	// 随机数生成器
	//rand.Seed(time.Now().UnixNano())

	for i := 0; i < iterations; i++ {
		// 随机选择一个 a，1 < a < n
		a := big.NewInt(0).Rand(rand.New(rand.NewSource(time.Now().UnixNano())), nMinusOne)
		a.Add(a, one)
		cmpA := new(big.Int).Set(a)
		// 检查 a^(n) ≡ a (mod n)
		if powerMod(a, new(big.Int).Set(n), new(big.Int).Set(n)).Cmp(cmpA) != 0 { // Use the copy of n for powerMod
			return false
		}
	}

	return true
}

// 我们使用费马小定理进行素数检查。a^(p-1) ≡ 1 (mod p) 这个实际上是Miller-Rabin检查，费马检查是a^(p) ≡ a (mod p)
// 费马小定理是基于数论的一个重要定理，它说明如果 p 是一个素数，a 是不可整除 p 的整数，那么 a^(p-1) ≡ 1 (mod p)。 （当 p 是一个素数时，费马小定理表明对于任意不可整除 p 的整数 a，它的幂 a^(p-1) 对 p 取模的结果等于 1。这可以表示为 a^(p-1) ≡ 1 (mod p)，其中 ≡ 表示同余关系，mod 表示取模运算）
// 这意味着如果 n 不是素数，那么对于任意 1 < a < n，a^(n-1) mod n ≠ 1。
func main() {
	//n := big.NewInt(13) // 待检查的数
	iterations := 1000 // 迭代次数

	carmichael := []*big.Int{
		big.NewInt(561),
		big.NewInt(1105),
		big.NewInt(1729),
		big.NewInt(2465),
		big.NewInt(2821),
		big.NewInt(6601),
		big.NewInt(13),
		big.NewInt(7),
		big.NewInt(3),
		big.NewInt(2),
		big.NewInt(1),
		big.NewInt(0)}

	for _, candidate := range carmichael {
		go parallelCheck(candidate, iterations)
	}

	time.Sleep(10 * time.Second)
}

func parallelCheck(candidate *big.Int, iterations int) {
	if isPrimeFermat(candidate, iterations) {
		fmt.Printf("Fermat %s 是素数\n", candidate.String())
	} else {
		fmt.Printf("Fermat %s 不是素数\n", candidate.String())
	}
	if isPrimeMillerRabin(candidate, iterations) {
		fmt.Printf("Miller-Rabin %s 是素数\n", candidate.String())
	} else {
		fmt.Printf("Miller-Rabin %s 不是素数\n", candidate.String())
	}
}
