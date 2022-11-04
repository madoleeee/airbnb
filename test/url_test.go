package test

import (
	"net/url"
	"testing"
)

func TestUrlDecode(t *testing.T) {
	s := "email=temp%40naver.com&password=temp%21&remember_me=true&from=email_login&origin_url=https%3A%2F%2Fwww.airbnb.co.kr%2F%3Fhas_logged_out%3D1&page_controller_action_pair="

	s, err := url.QueryUnescape(s)
	if err != nil {
		t.Error(err.Error())
	}

	t.Log(s)

}
