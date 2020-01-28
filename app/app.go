package app

import (
	"github.com/sentadmedia/account/app/entity"
	"github.com/sentadmedia/account/dep"
	"github.com/sentadmedia/elf/fw"
)

// Config represents require parameters for the backend APIs
type Config struct {
	LogLevel    fw.LogLevel
	ServiceName string
	GRpcAPIPort int
}

// Start launches the GraphQL & HTTP APIs
func Start(
	dbConfig fw.DBConfig,
	dbConnector fw.DBConnector,
	config Config,
	securityPolicy fw.SecurityPolicy,
) {
	db, err := dbConnector.Connect(dbConfig)
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&entity.Account{})

	gRPCService, err := dep.InitGRpcService(
		config.ServiceName,
		config.LogLevel,
		db,
		securityPolicy,
	)
	if err != nil {
		panic(err)
	}
	gRPCService.StartAndWait(config.GRpcAPIPort)
}
