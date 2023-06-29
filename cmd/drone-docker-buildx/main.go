package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/thegeeklab/drone-docker-buildx/plugin"
	"github.com/urfave/cli/v2"

	"github.com/thegeeklab/drone-plugin-lib/v2/drone"
	"github.com/thegeeklab/drone-plugin-lib/v2/urfave"
)

//nolint:gochecknoglobals
var (
	BuildVersion = "devel"
	BuildDate    = "00000000"
)

var ErrTypeAssertionFailed = errors.New("type assertion failed")

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

		cacheFrom, ok := ctx.Generic("cache-from").(*drone.StringSliceFlag)
		if !ok {
			return fmt.Errorf("%w: failed to read cache-from input", ErrTypeAssertionFailed)
		}

		settings.Build.CacheFrom = cacheFrom.Get()

		secrets, ok := ctx.Generic("secrets").(*drone.StringSliceFlag)
		if !ok {
			return fmt.Errorf("%w: failed to read secrets input", ErrTypeAssertionFailed)
		}

		settings.Build.Secrets = secrets.Get()

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
