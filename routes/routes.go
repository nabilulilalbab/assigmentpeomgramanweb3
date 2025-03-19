package routes

import (
	"embed"
	"net/http"

	"github.com/nabilulilalbab/TugasPemogreamanWeb3/controller"
)

func BookInfoRoutes(server *http.ServeMux, fs embed.FS)  {
  server.HandleFunc("/", controller.NewBookInfoController(fs))
}

