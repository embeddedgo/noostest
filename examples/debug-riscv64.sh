#!/bin/sh

name=$(basename $(pwd))

system='-machine virt -smp 2 -m 32'

qemu-system-riscv64 $system -s -S -nographic -monitor none -serial none --semihosting-config enable=on,target=native,userspace=on -bios $name.elf
