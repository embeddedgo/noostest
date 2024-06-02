#!/bin/sh

name=$(basename $(pwd))

system='-machine virt -cpu rv64,pmp=false,mmu=false,c=false -smp 2 -m 32'

opts='-nographic -monitor none -serial none --semihosting-config enable=on,target=native,userspace=on'

qemu-system-riscv64 $system -s -S $opts -bios $name.elf
