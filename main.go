// Section 3, 33

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=firstgo password=kiasu123")
	defer db.Close()

	if err != nil {
		db.Close()
	}

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		cards := newDeck()
		hand, remainingCards := deal(cards, 5)

		// cards.saveToFile("Cards")
		cards2 := newDeckFromFile("Cards")
		cards2.shuffle()

		c.JSON(200, gin.H{
			"message": "pong",
			"data": gin.H{
				"cards":          cards,
				"hand":           hand,
				"remainingCards": remainingCards,
				"string":         cards.toString(),
				"cards2":         cards2,
			},
		})
	})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
