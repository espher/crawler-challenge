package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	url := "https://www.ebay.com/sch/garlandcomputer/m.html"

	respuesta, err := http.Get(url)
	if err != nil {
		fmt.Println("Error al obtener la página:", err)
		return
	}
	defer respuesta.Body.Close()

	if respuesta.StatusCode != http.StatusOK {
		fmt.Println("No se pudo obtener la página. Código de estado:", respuesta.StatusCode)
		return
	}

	doc, err := goquery.NewDocumentFromReader(respuesta.Body)
	if err != nil {
		log.Fatalf("Error al crear el documento goquery: %v", err)
	}

	doc.Find("div.s-item__info").Each(func(index int, div *goquery.Selection) {
		// Obtener el texto dentro del div
		//texto := div.Text()
		//fmt.Printf("Texto dentro del div %d: %s\n", index+1, texto)
		t := doc.Find(".s-item__title").Text()
		fmt.Printf("title  %d: %s\n", index+1, t)
	})

}
