package main

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/virzz/virzz/common"
	"github.com/virzz/virzz/services/server/mariadb"
)

var mariadbCmd = &cobra.Command{
	Use:   "mariadb",
	Short: "Service Mariadb",
	RunE: func(cmd *cobra.Command, args []string) error {
		if cmd.Flags().NFlag() != 1 {
			return fmt.Errorf("one of flags(-m,-p,-s sql) is required")
		}
		r, err := ServiceMariadb()
		if err != nil {
			return err
		}
		return common.Output(string(r))
	},
}

var (
	migrate   bool
	procedure bool
	testSql   string
)

func init() {
	mariadbCmd.Flags().BoolVarP(&migrate, "migrate", "m", false, "Run auto migration for given models")
	mariadbCmd.Flags().BoolVarP(&procedure, "procedure", "p", false, "Print create procedure 'auto clean record' sql")
	mariadbCmd.Flags().StringVarP(&testSql, "sql", "s", "select 'success';", "Exec sql(one column) for test connect")
}

// ServiceMariadb -
func ServiceMariadb() (string, error) {

	err := mariadb.Connect(debugDatabase)
	if err != nil {
		return "", err
	}

	if migrate {
		// models.InitMariadb()
		err = mariadb.Migrate()
		if err != nil {
			return "", err
		}
	} else if procedure {
		return mariadb.Procedure(), nil
	} else if testSql != "" {
		res, err := mariadb.ExecSQL(testSql)
		if err != nil {
			return "", err
		}
		return res, nil
	}

	return "", nil
}
