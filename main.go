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

	dbConfig := fw.DBConfig{
		Host:     env.GetEnv("DB_HOST", "localhost"),
		Port:     mustInt(env.GetEnv("DB_PORT", "5432")),
		User:     env.GetEnv("DB_USER", "postgres"),
		Password: env.GetEnv("DB_PASSWORD", "password"),
		DbName:   env.GetEnv("DB_NAME", "account"),
	}

	dbConnector := dep.InitDBConnector()

	serviceName := env.GetEnv("SERVICE_NAME", "AccountService")

	isEncryptionEnabled := mustBool(env.GetEnv("ENABLE_ENCRYPTION", "false"))
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
		dbConfig,
		dbConnector,
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
