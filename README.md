# Text Validator

Validate the text whether legal or not

## Usage

### Startup

```bash
$ ./validator --port 19999 --dict=https://example.com/dict.txt

# or

$ ./validator --port 19999 --dict=./dict
```

### Request

```bash
$ curl http://127.0.0.1:8081/validate/text?query=legal
{"Code":0,"Message":"ok"} # http status code: 200


$ curl http://127.0.0.1:8081/validate/text?query=illegal
{"Code":-1,"Message":"illegal"} # http status code: 400
```



## Help

```
Usage:
  validator [flags]

Flags:
      --dict string   词典 (default "./dict.txt")
  -h, --help          help for validator
      --port int      服务器端口 (default 8081)
```

## Dict Rules

```
word1
word2
word3
```
