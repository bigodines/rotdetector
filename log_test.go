package rotdetector

import (
	"bytes"
	"strings"
	"testing"
)

func TestSetLogLevel(t *testing.T) {
	SetLogLevel(DEBUG)
	if logLevel != DEBUG {
		t.Errorf("Expected logLevel to be %v, got %v", DEBUG, logLevel)
	}

	SetLogLevel(INFO)
	if logLevel != INFO {
		t.Errorf("Expected logLevel to be %v, got %v", INFO, logLevel)
	}
}

func TestDebug(t *testing.T) {
	SetLogLevel(DEBUG)

	var buf bytes.Buffer
	logger.SetOutput(&buf)

	Debug("debug message")
	if !strings.Contains(buf.String(), "DEBUG") && !strings.Contains(buf.String(), "debug message") {
		t.Errorf("Expected 'DEBUG: debug message', got %v", buf.String())
	}
}

func TestInfo(t *testing.T) {
	SetLogLevel(INFO)

	var buf bytes.Buffer
	logger.SetOutput(&buf)

	Info("info message")
	if !strings.Contains(buf.String(), "INFO") && !strings.Contains(buf.String(), "info message") {
		t.Errorf("Expected 'INFO: info message', got %v", buf.String())
	}
}

func TestWarning(t *testing.T) {
	SetLogLevel(WARNING)

	var buf bytes.Buffer
	logger.SetOutput(&buf)

	Warning("warning message")
	if !strings.Contains(buf.String(), "WARNING") && !strings.Contains(buf.String(), "warning message") {
		t.Errorf("Expected 'WARNING: warning message', got %v", buf.String())
	}
}

func TestError(t *testing.T) {
	SetLogLevel(ERROR)

	var buf bytes.Buffer
	logger.SetOutput(&buf)

	Error("error message")
	if !strings.Contains(buf.String(), "ERROR") && !strings.Contains(buf.String(), "error message") {
		t.Errorf("Expected 'ERROR: error message', got %v", buf.String())
	}
}

func TestLogLevelFiltering(t *testing.T) {
	var buf bytes.Buffer
	logger.SetOutput(&buf)

	SetLogLevel(INFO)
	Debug("this should not appear")
	if strings.Contains(buf.String(), "DEBUG: this should not appear") {
		t.Errorf("Expected no debug message, got %v", buf.String())
	}

	SetLogLevel(WARNING)
	Info("this should not appear")
	if strings.Contains(buf.String(), "INFO: this should not appear") {
		t.Errorf("Expected no info message, got %v", buf.String())
	}

	SetLogLevel(ERROR)
	Warning("this should not appear")
	if strings.Contains(buf.String(), "WARNING: this should not appear") {
		t.Errorf("Expected no warning message, got %v", buf.String())
	}
}
