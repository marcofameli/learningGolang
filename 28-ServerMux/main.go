package main

import (
	"log"
	"net/http"
)

// / MUX SERVE PARA CRIAR MAIS DE UMA PORTA, EXEMPLO EMBAIXO 8080, 8081
//func main() {
//	mux := http.NewServeMux()
//	mux.HandleFunc("/", HomeHandler)
//	mux.Handle("/blog", blog{title: "MyBlog"})
//	http.ListenAndServe(":8080", mux)
//
//	mux2 := http.NewServeMux()
//	mux2.HandleFunc("/", func(w http.ResponseWriter, request *http.Request) {
//		w.Write([]byte("Hello World!"))
//	})
//	http.ListenAndServe(":8081", mux2)
//}
//
//func HomeHandler(w http.ResponseWriter, r *http.Request) {
//	w.Write([]byte("Hello World"))
//}
//
//type blog struct {
//	title string
//}
//
//func (b blog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
//	w.Write([]byte(b.title))
//}

func main() {
	fileServer := http.FileServer(http.Dir("./public"))
	mux := http.NewServeMux()
	mux.Handle("/", fileServer)
	mux.HandleFunc("/blog", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello from blog"))
	})
	log.Fatal(http.ListenAndServe(":8080", mux))
	// http://localhost:8080/
}
