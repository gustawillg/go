package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.Println("request iniciada")
	defer log.Println("request finalizada")
	select {
	case <-time.After(5 * time.Second):
		//imprime no comand line stdout
		log.Println("request processada com sucesso")
		//imprime no browser
		w.Write([]byte("request processada com sucesso"))
	case <-ctx.Done():
		//imprime no comand line stdout
		log.Println("request cancelada pelo client")
		//imprime no browser
		http.Error(w, "request cancelada pelo cliente", http.StatusRequestTimeout)
	}
}
