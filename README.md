# lisa

[![Build Status](https://travis-ci.org/miclle/lisa.svg?branch=master)](https://travis-ci.org/miclle/lisa)

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
```
--port, -p "8080"    Serving Static Files with HTTP used port.  
--dir,  -d "./"      Serving Static Files with HTTP in directory.  
--bind, -b "0.0.0.0" Serving Static Files with HTTP bind address.  
```

run `lisa s -h` get more info