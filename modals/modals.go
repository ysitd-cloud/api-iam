package modals

//go:generate  protoc -I . -I ../vendor --gogoslick_out=plugins=grpc:. ./modals.proto
