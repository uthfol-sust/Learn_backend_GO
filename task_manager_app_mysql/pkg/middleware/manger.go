package middleware

import "net/http"

type Middleware func(http.Handler) http.Handler

type Manger struct {
	gobalMiddlewares []Middleware
}

func (mngr *Manger) Use(middlewares ...Middleware) {
	mngr.gobalMiddlewares = append(mngr.gobalMiddlewares, middlewares...)
}

func (mngr *Manger) With(next http.Handler, middlewares ...Middleware) http.Handler {
	n := next

	for _, middleware := range middlewares {
		n = middleware(n)
	}

	for _, gobal := range mngr.gobalMiddlewares {
		n = gobal(n)
	}

	return n
}
