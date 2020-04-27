package router

import (
	"net/http"

	"github.com/products/core"
)



func About(res http.ResponseWriter, req *http.Request) {
	core.Render("views/layout.html", "views/about.html", "layout", res)
}
