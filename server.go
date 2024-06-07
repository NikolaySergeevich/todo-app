package todoapp

import (
	"context"
	"net/http"
	"time"
)

// Эта структура является абстракцией над структурой srver из пакета http.
// Возвращает метод стандартного http сервера ListenAndServe, который под копотом запускает бесконечный цикл for
// и слушает все входящие запросы для последующей обработки
type Server struct {
	httpServer *http.Server
}

// Запуск сервера
func (s *Server) Run(port string, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler: handler,
		MaxHeaderBytes: 1 << 20, // 1MB
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}
	return s.httpServer.ListenAndServe()
}

// Остановка работы
func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
