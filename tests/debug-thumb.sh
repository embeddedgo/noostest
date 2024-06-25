#!/bin/sh

system='-machine mps2-an500 -cpu cortex-m7'

qemu-system-arm $system -s -S -nographic -monitor none -serial none --semihosting-config enable=on,target=native,userspace=on -kernel tests.test
