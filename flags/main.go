package flags

import (
	"flag"
)

var (
	IsProduction *bool
)

func Init() {
	defer flag.Parse()
	IsProduction = flag.Bool("prod", false, "Run in production mode")
}
