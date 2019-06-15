package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/GuilhermeCaruso/bellt"
)

//main
func main() {

	router := bellt.NewRouter()

	router.HandleFunc("/contact/{id}/{user}", bellt.Use(
		exampleHandler,
		middlewareOne,
		middlewareTwo,
	), "GET")

	router.HandleFunc("/contact", bellt.Use(
		exampleNewHandler,
		middlewareOne,
		middlewareTwo,
	), "GET")

	router.HandleGroup("/api",
		router.SubHandleFunc("/check", bellt.Use(
			exampleNewHandler,
			middlewareOne,
			middlewareTwo,
		), "GET"),
		router.SubHandleFunc("/check/{id}/{user}", bellt.Use(
			exampleHandler,
			middlewareOne,
			middlewareTwo,
		), "GET"),
	)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func exampleHandler(w http.ResponseWriter, r *http.Request) {

	rv := bellt.RouteVariables(r)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf(`{"id": %v, "user": %v}`, rv.GetVar("user"), rv.GetVar("id"))))
}

func exampleNewHandler(w http.ResponseWriter, r *http.Request) {
	//rv := bellt.RouteVariables(r)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"msg": "Works"}`))
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