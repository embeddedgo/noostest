#!/bin/sh

export GOOS=noos
export GOARCH=riscv64
export GOFLAGS="-tags=noostest '-ldflags=-stripfn=1 -M=0x80000000:32M'"

go build $@
