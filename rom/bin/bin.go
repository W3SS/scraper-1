// Package bin decodes bin files.

package bin

import (
	"github.com/sselph/scraper/rom"
)

func init() {
	rom.RegisterFormat(".bin", rom.Noop)
}
