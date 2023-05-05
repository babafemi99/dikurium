package graph

import (
	"go.uber.org/zap"
	"test-dikurium/Cryptography"
	"test-dikurium/Repository"
	"test-dikurium/Token"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Repo          *Repository.GormRepo
	CryptoService Cryptography.BCryptMaker
	TokenService  Token.JWTMaker
	Logger        *zap.SugaredLogger
	// todo add logging functionality
}
