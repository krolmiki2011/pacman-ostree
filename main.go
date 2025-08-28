// Plik głowny

package main

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Core struct {
		Repo         string `yaml:"repo"`
		OSName       string `yaml:"osname"`
		Branch       string `yaml:"branch"`
		DefaultImage string `yaml:"default_image"`
	} `yaml:"core"`
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: ./main.go write-config [file]")
		return
	}

	cmd := os.Args[1]

	// Wczytaj config (lub stwórz pusty)
	cfg, err := loadConfigFile()
	if err != nil {
		log.Printf("No config file found, using default values")
	}

	switch cmd {
	case "write-config":
		outFile := "/etc/pacman-ostree.conf"
		if len(os.Args) >= 3 {
			outFile = os.Args[2]
		}
		cmdWriteConfig(cfg, outFile)
	default:
		fmt.Println("Unknown command:", cmd)
	}
}

// loadConfigFile wczytuje config z /etc/pacman-ostree.conf
func loadConfigFile() (Config, error) {
	cfg := Config{}
	data, err := os.ReadFile("/etc/pacman-ostree.conf")
	if err != nil {
		return cfg, err
	}
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return cfg, err
	}
	return cfg, nil
}

// writeConfigFile zapisuje Config do YAML
func writeConfigFile(cfg Config, path string) error {
	data, err := yaml.Marshal(&cfg)
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0o644)
}

// cmdWriteConfig wrapper do wywołania writeConfigFile
func cmdWriteConfig(cfg Config, path string) {
	if err := writeConfigFile(cfg, path); err != nil {
		log.Fatalf("Failed to write config: %v", err)
	}
	log.Printf("[✓] Configuration saved to %s", path)
}
