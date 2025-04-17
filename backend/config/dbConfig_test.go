package config

import (
	"os"
	"strings"
	"testing"
)

func TestMain(m *testing.M) {
	envContent := `
DB_HOST=localhost
DB_USER=testuser
DB_PASSWORD=testpass
DB_NAME=testdb
DB_PORT=5432
DB_SSLMODE=disable
DB_TIMEZONE=Asia/Shanghai
ENVIRONMENT=development
PORT=5050
SERVER_URL=http://0.0.0.0
JWT_SECRET=abcdefghijklmnopqrstuvwxyz1234567890abcd
JWT_EXPIRE=24h
JWT_REFRESH_SECRET=abcdefghijklmnopqrstuvwxyz0987654321abcd
JWT_REFRESH_EXPIRE=72h
OPENAI_API_KEY=dummy_openai_api_key
SMTP_HOST=smtp.example.com
SMTP_PORT=587
SMTP_USER=smtpuser
SMTP_NAME=smtpname
SMTP_PASSWORD=smtppass
SMTP_SSL=true
`
	if err := os.WriteFile("./.env", []byte(envContent), 0644); err != nil {
		os.Exit(1)
	}

	// Run tests.
	code := m.Run()

	_ = os.Remove("./.env")

	os.Exit(code)
}

func TestDbConfig(t *testing.T) {
	dsn := DbConfig()
	expectedSubstrings := []string{
		"host=localhost",
		"user=testuser",
		"password=testpass",
		"dbname=testdb",
		"port=5432",
		"sslmode=disable",
		"TimeZone=Asia/Shanghai",
	}
	for _, substr := range expectedSubstrings {
		if !strings.Contains(dsn, substr) {
			t.Errorf("DSN missing expected substring %q. Got: %s", substr, dsn)
		}
	}
}

func TestEnvironmentType(t *testing.T) {
	env := environmentType()
	if env != "development" {
		t.Errorf("Expected environment 'development', got %s", env)
	}
}

func TestServerPort(t *testing.T) {
	port := ServerPort()
	if port != "5050" {
		t.Errorf("Expected port '5050', got %s", port)
	}
}

func TestServerUrl(t *testing.T) {
	url := ServerUrl()
	expected := "http://0.0.0.0:5050"
	if url != expected {
		t.Errorf("Expected server URL %q, got %q", expected, url)
	}
}

func TestJwtInformation(t *testing.T) {
	jwtInfo := JwtInformation()
	if jwtInfo.secret == "" || jwtInfo.expire == "" || jwtInfo.refreshSecret == "" || jwtInfo.refreshExpire == "" {
		t.Error("Expected all JWT environment variables to be set")
	}
	if jwtInfo.secret != "abcdefghijklmnopqrstuvwxyz1234567890abcd" {
		t.Errorf("Unexpected JWT secret: %s", jwtInfo.secret)
	}
	if jwtInfo.expire != "24h" {
		t.Errorf("Unexpected JWT expire: %s", jwtInfo.expire)
	}
	if jwtInfo.refreshSecret != "abcdefghijklmnopqrstuvwxyz0987654321abcd" {
		t.Errorf("Unexpected JWT refresh secret: %s", jwtInfo.refreshSecret)
	}
	if jwtInfo.refreshExpire != "72h" {
		t.Errorf("Unexpected JWT refresh expire: %s", jwtInfo.refreshExpire)
	}
}

func TestGetOpenAiConfig(t *testing.T) {
	config := GetOpenAiConfig()
	if config.APIKey != "dummy_openai_api_key" {
		t.Errorf("Expected OPENAI_API_KEY 'dummy_openai_api_key', got %s", config.APIKey)
	}
	if config.Model != "gpt-3.5-turbo" {
		t.Errorf("Expected model 'gpt-3.5-turbo', got %s", config.Model)
	}
	if config.Temperature != 0.7 {
		t.Errorf("Expected temperature 0.7, got %v", config.Temperature)
	}
}

func TestGetSMTPConfig(t *testing.T) {
	smtpConfig := GetSMTPConfig()
	if smtpConfig.Host != "smtp.example.com" {
		t.Errorf("Expected SMTP_HOST 'smtp.example.com', got %s", smtpConfig.Host)
	}
	if smtpConfig.Port != 587 {
		t.Errorf("Expected SMTP_PORT 587, got %d", smtpConfig.Port)
	}
	if smtpConfig.User != "smtpuser" {
		t.Errorf("Expected SMTP_USER 'smtpuser', got %s", smtpConfig.User)
	}
	if smtpConfig.Name != "smtpname" {
		t.Errorf("Expected SMTP_NAME 'smtpname', got %s", smtpConfig.Name)
	}
	if smtpConfig.Password != "smtppass" {
		t.Errorf("Expected SMTP_PASSWORD 'smtppass', got %s", smtpConfig.Password)
	}
	if smtpConfig.SSL != true {
		t.Errorf("Expected SMTP_SSL true, got %v", smtpConfig.SSL)
	}
}
