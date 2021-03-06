package weibo

import (
	"os"
	"testing"
)

func TestStatusesMentions(t *testing.T) {
	appkey := os.Getenv("weibo_app_key")
	appsecret := os.Getenv("weibo_app_secret")
	username := os.Getenv("weibo_username")
	passwd := os.Getenv("weibo_passwd")
	redirecturi := os.Getenv("weibo_redirect_uri")
	weibo := New(appkey, appsecret, username, passwd, redirecturi)
	t.Log("PCLogin...")
	if err := weibo.PCLogin(); err != nil {
		t.Fatal(err)
	}

	t.Log("Authorize")
	code, err := weibo.Authorize()
	if err != nil {
		t.Fatal(err)
	}
	t.Log("AccessToken")
	token, err := weibo.AccessToken(code)
	if err != nil {
		t.Fatal(err)
	}
	sinceID := int64(0)
	maxID := int64(0)
	count := 50
	page := 1
	filterBySource := 0
	filterByAuthor := 0
	filterByType := 0
	resp, err := weibo.StatusesMentions(token.AccessToken, sinceID, maxID, count, page, filterBySource, filterByAuthor, filterByType)
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v", resp)
}
