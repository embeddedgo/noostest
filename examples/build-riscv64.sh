#!/bin/sh

export GOTARGET=noostest
export GOARCH=riscv64
export GOMEM=0x80000000:32M

emgo build
