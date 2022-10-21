package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/jetrica/cyoa/cyoaweb"
)

func main() {
	// Create flags for our optional variables
	port := flag.Int("port", 3000, "the port to start the CYOA web application on")
	filename := flag.String("file", "./cyoaweb/story.json", "the JSON file with the CYOA story")
	flag.Parse()
	fmt.Printf("Using the story in %s.\n", *filename)

	// Open the JSON file and parse the story in it.
	f, err := os.Open(*filename)
	if err != nil {
		panic(err)
	}
	story, err := cyoaweb.JsonStory(f)
	if err != nil {
		panic(err)
	}

	// Create our custom CYOA story handler
	h := cyoaweb.NewHandler(story)

	// Create a ServeMux to route our requests
	mux := http.NewServeMux()
	// This story handler is using a custom function and template
	// Because we use /story/ (trailing slash) all web requests
	// whose path has the /story/ prefix will be routed here.
	mux.Handle("/", h)
	// This story handler is using the default functions and templates
	// Because we use / (base path) all incoming requests not
	// mapped elsewhere will be sent here.
	fmt.Printf("Starting the server on port: %d\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), mux))
}
