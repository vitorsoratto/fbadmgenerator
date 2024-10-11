package config

var (
	logger *Logger
)

func GetLogger() *Logger {
	logger = NewLogger()
	return logger
}
