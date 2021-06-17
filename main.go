package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	fmt.Println("Read data")

	file, err := os.Open("all.html")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	doc, err := goquery.NewDocumentFromReader(file)
	if err != nil {
		log.Fatal(err)
	}

	// create csv
	fmt.Println("Create file")
	csvFile, err := os.Create("output.csv")
	if err != nil {
		log.Fatal("failed", err)
	}

	csvWriter := csv.NewWriter(csvFile)

	// extract fields
	doc.Find("ul > div").Each(func(i int, s *goquery.Selection) {
		title := s.Find("a.links").Text()
		description := s.Find(".paragraph").Text()
		youtube, _ := s.Find("a.links").Attr("href")
		_ = csvWriter.Write([]string{
			strconv.Itoa(i),
			title,
			youtube,
			description,
		})
	})

	fmt.Println("Complete")
	csvWriter.Flush()
	csvFile.Close()
}
