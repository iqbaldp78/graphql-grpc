package main

import (
	"log"

	"github.com/99designs/gqlgen/handler"
	"github.com/micro/go-micro/v2/client"
	"github.com/micro/go-micro/v2/web"

	"github.com/mbizmarket/dmp/graphql/graph"
	"github.com/mbizmarket/dmp/graphql/graph/generated"
	"github.com/mbizmarket/dmp/graphql/proto/users/pb"
)

func main() {

	// create new web service
	service := web.NewService(
		web.Name("go.dmp.service.api.graphql"),
		web.Version("latest"),
		web.Address(":8080"),
	)

	// initialise service
	if err := service.Init(); err != nil {
		log.Fatal(err)
	}

	// RPC client
	clientRadiance := pb.NewRadianceServicesService("go.dmp.service.grpc.radiance", client.DefaultClient)

	// register graphql handlers
	service.Handle("/", handler.Playground("GraphQL playground", "/query"))
	service.Handle("/query", handler.GraphQL(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{
		ResolverClientRadiance: clientRadiance,
	}})))

	// run service
	if err := service.Run(); err != nil {
		log.Fatal(err)

	}

}
