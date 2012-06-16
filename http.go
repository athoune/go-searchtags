package main

import (
	/*"encoding/json"*/
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func toInt(s []string) []uint64 {
	ints := make([]uint64, len(s))
	for i, value := range s {
		var parsed int64
		parsed, _ = strconv.ParseInt(value, 10, 64) //[FIXME] handle error
		ints[i] = uint64(parsed)
	}
	return ints
}

func doSearch(w http.ResponseWriter, req *http.Request) {
	q := req.URL.Query()["q"]
	h := w.Header()
	h.Set("Content-Type", "text/event-stream")
	h.Set("Connection", "keep-alive")
	w.WriteHeader(200)
	w.Write([]byte("["))
	docs.Score(
		NewDocument(toInt(strings.Split(q[0], ","))),
		0.2,
		func(docid uint32, score uint64) {
			w.Write([]byte(fmt.Sprintf("%d,", docid)))
		})
	w.Write([]byte("]"))
}

func startHttp() {
	http.HandleFunc("/search", doSearch)
	log.Printf("About to start http://localhost:8000")
	err := http.ListenAndServe("localhost:8000", nil)
	if err != nil {
		panic(err)
	}
}
