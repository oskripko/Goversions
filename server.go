package main
// START OMIT
import (
	"embed"
	"net/http"
)

//go:embed assets/*
var assets embed.FS // Loads the whole folder // HL

//go:embed index.html
var html []byte //Preloaded []byte for content writing // HL

func runServer() {
	fs := http.FileServer(http.FS(assets))
	http.Handle("/assets/", fs) // HL
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "text/html")
		w.Write(html) // HL
	})
	http.ListenAndServe(":8080", nil)
}
// END OMIT