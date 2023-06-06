package version

import (
	"fmt"
	"github.com/urfave/cli/v2"
)

var Version = "v1.0.8"

func RunVersion(c *cli.Context) error {
	fmt.Println(Version)
	return nil
}
