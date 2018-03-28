# OctetStream

## Usage

### Install dependencies

```bash
go get -u github.com/go-swagger/go-swagger/cmd/swagger
```

### Generate swagger 

```bash
$ make openapi-server-gen
```

### Build Service

```bash
$ make build
```

### Run Service
```bash
$ ./bin/octetstream
```

### Consume service with CURL
```bash
$ curl --output test  http://127.0.0.1:3000/file
```

---

## Profiling

## Install Apache Benchmark

```bash
$ sudo apt install apache2-utils
```

### Run Service
```bash
$ ./bin/octetstream
```

### Run Benchmark
- -k Enables HTTP keep-alive
- -c Number of concurrent requests
- -n Number of total requests to make

```bash
$ ab -k -c 8 -n 1000 http://127.0.0.1:3000/file?size=1000002
```


### Run pprof Memory
```bash
$ go tool pprof ./bin/octetstream http://localhost:8080/debug/pprof/heap
```

### Run pprof CPU
```bash
$ go tool pprof ./bin/octetstream http://127.0.0.1:8080/debug/pprof/profile
```

source: https://artem.krylysov.com/blog/2017/03/13/profiling-and-optimizing-go-web-applications/