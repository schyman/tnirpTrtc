package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"path"
	"strconv"
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
	match, _ := path.Match("*/chapter_versions/[0-9]*", req.URL.Path)
	if match {
		h.GetChapterVersions(res, req)
		return
	}

	// Unknwown request, return error
	http.Error(res, "Not Found", http.StatusNotFound)
}

func (h *RestServer) GetChapterVersions(res http.ResponseWriter, req *http.Request) {
	// Return error if request type is not GET
	if req.Method != "GET" {
		http.Error(res, "Only GET is allowed", http.StatusMethodNotAllowed)
		return
	}

	id := path.Base(req.URL.Path)
	chapterId, err := strconv.Atoi(id)
	// Return error if chapter id is not a integer
	if err != nil {
		http.Error(res, fmt.Sprintf("Invalid chapter id %q", id), http.StatusBadRequest)
		return
	}

	// Fetch chapter details from the database
	chapter, err := getChapter(chapterId)

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