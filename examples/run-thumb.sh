#!/bin/sh

name=$(basename $(pwd))

system='-machine mps2-an500 -cpu cortex-m7'

qemu-system-arm $system -nographic -monitor none -serial none --semihosting-config enable=on,target=native,userspace=on -kernel $name.elf
