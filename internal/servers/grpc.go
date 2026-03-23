package servers

import (
	"log"
	"net"

	"github.com/fpessoa64/desafio03_clean_arch/internal/handlers/grpc/service"
	"github.com/fpessoa64/desafio03_clean_arch/internal/usecase"
	orderpb "github.com/fpessoa64/desafio03_clean_arch/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Grpc struct {
	port string
}

func NewGrpc(port string) *Grpc {
	return &Grpc{port: port}
}

func (g *Grpc) Start(uc *usecase.OrderUsecase) error {
	lis, err := net.Listen("tcp", ":"+g.port)
	if err != nil {
		return err
	}
	srv := grpc.NewServer()
	orderpb.RegisterOrderServiceServer(srv, service.NewOrderServiceServer(uc))
	reflection.Register(srv)
	log.Printf("gRPC server running on :%s", g.port)
	return srv.Serve(lis)
}
