`jwt` command-line tool
=======================

This is a simple tool to sign, verify and show JSON Web Tokens from the command line.

```shell
# The following will create and sign a token
$ echo {\"foo\":\"bar\"} | ./jwt -key key/jwtRS256.key -alg RS256 -sign -

# then verify it and output the original claims:
$ echo {\"foo\":\"bar\"} | ./jwt -key key/jwtRS256.key -alg RS256 -sign - | ./jwt -key key/jwtRS256.key.pub -alg RS256 -verify -

# To simply display a token, use:
$ echo {\"foo\":\"bar\"} | ./jwt -key key/jwtRS256.key -alg RS256 -sign - | ./jwt -show -
```

> generate rsa256 key

```shell
# generate key
$ ssh-keygen -t rsa -b 2048 -f jwtRS256.key

$ openssl rsa -in jwtRS256.key -pubout -outform PEM -out jwtRS256.key.pub
```