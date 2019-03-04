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

func startRestServer(restConfig RestConfig) {
	restServer := &RestServer{
		ChapterHandler: new(ChapterHandler),
	}
	port := strconv.Itoa(restConfig.Port)

	fmt.Println("Starting Rest server at " + port)
	err := http.ListenAndServe(":"+port, restServer)
	if err != nil {
		log.Fatal(err)
	}
}

type RestServer struct {
	ChapterHandler *ChapterHandler
}

func (h *RestServer) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	var head string
	head, req.URL.Path = ShiftPath(req.URL.Path)
	if head == "chapter_versions" {
		h.ChapterHandler.ServeHTTP(res, req)
		return
	}
	http.Error(res, "Not Found", http.StatusNotFound)
}

type ChapterHandler struct {
}

func (h *ChapterHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	var head string
	head, req.URL.Path = ShiftPath(req.URL.Path)
	id, err := strconv.Atoi(head)
	if err != nil {
		http.Error(res, fmt.Sprintf("Invalid chapter id %q", head), http.StatusBadRequest)
		return
	}
	if req.Method != "GET" {
		http.Error(res, "Only GET is allowed", http.StatusMethodNotAllowed)
		return
	}

	chapter := Chapter{Versions: []ChapterVersion{}}
	err = getChapter(&chapter, id)
	if err != nil {
		http.Error(res, err.Error(), 500)
		return
	}

	response, err := json.Marshal(chapter)
	if err != nil {
		http.Error(res, err.Error(), 500)
		return
	}

	res.Header().Set("Content-Type", "application/json")
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
