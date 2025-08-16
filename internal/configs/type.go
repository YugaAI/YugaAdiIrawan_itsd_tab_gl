package configs

type (
	Config struct {
		Service  Service
		Database DatabaseConfig
	}

	Service struct {
		Port string
	}

	DatabaseConfig struct {
		DataSourceName string
	}
)
