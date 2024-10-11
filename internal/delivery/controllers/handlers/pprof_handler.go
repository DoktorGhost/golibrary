package handlers

import (
	"net/http"
	"net/http/pprof"
)

// PprofHandler обрабатывает запросы для профилирования
// @Summary Профилирование приложения
// @Description Возвращает информацию о профилировании для приложения
// @Tags Pprof
// @Accept json
// @Produce json
// @Router /debug/pprof/ [get]
// @Security BearerAuth
func PprofHandler(w http.ResponseWriter, r *http.Request) {
	http.HandlerFunc(pprof.Index)(w, r)

}
