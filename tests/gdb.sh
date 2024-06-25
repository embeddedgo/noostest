#!/bin/sh

gdb-multiarch --tui \
	-ex 'target extended-remote :1234' \
	-ex 'focus cmd' \
	tests.test
