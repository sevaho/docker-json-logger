# docker-json-logger

> A simple docker container that produces JSON logs

```shell
$ docker run -p3000:3000 sevaho/docker-example-json
```

```shell
$ curl localhost:3000
```

Should produce:

```
{
    "levelname": "INFO",
    "id": "e486e9b4-c71e-444f-9740-bf59111dae8e",
    "remote_ip": "172.17.0.1",
    "host": "localhost:3000",
    "method": "GET",
    "path": "/",
    "protocol": "http",
    "status_code": 200,
    "latency": 4.0973e-05,
    "tag": "request",
    "asctime": 1677082810,
    "caller": "/go/pkg/mod/github.com/edersohe/zflogger@v0.7.0/zflogger.go:104",
    "message": "success"
}
```
