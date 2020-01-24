package app

import (
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
	config Config,
	securityPolicy fw.SecurityPolicy,
) {
	gRPCService, err := dep.InitGRpcService(
		config.ServiceName,
		config.LogLevel,
		securityPolicy,
	)
	if err != nil {
		panic(err)
	}
	gRPCService.StartAndWait(config.GRpcAPIPort)
}
