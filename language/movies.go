package language

import (
	"bufio"
	"encoding/csv"
	"io"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"

	"github.com/olivia-ai/olivia/user"

	"github.com/olivia-ai/olivia/util"
)

// Movie is the serializer from res/datasets/movies.csv
type Movie struct {
	Name   string
	Genres []string
	Rating float64
}

var (
	// Initializes movies genres
	genres = []string{
		"Action", "Adventure", "Animation", "Children", "Comedy", "Crime", "Documentary", "Drama", "Fantasy",
		"Film-Noir", "Horror", "Musical", "Mystery", "Romance", "Sci-Fi", "Thriller", "War", "Western",
	}
	movies = SerializeMovies()
)

// SerializeMovies retrieves the content of res/datasets/movies.csv and serialize it
func SerializeMovies() (movies []Movie) {
	path := "res/datasets/movies.csv"
	bytes, err := os.Open(path)
	if err != nil {
		bytes, _ = os.Open("../" + path)
	}

	reader := csv.NewReader(bufio.NewReader(bytes))
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		// Convert the string to a float
		rating, _ := strconv.ParseFloat(line[3], 64)

		movies = append(movies, Movie{
			Name:   line[1],
			Genres: strings.Split(line[2], "|"),
			Rating: rating,
		})
	}

	return
}

// SearchMovie search a movie for a given genre
func SearchMovie(genre, userToken string) (output Movie) {
	for _, movie := range movies {
		userMovieBlacklist := user.GetUserInformation(userToken).MovieBlacklist
		// Continue if the movie is not from the request genre or if this movie has already been suggested
		if !util.Contains(movie.Genres, genre) || util.Contains(userMovieBlacklist, movie.Name) {
			continue
		}

		if reflect.DeepEqual(output, Movie{}) || movie.Rating > output.Rating {
			output = movie
		}
	}

	// Add the found movie to the user blacklist
	user.ChangeUserInformation(userToken, func(information user.Information) user.Information {
		information.MovieBlacklist = append(information.MovieBlacklist, output.Name)
		return information
	})

	return
}

// FindMoviesGenres returns an array of genres found in the entry string
func FindMoviesGenres(content string) (output []string) {
	for _, genre := range genres {
		if LevenshteinContains(strings.ToUpper(content), strings.ToUpper(genre), 2) {
			output = append(output, genre)
		}
	}

	return
}
