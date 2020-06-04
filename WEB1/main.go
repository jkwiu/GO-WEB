package main

import (
	"net/http"

	"github.com/GO-WEB/WEB1/myapp"
)

func main() {
	http.ListenAndServe(":3000", myapp.NewHttpHandler())
}
