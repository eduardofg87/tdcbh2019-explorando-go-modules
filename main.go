package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/GuilhermeCaruso/bellt"
	"github.com/qeesung/image2ascii/convert"
	_ "image/jpeg"
	_ "image/png"
)

func main() {

	router := bellt.NewRouter()

	router.HandleFunc("/gopher", bellt.Use(
		gopherHandler,
		middlewareOne,
		middlewareTwo,
	), "GET")

	router.HandleFunc("/tdc-logo", bellt.Use(
		tdcLogoHandler,
		middlewareOne,
		middlewareTwo,
	), "GET")

	log.Fatal(http.ListenAndServe(":8080", nil))
}


func gopherHandler(w http.ResponseWriter, r *http.Request) {
	convertOptions := convert.DefaultOptions
	convertOptions.FixedWidth = 100
	convertOptions.FixedHeight = 40

	// Create the image converter
	converter := convert.NewImageConverter()
	imageFilename := "go.png"
	fmt.Print(converter.ImageFile2ASCIIString(imageFilename, &convertOptions))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"msg": "Gopher printed!"}`))
}


func tdcLogoHandler(w http.ResponseWriter, r *http.Request) {
	convertOptions := convert.DefaultOptions
	convertOptions.FixedWidth = 100
	convertOptions.FixedHeight = 40

	// Create the image converter
	converter := convert.NewImageConverter()
	imageFilename := "tdc_logo.png"
	fmt.Print(converter.ImageFile2ASCIIString(imageFilename, &convertOptions))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"msg": "TDC Logo printed!"}`))
}
func middlewareOne(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Step One")
		next.ServeHTTP(w, r)
	}
}

func middlewareTwo(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Step Two")
		next.ServeHTTP(w, r)
	}
}