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

// 使用费马小定理检查是否为素数
func isPrime(n *big.Int, iterations int) bool {
	if n.Cmp(big.NewInt(2)) == 0 {
		return true
	}

	// 偶数
	if n.Bit(0) == 0 {
		return false
	}

	one := big.NewInt(1)
	nMinusOne := big.NewInt(0).Sub(n, one) // Create a separate copy of n - 1

	// 随机数生成器
	//rand.Seed(time.Now().UnixNano())

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

// 我们使用费马小定理进行素数检查。
// 费马小定理是基于数论的一个重要定理，它说明如果 p 是一个素数，a 是不可整除 p 的整数，那么 a^(p-1) ≡ 1 (mod p)。 （当 p 是一个素数时，费马小定理表明对于任意不可整除 p 的整数 a，它的幂 a^(p-1) 对 p 取模的结果等于 1。这可以表示为 a^(p-1) ≡ 1 (mod p)，其中 ≡ 表示同余关系，mod 表示取模运算）
// 这意味着如果 n 不是素数，那么对于任意 1 < a < n，a^(n-1) mod n ≠ 1。
func main() {
	n := big.NewInt(13) // 待检查的数
	iterations := 10    // 迭代次数

	if isPrime(n, iterations) {
		fmt.Printf("%s 是素数\n", n.String())
	} else {
		fmt.Printf("%s 不是素数\n", n.String())
	}
}
