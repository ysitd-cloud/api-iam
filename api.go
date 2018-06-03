package iam

//go:generate protoc -I . -I vendor --gogoslick_out=plugins=grpc:. ./service.proto
