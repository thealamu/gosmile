package main

import (
	"log"
	"net/http"
)

const jokeEndpoint = "https://official-joke-api.appspot.com/jokes/random"

// A Joke holds info about a joke
type Joke struct {
	ID                     int
	Type, Setup, Punchline string
}

func main() {
	response, err := http.Get(jokeEndpoint)
	if err != nil {
		log.Fatal(err)
	}
}
