#!/bin/sh

export GOOS=noos
export GOARCH=riscv64
export GOFLAGS="-tags=noostest '-ldflags=-stripfn=1 -M=0x80000000:32M'"
#export GOFLAGS="-tags=noostest '-ldflags=-stripfn=1 -M=0x80000000:32M' '-asmflags=all=-d=compressinstructions=0' '-gcflags=all=-d=compressinstructions=0'"

go build $@
