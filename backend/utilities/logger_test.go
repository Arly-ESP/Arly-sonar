package utilities

import (
	"bytes"
	"log"
	"os"
	"regexp"
	"strings"
	"testing"
	"time"
)

func TestCurrentEnv(t *testing.T) {
	os.Unsetenv("ENVIRONMENT")
	if env := currentEnv(); env != "development" {
		t.Errorf("Expected 'development' when ENVIRONMENT is unset, got %s", env)
	}

	os.Setenv("ENVIRONMENT", "production")
	if env := currentEnv(); env != "production" {
		t.Errorf("Expected 'production', got %s", env)
	}
}

func TestCurrentTime(t *testing.T) {
	nowStr := currentTime()
	_, err := time.Parse(time.RFC3339, nowStr)
	if err != nil {
		t.Errorf("currentTime did not return valid RFC3339 time: %v", err)
	}
}

func setupTestLogger() *bytes.Buffer {
	var buf bytes.Buffer
	logger = log.New(&buf, "ArlyAPI: ", log.LstdFlags|log.Lshortfile)
	return &buf
}

func TestLogInfo(t *testing.T) {
	os.Setenv("ENVIRONMENT", "development")
	buf := setupTestLogger()

	LogInfo("Test info message")
	output := buf.String()

	if !strings.Contains(output, "INFO: Test info message") {
		t.Errorf("Expected output to contain 'INFO: Test info message', got: %s", output)
	}
	re := regexp.MustCompile(`\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}`)
	if !re.MatchString(output) {
		t.Errorf("Expected output to contain a valid timestamp, got: %s", output)
	}
}

func TestLogError(t *testing.T) {
	buf := setupTestLogger()

	LogError("Test error message", nil)
	output := buf.String()

	expected := "ERROR: Test error message: <nil>"
	if !strings.Contains(output, expected) {
		t.Errorf("Expected output to contain %q, got: %s", expected, output)
	}
}

func TestLogPanic(t *testing.T) {
	buf := setupTestLogger()
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected LogPanic to panic, but it did not")
		} else {
			output := buf.String()
			if !strings.Contains(output, "PANIC: Test panic message:") {
				t.Errorf("Expected output to contain 'PANIC: Test panic message:', got: %s", output)
			}
		}
	}()
	LogPanic("Test panic message", nil)
	t.Error("LogPanic did not panic")
}
