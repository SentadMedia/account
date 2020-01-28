//+build wireinject

package dep

import (
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
	"github.com/sentadmedia/account/app/adapter/db"
	"github.com/sentadmedia/account/app/adapter/rpc"
	"github.com/sentadmedia/account/app/adapter/rpc/proto"
	"github.com/sentadmedia/account/app/usecase"
	"github.com/sentadmedia/account/app/usecase/accounts"
	"github.com/sentadmedia/account/app/usecase/repository"
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

func InitCommandFactory() fw.CommandFactory {
	wire.Build(
		wire.Bind(new(fw.CommandFactory), new(mdcli.CobraFactory)),
		mdcli.NewCobraFactory,
	)
	return mdcli.CobraFactory{}
}

// InjectEnvironment creates Environment with configured dependencies.
func InjectEnvironment() fw.Environment {
	wire.Build(
		wire.Bind(new(fw.Environment), new(mdenv.GoDotEnv)),
		mdenv.NewGoDotEnv,
	)
	return mdenv.GoDotEnv{}
}

var observabilitySet = wire.NewSet(
	wire.Bind(new(fw.Logger), new(mdlogger.Local)),
	mdlogger.NewLocal,
	mdtracer.NewLocal,
)

func InitGRpcService(
	name string,
	logLevel fw.LogLevel,
	sqlDB *gorm.DB,
	securityPolicy fw.SecurityPolicy,
) (mdservice.Service, error) {
	wire.Build(
		wire.Bind(new(fw.StdOut), new(mdio.StdOut)),
		wire.Bind(new(fw.ProgramRuntime), new(mdruntime.BuildIn)),
		wire.Bind(new(fw.Server), new(mdgrpc.GRpc)),
		wire.Bind(new(fw.GRpcAPI), new(rpc.AccountAPI)),
		wire.Bind(new(proto.AccountServer), new(rpc.AccountServer)),
		wire.Bind(new(accounts.Producer), new(accounts.ProducerPersist)),
		wire.Bind(new(accounts.Consumer), new(accounts.ConsumerPersist)),
		wire.Bind(new(repository.Account), new(db.AccountSQL)),

		observabilitySet,

		mdio.NewBuildInStdOut,
		mdruntime.NewBuildIn,
		mdtimer.NewTimer,
		mdgrpc.NewGRpc,
		mdservice.New,

		rpc.NewAccountServer,
		rpc.NewAccountAPI,
		usecase.NewUseCase,
		accounts.NewProducerPersist,
		accounts.NewConsumerPersist,
		db.NewAccountSQL,
	)
	return mdservice.Service{}, nil
}

// InitDBConnector creates DBConnector with configured dependencies.
func InitDBConnector() fw.DBConnector {
	wire.Build(
		wire.Bind(new(fw.DBConnector), new(mddb.GormPstgresConnector)),
		mddb.NewPostgresConnector,
	)
	return mddb.GormPstgresConnector{}
}
