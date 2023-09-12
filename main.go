package main

import (
	"context"
	"fxline/other"
	"fxline/route"
	"net"
	"net/http"

	"go.uber.org/fx"
	// "go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

func main() {
	fx.New(
		// replace fx default log
		// fx.WithLogger(func(log *zap.Logger) fxevent.Logger {
		// 	return &fxevent.ZapLogger{Logger: log}
		// }),
		fx.Provide(
			other.NewOtherOne,
			other.NewOtherTwo,

			NewHTTPServer,
			// handler.NewEchoHandler,
			// fx.Annotate(
			// 	route.NewEchoHandler,
			// 	fx.As(new(route.Route)),
			// 	fx.ResultTags(`name:"echo"`),
			// ),
			// fx.Annotate(
			// 	route.NewHelloHandler,
			// 	fx.As(new(route.Route)),
			// 	fx.ResultTags(`name:"hello"`),
			// ),

			fx.Annotate(
				route.NewEchoHandler,
				fx.As(new(route.Route)),
				fx.ResultTags(`group:"routes"`),
			),
			AsRoute(route.NewHelloHandler),

			// handler.NewServeMux,
			// route.NewServeMux,
			fx.Annotate(
				route.NewServeMux,
				// fx.ParamTags(`name:"echo"`, `name:"hello"`),
				fx.ParamTags(`group:"routes"`),
			),
			zap.NewExample,
		),
		// We used fx.Invoke to request that the HTTP server is always instantiated, even if none of the other components in the application reference it directly.
		fx.Invoke(func(*http.Server) {}, other.Hello2, other.Hello1, other.EchoName),
	).Run()
}

func NewHTTPServer(lc fx.Lifecycle, mux *http.ServeMux, log *zap.Logger) *http.Server {
	srv := &http.Server{Addr: ":8080", Handler: mux}
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			ln, err := net.Listen("tcp", srv.Addr)
			if err != nil {
				return err
			}
			log.Info("Starting HTTP server", zap.String("addr", srv.Addr))
			go srv.Serve(ln)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return srv.Shutdown(ctx)
		},
	})
	return srv
}

func AsRoute(f any) any {
	return fx.Annotate(
		f,
		fx.As(new(route.Route)),
		fx.ResultTags(`group:"routes"`),
	)
}
