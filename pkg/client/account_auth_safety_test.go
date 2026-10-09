package client

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/kataras/golog"
)

type authTransport func(*http.Request) (*http.Response, error)

func (f authTransport) RoundTrip(r *http.Request) (*http.Response, error) { return f(r) }
func TestAccountAuthenticationDoesNotLogCredentials(t *testing.T) {
	var logs bytes.Buffer
	original := golog.Default
	golog.Default = golog.New()
	golog.SetOutput(&logs)
	defer func() { golog.Default = original }()
	for _, accountUser := range []bool{false, true} {
		cfg := NewZSConfig("cloud.invalid", 8080, "zstack").LoginAccount("private-account", "private-password").Debug(false)
		if accountUser {
			cfg.authType = AuthTypeAccountUser
			cfg.accountUserName = "private-user"
		}
		cli := &ZSClient{ZSHttpClient: &ZSHttpClient{ZSConfig: cfg, httpClient: &http.Client{Transport: authTransport(func(r *http.Request) (*http.Response, error) {
			return &http.Response{StatusCode: 401, Header: make(http.Header), Body: io.NopCloser(strings.NewReader(`{"error":{"details":"private-password"}}`)), Request: r}, nil
		})}}}
		if _, err := cli.Login(context.Background()); err == nil {
			t.Fatal("expected login refusal")
		}
		cli.LoadSession("private-session")
		if err := cli.Logout(context.Background()); err == nil {
			t.Fatal("expected logout refusal")
		}
	}
	if logs.Len() != 0 {
		t.Fatal("authentication wrote a raw error to the SDK logger")
	}
}
func TestLogoutAcceptsEmpty204(t *testing.T) {
	cfg := NewZSConfig("cloud.invalid", 8080, "zstack").LoginAccount("admin", "password").Debug(false)
	cli := &ZSClient{ZSHttpClient: &ZSHttpClient{ZSConfig: cfg, httpClient: &http.Client{Transport: authTransport(func(r *http.Request) (*http.Response, error) {
		return &http.Response{StatusCode: 204, Header: make(http.Header), Body: io.NopCloser(strings.NewReader("")), Request: r}, nil
	})}}}
	cli.LoadSession("session")
	if err := cli.Logout(context.Background()); err != nil {
		t.Fatal(err)
	}
	if cli.sessionId != "" {
		t.Fatal("session not cleared")
	}
}
