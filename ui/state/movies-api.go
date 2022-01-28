package state

import (
	"github.com/vugu-examples/taco-store/internal/memstore"
	"github.com/machinebox/graphql"
	"context"
)

type MovieListAPi struct {
	MovieList []memstore.Movie
}

func (c *MovieListAPi) GetMovieList() ([]memstore.Movie, error) {
	graphqlClient := graphql.NewClient("http://localhost:8080/graphql")
	graphqlRequest := graphql.NewRequest(`
		movies {
			id
			title
			genre
			imgURL
			description
			releaseDate
			length
			likes
			comments
		}
	`)

	var graphqlResposne []memstore.Movie 
	err := graphqlClient.Run(context.Background(), graphqlRequest, &graphqlResposne)
	if err != nil {
		panic(err)
	}
	//fmt.Println(graphqlResposne)
	
	return graphqlResposne, err
}
