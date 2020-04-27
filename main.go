package main

import (
	"net/http"

	 "github.com/products/core"
	
	"github.com/products/router"
	
)

func main() {

	http.HandleFunc("/", router.Index)
	http.HandleFunc("/about", router.About)
	http.HandleFunc("/admin", router.Admin)
	
	core.Serve("public", "/public/", ":3000")

}
