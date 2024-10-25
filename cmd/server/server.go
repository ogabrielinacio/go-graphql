package main

import (
	"database/sql"
	"fmt"

	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	_ "github.com/lib/pq"
	"github.com/ogabrielinacio/go-graphql/graph"
	"github.com/ogabrielinacio/go-graphql/graph/generated"
	"github.com/ogabrielinacio/go-graphql/internal/database"
	"github.com/ogabrielinacio/go-graphql/utils"
)

func main() {
	connStr := utils.GetConnectionString()
	fmt.Println(connStr)

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatalf("failed to open database: %v", err)
	}

	defer db.Close()

	categoryDb := database.NewCategory(db)
	courseDb := database.NewCourse(db)

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{
		CategoryDB: categoryDb,
		CourseDB:   courseDb,
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Print("connect to http://localhost:8080/ for GraphQL playground")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
