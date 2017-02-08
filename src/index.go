package main

import (
	"flag"
	"fmt"
	"time"
	// "log"
	"net/http"
	// "os"
	// "path/filepath"

	"session"
)

func handler(w http.ResponseWriter, r *http.Request) {

	id := ""
	cookieFrom, err := r.Cookie("id")
	if err == nil {
		id = cookieFrom.Value
	}

	user := Session.Start().FindOrCreate(id)

	expiration := time.Now().Add(365 * 24 * time.Hour)
	cookie := http.Cookie{Name: "id", Value: user.ID, Expires: expiration}
	http.SetCookie(w, &cookie)
	fmt.Fprintf(w, "Hi there, I love %s! cookieFrom: %s, %T, ses: %s", r.URL.Path[1:], cookieFrom, cookieFrom, Session.Start())
}

func main() {
	var svar string
	flag.StringVar(&svar, "img", "./", "Dir with images")
	flag.Parse()

	fmt.Printf("Start with images: %s\n", svar)
	fmt.Printf("Start with index file: %s\n", svar+"html/index.html")

	// http.HandleFunc("/", handler)
	// Static pages
	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir(svar+"/html/"))))
	http.Handle("/html/", http.StripPrefix("/html/", http.FileServer(http.Dir(svar))))

	// Dinamic handlers
	http.HandleFunc("/start", handler)
	http.HandleFunc("/move", handler)

	http.ListenAndServe(":8080", nil)
}
