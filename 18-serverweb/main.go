package main

import "net/http"

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/login", login)
	http.HandleFunc("/saludo", saludo)

	http.ListenAndServe(":4008", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./index.html")
}

func login(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hola Mundo"))
}

func saludo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write([]byte(`{"mensaje":"Hola Mundo"}`))
}
