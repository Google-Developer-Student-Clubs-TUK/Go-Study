package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

var (
	linkList map[string]string
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	linkList = map[string]string{}

	http.HandleFunc("/addLink", addLink)
	http.HandleFunc("/short/", getLink)
	http.HandleFunc("/home", Home)

	log.Fatal(http.ListenAndServe(":9000", nil))
}

func addLink(w http.ResponseWriter, r *http.Request) {
	key, ok := r.URL.Query()["link"]
	fmt.Println(key, ok)
	if ok {
		if _, ok := linkList[key[0]]; !ok {
			fmt.Println(linkList)
			genString := fmt.Sprint(rand.Int63n(1000))
			linkList[genString] = key[0]
			w.Header().Set("Content-Type", "text/html")
			w.WriteHeader(http.StatusAccepted)
			linkString := fmt.Sprintf("<a href=\"http://localhost:9000/short/%s\">http://localhost:9000/short/%s</a>", genString, genString)
			fmt.Fprintf(w, "Added shortLink\n")
			fmt.Fprintf(w, linkString)
			return
		}
		w.WriteHeader(http.StatusConflict)
		fmt.Fprintf(w, "Already have this link")
		return
	}

	w.WriteHeader(http.StatusBadRequest)
	fmt.Fprintf(w, "Failed to add link")
	return
}

func getLink(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	pathArgs := strings.Split(path, "/")
	log.Printf("Redirected to: %s", linkList[pathArgs[2]])
	http.Redirect(w, r, linkList[pathArgs[2]], http.StatusPermanentRedirect)
	return
}

func Home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	log.Println("Get Home")
	if r.URL.Path != "/home" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	var response string
	for shortLink, link := range linkList {
		response += fmt.Sprintf("Link: <a href=\"http://localhost:8080/short/%s\">http://localhost:8080/short/%s</a> \t\t ShortLink: %s", shortLink, shortLink, link)
	}
	fmt.Fprintf(w, "<h2>Hello and Welcome to the Go URL Shortener!<h2><br>\n")
	fmt.Fprintf(w, response)
	return
}
