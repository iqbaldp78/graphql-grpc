package main

import (
	"log"

	"github.com/99designs/gqlgen/handler"
	"github.com/micro/go-micro/v2/client"
	"github.com/micro/go-micro/v2/web"

	"github.com/mbizmarket/dmp/graphql-aghanim/graph"
	"github.com/mbizmarket/dmp/graphql-aghanim/graph/generated"
	pb"github.com/mbizmarket/dmp/graphql-aghanim/proto/pb/rfqs"
)

func main() {

	// create new web service
	service := web.NewService(
		web.Name("go.dmp.service.api.graphql.aghanim"),
		web.Version("latest"),
		web.Address(":8083"),
	)

	// initialise service
	if err := service.Init(); err != nil {
		log.Fatal(err)
	}

	// RPC client
	clientAghanim := pb.NewAghanimServicesService("go.dmp.service.grpc.aghanim", client.DefaultClient)

	// register graphql handlers
	service.Handle("/", handler.Playground("GraphQL playground", "/query"))
	service.Handle("/query", handler.GraphQL(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{
		ResolverClientAghanim: clientAghanim,
	}})))

	// run service
	if err := service.Run(); err != nil {
		log.Fatal(err)

	}

}
