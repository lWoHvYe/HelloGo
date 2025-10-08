package main

import (
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// 请求前后的切入点
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		// 设置 example 变量
		c.Set("example", "12345")

		// 请求前

		c.Next()

		// 请求后
		latency := time.Since(t)
		log.Print(latency)

		// 获取发送的 status
		status := c.Writer.Status()
		log.Println(status)

		// Step2: Check if any errors were added to the context
		if len(c.Errors) > 0 {
			// Step3: Use the last error
			err := c.Errors.Last().Err

			// Step4: Respond with a generic error message
			c.JSON(http.StatusInternalServerError, map[string]any{
				"success": false,
				"message": err.Error(),
			})
		}
	}
}

func main() {
	r := gin.New()
	r.Use(Logger())
	r.Use(gin.Recovery())

	r.GET("/test", func(c *gin.Context) {
		example := c.MustGet("example").(string)
		p2, ok := c.GetQuery("p2") // 获取单值请求参数
		if !ok {
			c.Error(errors.New("p2 not found"))
			return
		}

		// 打印："12345"
		log.Println(example, p2)
	})

	// 监听并在 0.0.0.0:8080 上启动服务
	r.Run(":8080")
}
