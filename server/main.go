package main // import "github.com/obukhov/pantomime/server"

import (
	graphql "github.com/graph-gophers/graphql-go"
	relay "github.com/graph-gophers/graphql-go/relay"
	"github.com/obukhov/pantomime/server/resolver"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	schemaString, err := getSchema("../schema.graphql")
	if nil != err {
		log.Fatal(err)
	}

	schema := graphql.MustParseSchema(
		schemaString,
		&resolver.Resolver{},
		graphql.UseStringDescriptions(),
		graphql.UseFieldResolvers(),
		graphql.MaxParallelism(20),
	)

	http.Handle("/query", &relay.Handler{Schema: schema})

	log.Println("Start listening on 8083...")
	log.Fatal(http.ListenAndServe(":8083", nil))
}

func getSchema(path string) (string, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}

	return string(b), nil
}
