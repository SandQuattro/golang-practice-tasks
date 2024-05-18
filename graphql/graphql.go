package main

import (
	"encoding/json"
	"errors"
	"github.com/graphql-go/graphql"
	"net/http"
)

type User struct {
	ID    string
	Name  string
	Email string
}

var userType = graphql.NewObject(graphql.ObjectConfig{
	Name: "User",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.ID,
		},
		"name": &graphql.Field{
			Type: graphql.String,
		},
		"email": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var queryType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Query",
	Fields: graphql.Fields{
		"user": &graphql.Field{
			Type: userType,
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.ID),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				id, ok := p.Args["id"].(string)

				if !ok {
					return nil, errors.New("missing id argument")
				}

				user := User{id, "тестовый пользак", "test@test.ru"}
				// Здесь должен быть код для получения пользователя по id

				return user, nil
			},
		},
	},
})

var schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query: queryType,
})

func handleGraphQL(w http.ResponseWriter, r *http.Request) {
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: r.URL.Query().Get("query"),
	})

	json.NewEncoder(w).Encode(result)
}

func main() {
	http.HandleFunc("/graphql", handleGraphQL)
	http.ListenAndServe(":8000", nil)
}
