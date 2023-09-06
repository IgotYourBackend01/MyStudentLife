package tracker

import (
	"log"
	"net/http"
)

func HandelFunc() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", Home)
	mux.HandleFunc("/artist/", ArtistHandl)
	mux.HandleFunc("/search", SearhHandel)
	mux.HandleFunc("/filter", FilterHandler)
	mux.Handle("/style/", http.StripPrefix("/style/", http.FileServer(http.Dir("./web/style/"))))
	log.Println("Запуск веб-сервера на http://localhost:4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
