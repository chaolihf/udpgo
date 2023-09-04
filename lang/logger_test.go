package lang

import (
	"testing"
)

func TestLogger(t *testing.T) {
	logger := InitLogger()
	logger.Info("aa")
}

func TestProductLogger(t *testing.T) {
	logger := InitProductLogger("app.log", 1, 1, 1)
	logger.Info("info")
	logger.Debug("debug")
	logger.Error("error")
}

func TestLeft(t *testing.T) {
	var result string
	result = LeftString("abc", 1)
	if result != "a" {
		t.Fatal("fail 1")
	}
	result = LeftString("abc", 3)
	if result != "abc" {
		t.Fatal("fail 2")
	}
	result = LeftString("abc", 4)
	if result != "abc" {
		t.Fatal("fail 3")
	}
	result = LeftString("中ab吧c", 1)
	if result != "中" {
		t.Fatal("fail 4")
	}
	result = LeftString("a中b吧c", 3)
	if result != "a中b" {
		t.Fatal("fail 5")
	}
	result = LeftString("a中b吧c", 40)
	if result != "a中b吧c" {
		t.Fatal("fail 6")
	}
}

func TestFormatter(t *testing.T) {
	var formatter FormatStringer
	formatter.Parse("中{aa}dd")
	if formatter.GetParamNames()[0] != "aa" {
		t.Fatal("test 1")
	}
	formatter.Replace("aa", 0, "bb")
	if formatter.String() != "中bbdd" {
		t.Fatal("test 2")
	}
	formatter.Parse("中dd")
	if len(formatter.GetParamNames()) != 0 {
		t.Fatal("test 3")
	}
	formatter.Parse(`{中\\aa}dd{FF}dd{FF}`)
	if formatter.GetParamNames()[0] != `中\aa` {
		t.Fatal("test 4")
	}
	if formatter.GetParamNames()[1] != `FF` {
		t.Fatal("test 5")
	}
	formatter.Replace(`中\aa`, 0, "bb")
	formatter.ReplaceAll("FF", "开始")
	if formatter.String() != "bbdd开始dd开始" {
		t.Fatal("test 6")
	}
}
