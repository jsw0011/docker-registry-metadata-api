package configuration

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type ConfDB struct {
	User     string `yaml:"User"`
	Pass     string `yaml:"Pass"`
	Host     string `yaml:"Host"`
	Port     uint   `yaml:"Port"`
	SSL      bool   `yaml:"SSL"`
	Database string `yaml:"Database"`
}

type ConfRegistry struct {
	User string `yaml:"User"`
	Pass string `yaml:"Pass"`
	Host string `yaml:"Host"`
	Port uint   `yaml:"Port"`
	SSL  bool   `yaml:"SSL"`
}

type ConfGeneral struct {
	LogLevel     string `yaml:"LogLevel"`
	LogFile      string `yaml:"LogFile"`
	LogToConsole bool   `yaml:"LogToConsole"`
	Port         uint   `yaml:"Port"`
}

type AppConfiguration struct {
	Database ConfDB       `yaml:"Database"`
	Registry ConfRegistry `yaml:"Registry"`
	General  ConfGeneral  `yaml:"General"`
}

func ParseConfiguration(filePath *string, conf *AppConfiguration) {

	yfile, err := os.ReadFile(*filePath)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to read configuration file: %v", err)
		os.Exit(3)
	}

	err = yaml.Unmarshal(yfile, &conf)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to parse configuration file: %v", err)
		os.Exit(3)
	}

	// verify not empty values and set default values if empty
	if conf.Database.User == "" {
		fmt.Fprintf(os.Stdout, "[Configuration] Warning: Database.User is empty!\n")
	}
	if conf.Database.Pass == "" {
		fmt.Fprintf(os.Stdout, "[Configuration] Warning: Database.Pass is empty!\n")
	}
	if conf.Database.Host == "" {
		fmt.Fprintf(os.Stderr, "[Configuration] Error: Missing Database.Host!\n")
		os.Exit(3)
	}
	if conf.Database.Port == 0 {
		fmt.Fprintf(os.Stdout, "[Configuration] Error: Missing Database.Port!\n")
		os.Exit(3)
	}
	if conf.Database.Database == "" {
		fmt.Fprintf(os.Stdout, "[Configuration] Error: Missing Database.Database (Database name)!\n")
		os.Exit(3)
	}

	if conf.Registry.User == "" {
		fmt.Fprintf(os.Stdout, "[Configuration] Warning: Registry.User is empty!\n")
	}
	if conf.Registry.Pass == "" {
		fmt.Fprintf(os.Stdout, "[Configuration] Warning: Registry.Pass is empty!\n")
	}
	if conf.Registry.Host == "" {
		fmt.Fprintf(os.Stderr, "[Configuration] Error: Missing Registry.Host!\n")
		os.Exit(3)
	}
	if conf.Registry.Port == 0 {
		fmt.Fprintf(os.Stdout, "[Configuration] Error: Missing Registry.Port!\n")
		os.Exit(3)
	}

	if conf.General.LogLevel == "" {
		fmt.Fprintf(os.Stdout, "[Configuration] Warning: General.LogLevel is empty! (DEBUG, NOTICE, INFO, WARNING, ERROR, CRITICAL, )\n")
	}
	if conf.General.LogFile == "" && !conf.General.LogToConsole {
		fmt.Fprintf(os.Stdout, "[Configuration] Warning: Using default log file \"api.log\" (General.LogFile)\n")
	}
	if conf.General.Port == 0 {
		fmt.Fprintf(os.Stdout, "[Configuration] Error: Missing General.Port!\n")
		os.Exit(3)
	}
}
