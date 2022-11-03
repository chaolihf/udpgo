package lang

import "testing"

func TestLogger(t *testing.T) {
	logger := InitLogger()
	logger.Info("aa")
}

func TestLeft(t *testing.T) {
	var result string
	result = LeftString("abc", 1)
	if "a" != result {
		t.Fatal("fail 1")
	}
	result = LeftString("abc", 3)
	if "abc" != result {
		t.Fatal("fail 2")
	}
	result = LeftString("abc", 4)
	if "abc" != result {
		t.Fatal("fail 3")
	}
	result = LeftString("中ab吧c", 1)
	if "中" != result {
		t.Fatal("fail 4")
	}
	result = LeftString("a中b吧c", 3)
	if "a中b" != result {
		t.Fatal("fail 5")
	}
	result = LeftString("a中b吧c", 40)
	if "a中b吧c" != result {
		t.Fatal("fail 6")
	}
}
