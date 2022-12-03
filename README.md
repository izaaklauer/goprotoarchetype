# goprotoarchetype

description


### Local testing

Start your server with `make serve`

```
$ make serve

go run cmd/goprotoarchetype-api/main.go config/config_local.hcl
starting goprotoarchetype.......
2022/12/02 20:58:46 Serving on "localhost:8080"

```

Then send a request via grpcurl:

```
$ grpcurl -plaintext -d '{"message": "hello from local development"}' localhost:8080 goprotoarchetype.v1.GoprotoarchetypeService/HelloWorld

{
  "configMessage": "hello from goprotoarchetype",
  "requestMessage": "hello from local development",
  "now": "2022-12-03T02:01:19.505743Z"
}
```
