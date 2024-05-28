#!/bin/sh

name=$(basename $(pwd))

# -serial stdio

qemu-system-arm -machine mps2-an500  -cpu cortex-m7 -nographic -monitor none -serial none --semihosting-config enable=on,target=native,userspace=on -kernel $name.elf
