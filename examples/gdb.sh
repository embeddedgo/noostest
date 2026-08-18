#!/bin/sh

name=$(basename $(pwd))

gdb-multiarch --write --tui \
	-ex 'target extended-remote :1234' \
	-ex 'layout split' \
	-ex 'focus cmd' \
	$name.elf
