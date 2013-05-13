## Get a web application up and running with go

###Steps

- For this project, you only need contents of this directory
- set `$GOPATH` to value applicable to your environment
- `go get github.com/hoise/webgo`
- set appropriate vhost entry in nginx (see `sample-nginx.conf`)
- `/etc/init.d/nginx restart`
- `go run samplewebapp.go`
- `curl -XGET http://{{VHOST}}/world`

#####Output:  `hello world`

---

Curiosity got the best of me so I ran a simple benchmark.
This was run locally (without nginx / fcgi).  I ran this
on a "large" instance on AWS, but keep in mind, `ab` was
running locally with the application.  Memory of the process
did not budge, and 1x CPU was being utilized almost 100%.

<code>
Finished 100000 requests


Server Software:        web.go
Server Hostname:        localhost
Server Port:            9999

Document Path:          /test
Document Length:        10 bytes

Concurrency Level:      100
Time taken for tests:   17.366 seconds
Complete requests:      100000
Failed requests:        0
Write errors:           0
Total transferred:      16100000 bytes
HTML transferred:       1000000 bytes
Requests per second:    5758.53 [#/sec] (mean)
Time per request:       17.366 [ms] (mean)
Time per request:       0.174 [ms] (mean, across all concurrent requests)
Transfer rate:          905.39 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.8      0      14
Processing:     2   17   3.1     16      38
Waiting:        1   17   3.0     16      38
Total:          7   17   2.9     16      38

Percentage of the requests served within a certain time (ms)
  50%     16
  66%     18
  75%     20
  80%     20
  90%     21
  95%     22
  98%     23
  99%     26
 100%     38 (longest request)

</code>
