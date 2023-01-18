# grpc
golang microservice grpc

# Run project
´go run folder_name/main.go'

# Install PROTO BUFFER

´go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26 GENERADOR PROTO BUFFER GO
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1 GRPC GO´

# PROTO BUFFER IS NOT WORKING

´export GOPATH=$HOME/go                                                                                                  
export PATH=$PATH:$GOPATH/bin´

then run command ´protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/student.proto´

# DEPENDENCIES

Postgresql go get github.com/lib/pq
GRPC go get google.golang.org/grpc