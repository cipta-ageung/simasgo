package ms

import (
	"context"
	"log"
	"os"

	"github.com/cipta-ageung/simasgo/protobuf/db"
	"github.com/micro/go-micro/config"
)

// ConnectionService : struct
type ConnectionService struct{}

// SetupDb : method
func (g *ConnectionService) SetupDb(ctx context.Context, req *db.ServiceApp, rsp *db.ServiceDb) error {

	// clear or empty first for new request
	//rsp = &db.ServiceDb{}

	// load data from file configuration
	config.LoadFile(os.Getenv("HOME") + "/go/src/simas/simasgo/config/db.json")

	// set data configuration into type struct ServiceDb
	config.Get("services", req.Svc).Scan(&rsp)
	// check data struct or response
	if rsp != nil {

		// set data response string connection
		rsp.ConnectionDb = "postgres://" + rsp.GetUser() + ":" + rsp.GetPassword() + "@" + rsp.GetHost() +
			":" + rsp.GetPort() + "/" + rsp.GetDbname() + "?sslmode=disable"
	}

	log.Print(rsp)

	return nil
}
