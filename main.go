package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/alecthomas/kong"
)

type converterName string
type cliConverter struct {
	Converter
}

type clioption struct {
	Input  string        `arg name:"input" type:"path" help:"input file"`
	Type   converterName `arg name:"type" help:"convert type"`
	Output string        `help:"output file" short:"o"`
}

var (
	cliargument struct {
		Encode struct {
			clioption
		} `cmd help:"Encode file."`

		Decode struct {
			clioption
		} `cmd help:"Decode file"`
	}
)

func (option *clioption) readData() ([]byte, error) {
	if option.Input != "" {
		return ioutil.ReadFile(option.Input)
	}

	return ioutil.ReadAll(os.Stdin)
}

func (converter cliConverter) convert(option clioption, reverse bool) error {
	if data, err := option.readData(); err == nil {
		if data, err = converter.Convert(data, reverse); err == nil {
			return option.saveData(data)
		}
		return err
	} else {
		return err
	}
}

func (option *clioption) saveData(image []byte) error {
	output := os.Stdout
	var err error

	if option.Output != "" {
		output, err = os.Create(option.Output)
		defer output.Close()

		if err != nil {
			return err
		}
	}

	_, err = output.Write(image)
	return err
}

/*CreateConverter method is used to create the convertor based on the type provided in CLI argument. (default value is base64)*/
func (converterName converterName) CreateConverter() (cliConverter, error) {
	conv, err := Resolve(string(converterName))
	return cliConverter{Converter: conv}, err
}

func main() {
	ctx := kong.Parse(&cliargument)
	command := ctx.Command()
	var direction bool
	var option clioption
	switch command {
	case "encode <input> <type>":
		option = cliargument.Encode.clioption
		direction = true
	case "decode <input> <type>":
		option = cliargument.Decode.clioption
		direction = false
	default:
		fmt.Printf(ctx.Command())
	}
	if conv, err := option.Type.CreateConverter(); err == nil {
		err = conv.convert(option, direction)
		if err != nil {
			fmt.Printf(err.Error())
		}
	} else {
		fmt.Printf(err.Error())
	}
}
