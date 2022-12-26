package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

var client *http.Client

type CatFact struct {
	Fact   string `json:"fact"`
	Length int    `json:"length"`
}

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func getJson(url string, target interface{}) error {
	resp, err := client.Get(url)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(target)
}
func (a *App) GetCatFact() string {
	url := "https://catfact.ninja/fact"
	client = &http.Client{Timeout: 10 * time.Second}

	var catFact CatFact
	// return url
	err := getJson(url, &catFact)
	if err != nil {
		return fmt.Sprintf("error getting cat fact: %s\n", err.Error())
	} else {
		return fmt.Sprintf("A super interesting Cat Fact: %s\n", catFact.Fact)
	}
}
