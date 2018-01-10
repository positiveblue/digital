package core

import (
	"time"
)

type DigitalCoreError struct {
	message   string
	function  string
	timestamp time.Time
}
