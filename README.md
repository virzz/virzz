# Virzz

[![Build](https://github.com/virzz/virzz/actions/workflows/virzz.yml/badge.svg)](https://github.com/virzz/virzz/actions/workflows/virzz.yml) [![Build Release](https://github.com/virzz/virzz/actions/workflows/virzz_release.yml/badge.svg)](https://github.com/virzz/virzz/actions/workflows/virzz_release.yml)

## Install

`brew install virzz/virzz/<formula>` || `brew tap virzz/virzz; brew install <formula>`

### Formulae

- God `brew install virzz/virzz/god` || `brew tap virzz/virzz; brew install god`

## God - CLI 命令行小工具

```
NAME:
   god - The Cyber Swiss Army Knife for terminal

USAGE:
   god [global options] command [command options] [arguments...]

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
   Misc:
     domain      Some tools for Domain/SubDomain
     netool      Net utils for IP/Port
     hashpow     Brute Hash Power of Work with md5/sha1
     qrcode, qr  A qrcode tool for terminal
     parser      Parse some file
   Web:
     githack       A `.git` folder disclosure exploit
     gopher        Generate Gopher Exp
     jwttool, jwt  A jwt tool with Print/Crack/Modify

GLOBAL OPTIONS:
   --help, -h     show help (default: false)
   --version, -v  print the version (default: false)
```

## Public Projects

- githack
- gopher
- hashpow
- jwttool
- parser

## githack

```
NAME:
   githack - A `.git` folder disclosure exploit

USAGE:
   githack [global options] [arguments...]

AUTHOR:
   陌竹(@mozhu1024) <mozhu233@outlook.com>

GLOBAL OPTIONS:
   --help, -h  show help (default: false)
```

## gopher

```
NAME:
   gopher - A jwt tool with Print/Crack/Modify

USAGE:
   gopher [global options] command [command options] [arguments...]

AUTHOR:
   陌竹(@mozhu1024) <mozhu233@outlook.com>

COMMANDS:
   JWT:
     print, p   Print jwt pretty
     modify, m  Modify jwt
     crack, c   Crack jwt
     create, n  Create jwt

GLOBAL OPTIONS:
   --help, -h  show help (default: false)
```

## hashpow

```
NAME:
   jwttool - Brute Hash Power of Work with md5/sha1

USAGE:
   jwttool [global options] [arguments...]

AUTHOR:
   陌竹(@mozhu1024) <mozhu233@outlook.com>

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
     print, p   Print jwt pretty
     modify, m  Modify jwt
     crack, c   Crack jwt
     create, n  Create jwt

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

