package scraper

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"strconv"
	"sync"
	"time"
)

const baseUrl = "https://scrapeme.live/shop"

type Pokemon struct {
	Name string
}

func (p *Pokemon) Set(Name string) {
	p.Name = Name
}

type PokemonList struct {
	mu      sync.Mutex
	pokemon []Pokemon
}

func (pm *PokemonList) Add(p Pokemon) {
	pm.mu.Lock()
	defer pm.mu.Unlock()

	pm.pokemon = append(pm.pokemon, p)
}

func (pm *PokemonList) Print() {
	pm.mu.Lock()
	defer pm.mu.Unlock()

	for i, p := range pm.pokemon {
		fmt.Printf("%d: %s\n", i, p.Name)
	}
}

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
