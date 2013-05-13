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

