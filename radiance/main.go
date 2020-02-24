package main

import (
	"context"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/micro/cli/v2"
	"github.com/micro/go-micro/v2"

	"github.com/mbizmarket/dmp/radiance/proto/users/pb"
	"github.com/mbizmarket/dmp/radiance/repository"
	"github.com/mbizmarket/dmp/radiance/utils"
)

// Setup and the client
func runClient(service micro.Service) {
	srv := pb.NewRadianceServicesService("go.dmp.service.grpc.radiance", service.Client())
	fmt.Println(srv)

	in := &empty.Empty{}
	rsp, err := srv.GetAllUsers(context.Background(), in)
	if err != nil {
		fmt.Println("err GetAllUsers", err)
		return
	}

	// Print response
	fmt.Println("rsp", rsp)
}

func main() {

	// Create a new service. Optionally include some options here.
	srv := micro.NewService(
		micro.Name("go.dmp.service.grpc.radiance"),

		// Setup some flags. Specify --run_client to run the client
		// Add runtime flags
		// We could do this below too
		micro.Flags(&cli.BoolFlag{
			Name:  "run_client",
			Usage: "Launch the client",
		}),
	)

	// Init will parse the command line flags.
	srv.Init(
		// Add runtime action
		// We could actually do this above
		micro.Action(
			func(c *cli.Context) (err error) {
				if c.Bool("run_client") {
					runClient(srv)
					os.Exit(0)
				}
				return
			}),
	)

	utils.SetConfigViper("config", "./")

	db := utils.GetDBConnection()
	defer db.Close()

	// Register handler
	pb.RegisterRadianceServicesHandler(srv.Server(), &repository.Service{DB: db})

	// Run the server
	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
