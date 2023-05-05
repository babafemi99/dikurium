package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"net/http"
	"os"
	"test-dikurium/Configuration"
	"test-dikurium/Cryptography"
	"test-dikurium/Repository"
	"test-dikurium/Token"
	"test-dikurium/graph"
	"test-dikurium/middlewares"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

func main() {
	logger := initLogger()
	defer func() {
		_ = logger.Sync()
	}()

	// Load Configurations
	config, err := Configuration.NewConfig(".")
	if err != nil {
		logger.Errorw("failed to load config file", "error", err)
		os.Exit(1)
	}

	logger.Infow("Configuration file loaded successfully", "port", config.DBPassword)

	port := config.Port
	if port == "" {
		port = "7500"
	}

	dbName := config.DBName
	if dbName == "" {
		dbName = "mydb"
	}

	dbHost := config.DBHost
	if dbHost == "" {
		log.Println("here")
		dbHost = "localhost"
	}

	dbPassword := config.DBPassword
	if dbPassword == "" {
		dbPassword = "password"
	}

	dbPort := config.DBPort
	if dbPort == "" {
		dbPort = "5432"
	}

	dbUser := config.DBUser
	if dbUser == "" {
		dbUser = "postgres"
	}

	secretKey := config.JWTSecretKey
	if secretKey == "" {
		secretKey = "mysmallsecret"
	}

	// Load Database
	log.Println(dbHost, dbPort, dbName, dbPassword, secretKey)

	// create Repository Service
	repo := Repository.NewGormRepo(dbUser, dbPassword, dbHost, dbPort, dbName)
	logger.Info("Database connected successfully")

	// create Token Service
	tokenSrv := Token.NewJWTMaker(secretKey)

	//cryptography Service
	cryptoSrv := Cryptography.NewCryptoSrv()

	// create middleware service
	authMiddlewares := middlewares.NewAuthMiddlewares(tokenSrv)

	router := chi.NewRouter()
	router.Use(authMiddlewares.Authenticate)
	router.Use(middleware.Recoverer)

	c := graph.Config{Resolvers: &graph.Resolver{
		Repo:          repo,
		CryptoService: cryptoSrv,
		TokenService:  tokenSrv,
		Logger:        logger,
	}}
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(c))

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", graph.DataloaderMiddleware(repo.Client, srv))

	logger.Info("connect to http://localhost:%s/ for GraphQL playground", port)
	logger.Fatal(http.ListenAndServe(":"+port, router))
}

func initLogger() *zap.SugaredLogger {

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoder := zapcore.NewJSONEncoder(encoderConfig)

	core := zapcore.NewCore(
		encoder,
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout)),
		zap.NewAtomicLevelAt(zap.InfoLevel),
	)

	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1), zap.Development(), zap.WrapCore(func(core zapcore.Core) zapcore.Core {
		return zapcore.NewSamplerWithOptions(core, time.Second, 10, 100)
	}))

	sugarLogger := logger.Sugar()
	return sugarLogger

}
