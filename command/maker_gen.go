package command

import (
	"bytes"
	"io"
	"log/slog"
	"os"

	"github.com/dondrozd/maker-gen/model"
	"github.com/dondrozd/maker-gen/parser"
	"github.com/dondrozd/maker-gen/processor"
	"github.com/dondrozd/maker-gen/renderer"
	"golang.org/x/tools/imports"
)

func GenerateMaker(cmd model.GenerateParams) error {
	slog.Info("source file: " + cmd.SrcFile)
	slog.Info("destination file: " + cmd.DestFile)
	slog.Info("generate maker: " + cmd.StructName)

	var err error
	// parse the input file
	fileModel, err := parser.MakerParse(cmd.SrcFile)
	if err != nil {
		return err
	}
	// process the details from the input go file in to what we need to render the maker
	makerModel, err := processor.PublicProc(fileModel, cmd)
	if err != nil {
		return err
	}
	// choose how to output the generated maker
	// since no destination file was set we will just dump the file to std out
	// create file
	// render the new maker file
	err = generateMakerFile(cmd, makerModel)
	if err != nil {
		return err
	}
	//

	return nil
}

func generateMakerFile(cmd model.GenerateParams, makerModel model.MakerModel) error {
	var err error
	// render file
	buf := new(bytes.Buffer)
	err = renderer.RenderMaker(makerModel, buf)
	if err != nil {
		return err
	}
	// this tool will remove unnecessary imports and format the code
	formattedBytes, err := imports.Process("", buf.Bytes(), nil)
	if err != nil {
		return err
	}
	// decide on the desired output
	var out io.WriteCloser
	if len(cmd.DestFile) == 0 {
		out = os.Stdout
	} else {
		out, err = os.Create(cmd.DestFile)
		if err != nil {
			return err
		}
	}
	defer out.Close()
	// write bytes to desired output
	_, err = out.Write(formattedBytes)

	return err
}
