#!/bin/sh

name=$(basename $(pwd))

gdb-multiarch --tui -ex 'target extended-remote :1234' $name.elf
