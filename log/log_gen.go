package log

// CODE GENERATED AUTOMATICALLY
// THIS FILE MUST NOT BE EDITED BY HAND

import (
	"fmt"

	"github.com/ldez/go-git-cmd-wrapper/types"
)

// MaxCount Limit the number of commits to output.
// -<number>, -n <number>, --max-count=<number>
func MaxCount(number string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions(fmt.Sprintf("--max-count=%s", number))
	}
}

// Pretty Pretty-print the contents of the commit logs in a given format, where <format> can be one of oneline, short, medium, full, fuller, email, raw, format:<string> and tformat:<string>. When <format> is none of the above, and has %placeholder in it, it acts as if --pretty=tformat:<format> were given.
//
//  See the "PRETTY FORMATS" section for some additional details for each format. When =<format> part is omitted, it defaults to medium.
//
//  Note: you can specify the default pretty format in the repository configuration (see git-config(1)).
// --pretty[=<format>], --format=<format>
func Pretty(format string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		if len(format) == 0 {
			g.AddOptions("--pretty")
		} else {
			g.AddOptions(fmt.Sprintf("--pretty=%s", format))
		}
	}
}
