package basicchi

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
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
	ctx := r.Context()
	w.Write(([]byte)("ReqID:\t"))
	w.Write(([]byte)(ctx.Value(middleware.RequestIDKey).(string)))
	w.Write(([]byte)("\nHello Template\n"))
}
func (sis *ServeInterState) general(w http.ResponseWriter, r *http.Request) {
	w.Write(([]byte)("Hello From General\n"))
}

/*
http://127.0.0.1/__ELSE__
*/
func (sis *ServeInterState) serve__ELSE__(w http.ResponseWriter, r *http.Request) {
	log.Println("else")
	//	time.Sleep(time.Second * 2)
	//w.WriteHeader(http.StatusFound)
	//	w.Write(([]byte)("Hello From Else. You 'll soon be redirected\n"))
	http.Redirect(w, r, "/general", http.StatusFound)

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
	log.Println("panicMiddleware init")
	fn := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("recovered", err)
				http.Error(w, "Internal server error", 500)
			}
		}()
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
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

	chiRouter := chi.NewRouter()
	sis := NewServeInterState(ctx)
	chiRouter.Get("/", sis.serveRoot)
	chiRouter.Get("/else", sis.serve__ELSE__)
	mdw := loggingMiddleware(chiRouter)
	mdw = panicMiddleware(mdw)
	/* === other implementation of panic middleware === */
	//beforeMdwChiRouter.Use(middleware.Recoverer)
	beforeMdwChiRouter := chi.NewRouter()
	beforeMdwSIS := NewServeInterState(context.TODO())
	beforeMdwChiRouter.Use(middleware.RequestID)
	beforeMdwChiRouter.Handle("/", mdw)
	beforeMdwChiRouter.Handle("/else", mdw)
	beforeMdwChiRouter.HandleFunc("/general", beforeMdwSIS.general)
	/*  ===  cannot use after __ANY__ .Handle === */
	//beforeMdwChiRouter.Use(middleware.RequestID)

	server := http.Server{

		Addr:              addr,
		Handler:           beforeMdwChiRouter,
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
