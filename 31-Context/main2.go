package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.Println("req iniciado")
	defer log.Println("req finalizada")
	select {
	case <-time.After(5 * time.Second):
		// imprime no comand line
		log.Println("request processada com sucesso")
		//printa no browser
		w.Write([]byte("Req processada com sucesso"))
	case <-ctx.Done():
		// imprime no comand line
		log.Println("request cancelada pelo client")
	}
}
