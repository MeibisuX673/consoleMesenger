package connection

import (
	"flag"
)

func GetConnection() string {
	return flag.Arg(0)
}
