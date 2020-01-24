//+build wireinject

package dep

import (
	"github.com/google/wire"
	"github.com/sentadmedia/account/app/adapter/rpc"
	"github.com/sentadmedia/account/app/adapter/rpc/proto"
	"github.com/sentadmedia/account/app/usecase"
	"github.com/sentadmedia/elf/fw"
	"github.com/sentadmedia/elf/modern/mdcli"
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
	securityPolicy fw.SecurityPolicy,
) (mdservice.Service, error) {
	wire.Build(
		wire.Bind(new(fw.StdOut), new(mdio.StdOut)),
		wire.Bind(new(fw.ProgramRuntime), new(mdruntime.BuildIn)),
		wire.Bind(new(fw.Server), new(mdgrpc.GRpc)),
		wire.Bind(new(fw.GRpcAPI), new(rpc.AccountAPI)),
		wire.Bind(new(proto.AccountAPIServer), new(rpc.AccountAPIServer)),
		// wire.Bind(new(keys.Producer), new(keys.ProducerPersist)),
		// wire.Bind(new(keys.Consumer), new(keys.ConsumerCached)),
		// wire.Bind(new(gen.Generator), new(gen.Alphabet)),
		// wire.Bind(new(repo.AvailableKey), new(db.AvailableKeySQL)),
		// wire.Bind(new(repo.AllocatedKey), new(db.AllocatedKeySQL)),

		observabilitySet,

		mdio.NewBuildInStdOut,
		mdruntime.NewBuildIn,
		mdtimer.NewTimer,
		mdgrpc.NewGRpc,
		mdservice.New,

		rpc.NewAccountAPIServer,
		rpc.NewAccountAPI,
		usecase.NewUseCase,
		// provider.NewHTML,
		// provider.NewEmailNotifier,
		// keys.NewProducerPersist,
		// provider.NewConsumer,
		// keys.NewConsumerPersist,
		// db.NewAvailableKeySQL,
		// db.NewAllocatedKeySQL,
		// gen.NewAlphabet,
		// gen.NewBase62,
	)
	return mdservice.Service{}, nil
}
