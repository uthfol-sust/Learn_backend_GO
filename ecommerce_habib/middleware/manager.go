package middleware

import "net/http"

type Middlewares func(http.Handler) http.Handler

type Manager struct{
	gobalMiddleware []Middlewares
}


func NewManager() *Manager{
	return &Manager{
		gobalMiddleware: make([]Middlewares, 0),
	}
}

func (mngr *Manager) Use(Middlewares... Middlewares){
	mngr.gobalMiddleware =append(mngr.gobalMiddleware, Middlewares...)
}

func (mngr *Manager)With(next http.Handler ,middlewares... Middlewares ) http.Handler{
  n := next

  for _ , middleware := range mngr.gobalMiddleware{
     n = middleware(n)
  }

  for _ , middleware := range middlewares{
     n = middleware(n)
  }
  
  return n
}