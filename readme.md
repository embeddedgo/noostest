## Test your Embedded Go code on your PC

### Prerequisites

1. Go complier.

   You can download it from [go.dev/dl](https://go.dev/dl/).

2. Git command.

   To instll git on Linux use the package manager provided by your Linux distribution (apt, pacman, rpm, ...).

   Windows users may check the [git for Windows](https://gitforwindows.org/) website.

   The Mac users may use the git command provided by the [Xcode](https://developer.apple.com/xcode/) commandline tools. Another way is to use the [Homebrew](https://brew.sh/) package manager.

3. QEMU

For ARM-M (Cortex-M, GOARCH=thumb) install the qemu-system-arm. For RISCV64 (GOARCH=riscv64) instal the qemu-system-riscv64.

### Getting started

1. Install the Embedded Go toolchain.

   Make sure the `$GOPATH/bin` directory is in your `PATH`, as tools installed with the `go install` command will be placed here. If you didn't set the `GOPATH` environment variable manually you can find its default value using the `go env GOPATH` command.

   Then install the Embedded Go toolchain using the following two commands:

   ```sh
   go install github.com/embeddedgo/dl/go1.24.5-embedded@latest
   go1.24.5-embedded download
   ```

2. Install egtool.

   ```sh
   go install github.com/embeddedgo/tools/egtool@latest
   ```

3. Run examples

   ```sh
   $ GOENV=noos_thumb.env egtool run ./examples/time
   1970-01-01 00:00:00.013007 +0000 UTC m=+0.013007000
   1970-01-01 00:00:01.036484 +0000 UTC m=+1.036484000
   1970-01-01 00:00:02.040499 +0000 UTC m=+2.040499000
   ^C
   $ GOENV=noos_riscv64.env egtool run ./examples/time
   1970-01-01 00:00:00.0674071 +0000 UTC m=+0.067407100
   1970-01-01 00:00:01.0792119 +0000 UTC m=+1.079211900
   1970-01-01 00:00:02.0811142 +0000 UTC m=+2.081114200
   ^C
   $ GOENV=../noos_thumb.env egtool run ./examples/semihostfs 1 2
   Hello over stderr!
   os.Args: []string{"/tmp/go-build1014682335/b001/exe/semihostfs.elf", "1", "2"}

   Enter a number: 12
   Writing 12 to the /x.txt file.
   Reading from the /x.txt file:
   12
   Delete /x.txt
   Exit
   $
   ```

4. Run tests

   ```sh
   $ GOENV=noos_riscv64.env egtool test ./tests
   ok      github.com/embeddedgo/noostest/tests    0.169s
   ```
