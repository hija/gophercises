package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"text/template"
)

var jsonstory story

type story map[string]storyPart

type storyPart struct {
	Title   string        `json:"title"`
	Story   []string      `json:"story"`
	Options []storyOption `json:"options"`
}

type storyOption struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

func loadStory() story {

	var story story
	jsonFile, err := os.Open("gopher.json")
	if err != nil {
		fmt.Println(err)
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &story)
	return story

}

func storyHandler(w http.ResponseWriter, req *http.Request) {
	// Check if we know the part of the story
	if storypart, ok := jsonstory[req.RequestURI[1:]]; ok {
		// Parse storypart
		tmpl, err := template.ParseFiles("templates/story.html")
		if err != nil {
			fmt.Println(err)
		}
		tmpl.Execute(w, storypart)
	} else {
		fmt.Println("Unknown RequestURI", req.RequestURI[1:], "--> Redirect to intro!")
		http.Redirect(w, req, "/intro", 301)
	}
}

func main() {
	// Load json
	jsonstory = loadStory()

	fmt.Println("Hallo Welt :)")
	http.Handle("/css/", http.FileServer(http.Dir("templates/")))
	http.Handle("/js/", http.FileServer(http.Dir("templates/")))
	http.HandleFunc("/", storyHandler)
	http.HandleFunc("/intro", storyHandler)
	http.ListenAndServe("localhost:80", nil)
}
