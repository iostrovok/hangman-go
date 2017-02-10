package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"game"
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

func initGame(w http.ResponseWriter, r *http.Request) *Game.Game {

	id := ""
	cookieFrom, err := r.Cookie("id")
	if err == nil {
		id = cookieFrom.Value
	}

	mygame := Session.Start().FindOrCreate(id)

	expiration := time.Now().Add(24 * time.Hour)
	cookie := http.Cookie{Name: "id", Value: mygame.ID, Expires: expiration}
	http.SetCookie(w, &cookie)
	return mygame
}

func newGame(w http.ResponseWriter, r *http.Request) {
	mygame := initGame(w, r)
	words := Words.Get()
	res := mygame.NewWord(words)
	printData(w, res)
}

func move(w http.ResponseWriter, r *http.Request) {
	mygame := initGame(w, r)
	letter := r.FormValue("letter")

	res := mygame.Move(letter)
	printData(w, res)
}

func userInfo(w http.ResponseWriter, r *http.Request) {
	mygame := initGame(w, r)
	printData(w, mygame.Info())
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
