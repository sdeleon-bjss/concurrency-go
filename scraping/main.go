package main

import (
	"github.com/sdeleon-bjss/scraping/scraper"
	"sync"
)

var (
	wg   sync.WaitGroup
	data = scraper.PokemonList{}
)

func main() {
	println("Beginning scraper...")

	// 48 pagination pages in total
	for i := range 48 {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()
			p := scraper.Scrape(i)

			for _, pokemon := range p {
				data.Add(scraper.Pokemon{Name: pokemon})
			}
		}(i)
	}

	wg.Wait()
	data.Print()
}
