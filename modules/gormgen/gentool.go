// Fork from: https://github.com/go-gorm/gen/blob/master/tools/gentool/gentool.go
package gormgen

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/driver/sqlserver"
	"gorm.io/gen"
	"gorm.io/gorm"
)

type DBType string

const (
	dbMySQL     DBType = "mysql"
	dbPostgres  DBType = "postgres"
	dbSQLite    DBType = "sqlite"
	dbSQLServer DBType = "sqlserver"
)

// ConfigParams is command line parameters
type ConfigParams struct {
	DSN               string           `yaml:"dsn"`               // consult[https://gorm.io/docs/connecting_to_the_database.html]"
	DB                string           `yaml:"db"`                // input mysql or postgres or sqlite or sqlserver. consult[https://gorm.io/docs/connecting_to_the_database.html]
	Tables            []string         `yaml:"tables"`            // enter the required data table or leave it blank
	OnlyModel         bool             `yaml:"onlyModel"`         // only generate model
	OutPath           string           `yaml:"outPath"`           // specify a directory for output
	OutFile           string           `yaml:"outFile"`           // query code file name, default: gen.go
	WithUnitTest      bool             `yaml:"withUnitTest"`      // generate unit test for query code
	ModelPkgName      string           `yaml:"modelPkgName"`      // generated model code's package name
	FieldNullable     bool             `yaml:"fieldNullable"`     // generate with pointer when field is nullable
	FieldWithIndexTag bool             `yaml:"fieldWithIndexTag"` // generate field with gorm index tag
	FieldWithTypeTag  bool             `yaml:"fieldWithTypeTag"`  // generate field with gorm column type tag
	FieldSignable     bool             `yaml:"fieldSignable"`     // detect integer field's unsigned type, adjust generated data type
	Mode              gen.GenerateMode `yaml:"-"`                 // generate mode
}

// connectDB choose db type for connection to database
func connectDB(t DBType, dsn string) (*gorm.DB, error) {
	if dsn == "" {
		return nil, fmt.Errorf("dsn cannot be empty")
	}
	switch t {
	case dbMySQL:
		return gorm.Open(mysql.Open(dsn))
	case dbPostgres:
		return gorm.Open(postgres.Open(dsn))
	case dbSQLite:
		return gorm.Open(sqlite.Open(dsn))
	case dbSQLServer:
		return gorm.Open(sqlserver.Open(dsn))
	default:
		return nil, fmt.Errorf("unknow db %q (support mysql || postgres || sqlite || sqlserver for now)", t)
	}
}

// genModels is gorm/gen generated models
func genModels(g *gen.Generator, db *gorm.DB, tables []string) (models []interface{}, err error) {
	var tablesList []string
	if len(tables) == 0 {
		// Execute tasks for all tables in the database
		tablesList, err = db.Migrator().GetTables()
		if err != nil {
			return nil, fmt.Errorf("GORM migrator get all tables fail: %w", err)
		}
	} else {
		tablesList = tables
	}

	// Execute some data table tasks
	models = make([]interface{}, len(tablesList))
	for i, tableName := range tablesList {
		models[i] = g.GenerateModel(tableName)
	}
	return models, nil
}

// loadConfigFile load config file from path
func loadConfigFile(path string) (*ConfigParams, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close() // nolint
	var yamlConfig ConfigParams
	if cmdErr := yaml.NewDecoder(file).Decode(&yamlConfig); cmdErr != nil {
		return nil, cmdErr
	}
	return &yamlConfig, nil
}

func Generate(config *ConfigParams) error {
	db, err := connectDB(DBType(config.DB), config.DSN)
	if err != nil {
		return err
	}
	g := gen.NewGenerator(gen.Config{
		OutPath:           config.OutPath,
		OutFile:           config.OutFile,
		ModelPkgPath:      config.ModelPkgName,
		WithUnitTest:      config.WithUnitTest,
		FieldNullable:     config.FieldNullable,
		FieldWithIndexTag: config.FieldWithIndexTag,
		FieldWithTypeTag:  config.FieldWithTypeTag,
		FieldSignable:     config.FieldSignable,
		Mode:              config.Mode,
	})

	g.UseDB(db)

	models, err := genModels(g, db, config.Tables)
	if err != nil {
		return err
	}

	if !config.OnlyModel {
		g.ApplyBasic(models...)
	}

	g.Execute()
	return nil
}
