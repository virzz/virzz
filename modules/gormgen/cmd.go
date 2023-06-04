package gormgen

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/urfave/cli/v3"
	"gorm.io/gen"
)

//go:embed gen.yaml
var genTmpl string

var Cmd = &cli.Command{
	Name:    "gormgen",
	Aliases: []string{"gorm"},
	Usage:   "Gen Tool For Gorm",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "config",
			Aliases: []string{"c"},
			Usage:   "is path for gen.yml",
		},
		&cli.StringFlag{
			Name:  "dsn",
			Usage: "consult[https://gorm.io/docs/connecting_to_the_database.html]",
		},
		&cli.StringFlag{
			Name:  "db",
			Usage: "input mysql|postgres|sqlite|sqlserver.",
			Value: "mysql",
		},
		&cli.StringFlag{
			Name:  "tables",
			Usage: "input tables name, split by ','",
		},
		&cli.BoolFlag{
			Name:  "onlyModel",
			Usage: "only generate models (without query file)",
		},
		&cli.StringFlag{
			Name:  "outPath",
			Usage: "specify a directory for output",
			Value: "./query",
		},
		&cli.StringFlag{
			Name:  "outFile",
			Usage: "query code file name",
			Value: "gen.go",
		},
		&cli.StringFlag{
			Name:  "modelPkgName",
			Usage: "generated model code's package name",
			Value: "model",
		},
		&cli.BoolFlag{
			Name:  "withUnitTest",
			Usage: "generate unit test for query code",
		},
		&cli.BoolFlag{
			Name:  "fieldNullable",
			Usage: "generate with pointer when field is nullable",
		},
		&cli.BoolFlag{
			Name:  "fieldWithIndexTag",
			Usage: "generate field with gorm index tag",
			Value: true,
		},
		&cli.BoolFlag{
			Name:  "fieldWithTypeTag",
			Usage: "generate field with gorm column type tag",
			Value: true,
		},
		&cli.BoolFlag{
			Name:  "fieldSignable",
			Usage: "detect integer field's unsigned type, adjust generated data type",
		},
		&cli.BoolFlag{
			Name:  "WithDefaultQuery",
			Usage: "create default query in generated code",
		},
		&cli.BoolFlag{
			Name:  "WithoutContext",
			Usage: "generate code without context constrain",
		},
		&cli.BoolFlag{
			Name:  "WithQueryInterface",
			Usage: "generate code with exported interface object",
		},
		&cli.BoolFlag{
			Name:  "template",
			Usage: "generate config template",
		},
	},
	Action: func(c *cli.Context) (err error) {
		if c.Bool("template") {
			_, err = fmt.Println(genTmpl)
			return
		}
		genPath := c.String("config")
		var conf *ConfigParams
		if genPath != "" {
			conf, err = loadConfigFile(genPath)
			if err != nil {
				return err
			}
		} else {
			conf.DB = c.String("db")
			conf.DSN = c.String("dsn")
			conf.OnlyModel = c.Bool("onlyModel")
			conf.OutPath = c.String("outPath")
			conf.OutFile = c.String("outFile")
			conf.ModelPkgName = c.String("modelPkgName")
			conf.WithUnitTest = c.Bool("withUnitTest")
			conf.FieldNullable = c.Bool("fieldNullable")
			conf.FieldWithIndexTag = c.Bool("fieldWithIndexTag")
			conf.FieldWithTypeTag = c.Bool("fieldWithTypeTag")
			conf.FieldSignable = c.Bool("fieldSignable")
			conf.Tables = strings.Split(c.String("tables"), ",")
		}
		if c.Bool("withDefaultQuery") {
			conf.Mode |= gen.WithDefaultQuery
		}
		if c.Bool("withoutContext") {
			conf.Mode |= gen.WithoutContext
		}
		if c.Bool("withQueryInterface") {
			conf.Mode |= gen.WithQueryInterface
		}
		return Generate(conf)
	},
}
