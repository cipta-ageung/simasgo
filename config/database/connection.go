package database

import (
	"context"
	"log"

	"github.com/jinzhu/gorm"

	// driver postgres
	dbpb "github.com/cipta-ageung/simasgo/protobuf/db"
	microclient "github.com/micro/go-micro/client"
)

// Pgdb struct type
type Pgdb struct {
	*gorm.DB
}

// Connect database setup configuration
func Connect(connectionDb string) (*Pgdb, error) {
	db, err := gorm.Open("postgres", connectionDb)
	defer db.Close()
	if err != nil {
		return nil, err
	}

	return &Pgdb{db}, nil
}

// ConnectDb database setup configuration
func ConnectDb(svcName string) *Pgdb {

	/* use the generated client stub */
	serviceDb := dbpb.NewConnectionService(svcName, microclient.DefaultClient)

	svcDb, err := serviceDb.SetupDb(context.TODO(), &dbpb.ServiceApp{Svc: "go.micro.srv.simasuser"})
	if err != nil || svcDb == nil {
		log.Fatalf("cannot setup database")
	}

	log.Println(svcDb.ConnectionDb)
	pgdb, err := Connect(svcDb.ConnectionDb)
	if err != nil {
		log.Printf("cannot connect database")
	}

	return pgdb
}
