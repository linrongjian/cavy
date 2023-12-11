package retry

import (
	"errors"
	"testing"
	"time"

	"github.com/linrongjian/cavy/common/xlog"
)

func TestRetry(t *testing.T) {
	err := Retry(func() error {
		xlog.Warnf("retry func")
		return errors.New("please retry")
	}, 3, 2*time.Second)
	if err != nil {
		xlog.Error(err)
	}
}
