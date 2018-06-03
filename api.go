package iam

//go:generate protoc -I . -I $GOPATH/src --gogoslick_out=plugins=grpc:. ./service.proto
