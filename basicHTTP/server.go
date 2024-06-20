package basichttp

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

// Константа для хранения адреса в контексте
const (
	addressKey = "Address"
)

// Структура Starter для хранения адреса
type Starter struct {
	Address string
}

// Функция создания нового Starter из контекста
func NewStarter(ctx context.Context) (*Starter, error) {
	// Получение адреса из контекста
	addr, ok := ctx.Value(addressKey).(string)
	if !ok {
		// Возврат ошибки, если адрес не найден в контексте
		return nil, fmt.Errorf("address not found in context")
	}
	return &Starter{Address: addr}, nil
}

// Структура ServeInterState для обработки запросов
type ServeInterState struct{}

// Функция создания нового ServeInterState
func NewServeInterState() *ServeInterState {
	return &ServeInterState{}
}

// Обработка корневого запроса
func (sis *ServeInterState) serveRoot(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Template\n"))
}

// Обработка общего запроса
func (sis *ServeInterState) general(w http.ResponseWriter, r *http.Request) {

}

// Middleware для логирования запросов
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Логирование запроса
		logrus.Infof("Request: %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

// Middleware для перехвата паники
func panicMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			// Перехват паники и логирование ошибки
			if err := recover(); err != nil {
				logrus.Errorf("Recovered from panic: %v", err)
				http.Error(w, "Internal server error", 500)
			}
		}()
		next.ServeHTTP(w, r)
	})
}

// Структура ServerConfig для конфигурации сервера
type ServerConfig struct {
	Address           string
	ReadTimeout       time.Duration
	ReadHeaderTimeout time.Duration
	WriteTimeout      time.Duration
	IdleTimeout       time.Duration
}

// Функция запуска сервера
func (starter *Starter) StartServer(ctx context.Context, config ServerConfig) {
	// Создание ServeMux для обработки запросов
	stdlibMux := http.NewServeMux()
	sis := NewServeInterState()
	stdlibMux.HandleFunc("/", sis.serveRoot)
	stdlibMux.HandleFunc("/__ELSE__", sis.serve__ELSE__)
	mdw := loggingMiddleware(stdlibMux)
	mdw = panicMiddleware(mdw)
	generalMux := http.NewServeMux()
	generalMux.Handle("/", mdw)

	// Создание сервера
	server := http.Server{
		Addr:              config.Address,
		Handler:           generalMux,
		ReadTimeout:       config.ReadTimeout,
		ReadHeaderTimeout: config.ReadHeaderTimeout,
		WriteTimeout:      config.WriteTimeout,
		IdleTimeout:       config.IdleTimeout,
		BaseContext:       func(l net.Listener) context.Context { return ctx },
	}

	// Запуск сервера в отдельной горутине
	go func() {
		<-ctx.Done()
		server.Shutdown(ctx)
		logrus.Info("Shutdown by context.Done call")
	}()

	// Логирование запуска сервера
	logrus.Infof("Start Serving at: %s", config.Address)
	server.ListenAndServe()
	logrus.Info("Done Serving")
}
