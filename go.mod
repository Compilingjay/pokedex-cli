module github.com/Compilingjay/pokedexcli

go 1.23.0

require pokeapi v0.0.0

require pokecache v0.0.0 // indirect

replace (
	pokeapi v0.0.0 => ./internal/pokeapi
	pokecache v0.0.0 => ./internal/pokecache
)
