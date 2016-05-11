# lisa

[![Build Status](https://travis-ci.org/miclle/lisa.svg?branch=master)](https://travis-ci.org/miclle/lisa)
[![Go Report Card](https://goreportcard.com/badge/github.com/miclle/lisa)](https://goreportcard.com/report/github.com/miclle/lisa)

Lisa is a development tool for the Go Programming Language. It automatically runs your command when it detects file system changes.
http://miclle.me/lisa/

## Installation

Assuming you have a working Go environment and `GOPATH/bin` is in your `PATH`, `lisa` is a breeze to install:

```
go get github.com/miclle/lisa
```

Then verify that `lisa` was installed correctly:

```
lisa -h
```

## Commands

### server, s

Serving Static Files with HTTP

```
lisa s
```

OPTIONS:

flag   | short | default | description
-------|-------|---------|---------------------------------------------
--port | -p    | 8080    | Serving Static Files with HTTP used port.
--dir  | -d    | ./      | Serving Static Files with HTTP in directory.
--bind | -b    | 0.0.0.0 | Serving Static Files with HTTP bind address.

run `lisa s -h` get more info

### watch, w

Starting a file system watcher then execute a command

```
lisa w
```

OPTIONS:

flag      | short | default | description
----------|-------|---------|---------------------------------------------
--command | -c    | 			  | Execute the command when the directory files modified.
--path    | -p    | ./      | Watching the directory or file.
--event   | -e    | create,rename,write,remove | Execute the command when the events was trigger: <br /> create,rename,write,remove,chmod
--delay   | -d    | 1000    | Execute the command after the number of milliseconds.

run `lisa w -h` get more info

### home

Go to the document website

```
lisa home
```