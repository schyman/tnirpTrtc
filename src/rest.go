package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"path"
	"strconv"
	"strings"
)

// Start the rest server and listen to port specified in config file
func startRestServer(restConfig RestConfig) {
	restServer := new(RestServer)
	port := strconv.Itoa(restConfig.Port)

	fmt.Println("Starting Rest server at " + port)
	err := http.ListenAndServe(":"+port, restServer)
	if err != nil {
		log.Fatal(err)
	}
}

type RestServer struct {
}

// Request Router
func (h *RestServer) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	// Return error if request type is not GET
	if req.Method != "GET" {
		http.Error(res, "Only GET is allowed", http.StatusMethodNotAllowed)
		return
	}

	match, _ := path.Match("*/chapter_versions/[0-9]*", req.URL.Path)
	if match {
		GetChapterVersions(res, req)
		return
	}

	// Unknwown request, return error
	http.Error(res, "Not Found", http.StatusNotFound)
}

func GetChapterVersions(res http.ResponseWriter, req *http.Request) {
	var head string

	id, err := strconv.Atoi(path.Base(req.URL.Path))
	// Return error if chapter id is not a integer
	if err != nil {
		http.Error(res, fmt.Sprintf("Invalid chapter id %q", head), http.StatusBadRequest)
		return
	}

	// Initialize chapter object with empty chapter versions object
	// else 'null' instead of '[]' would be returned as value of 'Versions' field if it is empty
	chapter := Chapter{Versions: []ChapterVersion{}}

	// Fetch chapter details from the database
	err = getChapter(&chapter, id)

	// Return error if there is some issue fetching details from the database
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	// Convert chapter object into json string
	response, err := json.Marshal(chapter)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	// Send response
	fmt.Fprintf(res, string(response))
}

// ShiftPath splits off the first component of p, which will be cleaned of
// relative components before processing. head will never contain a slash and
// tail will always be a rooted path without trailing slash.
func ShiftPath(p string) (head, tail string) {
	p = path.Clean("/" + p)
	i := strings.Index(p[1:], "/") + 1
	if i <= 0 {
		return p[1:], "/"
	}
	return p[1:i], p[i:]
}
