package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// Constants
const portNum string = ":6969"

// Create a struct containing all the current state.
type State struct {
	currentSplit int
	currentSplitName string
	lastSplitTime int
	lastSplitDiff int //time of previous split went compared to best of previous split
	currentTotalTime int
	currentTotalTimeDiff int //current time compared to best whole run current time at current split
}

var state = 0


// Update state on HTTP call. Takes struct as input.
func Hello (w http.ResponseWriter, r *http.Request) {
	state = state + 1
	log.Printf("Updated state to %d", state)
	fmt.Fprintf(w, "Current split: %d", state)
}

func eventsHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Expose-Headers", "Content-Type")
   
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
   
	// Simulate sending events (you can replace this with real data)
	for i := 0; i < 10; i++ {
	 fmt.Fprintf(w, "data: %s\n\n", fmt.Sprintf("Event %d", i))
	 log.Println("Event sent")
	 time.Sleep(2 * time.Second)
	 w.(http.Flusher).Flush()
	}
   
	// Simulate closing the connection
	closeNotify := r.Context().Done()
	<-closeNotify
}

func main() {

	fmt.Println("Hello world!")
	log.Println("starting http server.")

	http.HandleFunc("/boi", Hello)
	http.HandleFunc("/events", eventsHandler)
	http.Handle("/", http.FileServer(http.Dir("./static/")))
	
	log.Println("Started on port", portNum)

	err := http.ListenAndServe(portNum, nil)
	if err != nil  {
		log.Fatal(err)
	}

}