package log

import "github.com/crazy-choose/helper/usage"

func defLogName() string {
	return usage.HomeDir() + "/logs/" + usage.ExeName() + ".log"
}
