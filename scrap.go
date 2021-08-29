package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gocolly/colly"
)

type Fact struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
}

func Facts(c *gin.Context) {
	allFacts := make([]Fact, 0)

	// A collector
	collector := colly.NewCollector(
		colly.AllowedDomains("factretriever.com", "www.factretriever.com"),
	)

	// Working on DOM
	collector.OnHTML(".factsList li", func(h *colly.HTMLElement) {
		faciId, err := strconv.Atoi(h.Attr("id"))

		if(err != nil){
			log.Println("Err 1 : Could not get Id...")
		}

		factDescription := h.Text


		fact := Fact{
			ID: faciId,
			Description: factDescription,
		}

		allFacts = append(allFacts, fact)

	})

	// Sanity Check
	collector.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting : ", r.URL.String())
	})


	collector.Visit("https://www.factretriever.com/area-51-facts")

	enc := json.NewEncoder(c.Writer)

	enc.SetIndent("", " ")
	enc.Encode(allFacts)

}