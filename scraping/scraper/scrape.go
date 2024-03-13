// This package demonstrates scraping and go docs (synopsis test)
//
// The scraper package is used to scrape a website for Pokemon data
package scraper

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"strconv"
	"sync"
	"time"
)

// baseUrl represents the base url for the website to scrape
const baseUrl = "https://scrapeme.live/shop"

// Pokemon represents a Pokemon entity
type Pokemon struct {
	Name string
}

// Set method sets the name of the Pokemon
func (p *Pokemon) Set(Name string) {
	p.Name = Name
}

// PokemonList represents a list of Pokemon and a mutex to lock the list
type PokemonList struct {
	mu      sync.Mutex
	pokemon []Pokemon
}

// Add method adds a Pokemon to the list
func (pm *PokemonList) Add(p Pokemon) {
	pm.mu.Lock()
	defer pm.mu.Unlock()

	pm.pokemon = append(pm.pokemon, p)
}

// Print method prints the list of Pokemon
func (pm *PokemonList) Print() {
	pm.mu.Lock()
	defer pm.mu.Unlock()

	for i, p := range pm.pokemon {
		fmt.Printf("%d: %s\n", i, p.Name)
	}
}

// Scrape function scrapes the website for Pokemon data
//
// ...multi line comment testing...
//
// - It takes an integer as an argument, representing the page number to scrape
func Scrape(i int) []string {
	var currentUrl string

	if i == 0 {
		currentUrl = baseUrl
	} else {
		currentUrl = baseUrl + "/page/" + strconv.Itoa(i)
	}

	res, err := http.Get(currentUrl)
	if err != nil {
		fmt.Printf("error fetching url for page: %d\t err: %v\n", i, err)
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		panic(fmt.Errorf("status code error: %d %s\t on page number %d", res.StatusCode, res.Status, i))
	}
	fmt.Println("scraping page number: ", i)

	time.Sleep(1 * time.Second)

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		panic(err)
	}

	var pokemonOnPage []string

	doc.Find("#main").Each(func(i int, s *goquery.Selection) {
		s.Find("ul").Each(func(i int, s *goquery.Selection) {
			s.Find("li h2").Each(func(i int, s *goquery.Selection) {
				pokemonOnPage = append(pokemonOnPage, s.Text())
			})
		})
	})

	return pokemonOnPage
}

// Notes:
// - net/http for making requests
// - goquery is a cool pkg for parsing html. Able to walk dom similar to jQuery

// After initial naive implementation - expanded the program to use:
// - mutexes to lock the shared data
// - waitgroups to wait for all goroutines to finish
// - a struct to store the data, with methods to add and print the data
// - abstracted the scraping logic into a function for a cleaner main
// godoc for documenting the code
