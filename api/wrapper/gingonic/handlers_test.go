package gingonic

import (
	"github.com/gin-gonic/gin"
	"github.com/zibilal/simpleapi/api/v3/handler"
	"io/ioutil"
	"net/http/httptest"
	"testing"
)

const (
	failed  = "\u2717"
	success = "\u2713"
)

func TestPingApi(t *testing.T) {
	t.Log("Testing PingAPI handler")
	{
		w := httptest.NewRecorder()

		ctx, _ := gin.CreateTestContext(w)
		apiCtx := WrapGinContext(ctx)

		err := handler.PingApi(apiCtx)
		if err != nil {
			t.Fatalf("%s expected error nil, got %s", failed, err.Error())
		}

		ctx, found := apiCtx.UnwrapContext().(*gin.Context)

		if !found {
			t.Fatalf("%s expected type is *gin.Context", failed)
		}

		b, err := ioutil.ReadAll(w.Body)
		if err != nil {
			t.Fatalf("%s expected body read successfully, %s", failed, err.Error())
		}

		expectedResult := `{"code":1000,"message":"Without Middleware"}`
		if expectedResult == string(b) {
			t.Logf("%s expected result = [%s]", success, expectedResult)
		} else {
			t.Fatalf("%s expected result = [%s], got [%s]", failed, expectedResult, string(b))
		}
	}
}

func TestMiddlewareTest(t *testing.T) {
	t.Log("Testing MiddlewareTest handler")
	{
		w := httptest.NewRecorder()

		ctx, _ := gin.CreateTestContext(w)
		aCtx := WrapGinContext(ctx)
		err := handler.MiddlewareTest(aCtx)
		if err != nil {
			t.Fatalf("%s expected error nil, got %s", failed, err.Error())
		}

		data := aCtx.Get("MID")
		if data == nil {
			t.Fatalf("%s expected data not nil", failed)
		}

		if data.(string) == "here" {
			t.Logf("%s expected data == here", success)
		}
	}
}
