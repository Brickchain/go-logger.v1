package logger

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"testing"
)

func init() {
	out := DummyWriter{}
	SetOutput(&out)

	SetFormatter("json")
}

func TestInit(t *testing.T) {
	l := GetLogger()
	if l.Logger == nil {
		t.Error("Logger not initialized")
	}
}

func TestAddContext(t *testing.T) {
	AddContext("test", "stuff")
	l := GetLogger()

	if l.Data["test"] == nil {
		t.Error("Added context not present in Fields")
	}
}

func TestSetOutput(t *testing.T) {
	orgWriter := GetLogger().Logger.Out
	out := DummyWriter{}
	SetOutput(&out)
	Info("test")

	data := parseJSON(out.GetBuffer())

	if data["msg"] != "test" {
		t.Error("Written data not same after fetch")
	}

	SetOutput(orgWriter)
}

func TestSetFormatter(t *testing.T) {
	orgWriter := GetLogger().Logger.Out
	out := DummyWriter{}
	SetOutput(&out)
	SetFormatter("text")

	Info("Some test string")

	str := string(out.GetBuffer())
	if !strings.Contains(str, "Some test string") {
		t.Error("Written data not same after fetch")
	}

	SetFormatter("json")
	SetOutput(orgWriter)
}

func TestDebug(t *testing.T) {
	orgWriter := GetLogger().Logger.Out
	out := DummyWriter{}
	SetOutput(&out)
	SetLevel("debug")

	Debug("test")

	data := parseJSON(out.GetBuffer())

	if data["level"] != "debug" {
		t.Error("Message not written with debug level")
	}

	SetLevel("info")
	SetOutput(orgWriter)
}

func TestWithField(t *testing.T) {
	orgWriter := GetLogger().Logger.Out
	out := DummyWriter{}
	SetOutput(&out)

	localLogger := WithField("testkey", "testvalue")

	localLogger.Info("test")

	data := parseJSON(out.GetBuffer())

	if data["testkey"] == nil {
		t.Error("Message does not have testkey set")
	}

	SetLevel("info")
	SetOutput(orgWriter)
}

func TestWarn(t *testing.T) {
	orgWriter := GetLogger().Logger.Out
	out := DummyWriter{}
	SetOutput(&out)
	SetLevel("warn")

	Warn("test")

	data := parseJSON(out.GetBuffer())

	if data["level"] != "warning" {
		t.Error("Message not written with warning level")
	}

	SetLevel("info")
	SetOutput(orgWriter)
}

func TestError(t *testing.T) {
	orgWriter := GetLogger().Logger.Out
	out := DummyWriter{}
	SetOutput(&out)
	SetLevel("error")

	Error("test")

	data := parseJSON(out.GetBuffer())

	if data["level"] != "error" {
		t.Error("Message not written with error level")
	}

	SetLevel("info")
	SetOutput(orgWriter)
}

func TestErrorf(t *testing.T) {
	orgWriter := GetLogger().Logger.Out
	out := DummyWriter{}
	SetOutput(&out)
	SetLevel("error")

	Errorf("Msg: %s", "test")

	data := parseJSON(out.GetBuffer())

	if data["level"] != "error" {
		t.Error("Message not written with error level")
	}

	if data["msg"] != "Msg: test" {
		t.Error("Formatting incorrect")
	}

	SetLevel("info")
	SetOutput(orgWriter)
}

func TestInfof(t *testing.T) {
	orgWriter := GetLogger().Logger.Out
	out := DummyWriter{}
	SetOutput(&out)
	SetLevel("info")

	Infof("Msg: %s", "test")

	data := parseJSON(out.GetBuffer())

	if data["level"] != "info" {
		t.Error("Message not written with info level")
	}

	if data["msg"] != "Msg: test" {
		t.Error("Formatting incorrect")
	}

	SetLevel("info")
	SetOutput(orgWriter)
}

func TestWarningf(t *testing.T) {
	orgWriter := GetLogger().Logger.Out
	out := DummyWriter{}
	SetOutput(&out)
	SetLevel("warn")

	Warningf("Msg: %s", "test")

	data := parseJSON(out.GetBuffer())

	if data["level"] != "warning" {
		t.Error("Message not written with warning level")
	}

	if data["msg"] != "Msg: test" {
		t.Error("Formatting incorrect")
	}

	SetLevel("info")
	SetOutput(orgWriter)
}

func TestDebugf(t *testing.T) {
	orgWriter := GetLogger().Logger.Out
	out := DummyWriter{}
	SetOutput(&out)
	SetLevel("debug")

	Debugf("Msg: %s", "test")

	data := parseJSON(out.GetBuffer())

	if data["level"] != "debug" {
		t.Error("Message not written with debug level")
	}

	if data["msg"] != "Msg: test" {
		t.Error("Formatting incorrect")
	}

	SetLevel("info")
	SetOutput(orgWriter)
}

func TestUnknownLogLevel(t *testing.T) {
	SetLevel("abcd")

	if GetLoglevel() != "info" {
		t.Error("Level not set to info when given an unknown level")
	}

	SetLevel("info")
}

func TestWithFields(t *testing.T) {
	_ = WithFields(Fields{"test": "ok"})
}

func parseJSON(data []byte) map[string]interface{} {
	d := make(map[string]interface{})
	_ = json.Unmarshal(data, &d)

	return d
}

type DummyWriter struct {
	buffer []byte
}

func (w *DummyWriter) Write(p []byte) (n int, err error) {
	w.buffer = p

	return len(p), nil
}

func (w *DummyWriter) GetBuffer() []byte {
	return w.buffer
}

func TestForContext(t *testing.T) {
	orgWriter := GetLogger().Logger.Out
	defer SetOutput(orgWriter)
	out := DummyWriter{}
	SetOutput(&out)
	SetLevel("info")

	ctx := context.WithValue(context.Background(), 0, "abc")
	l := ForContext(ctx)
	l.Info("test stuff")

	fmt.Println(string(out.GetBuffer()))
}
