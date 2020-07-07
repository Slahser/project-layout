package opkg

import "flag"

var (
	fVerbose bool
	fVersion bool
)

func init() {
	flag.BoolVar(&fVerbose, "v", false, "Do not compact consecutive bytes of fields")
	flag.BoolVar(&fVersion, "version", false, "Print version and exit")
}
