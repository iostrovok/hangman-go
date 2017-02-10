package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"session"
	"words"
)

func printData(w http.ResponseWriter, data map[string]interface{}) {

	res, err := json.Marshal(data)
	if err != nil {
		log.Printf("printData: %s\n", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

func initUserGame(w http.ResponseWriter, r *http.Request) *Session.UserGame {

	id := ""
	cookieFrom, err := r.Cookie("id")
	if err == nil {
		id = cookieFrom.Value
	}

	user := Session.Start().FindOrCreate(id)

	expiration := time.Now().Add(24 * time.Hour)
	cookie := http.Cookie{Name: "id", Value: user.ID, Expires: expiration}
	http.SetCookie(w, &cookie)
	return user
}

func newGame(w http.ResponseWriter, r *http.Request) {
	user := initUserGame(w, r)
	words := Words.Get()
	res := user.NewWord(words)
	printData(w, res)
}

func move(w http.ResponseWriter, r *http.Request) {
	user := initUserGame(w, r)
	letter := r.FormValue("letter")

	res := user.Move(letter)
	printData(w, res)
}

func userInfo(w http.ResponseWriter, r *http.Request) {
	user := initUserGame(w, r)
	printData(w, user.Info())
}

func main() {
	var dasedir string
	var port int
	flag.StringVar(&dasedir, "dir", "./", "Root dir")
	flag.IntVar(&port, "port", 19720, "Port")
	flag.Parse()

	dasedir = strings.TrimRight(dasedir, "/")

	// Load words from txt file
	Words.Init(dasedir + "/words.txt")

	fmt.Printf("Port: %d\n", port)
	fmt.Printf("Base dir: %s\n", dasedir)
	fmt.Printf("Start with images: %s\n", dasedir+"/html/")
	fmt.Printf("Start with css: %s\n", dasedir+"/html/"+"css/")
	fmt.Printf("Start with index file: %s\n", dasedir+"/html/"+"html/index.html")

	// Static pages
	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir(dasedir+"/html/html/"))))
	http.Handle("/html/", http.StripPrefix("/html/", http.FileServer(http.Dir(dasedir+"/html/"))))

	// Dinamic handlers
	http.HandleFunc("/start", newGame)
	http.HandleFunc("/move", move)
	http.HandleFunc("/user_info", userInfo)

	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
