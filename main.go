// Section 4, 41

package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

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

	r.GET("/structs", func(c *gin.Context) {
		name := person{
			firstName: "Ivan",
			lastName:  "Satya",
			contact: contactInfo{
				email:   "ivansatya10@gmail.com",
				zipCode: 60187,
			},
		}

		name.firstName = "Satya"
		name.lastName = "Putra"

		fmt.Println(name)
		fmt.Printf("%+v", name)

		c.JSON(200, gin.H{
			"message": name,
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
