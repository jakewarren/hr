package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	flag "github.com/spf13/pflag"

	"github.com/arsham/rainbow/rainbow"
	"github.com/jakewarren/hr"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	noColor := envDisableColor()

	disableColor := flag.Bool("nocolor", false, "disable color")
	flag.Parse()

	// default line setup
	line := "-|#|-"

	if *disableColor {
		noColor = true
	}

	if flag.NArg() >= 1 {
		line = flag.Arg(0)

		// TODO: read from config file
		if flag.Arg(0) == "cut" {
			line = `---- ✂ `
		}
	}

	for _, l := range strings.Split(line, "|") {

		hr := hr.HorizontalRule(l)

		if !noColor {
			rb := rainbow.Light{
				Reader: strings.NewReader(hr),
				Writer: os.Stdout,
			}
			rb.Paint()
		} else {
			fmt.Print(hr)
		}

	}
}

func envDisableColor() bool {
	// check for the existance of NO_COLOR to satisfy the nocolor standard http://no-color.org
	_, exists := os.LookupEnv("NO_COLOR")

	if exists {
		return true
	}

	// disable color if terminal is set to dumb
	val, _ := os.LookupEnv("TERM")

	if val == "DUMB" {
		return true
	}

	return false
}
