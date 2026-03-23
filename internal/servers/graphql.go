package servers

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/fpessoa64/desafio03_clean_arch/graph"
	"github.com/fpessoa64/desafio03_clean_arch/internal/usecase"
)

type GraphQL struct {
	port string
}

func NewGraphQL(port string) *GraphQL {
	return &GraphQL{port: port}
}

func (g *GraphQL) Start(uc *usecase.OrderUsecase) error {
	srv := handler.New(graph.NewExecutableSchema(graph.Config{
		Resolvers: &graph.Resolver{UC: uc},
	}))
	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})
	srv.Use(extension.Introspection{})

	mux := http.NewServeMux()
	mux.Handle("/", playground.Handler("GraphQL Orders", "/query"))
	mux.Handle("/query", srv)

	log.Printf("GraphQL server running on :%s", g.port)
	return http.ListenAndServe(":"+g.port, mux)
}
