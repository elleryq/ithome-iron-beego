package checks

import (
	"errors"
	"math/rand"
)

type DatabaseCheck struct {
}

func (dc *DatabaseCheck) Check() error {
	if dc.isConnected() {
		return nil
	} else {
		return errors.New("can't connect database")
	}
}

func (dc *DatabaseCheck) isConnected() bool {
	x := rand.Intn(100)
	if x%2 == 1 {
		return false
	}
	return true
}
