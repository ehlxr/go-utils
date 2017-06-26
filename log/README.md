# log

golang log library base [logrus](https://github.com/sirupsen/logrus) and [logrus-prefixed-formatter](https://github.com/x-cray/logrus-prefixed-formatter)

# Usage

## 1. Use The glide Package Management

### install [glide](https://github.com/Masterminds/glide#install)

```bash
$ go get github.com/Masterminds/glide

$ cd $GOPATH/src/github.com/Masterminds/glide

$ make setup
```
or

```bash
# Mac OS
$ brew install glide

# Mac or Linux
$ curl https://glide.sh/get | sh
```
[Binary packages](https://github.com/Masterminds/glide/releases) are available for Mac, Linux and Windows.

### install log

```bash
$ go get -u github.com/ehlxr/go-utils

$ cd $GOPATH/src/github.com/ehlxr/go-utils/log

$ glide install
```

## 2. Manually add Dependencies

### add dependencies

```bash
$ go get github.com/sirupsen/logrus
$ go get github.com/x-cray/logrus-prefixed-formatter
```

### install log

```bash
$ go get -u github.com/ehlxr/go-utils
```


