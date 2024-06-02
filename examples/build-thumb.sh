#!/bin/sh

export GOTARGET=noostest
export GOARCH=thumb
export GOARM=7,hardfloat
export GOMEM=0x60000000:16M,0x20000000:4M
export GOTEXT=0x00000000:4M

emgo build

