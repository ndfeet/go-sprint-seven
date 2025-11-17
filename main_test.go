package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMainHandlerAllScenarios(t *testing.T) {
	handler := http.HandlerFunc(mainHandle)

	// Сценарий 1: корректный запрос
	req1 := httptest.NewRequest(http.MethodGet, "/cafe?count=2&city=moscow", nil)
	rec1 := httptest.NewRecorder()
	handler.ServeHTTP(rec1, req1)
	assert.Equal(t, http.StatusOK, rec1.Code)
	assert.NotEmpty(t, rec1.Body.String())

	// Сценарий 2: неверный город
	req2 := httptest.NewRequest(http.MethodGet, "/cafe?count=3&city=london", nil)
	rec2 := httptest.NewRecorder()
	handler.ServeHTTP(rec2, req2)
	require.Equal(t, http.StatusBadRequest, rec2.Code)
	assert.Equal(t, "wrong city value", rec2.Body.String())

	// Сценарий 3: count больше количества кафе
	req3 := httptest.NewRequest(http.MethodGet, "/cafe?count=999&city=moscow", nil)
	rec3 := httptest.NewRecorder()
	handler.ServeHTTP(rec3, req3)
	require.Equal(t, http.StatusOK, rec3.Code)
	cafes := strings.Split(rec3.Body.String(), ",")
	assert.Len(t, cafes, 4)
}
