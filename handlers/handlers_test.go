package handlers

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type testRepo struct {
	DB *gorm.DB
}

type postData struct {
	key   string
	value string
}

var getContentTest = []struct {
	Title      string
	Content    string
	Comment    string
	url        string
	statuscode int
}{
	{
		Title:      "Get Content",
		Content:    "this is the content",
		Comment:    "nice post",
		url:        "/get-all-contents",
		statuscode: http.StatusAccepted,
	},
	{
		Title:      "Unit Test",
		Content:    "Golang test helps alot",
		Comment:    "this is a cool post",
		url:        "/get-all-contents",
		statuscode: http.StatusAccepted,
	},
}

func getCtx(req *http.Request) context.Context

func Test_GetContents(t *testing.T)  {
	for _, e := range getContentTest {
		req, _ := http.NewRequest("GET", e.url, nil)
		ctx := getCtx(req)
		req = req.WithContext(ctx)
		req.RequestURI = e.url //uri encodes the data being sent using http forms 

		rt := httptest.NewRecorder()
		handler := http.HandlerFunc(Repo.GetContent())
	}}

