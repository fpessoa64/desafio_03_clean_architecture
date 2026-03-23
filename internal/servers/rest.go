package servers

import (
	"log"
	"net/http"

	"github.com/fpessoa64/desafio03_clean_arch/internal/handlers/rest"
	"github.com/fpessoa64/desafio03_clean_arch/internal/handlers/rest/routes"
	"github.com/fpessoa64/desafio03_clean_arch/internal/usecase"
	"github.com/go-chi/chi/v5"
)

type Rest struct {
	port string
}

func NewRest(port string) *Rest {
	return &Rest{port: port}
}

func (rs *Rest) Start(uc *usecase.OrderUsecase) error {
	h := rest.NewHandler(uc)
	r := chi.NewRouter()
	routes.RegisterOrderRoutes(r, h)
	log.Printf("REST API running on :%s", rs.port)
	return http.ListenAndServe(":"+rs.port, r)
}
