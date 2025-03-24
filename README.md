![Ash Ketchum](https://www.cnet.com/a/img/resize/4b1ab53ded9295c6eb752134fc178251778e10dd/hub/2022/11/11/61b6d3b1-5b59-4564-812b-d9af287d92de/ash-ketchum-champ.jpg?auto=webp&fit=crop&height=675&width=1200)

# Pokedex CLI Application

Thank you [Boot.dev](https://www.boot.dev/) for this fun and educational project 

This project is a command-line interface (CLI) application that interacts with the [PokeAPI](https://pokeapi.co/) to provide information about Pokémon, locations, and encounters. It also includes features like exploring locations, catching Pokémon, and managing a personal Pokedex.

## What I Learned and Liked About This Project

Coding this project really taught me a lot. It solidified my understanding of APIs, HTTP requests, and significantly boosted my confidence in working with JSON data.

Implementing the basic caching system was particularly interesting because, as simple as it may sound, it was the first time I added a feature that wasn't directly related to the application's core functionality. Instead, the caching system served as an optimization layer, improving the program's performance by reducing redundant processing and network calls. This experience helped me better appreciate the importance of auxiliary features.

Tests also stop feeling like arcane magic that only "Real Developer" can create.

## Features

- **Explore Locations**: Fetch and display Pokémon encounters in specific locations.
- **Catch Pokémon**: Attempt to catch Pokémon with a success rate based on their base experience.
- **Inspect Pokémon**: View detailed stats and information about Pokémon in your Pokedex.
- **Navigate Locations**: Move forward or backward through location maps.
- **Caching**: Implements an in-memory cache to reduce API calls and improve performance.
- **Help Command**: Lists all available commands and their descriptions.
- **Exit Command**: Gracefully exits the application.

## Commands

| Command   | Description                                                                 |
|-----------|-----------------------------------------------------------------------------|
| `help`    | Lists all available commands.                                              |
| `exit`    | Exits the application.                                                     |
| `map`     | Fetches and displays the next 20 locations.                                |
| `mapb`    | Fetches and displays the previous 20 locations.                            |
| `explore` | Explores a location and lists Pokémon encounters.                          |
| `catch`   | Attempts to catch a Pokémon by name.                                       |
| `inspect` | Displays detailed information about a Pokémon in your Pokedex.            |
| `pokedex` | Lists all Pokémon currently in your Pokedex.                               |

## How to Run

1. Clone the repository:
   ```bash
   git clone https://github.com/maxBRT/pokedex.git
   cd pokedex
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Run the application:
   ```bash
   go run main.go
   ```

## Dependencies

- [Go](https://golang.org/) (version 1.23.7 or higher)



