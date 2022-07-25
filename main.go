package main

import (
	"net/http"

	"curso_api_golang/routes"
)

func main() {
	routes.CarregaRotas()
	http.ListenAndServe(":8080", nil)
}
