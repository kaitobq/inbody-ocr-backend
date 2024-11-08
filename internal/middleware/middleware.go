package middleware

import (
	"inbody-ocr-backend/internal/domain/repository"
	"inbody-ocr-backend/internal/domain/service"
)

type Middleware struct {
	API       *API
	CORS      *CORS
	UserAgent *UserAgent
	Logging   *Logging
}

func NewMiddleware(tokenService service.TokenService, userRepo repository.UserRepository) *Middleware {
	return &Middleware{
		API:       NewAPI(tokenService, userRepo),
		CORS:      NewCORS(),
		UserAgent: NewUserAgent(),
		Logging:   NewLogging(),
	}
}
