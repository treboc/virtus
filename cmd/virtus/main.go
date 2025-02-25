package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/treboc/virtus/internal/config"
)

var (
	cfg        *config.Config
	configPath string
)

func main() {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	var rootCmd = &cobra.Command{
		Use:   "virtus",
		Short: "Virtus - Forge your path to excellence",
		Long: `Virtus helps you build and maintain virtuous habits through an elegant CLI interface.
	Track your daily practices, weekly routines, and watch your excellence grow over time.`,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			loadedCfg, err := config.Load(configPath)
			if err != nil {
				return err
			}
			cfg = loadedCfg

			return cfg.EnsurePaths()
		},
	}

	defaultConfigPath := filepath.Join(home, ".virtus", "config.yaml")

	rootCmd.PersistentFlags().StringVarP(&configPath, "config", "c", defaultConfigPath, "Path to configuration file")

	rootCmd.AddCommand(
		// newHabitCmd(),
		// newListCmd(),
		// newCompleteCmd(),
		// newMoodCmd(),
		newConfigCmd(),
	)

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

func newHabitAddCmd() *cobra.Command {
	var title, note, trackingType, trackingConfig string

	cmd := &cobra.Command{
		Use:   "add",
		Short: "Add a new habit to track",
		RunE: func(cmd *cobra.Command, args []string) error {
			// TODO: Implement habit creation
			return nil
		},
	}

	cmd.Flags().StringVarP(&title, "title", "t", "", "Habit title")
	cmd.Flags().StringVarP(&note, "note", "n", "", "Habit note")
	cmd.Flags().StringVarP(&trackingType, "type", "y", "daily", "Tracking type (daily, weekly, monthly, interval)")
	cmd.Flags().StringVarP(&trackingConfig, "config", "c", "", "Tracking configuration")

	return cmd
}

func newConfigCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "config",
		Short: "Manage configuration",
		RunE: func(cmd *cobra.Command, args []string) error {
			return cfg.Save(configPath)
		},
	}
}
