package main

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/utilsgo/husky/internal/version"
	"github.com/utilsgo/husky/pkg/husky"
	"github.com/utilsgo/husky/pkg/log"
)

var (
	logger      = log.Logger
	loggerV     = 0
	projectRoot = husky.ResolveGitRoot()
	theHusky    *husky.Husky
)

var CmdRoot = &cobra.Command{
	Use:     "husky",
	Short:   "husky " + version.Version,
	Version: version.Version,
}

func init() {
	CmdRoot.PersistentFlags().IntVarP(&loggerV, "verbose", "v", 0, "log level")

	cobra.OnInitialize(func() {
		log.SetVerbosity(loggerV)
		theHusky = husky.HuskyFrom(log.WithLogger(logger)(context.Background()), projectRoot)
	})
}
