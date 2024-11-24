package main

import (
	"fmt"
	"io"
	"log"
	"os"

	parsers "github.com/dondrozd/maker-gen/parser"
	"github.com/dondrozd/maker-gen/processor"
	"github.com/dondrozd/maker-gen/renderer"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:    "generate",
				Aliases: []string{"gen", "g"},
				Usage:   "generate",
				Action:  genMaker,
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func genMaker(cCtx *cli.Context) error {
	structName := cCtx.Args().First()

	fmt.Println("generate maker: ", structName)
	var readCloser io.WriteCloser
	var err error
	fileModel, err := parsers.MakerParse("example/example_1.go")
	if err != nil {
		return err
	}
	makerModel, err := processor.PublicProc(fileModel, structName)
	if err != nil {
		return err
	}
	fmt.Println("file model:\n", fileModel)
	if true {
		readCloser = os.Stdout
	} else {
		readCloser, err = os.Create("resources/scatch/myfile_gen.go")
		if err != nil {
			return err
		}
	}
	defer readCloser.Close()

	return renderer.RenderMaker(makerModel, readCloser)
}
