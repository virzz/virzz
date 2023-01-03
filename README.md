# Virzz

![Build](https://github.com/virzz/virzz/workflows/Build/badge.svg)

## Projects

- githack
- gopher
- hashpow
- jwttool
- parser

## githack

```
A Git source leak exploit tool that restores the entire Git repository, including data from stash, for white-box auditing and analysis of developers' mind

Usage:
  githack [flags]
  githack [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  version     Print the version

Flags:
  -d, --delay int   Request delay (N times one second)
  -h, --help        help for githack
  -l, --limit int   Request limit (N times one second) (default 10)

Use "githack [command] --help" for more information about a command.
```

## gopher

```
Generate Gopher Exp

Usage:
  gopher [flags]
  gopher [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  fcgi        Gopher Exp FastCGI
  help        Help about any command
  listen      Gopher Exp By Listen
  post        Gopher Exp HTTP POST
  redis       Gopher Exp Redis
  upload      Gopher Exp HTTP Upload
  version     Print the version

Flags:
  -f, --filename string   Filename
  -h, --help              help for gopher
  -e, --urlencode count   URL Encode (-e , -ee -eee)

Use "gopher [command] --help" for more information about a command.
```

## hashpow

```
A tool for ctfer which make hash collision faster

Usage:
  hashpow [flags]
  hashpow [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  version     Print the version

Flags:
  -c, --code string     part of hash code
  -t, --hash string     hash type : md5 sha1 (default "md5")
  -h, --help            help for hashpow
  -i, --pos int         starting position of hash
  -p, --prefix string   text prefix
  -s, --suffix string   text suffix

Use "hashpow [command] --help" for more information about a command.
```

## jwttool

```
A jwt tool with Print/Crack/Modify

Usage:
  jwttool [flags]
  jwttool [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  jwtc        JWT Crack
  jwtm        JWT Modify
  jwtp        JWT Print
  version     Print the version

Flags:
  -h, --help   help for jwttool

Use "jwttool [command] --help" for more information about a command.
```

## parser

```
Parse some file

Usage:
  parser [flags]
  parser [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  tcp         Parse /proc/net/tcp
  version     Print the version

Flags:
  -h, --help   help for parser

Use "parser [command] --help" for more information about a command.
```

