package main

import (
	"fmt"
	"github.com/jessevdk/go-flags"
	"golang.org/x/crypto/ssh/terminal"
	"strconv"
	"os"
	"runtime"
	"regexp"
	"strings"
	"io/ioutil"
)

var BuildVersion string

type Options struct {
	Version  bool   `long:"version" description:"print version"`
}

func main() {
	args, opts := getArgs()
	if opts.Version {
		printVersion()
		os.Exit(0)
	}

	var hexip string
	if len(args) > 0 {
		hexip = args[0]
	} else {
		if terminal.IsTerminal(0) {
			printError()
		} else {
			b, err := ioutil.ReadAll(os.Stdin)
			if err != nil {
				os.Exit(1)
			}
			hexip = strings.TrimRight(string(b), "\r\n")
		}
	}

	r := regexp.MustCompile(`[0123456789abcdef]{2}`)
	splitedHexIps := r.FindAllStringSubmatch(hexip, -1)

	if len(hexip) != 8 || len(splitedHexIps) != 4 {
		printError()
	}

	fmt.Printf("%d.%d.%d.%d\n", decodeHex(splitedHexIps[0][0]),
								decodeHex(splitedHexIps[1][0]),
								decodeHex(splitedHexIps[2][0]),
								decodeHex(splitedHexIps[3][0]))
}

func getArgs() ([]string, Options) {
	opts := Options{}
	args, err := flags.Parse(&opts)
	if err != nil {
		os.Exit(1)
	}
	return args, opts
}

func printVersion() {
	fmt.Printf(`hexip %s
Compiler: %s %s
`,
		BuildVersion,
		runtime.Compiler,
		runtime.Version())
}

func printError() {
	fmt.Println("Need hexip like [0a004f79]")
	fmt.Fprintf(os.Stderr, "Need hexip like [a000001]\n")
	os.Exit(1)
}

func decodeHex (hexString string) int64 {
	i, _ := strconv.ParseInt(hexString, 16, 10)
	return i
}
