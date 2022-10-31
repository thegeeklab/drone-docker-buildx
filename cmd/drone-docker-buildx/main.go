package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/thegeeklab/drone-docker-buildx/plugin"
	"github.com/urfave/cli/v2"

	"github.com/thegeeklab/drone-plugin-lib/v2/drone"
	"github.com/thegeeklab/drone-plugin-lib/v2/urfave"
)

var (
	BuildVersion = "devel"
	BuildDate    = "00000000"
)

func main() {
	settings := &plugin.Settings{}

	if _, err := os.Stat("/run/drone/env"); err == nil {
		_ = godotenv.Overload("/run/drone/env")
	}

	cli.VersionPrinter = func(c *cli.Context) {
		fmt.Printf("%s version=%s date=%s\n", c.App.Name, c.App.Version, BuildDate)
	}

	app := &cli.App{
		Name:    "drone-docker-buildx",
		Usage:   "build docker container with DinD and buildx",
		Version: BuildVersion,
		Flags:   append(settingsFlags(settings, urfave.FlagsPluginCategory), urfave.Flags()...),
		Action:  run(settings),
	}

	if err := app.Run(os.Args); err != nil {
		logrus.Fatal(err)
	}
}

func run(settings *plugin.Settings) cli.ActionFunc {
	return func(ctx *cli.Context) error {
		urfave.LoggingFromContext(ctx)

		settings.Build.CacheFrom = ctx.Generic("cache-from").(*drone.StringSliceFlag).Get()

		plugin := plugin.New(
			*settings,
			urfave.PipelineFromContext(ctx),
			urfave.NetworkFromContext(ctx),
		)

		if err := plugin.Validate(); err != nil {
			return fmt.Errorf("validation failed: %w", err)
		}

		if err := plugin.Execute(); err != nil {
			return fmt.Errorf("execution failed: %w", err)
		}

		return nil
	}
}
