package basichttp

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"time"
)

type Starter struct {
	Address string
}

func NewStarter(ctx context.Context) Starter {
	starter := Starter{}
	ctxv := ctx.Value("Address")
	addr, ok := ctxv.(string)
	if ok {
		starter.Address = addr
	}
	return starter
}

type ServeInterState struct {
}

func NewServeInterState(ctx context.Context) *ServeInterState {
	fromCtx := ctx.Value("__INTERSTATE_KEY__")
	if fromCtx != nil {
		/* logic */
	}
	return &ServeInterState{}
}

/*
http://127.0.0.1/
*/
func (sis *ServeInterState) serveRoot(w http.ResponseWriter, r *http.Request) {
	w.Write(([]byte)("Hello Template\n"))
}
func (sis *ServeInterState) general(w http.ResponseWriter, r *http.Request) {

}

/*
http://127.0.0.1/__ELSE__
*/
func (sis *ServeInterState) serve__ELSE__(w http.ResponseWriter, r *http.Request) {

}
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		/* logging */
		/*
			ctx := r.Context()
			next.ServeHTTP(w,r.WithContext(ctx))
		*/
		next.ServeHTTP(w, r)
	})

}
func panicMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("recovered", err)
				http.Error(w, "Internal server error", 500)
			}
		}()
		next.ServeHTTP(w, r)
	})
}

// no Request , Reply, error -- to start as goroutine
func (starter *Starter) StartServer(ctx context.Context) {
	addr := starter.Address
	if addr == "" {
		addr, _ = ctx.Value("__ADDRESS__").(string)
		if addr == "" {
			log.Fatal("Cannot get addres from context:\t __ADDRESS__ ")
		}
		log.Fatal("Cannot get addres:\t __ADDRESS__ ")
	}
	stdlibMux := http.NewServeMux()
	sis := NewServeInterState(ctx)
	stdlibMux.HandleFunc("/", sis.serveRoot)
	stdlibMux.HandleFunc("/__ELSE__", sis.serve__ELSE__)
	mdw := loggingMiddleware(stdlibMux)
	mdw = panicMiddleware(mdw)
	generalMux := http.NewServeMux()
	generalMux.Handle("/", mdw)

	server := http.Server{

		Addr:              addr,
		Handler:           generalMux,
		ReadTimeout:       10 * time.Second,
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       10 * time.Second,
		BaseContext:       func(l net.Listener) context.Context { return ctx },

		/*
		   ConnContext: func(ctx context.Context, c net.Conn) context.Context {
		   			// optional
		   			return context.WithValue(ctx, __KEY__, __VALUE__)
		   		},
		*/
	}
	log.Println("Start Serving at:", addr)
	go func() {
		<-ctx.Done()
		server.Shutdown(ctx)
		log.Println("Shutdown by context.Done call")
	}()
	server.ListenAndServe()
	log.Println("Done Serving")

}
