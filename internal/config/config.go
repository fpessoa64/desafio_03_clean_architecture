package config

import "os"

type Config struct {
	DBDSN       string
	RESTPort    string
	GRPCPort    string
	GraphQLPort string
}

func Load() *Config {
	restPort := os.Getenv("REST_PORT")
	if restPort == "" {
		restPort = "8080"
	}
	grpcPort := os.Getenv("GRPC_PORT")
	if grpcPort == "" {
		grpcPort = "50051"
	}
	graphqlPort := os.Getenv("GRAPHQL_PORT")
	if graphqlPort == "" {
		graphqlPort = "8081"
	}
	return &Config{
		DBDSN:       os.Getenv("DB_DSN"),
		RESTPort:    restPort,
		GRPCPort:    grpcPort,
		GraphQLPort: graphqlPort,
	}
}
