package main

import (
	"strconv"

	"github.com/sentadmedia/account/app"
	"github.com/sentadmedia/account/cmd"
	"github.com/sentadmedia/account/dep"
	"github.com/sentadmedia/elf/fw"
)

func main() {
	env := dep.InjectEnvironment()
	env.AutoLoadDotEnvFile()

	serviceName := env.GetEnv("SERVICE_NAME", "Kgs")

	isEncryptionEnabled := mustBool(env.GetEnv("ENABLE_ENCRYPTION", ""))
	certFilePath := env.GetEnv("CERT_FILE_PATH", "")
	keyFilePath := env.GetEnv("KEY_FILE_PATH", "")

	gRPCAPIPort := mustInt(env.GetEnv("GRPC_API_PORT", "9000"))

	config := app.Config{
		LogLevel:    fw.LogInfo,
		ServiceName: serviceName,
		GRpcAPIPort: gRPCAPIPort,
	}

	securityPolicy := fw.SecurityPolicy{
		IsEncrypted:         isEncryptionEnabled,
		CertificateFilePath: certFilePath,
		KeyFilePath:         keyFilePath,
	}

	rootCmd := cmd.NewRootCmd(
		config,
		securityPolicy,
	)
	cmd.Execute(rootCmd)
}

func mustInt(numStr string) int {
	num, err := strconv.Atoi(numStr)
	if err != nil {
		panic(err)
	}
	return num
}

func mustBool(boolStr string) bool {
	boolean, err := strconv.ParseBool(boolStr)
	if err != nil {
		panic(err)
	}
	return boolean
}
