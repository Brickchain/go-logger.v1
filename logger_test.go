package logger

import (
	"testing"
	"strings"
	"gitlab.brickchain.com/brickchain/oauth2-idp/logger"
	"encoding/json"
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
	org_writer := logger.GetLogger().Logger.Out
	out := DummyWriter{}
	SetOutput(&out)
	Info("test")

	data := parseJson(out.GetBuffer())

	if data["msg"] != "test" {
		t.Error("Written data not same after fetch")
	}

	SetOutput(org_writer)
}

func TestSetFormatter(t *testing.T) {
	org_writer := logger.GetLogger().Logger.Out
	out := DummyWriter{}
	SetOutput(&out)
	SetFormatter("text")

	Info("Some test string")

	str := string(out.GetBuffer())
	if !strings.Contains(str, "Some test string") {
		t.Error("Written data not same after fetch")
	}

	SetFormatter("json")
	SetOutput(org_writer)
}

func TestDebug(t *testing.T) {
	org_writer := logger.GetLogger().Logger.Out
	out := DummyWriter{}
	SetOutput(&out)
	SetLevel("debug")

	Debug("test")

	data := parseJson(out.GetBuffer())

	if data["level"] != "debug" {
		t.Error("Message not written with debug level")
	}

	SetLevel("info")
	SetOutput(org_writer)
}

func parseJson(data []byte) map[string]interface{} {
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