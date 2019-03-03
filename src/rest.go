package main

import (
	"fmt"
	"io/ioutil"
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

	go func() {
		err := http.ListenAndServe(":"+port, restServer)
		if err != nil {
			log.Fatal(err)
		}
	}()

	fmt.Println("Started Rest server at " + port)

	requestServer()
}

func requestServer() {
	resp, err := http.Get("http://localhost:3001/chapter_versions/311111?skjdfn")
	fmt.Println(err)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Printf("\nWebserver said: `%s`", string(body))
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

	// h.handleGet(id)
	http.Error(res, fmt.Sprintf("actually it was successful %d", id), http.StatusBadRequest)
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
