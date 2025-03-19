package controller

import (
	"embed"
	"html/template"
	"net/http"

	"github.com/nabilulilalbab/TugasPemogreamanWeb3/entity"
)

func NewBookInfoController(views embed.FS) func (w http.ResponseWriter, r *http.Request) {
  return func(w http.ResponseWriter, r *http.Request) {
   t := template.Must(template.ParseFS(views, "views/*.html"))
    buku1 := entity.Book{
      ID: 1,
      Title: "clean Code",
      Author: "Jarwo",
      Price: 8000.00,
      Publisher: entity.Publisher{
        Name: "Nabiel",
        City : "Demak", 
      },
    }
    t.ExecuteTemplate(w, "bookInfo.html", buku1)
  }
}
