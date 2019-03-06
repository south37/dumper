# dumper
Simple server implementation for dumping request headers

## Usage

Run commands shown below.

```
$ docker run -p 8080:8080 south37/dumper
```

Send requests with curl.

```console
$ curl -v http://localhost:8080/
*   Trying ::1...
* TCP_NODELAY set
* Connected to localhost (::1) port 8080 (#0)
> GET / HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.54.0
> Accept: */*
>
< HTTP/1.1 200 OK
< Date: Wed, 06 Mar 2019 14:53:32 GMT
< Content-Length: 14
< Content-Type: text/plain; charset=utf-8
<
* Connection #0 to host localhost left intact
Hello, dumper!%
```

Then, running server prints request headers.

```console
$ docker run -p 8080:8080 south37/dumper
=> Server starting on :8080
2019/03/06 14:53:32 GET / HTTP/1.1
2019/03/06 14:53:32 Host: localhost:8080
2019/03/06 14:53:32 User-Agent: curl/7.54.0
2019/03/06 14:53:32 Accept: */*
.
.
.
```
