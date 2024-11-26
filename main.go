package main

import (
	"log"
	"log/slog"
	"os"

	"github.com/dondrozd/maker-gen/command"
	"github.com/dondrozd/maker-gen/model"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Usage: "This tool is used to generate templates.  This is intended to be used for unit tests to build data.  Instead of each unit test building its own data you can build templates that are reusable.",
		Commands: []*cli.Command{
			{
				Name:    "generate",
				Aliases: []string{"gen", "g"},
				Usage:   "generate",
				Action:  genMaker,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "src",
						Aliases:  []string{"s"},
						Usage:    "file name that contains the struct",
						Required: true,
					},
					&cli.StringFlag{
						Name:    "dst",
						Aliases: []string{"d"},
						Usage:   "file name that to generate",
					},
					&cli.StringFlag{
						Name:    "withPrefix",
						Aliases: []string{"wp"},
						Usage:   "prefix for each with function.  Example: Each With{withPrefix}PopertyName(value)",
					},
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func genMaker(cCtx *cli.Context) error {
	structName := cCtx.Args().First()
	srcFile := cCtx.String("src")
	destFile := cCtx.String("dst")
	withPrefix := cCtx.String("withPrefix")
	if destFile == "" {
		slog.SetLogLoggerLevel(slog.LevelError)
	}

	return command.GenerateMaker(model.GenerateParams{
		SrcFile:    srcFile,
		DestFile:   destFile,
		WithPrefix: withPrefix,
		StructName: structName,
	})
}
