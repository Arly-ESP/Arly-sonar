package templates_test

import (
	"errors"
	"html/template"
	"io"
	"reflect"
	"strings"
	"testing"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/arly/arlyApi/templates"
	"github.com/stretchr/testify/assert"
)

func TestGenerateVerificationCodeEmail_Success(t *testing.T) {
	code := "ABC123"
	htmlOutput, err := templates.GenerateVerificationCodeEmail(code)
	assert.NoError(t, err)
	// Check that the output contains the verification code.
	assert.True(t, strings.Contains(htmlOutput, code), "Output should contain the verification code")
	// Check for some known substrings.
	assert.True(t, strings.Contains(htmlOutput, "<html>"), "Output should be valid HTML")
	assert.True(t, strings.Contains(htmlOutput, "support@arly.com"), "Output should mention support email")
}

func TestGenerateVerificationCodeEmail_ExecuteError(t *testing.T) {
	// Patch the Execute method of *template.Template to force an error.
	patch := gomonkey.ApplyMethod(reflect.TypeOf(&template.Template{}), "Execute", func(t *template.Template, w io.Writer, data interface{}) error {
		return errors.New("forced execute error")
	})
	defer patch.Reset()

	htmlOutput, err := templates.GenerateVerificationCodeEmail("ABC123")
	assert.Error(t, err)
	assert.Equal(t, "forced execute error", err.Error())
	assert.Equal(t, "", htmlOutput)
}
