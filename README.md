# OctetStream

Generate swagger 

```bash
$ make openapi-server-gen
```

Build Service

```bash
$ make build
```

Run Service
```bash
$ ./bin/octetstream
```

Consume service with CURL
```bash
$ curl --output test  http://127.0.0.1:3000/file
```