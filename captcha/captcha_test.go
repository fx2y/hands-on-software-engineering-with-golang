package captcha_test

import (
	"github.com/fx2y/hands-on-software-engineering-with-golang/captcha"
	"image"
	"testing"
)

func TestChallengeUserSuccess(t *testing.T) {
	got := captcha.ChallengeUser(stubChallenger("42"), stubPrompter("42"))
	if got != true {
		t.Fatal("expected ChallengeUser to return true")
	}
}

type stubChallenger string

func (c stubChallenger) Challenge() (image.Image, string) {
	return image.NewRGBA(image.Rect(0, 0, 100, 100)), string(c)
}

type stubPrompter string

func (p stubPrompter) Prompt(_ image.Image) string {
	return string(p)
}
