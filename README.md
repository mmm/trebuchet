# SciML Playground - Trebuchet

Communicate trebuchet model params and inferences over grpc.

This is intended as a playground to work with differentiable programming models
where we use a combo of DL and domain knowledge to solve scientific problems.

This uses docker-compose atm, but can be easily terraformed to use cloud run
and/or k8s.


## Background

- A soldier has a trebuchet, the distance to a target, and the current wind
  conditions
- A sorcerer accepts that info and returns a firing solution to the soldier,
  which consists of an angle of release and projectile mass.
- Profit

## Setup

You need a working go environment...

on cloudshell, you'll need to add `protoc` to compile protobufs

    go get -u github.com/golang/protobuf/protoc-gen-go

then just run `make` or `make clean` from the top-level directory of this repo.


## Usage

Spin up the local cluster

    docker-compose up -d
    docker-compose ps

then check out the logs

    docker-compose logs sorcerer

and the more interesting

    docker-compose logs soldier

and you can re-run the client multiple times

    docker-compose up soldier


## Tests

TBD


## Development

To build, you can run

    docker-compose build

to compile the go code as well as build docker images.

Within each project directory, you can riff on code

    cd sorcerer
    vi main.go
    go build

and then run via

    docker-compose build

If you ever muck about with your protobufs, you'll have to do a full

    make clean
    make
    docker-compose build

cycle.


