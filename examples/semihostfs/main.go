// Copyright 2024 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"io"
	"os"

	_ "github.com/embeddedgo/noostest/init"
)

func fatalErr(what string, err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", what, err)
		os.Exit(1)
	}
}

func main() {
	fmt.Fprintln(os.Stderr, "Hello over stderr!")
	fmt.Printf("os.Args: %#v\n\n", os.Args)

	var x float64
	fmt.Print("Enter a number: ")
	fmt.Scanf("%g", &x)

	fname := "/x.txt"

	fmt.Println("Writing", x, "to the", fname, "file.")
	f, err := os.Create(fname)
	fatalErr("create", err)
	_, err = fmt.Fprintln(f, x)
	fatalErr("write", err)
	err = f.Close()
	fatalErr("close", err)

	fmt.Println("Reading from the", fname, "file:")
	f, err = os.Open(fname)
	fatalErr("open", err)
	_, err = io.Copy(os.Stdout, f)
	fatalErr("copy", err)
	err = f.Close()
	fatalErr("close", err)

	fmt.Println("Delete", fname)
	err = os.Remove(fname)
	fatalErr("remove", err)

	fmt.Println("Exit")
}
