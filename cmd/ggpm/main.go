package main

import (
	"fmt"
	"log/slog"
	"os"
	"os/exec"
	"time"

	"github.com/briandowns/spinner"
	"github.com/spf13/viper"
)

type Package struct {
	Module string  `mapstructure:"module"`
	Tags   *string `mapstructure:"tags"`
}

func main() {
	viper.SetConfigName(".ggpm")

	if home, err := os.UserHomeDir(); err != nil {
		viper.AddConfigPath(home)
	}
	viper.AddConfigPath(".")
	viper.AddConfigPath("$HOME")

	if err := viper.ReadInConfig(); err != nil {
		slog.Error("failed to read config", "err", err)
		os.Exit(1)
	}

	var packages []Package

	if err := viper.UnmarshalKey("packages", &packages); err != nil {
		slog.Error("failed to unmarshal packages", "err", err)
		os.Exit(1)
	}

	s := spinner.New(spinner.CharSets[14], 100*time.Millisecond)
	for _, pkg := range packages {
		args := []string{"install"}

		if pkg.Tags != nil {
			args = append(args, "-tags", *pkg.Tags)
		}

		args = append(args, pkg.Module)

		s.Suffix = fmt.Sprintf(" Installing: %s", pkg.Module)
		s.Start()

		cmd := exec.Command("go", args...)
		output, err := cmd.CombinedOutput()
		// Stop the spinner before printing the final status
		s.Stop()
		if err != nil {
			slog.Error("installation failed",
				"package", pkg.Module,
				"err", err,
				"output", string(output),
			)
		} else {
			args := []any{"package", pkg.Module}
			if pkg.Tags != nil {
				args = append(args, "tags", *pkg.Tags)
			}
			slog.Info("successfully installed", args...)
		}
	}
}
