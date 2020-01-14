package main

import (
	"log"
	"time"

	_ "github.com/jinzhu/gorm/dialects/postgres"

	db "github.com/cipta-ageung/simasgo/config/database"
	user "github.com/cipta-ageung/simasgo/protobuf/user"
	ms "github.com/cipta-ageung/simasgo/services/user/ms"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/etcdv3"
)

func main() {

	log.Println("starting services user . . .")

	reg := etcdv3.NewRegistry(func(op *registry.Options) {
		op.Addrs = []string{"http://127.0.0.1:2379"}
	})

	service := micro.NewService(
		micro.Name("go.micro.srv.simasuser"),
		micro.Registry(reg),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
		micro.Version("latest"),
		micro.Metadata(map[string]string{
			"type": "simas_user",
		}),
	)

	service.Init()
	user.RegisterUserServiceHandler(service.Server(), &ms.UserService{})

	db.ConnectDb("go.micro.srv.simasdb")
	/* Run server */
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}

}
