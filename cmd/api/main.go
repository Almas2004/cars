package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"hw2.com/internal/config"
	"hw2.com/internal/db"
)

type CreateCarRequest struct {
	Brand string `json:"brand"`
	Model string `json:"model"`
	Price float64 `json:"price"`
}

func main() {
	cfg := config.Load()
	pool := db.NewPool(cfg.DBUrl)
	defer pool.Close()

	r := gin.Default()

	r.POST("/cars", func(c *gin.Context) {
		var req CreateCarRequest
		if err := c.ShouldBind(&req); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		var id int64
		err := pool.QueryRow(c.Request.Context(),
		`INSERT INTO cars (brand, model, price) VALUES ($1, $2, $3) RETURNING id`,
			req.Brand, req.Model, req.Price,
	).Scan(&id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, gin.H{"id": id})
	})

	r.GET("/cars", func(c *gin.Context) {
		rows, err := pool.Query(
			c.Request.Context(),
			`SELECT id, brand, model, price, created_at FROM cars ORDER BY id DESC LIMIT 50`,
		)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		defer rows.Close()

		cars := make([]gin.H, 0)

		for rows.Next() {
			var id int64
			var brand string
			var model string
			var price float64
			var createdAt any

			if err := rows.Scan(&id, &brand, &model, &price, &createdAt); err != nil {
				c.JSON(500, gin.H{"error": err.Error()})
				return
			}

			cars = append(cars, gin.H{
				"id":         id,
				"brand":      brand,
				"model":      model,
				"price":      price,
				"created_at": createdAt,
			})
		}

		c.JSON(200, cars)
	})


	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "OK"})
	})

	r.Run(":8686")
} 