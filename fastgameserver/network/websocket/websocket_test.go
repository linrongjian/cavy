package ws

import (
	"sync"
	"testing"
	"trainserver/util/logger"
)

func Test(t *testing.T) {

	var sm sync.Map

	sm.Delete("a")

	logger.Info("ok")
}
