package lang

import "testing"

func TestLogger(t *testing.T) {
	logger := InitLogger()
	logger.Info("aa")
}
