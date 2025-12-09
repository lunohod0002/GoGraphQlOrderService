package main

import (
	"OzonOrderService/graph"
	"OzonOrderService/internal/repositories"
	"OzonOrderService/internal/services"

	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	_ "github.com/lib/pq"
	"github.com/vektah/gqlparser/v2/ast"
)

const defaultPort = "8081"

type App struct {
	GQLResolver *graph.Resolver
}

func NewApp(db *sql.DB) *App {
	productRepo := repositories.NewProductRepository(db)
	productService := services.NewProductService(productRepo)
	userRepo := repositories.NewUserRepository(db)
	cartRepo := repositories.NewCartRepository(db)

	userService := services.NewUserService(cartRepo, userRepo)
	cartService := services.NewCartService(cartRepo)
	resolver := graph.NewResolver(productService, cartService, userService)

	return &App{
		GQLResolver: resolver,
	}
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	db, err := sql.Open("postgres", "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable")
	if err != nil {
		panic(err)
	}
	app := NewApp(db)
	srv := handler.New(graph.NewExecutableSchema(graph.Config{Resolvers: app.GQLResolver}))

	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})

	srv.SetQueryCache(lru.New[*ast.QueryDocument](1000))

	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New[string](100),
	})

	http.Handle("/", playground.Handler("GraphQL playground", "/graphql"))
	http.Handle("/graphql", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
