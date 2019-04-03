// Section 3, 29

package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		cards := newDeck()
		hand, remainingCards := deal(cards, 5)

		// cards.saveToFile("Cards")
		cards2 := newDeckFromFile("Cards")

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
