//go:build ignore

package main

import (
	"fmt"
	"log"
	"net/http"
)

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	p, _ := loadPage(title)
	fmt.Fprintf(w, "<body style='background: black; color:white; font-size:3rem;text-align: center;'><h1>%s</h1><div>%s</div></body>", p.Title, p.Body)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, %s", r.URL.Path[1:])
}

func main() {
	fmt.Println("go to localhost:3000")
	http.HandleFunc("/", handler)
	http.HandleFunc("/view/", viewHandler)
	log.Fatal(http.ListenAndServe(":3000", nil))
}
