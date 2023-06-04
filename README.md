# Virzz

[![Build](https://github.com/virzz/virzz/actions/workflows/virzz.yml/badge.svg)](https://github.com/virzz/virzz/actions/workflows/virzz.yml) [![Build Release](https://github.com/virzz/virzz/actions/workflows/virzz_release.yml/badge.svg)](https://github.com/virzz/virzz/actions/workflows/virzz_release.yml)

## Install

`brew install virzz/virzz/<formula>` || `brew tap virzz/virzz; brew install <formula>`

### Formulae

- Enyo `brew install virzz/virzz/enyo` || `brew tap virzz/virzz; brew install enyo`

## Enyo - CLI 命令行小工具

```
NAME:
   enyo - The Cyber Swiss Army Knife for terminal

USAGE:
   enyo [global options] command [command options] [arguments...]

VERSION:
   latest

AUTHOR:
   陌竹(@mozhu1024) <mozhu233@outlook.com>

COMMANDS:
   Crypto:
     basex      Base 16/32/58/62/64/85/91/92/100 Encode/Decode
     basic      Some basic encodings
     classical  Classical cryptography
     hash       Hash Function
     bcrypt     Bcrypt Generate/Compare
     hashpwd    A tool for query password hash offline
   GitHub:
     gh-mozhu, ghext  A little toolkit using GitHub API
   Misc:
     domain         Some tools for Domain/SubDomain
     netool         Net utils for IP/Port
     hashpow        Brute Hash Power of Work with md5/sha1
     qrcode, qr     A qrcode tool for terminal
     parser         Parse some file
     resh, reshell  Reverse Shell Template Generator
   Web:
     githack       A `.git` folder disclosure exploit
     gopher        Generate Gopher Exp
     jwttool, jwt  A jwt tool with Print/Crack/Modify

GLOBAL OPTIONS:
   --help, -h     show help (default: false)
   --version, -v  print the version (default: false)
```

## Public Projects

- gh-mozhu
- githack
- gopher
- gostrip
- hashpow
- hashpwd
- jwttool
- parser

## gh-mozhu

```
NAME:
   gh-mozhu - A little toolkit using GitHub API

USAGE:
   gh-mozhu [global options] command [command options] [arguments...]

AUTHOR:
   陌竹(@mozhu1024) <mozhu233@outlook.com>

COMMANDS:
   install   Install this
   orgs      List organizations for the authenticated user
   transfer  Transfer a repository
   Ext:
     commit, gcmt, c  Generate Commit Message

GLOBAL OPTIONS:
   --help, -h  show help (default: false)
```

## githack

```
NAME:
   githack - A `.git` folder disclosure exploit

USAGE:
   githack [global options] [arguments...]

AUTHOR:
   陌竹(@mozhu1024) <mozhu233@outlook.com>

GLOBAL OPTIONS:
   --limit value, -l value    Request limit (default: 10)
   --delay value, -d value    Request delay (default: 0)
   --timeout value, -t value  Request timeout (default: 10)
   --help, -h                 show help (default: false)
```

## gopher

```
NAME:
   gopher - Generate Gopher Exp

USAGE:
   gopher [global options] command [command options] [arguments...]

AUTHOR:
   陌竹(@mozhu1024) <mozhu233@outlook.com>

COMMANDS:
   HTTP:
     post    HTTP Post
     upload  HTTP Upload
   Other:
     fastcgi, fcgi  FastCGI
   Redis:
     listen    By Listen redis-cli command
     write     Redis Write File
     webshell  Redis Write Webshell
     write     Redis Write Crontab
     reverse   Redis Write File

GLOBAL OPTIONS:
   --urlencode value, -e value  Urlencode count (default: 0)
   --help, -h                   show help (default: false)
```

## gostrip

```
NAME:
   gostrip - Strip golang binary file

USAGE:
   gostrip [global options] [arguments...]

AUTHOR:
   陌竹(@mozhu1024) <mozhu233@outlook.com>

GLOBAL OPTIONS:
   --force, -f  Force to strip 'Go Struct Name' (未完全测试,请谨慎使用) (default: false)
   --help, -h   show help (default: false)
```

## hashpow

```
NAME:
   hashpow - Brute Hash Power of Work with md5/sha1

USAGE:
   hashpow [global options] [arguments...]

AUTHOR:
   陌竹(@mozhu1024) <mozhu233@outlook.com>

GLOBAL OPTIONS:
   --code value, -c value    Request code
   --pos value, -i value     Starting position of hash (default: 0)
   --prefix value, -p value  Hash prefix
   --suffix value, -s value  Hash suffix
   --method value, -m value  Hash method: <sha1|md5> (default: "md5")
   --help, -h                show help (default: false)
```

## hashpwd

```
NAME:
   hashpwd - A tool for query password hash offline

USAGE:
   hashpwd [global options] command [command options] [arguments...]

AUTHOR:
   陌竹(@mozhu1024) <mozhu233@outlook.com>

COMMANDS:
   generate, g  Generate password hash form password dict
   lookup, l    Generate password hash form password dict

GLOBAL OPTIONS:
   --help, -h  show help (default: false)
```

## jwttool

```
NAME:
   jwttool - A jwt tool with Print/Crack/Modify

USAGE:
   jwttool [global options] command [command options] [arguments...]

AUTHOR:
   陌竹(@mozhu1024) <mozhu233@outlook.com>

COMMANDS:
   JWT:
     jwtp, print, p             Print jwt pretty
     jwtm, modify, m            Modify jwt
     jwtc, crack, c             Crack jwt
     jwtg, generate, create, n  Create/Generate jwt

GLOBAL OPTIONS:
   --help, -h  show help (default: false)
```

## parser

```
NAME:
   parser - Parse some file

USAGE:
   parser [global options] command [command options] [arguments...]

AUTHOR:
   陌竹(@mozhu1024) <mozhu233@outlook.com>

COMMANDS:
   procnet, net  Parse /proc/net/tcp|udp
   dsstore       .DS_Store Parser

GLOBAL OPTIONS:
   --help, -h  show help (default: false)
```

