package main

import (
	"fmt"
	"log"
	"net/http"
)

func abuoutHandler(w http.ResponseWriter, r *http.Request) {
	p := "./template/about.html"
	if r.URL.Path != "/about" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return

	} else {
		http.ServeFile(w, r, p)
	}
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	p := "./template/contact.html"
	if r.URL.Path != "/contact" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return

	} else {
		http.ServeFile(w, r, p)
	}
}
func homeHandler(w http.ResponseWriter, r *http.Request) {
	p := "./template/home.html"
	if r.URL.Path != "/home" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	} else {
		http.ServeFile(w, r, p)
	}
}

func main() {
	fmt.Println("hello world")
	fileServer := http.FileServer(http.Dir("./template"))
	http.Handle("/", fileServer)
	http.HandleFunc("/home", homeHandler)
	http.HandleFunc("/about", abuoutHandler)
	http.HandleFunc("/contact", contactHandler)

	fmt.Println("string server at port 6060")

	if err := http.ListenAndServe(":6060", nil); err != nil {
		log.Fatal(err)
	}

}
