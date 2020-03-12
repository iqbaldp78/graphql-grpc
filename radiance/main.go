package main

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gogo/protobuf/jsonpb"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/micro/cli/v2"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/client"

	pbRfqs "github.com/mbizmarket/dmp/radiance/proto/pb/rfqs"
	pb "github.com/mbizmarket/dmp/radiance/proto/pb/user"
	"github.com/mbizmarket/dmp/radiance/repository"
	"github.com/mbizmarket/dmp/radiance/utils"
)

// Setup and the client
func runClient(service micro.Service) {
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Hour)
	defer cancel()
	srv := pb.NewRadianceServicesService("go.dmp.service.grpc.radiance", service.Client())
	fmt.Println(srv)

	in := &empty.Empty{}
	rsp, err := srv.GetAllUsers(ctx, in)
	if err != nil {
		fmt.Println("err GetAllUsers", err)
		return
	}

	// Print response
	// fmt.Println("rsp", rsp)

	// =========== as json string
	var buf bytes.Buffer
	err1 := (&jsonpb.Marshaler{}).Marshal(&buf, rsp)
	if err1 != nil {
		fmt.Println(err1.Error())
		os.Exit(0)
	}
	jsonString := buf.String()
	fmt.Printf("# ==== As JSON String\n       %v \n", jsonString)
}

func main() {

	// Create a new service. Optionally include some options here.
	srv := micro.NewService(
		micro.Name("go.dmp.service.grpc.radiance"),
		micro.RegisterTTL(time.Second*300),
		micro.RegisterInterval(time.Second*300),

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

	// RPC client
	clientAghanim := pbRfqs.NewAghanimServicesService("go.dmp.service.grpc.aghanim", client.DefaultClient)

	// Register handler
	pb.RegisterRadianceServicesHandler(srv.Server(), &repository.Service{DB: db, ClientAghanim: clientAghanim})

	// Run the server
	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
