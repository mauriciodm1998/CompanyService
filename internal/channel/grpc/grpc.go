package grpc

import (
	"context"
	"log"
	"net"

	"github.com/mauriciodm1998/CompanyService/internal/canonical"
	"github.com/mauriciodm1998/CompanyService/internal/service"
	"github.com/mauriciodm1998/CompanyService/pb"

	grpc "google.golang.org/grpc"
)

type grpcChan struct {
	service service.Service
	pb.UnimplementedCompanyServerServer
}

func New() *grpcChan {
	return &grpcChan{}
}

func (g *grpcChan) Start() {
	server := grpc.NewServer()
	listener, err := net.Listen("tcp", ":50055")
	if err != nil {
		log.Fatal("Cannot open the grpc server")
	}

	pb.RegisterCompanyServerServer(server, g)

	server.Serve(listener)
}

func (g *grpcChan) Get(ctx context.Context, request *pb.Id) (*pb.Company, error) {
	company, err := g.service.Get(ctx, request.Id)
	if err != nil {
		return nil, err
	}

	companyReturn := pb.Company{
		Id:   company.Id,
		Name: company.Name,
	}

	return &companyReturn, nil
}

func (g *grpcChan) Create(ctx context.Context, request *pb.Company) (string, error) {

	company := canonical.Company{
		Id:   "",
		Name: request.Name,
	}

	id, err := g.service.Create(ctx, company)
	if err != nil {
		return "", err
	}

	return id, nil
}
