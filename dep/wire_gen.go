// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package dep

import (
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
	"github.com/sentadmedia/account/app/adapter/db"
	"github.com/sentadmedia/account/app/adapter/rpc"
	"github.com/sentadmedia/account/app/usecase"
	"github.com/sentadmedia/account/app/usecase/accounts"
	"github.com/sentadmedia/elf/fw"
	"github.com/sentadmedia/elf/modern/mdcli"
	"github.com/sentadmedia/elf/modern/mddb"
	"github.com/sentadmedia/elf/modern/mdenv"
	"github.com/sentadmedia/elf/modern/mdgrpc"
	"github.com/sentadmedia/elf/modern/mdio"
	"github.com/sentadmedia/elf/modern/mdlogger"
	"github.com/sentadmedia/elf/modern/mdruntime"
	"github.com/sentadmedia/elf/modern/mdservice"
	"github.com/sentadmedia/elf/modern/mdtimer"
	"github.com/sentadmedia/elf/modern/mdtracer"
)

// Injectors from wire.go:

func InitCommandFactory() fw.CommandFactory {
	cobraFactory := mdcli.NewCobraFactory()
	return cobraFactory
}

func InjectEnvironment() fw.Environment {
	goDotEnv := mdenv.NewGoDotEnv()
	return goDotEnv
}

func InitGRpcService(name string, logLevel fw.LogLevel, sqlDB *gorm.DB, securityPolicy fw.SecurityPolicy) (mdservice.Service, error) {
	stdOut := mdio.NewBuildInStdOut()
	timer := mdtimer.NewTimer()
	buildIn := mdruntime.NewBuildIn()
	local := mdlogger.NewLocal(name, logLevel, stdOut, timer, buildIn)
	accountSQL := db.NewAccountSQL(sqlDB, local)
	producerPersist := accounts.NewProducerPersist(accountSQL, local)
	consumerPersist := accounts.NewConsumerPersist(accountSQL)
	useCase := usecase.NewUseCase(local, producerPersist, consumerPersist)
	accountServer := rpc.NewAccountServer(useCase, local)
	accountAPI := rpc.NewAccountAPI(accountServer)
	gRpc, err := mdgrpc.NewGRpc(accountAPI, securityPolicy)
	if err != nil {
		return mdservice.Service{}, err
	}
	service := mdservice.New(name, gRpc, local)
	return service, nil
}

func InitDBConnector() fw.DBConnector {
	gormPstgresConnector := mddb.NewPostgresConnector()
	return gormPstgresConnector
}

// wire.go:

var observabilitySet = wire.NewSet(wire.Bind(new(fw.Logger), new(mdlogger.Local)), mdlogger.NewLocal, mdtracer.NewLocal)
