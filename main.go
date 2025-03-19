package main

import (
	"embed"
	"net/http"

	"github.com/nabilulilalbab/TugasPemogreamanWeb3/routes"
)

//go:embed views/*.html
var Views embed.FS




func main()  {
  mux := http.NewServeMux()
  routes.BookInfoRoutes(mux, Views)
  if err := http.ListenAndServe(":8080", mux); err != nil {
    panic(err)
  }
}
