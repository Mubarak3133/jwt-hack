# jwt-hack

## Usage
```
   d8p 8d8   d88 888888888          888  888 ,8b.     doooooo 888  ,dP
   88p 888,o.d88    '88d     ______ 88888888 88'8o    d88     888o8P'
   88P 888P`Y8b8   '888      XXXXXX 88P  888 88PPY8.  d88     888 Y8L
88888' 88P   YP8 '88p               88P  888 8b   `Y' d888888 888  `8p
-------------------------
Hack the JWT(JSON Web Token) | by @hahwul | v0.0.1

Usage:
  jwt-hack [command]

Available Commands:
  crack       Cracking JWT Token
  decode      Decode JWT to JSON
  encode      Encode json to JWT
  help        Help about any command
  payload     Generate JWT Attack payloads

Flags:
  -h, --help   help for jwt-hack
```

## Encode mode(JSON to JWT)
```
$ jwt-hack encode '{"test":"1234"}' --secret=asdf
```
## Decode mode(JWT to JSON)
```
$ jwt-hack decode eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c
```
## Crack mode(Dictionary attack / BruteForce)
```
$ jwt-hack crack -w 1.lst eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.cq-uoLxOu3V4RjxnbUAFZ36aSZ24BXiAH8RFDYVA6XU
```

## Payload mode(Alg none attack, etc..)
```
$ jwt-hack payload eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c
```
