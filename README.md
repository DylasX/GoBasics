# GO DOCUMENTATION

1. https://golang.org/pkg/
2. https://pkg.go.dev/


#Si los modulos no los encuentra
go env -w GO111MODULE=off

GO111MODULE=on will force using Go modules even if the project is in your GOPATH. Requires go.mod to work.
GO111MODULE=off forces Go to behave the GOPATH way, even outside of GOPATH.
GO111MODULE=auto is the default mode. In this mode, Go will behave
similarly to GO111MODULE=on when you are outside of GOPATH,
similarly to GO111MODULE=off when you are inside the GOPATH even if a go.mod is present.