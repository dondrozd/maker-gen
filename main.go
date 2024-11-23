package main

import (
	parsers "dondrozd/maker-gen/parser"
	"dondrozd/maker-gen/renderer"
	"fmt"
	"io"
	"log"
	"os"

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
	fmt.Println("generate maker: ", cCtx.Args().First())
	var readCloser io.WriteCloser
	var err error
	fileModel, err := parsers.MakerParse("example/example_1.go")
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

	return renderer.RenderMaker(fileModel, readCloser)
}
