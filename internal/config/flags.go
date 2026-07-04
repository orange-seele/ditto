package config

import (
	"flag"
	"fmt"
	"os"
)

func LoadWithFlags() Config {
	lock := flag.Bool("lock", false, "lock settings toggles")
	flag.Bool("unlock", false, "unlock settings toggles")
	flag.Parse()

	cfg := LoadConfig()
	flag.Visit(func(f *flag.Flag) {
		switch f.Name {
		case "lock":
			cfg.Locked = *lock
		case "unlock":
			cfg.Locked = false
		}
		if err := SaveConfig(cfg); err != nil {
			fmt.Fprintf(os.Stderr, "failed to save config: %v\n", err)
		}
	})
	return cfg
}
