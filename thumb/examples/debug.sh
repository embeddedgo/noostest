#!/bin/sh

name=$(basename $(pwd))

qemu-system-arm -machine mps2-an500 -cpu cortex-m7 -gdb tcp::3333 -S -nographic -monitor none -serial none --semihosting-config enable=on,target=native,userspace=on -kernel $name.elf