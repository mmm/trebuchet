# SciML Playground - Trebuchet

Communicate trebuchet model params and inferences over grpc.

This is intended as a playground to work with differentiable programming models
where we use a combo of DL and domain knowledge to solve scientific problems.

This uses docker-compose atm, but can be easily terraformed to use cloud run
and/or k8s.


## You need a working go environment...

on cloudshell, you'll need to add `protoc` to compile protobufs

    go get -u github.com/golang/protobuf/protoc-gen-go

then just run `make` or `make clean` from the top-level directory of this repo.
