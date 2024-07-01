package basichttp

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/sirupsen/logrus" // можно попробовать использовать логрус для логирования
)

// константа для хранения адреса
const (
	addressKey = "Address"
)

// структура для хранения адреса
type Starter struct {
	Address string
}

// функция создания нового "стартера" из контекста
func NewStarter(ctx context.Context) (*Starter, error) {
	// получение адреса
	addr, ok := ctx.Value(addressKey).(string)
	if !ok {
		// отлавливаем ошибку
		return nil, fmt.Errorf("address not found in context")
	}
	return &Starter{Address: addr}, nil
}

// структура для обработки запросов
type ServeInterState struct {
	//...
}

// функция создания нового запроса
func NewServeInterState() *ServeInterState {
	return &ServeInterState{}
}

// обработка корневого запроса
func (sis *ServeInterState) serveRoot(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Template\n"))
}

// обработка генерал запроса
func (sis *ServeInterState) general(w http.ResponseWriter, r *http.Request) {

}

// мидлы для логирования запросов
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// логирование
		logrus.Infof("Request: %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

// мидлы для отвода паники
func panicMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			// отводим панику
			if err := recover(); err != nil {
				logrus.Errorf("Recovered from panic: %v", err)
				http.Error(w, "Internal server error", 500)
			}
		}()
		next.ServeHTTP(w, r)
	})
}

// структура для конфига сервера
type ServerConfig struct {
	Address           string
	ReadTimeout       time.Duration
	ReadHeaderTimeout time.Duration
	WriteTimeout      time.Duration
	IdleTimeout       time.Duration
}

// запуск сервера
func (starter *Starter) StartServer(ctx context.Context, config ServerConfig) {
	// обработка запросов
	stdlibMux := http.NewServeMux()
	sis := NewServeInterState()
	stdlibMux.HandleFunc("/", sis.serveRoot)
	stdlibMux.HandleFunc("/__ELSE__", sis.serve__ELSE__)
	mdw := loggingMiddleware(stdlibMux)
	mdw = panicMiddleware(mdw)
	generalMux := http.NewServeMux()
	generalMux.Handle("/", mdw)

	// создание сервера
	server := http.Server{
		Addr:              config.Address,
		Handler:           generalMux,
		ReadTimeout:       config.ReadTimeout,
		ReadHeaderTimeout: config.ReadHeaderTimeout,
		WriteTimeout:      config.WriteTimeout,
		IdleTimeout:       config.IdleTimeout,
		BaseContext:       func(l net.Listener) context.Context { return ctx },
	}

	// запуск(в горутине ,если можно по другому то окэй)
	go func() {
		<-ctx.Done()
		server.Shutdown(ctx)
		logrus.Info("Shutdown by context.Done call")
	}()

	// логи для запуска
	logrus.Infof("Start Serving at: %s", config.Address)
	server.ListenAndServe()
	logrus.Info("Done Serving")
}
