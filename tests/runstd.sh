#!/bin/sh

export GOTARGET=noostest

cd $(emgo env GOROOT)/src

for GOARCH in thumb riscv64; do
	export GOARCH
	echo "####  GOARCH=$GOARCH  ####"
	for d in $(find -type d '!' -name .); do
		[ "$(find $d -maxdepth 1 -name '*_test.*')" ] && emgo test $d
	done
done

