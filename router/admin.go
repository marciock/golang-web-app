package router

import (
	"net/http"

	"github.com/products/core"
)



func Admin(res http.ResponseWriter, req *http.Request) {
	core.Render("views/layout.html", "views/admin/menu.html", "layout", res)
}
