package main

import (
	"log"
	"time"

	"github.com/cipta-ageung/simasgo/protobuf/db"
	ms "github.com/cipta-ageung/simasgo/services/db/ms"

	_ "github.com/lib/pq"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/eureka"
)

func main() {

	reg := eureka.NewRegistry(func(op *registry.Options) {
		op.Addrs = []string{"http://admin:Metal7390@localhost:8761/eureka"}
	})
	// Initiate Service
	service := micro.NewService(
		micro.Name("go.micro.srv.simasdb"),
		micro.Registry(reg),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
		micro.Version("latest"),
		micro.Metadata(map[string]string{
			"type": "simas_db",
		}),
	)
	service.Init()

	// Register handler
	db.RegisterConnectionServiceHandler(service.Server(), new(ms.ConnectionService))

	// Run server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
