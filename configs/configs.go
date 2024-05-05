package configs

// struct to hold cli config provided
type CliConfig struct {
	LogLevel                        string
	TerminationGracePeriodInSeconds int
}

func LoadArguments() CliConfig {
	var config CliConfig
	return config
}
