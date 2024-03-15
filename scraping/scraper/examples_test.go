package scraper_test

// Notes - this file is testing how to use Examples in godoc

import (
	"fmt"
	"github.com/sdeleon-bjss/scraping/scraper"
)

func ExamplePokemon_Set() {
	p := &scraper.Pokemon{}
	p.Set("Pikachu")
	fmt.Println(p.Name)
	// Output: Pikachu
}

func ExamplePokemonList_Add() {
	pm := &scraper.PokemonList{}
	pm.Add(scraper.Pokemon{Name: "Pikachu"})
	pm.Print()
	// Output: 0: Pikachu
}

func ExampleScrape() {
	p := scraper.Scrape(1)
	fmt.Println(p)
	// Output: [Bulbasaur IvySaur Venusaur Charmander Charmeleon Charizard Squirtle Wartortle Blastoise Caterpie Metapod Butterfree Weedle Kakuna Beedrill Pidgey]
}

func ExamplePokemonList_Print() {
	pm := &scraper.PokemonList{}
	pm.Add(scraper.Pokemon{Name: "Pikachu"})
	pm.Print()
	// Output: 0: Pikachu
}

// This is an example of how the output will look with several Pokemon
func ExamplePokemonList_Print_several() {
	pm := &scraper.PokemonList{}
	pm.Add(scraper.Pokemon{Name: "Pikachu"})
	pm.Add(scraper.Pokemon{Name: "Bulbasaur"})
	pm.Add(scraper.Pokemon{Name: "Charmander"})
	pm.Print()
	// Output: 0: Pikachu
	// 1: Bulbasaur
	// 2: Charmander
}
