package main // import "github.com/obukhov/pantomime/server"

import (
	"fmt"
	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"github.com/joho/godotenv"
	"github.com/obukhov/pantomime/server/resolver"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	schemaString, err := getSchema("./schema.graphql")
	if nil != err {
		log.Fatal(err)
	}

	err = godotenv.Load("./.env", "./.env.dist")
	if err != nil {
		log.Print("Could not load .env file")
	} else {
		log.Print("Loaded .env file")
	}

	schema := graphql.MustParseSchema(
		schemaString,
		&resolver.Resolver{},
		graphql.UseStringDescriptions(),
		graphql.UseFieldResolvers(),
		graphql.MaxParallelism(20),
	)

	http.Handle("/query", &relay.Handler{Schema: schema})

	serverPort := os.Getenv("SERVER_PORT")
	log.Println(fmt.Sprintf("Start listening on %s...", serverPort))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", serverPort), nil))
}

func getSchema(path string) (string, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}

	return string(b), nil
}
