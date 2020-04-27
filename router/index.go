package router

import (
	"net/http"

	"github.com/products/core"
)



func Index(res http.ResponseWriter, req *http.Request) {
	core.Render("views/layout.html", "views/home.html", "layout", res)
}
