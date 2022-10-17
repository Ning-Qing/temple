package test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Ning-Qing/temple/config"
	"github.com/Ning-Qing/temple/global"
	"github.com/Ning-Qing/temple/router"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
)

func MockServer(configPath string) *gin.Engine {
	global.GlobalSettings = config.InitConfig(configPath)
	return router.InitRouter()
}

func TestPing(t *testing.T) {
	router := MockServer("../config.yaml")

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "pong", w.Body.String())
}
