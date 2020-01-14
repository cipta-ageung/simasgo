# simasgo
simas repository go microservices

## if using etcd
`go get -v go.etcd.io/etcd`

`go get -v go.etcd.io/etcd/etcdctl`

### run plugin :

`etcd`

`micro --registry=etcd api --handler=rpc --enable_rpc`

`micro --registry=etcd web`

### clone repository

`git clone http://github.com/cipta-ageung/simasgo`

### run module vendor

`go mod init github.com/cipta-ageung/simasgo`

`go mod vendor`

`go mod tidy`

### run services :

cd `~/go/src/simas/simasgo/services/db`

run : `go run main.go`

cd `~/go/src/simas/simasgo/services/user`

run : `go run main.go`
