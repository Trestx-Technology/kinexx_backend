package db

import (
	trestCommon "github.com/Trestx-technology/trestx-common-go-lib"
	"github.com/ryanbradynd05/go-tmdb"
	"github.com/spf13/viper"
)

var tmdbAPI *tmdb.TMDb

func init() {
	trestCommon.LoadConfig()
	apiKey := viper.GetString("tmdb.key")
	config := tmdb.Config{
		APIKey:   apiKey,
		Proxies:  nil,
		UseProxy: false,
	}

	tmdbAPI = tmdb.Init(config)
}

func FindMovies(query string, page string) (*tmdb.MovieSearchResults, error) {
	option := map[string]string{"page": page}
	return tmdbAPI.SearchMovie(query, option)
}
func GetMovie(id []int, page string) ([]*tmdb.Movie, error) {
	option := map[string]string{"page": page}
	var movies []*tmdb.Movie
	for i := range id {
		movie, err := tmdbAPI.GetMovieInfo(id[i], option)
		if err == nil {
			movies = append(movies, movie)
		}
	}

	return movies, nil
}
