package lang

import "testing"

func TestLogger(t *testing.T) {
	InitLogger()
	logger.Info("aa")
}
