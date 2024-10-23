# FlickSearch

**FlickSearch** is a simple command-line tool to search for movies and open streaming links directly from the terminal. The tool fetches movie data from the OMDB API and allows users to select a movie, then automatically opens the movie's streaming link in their default browser.

## Features

- **Search for movies by title** from the command line.
- **Select a movie** from the search results.
- **Automatically opens the selected movie's streaming link** in your default browser.
- Works on **Linux**, **macOS**, and **Windows**.

## Installation

1. **Clone the repository**:

   ```bash
   git clone https://github.com/stablelane/FlickSearch.git  

2. **Navigate to the project directory**:

   ```bash
   cd FlickSearch

3. **Build the project**:

   ```bash
   go build
   
4. **Run the tool**:

   ```bash
   ./flicksearch <movie-name>
