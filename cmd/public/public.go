package public

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v3"
	"github.com/virzz/logger"
	"github.com/virzz/virzz/common"
	"github.com/virzz/virzz/utils/tongji"
)

func RunCliApp(cmd *cli.Command, name, version string) {

	app := &cli.App{
		Authors:         []any{fmt.Sprintf("%s <%s>", common.Author, common.Email)},
		Name:            name,
		Usage:           cmd.Usage,
		Version:         fmt.Sprintf("revision: %s", version),
		HideVersion:     true,
		HideHelpCommand: true,
		Suggest:         true,
		Commands:        cmd.Commands,
	}

	// Add Tongji
	go tongji.Tongji("http://god.tool.virzz.com", name, app.Version)

	if err := app.Run(os.Args); err != nil {
		logger.Error(err)
		os.Exit(1)
	}
}
