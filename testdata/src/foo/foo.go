package foo

import (
	"log/slog"
	"os"

	"github.com/Vladroon22/linters-check/testdata/src/go.uber.org/zap"
)

func main() {
	var (
		token      = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9"
		password   = "secret123"
		apiKey     = "sk-1234567890abcdef"
		creditCard = "4111-1111-1111-1111"
		authHeader = "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9"
	)

	zapper := zap.L()

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	// ============ ТЕСТ 1: РЕГИСТР ПЕРВОГО СИМВОЛА ============

	zapper.Info("starting server on port")
	zapper.Info("initializing database connection")
	slog.Info("processing incoming request")
	slog.Info("user authenticated successfully")
	slog.Debug("cache miss for key user123")
	slog.Warn("connection pool exhausted")
	slog.Error("failed to connect to database")

	zapper.Info("Starting server on port 8080")
	zapper.Info("Server started successfully")
	slog.Info("Failed to connect to database")
	slog.Info("User logged in")
	slog.Warn("Connection timeout")
	slog.Error("Database connection lost")
	slog.Info("Processing completed")

	slog.Info("12345 is the port number")
	slog.Info("starting with dash")
	slog.Info("(parenthesized message)")
	slog.Info("[INFO] server started")

	// ============ ТЕСТ 2: ТОЛЬКО АНГЛИЙСКИЙ ЯЗЫК ============

	zapper.Info("server started successfully")
	zapper.Info("user authentication failed")
	slog.Info("cache miss for key: user_123")
	slog.Info("request processed in 150ms")
	slog.Info("database query executed")

	slog.Info("запуск сервера")
	slog.Info("ошибка подключения к базе данных")
	zapper.Info("сервер запущен")
	slog.Info("服务器启动成功")
	slog.Info("サーバーが起動しました")
	slog.Info("서버가 시작되었습니다")
	slog.Info("serveur démarré")
	slog.Info("servidor iniciado")
	slog.Info("Server gestartet")

	slog.Info("server запущен successfully")
	slog.Info("ошибка connection")
	zapper.Info("error подключения")

	// ============ ТЕСТ 3: СПЕЦСИМВОЛЫ И ЭМОДЗИ ============

	zapper.Info("server started successfully")
	slog.Info("request processed in 150ms")
	slog.Info("user id 12345 logged in")

	slog.Info("server started!!!")
	slog.Info("connection failed???")
	zapper.Info("server started...")
	slog.Info("===== server started =====")
	slog.Info("*** important ***")
	slog.Info(">>>> warning <<<<")
	slog.Info("server [started]")
	slog.Info("server {started}")
	slog.Info("server (started)!!!")
	slog.Info("path: /api/v1/users")
	slog.Info("duration: 1m30s")
	slog.Info("progress: 75%")

	slog.Info("server started 🚀")
	slog.Info("connection failed 😢")
	zapper.Info("success! 🎉")
	slog.Info("warning ⚠️")
	slog.Info("task completed ✅")
	slog.Info("error ❌")
	slog.Info("heart ❤️")
	slog.Info("thumbs up 👍")
	slog.Info("fire 🔥")
	slog.Info("rocket launch 🚀🔥")
	slog.Info("👨‍👩‍👧‍👦 family")
	slog.Info("⌛ timeout")
	slog.Info("⭐ priority")

	// ============ ТЕСТ 4: ЧУВСТВИТЕЛЬНЫЕ ДАННЫЕ ============

	zapper.Info("user authenticated successfully")
	slog.Info("database connection established")
	slog.Info("request completed in 250ms")
	slog.Info("cache updated for user 12345")

	slog.Info("user password: " + password)
	slog.Info("password=" + password)
	slog.Info("passwd: " + password)
	zapper.Info("pwd: " + password)
	slog.Info("Password: " + password)

	slog.Info("api_key=" + apiKey)
	slog.Info("apikey: " + apiKey)
	zapper.Info("API Key: " + apiKey)
	slog.Info("api key: " + apiKey)

	slog.Info("token: " + token)
	slog.Info("jwt: " + token)
	zapper.Info("bearer " + token)
	slog.Info("access_token: " + token)
	slog.Info("refresh_token: " + token)
	slog.Info("auth: " + authHeader)

	slog.Info("secret: mysecretvalue")
	slog.Info("private key: somekey")
	zapper.Info("client_secret: secret123")
	slog.Info("credentials: user:pass")

	zapper.Debug("debug message with password: " + password)
	zapper.Info("info message with token: " + token)
	zapper.Warn("warn message with api_key: " + apiKey)
	zapper.Error("error message with credit card: " + creditCard)
}
