package logger

import (
	"log"
	"os"
	"time"
	"io"
)

var customLogger *log.Logger
func SetupLogger() {
	location, err := time.LoadLocation("America/Sao_Paulo")
	if err != nil {
		log.Fatalf("Erro ao carregar localização: %v", err)
	}

	logFile, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("Erro ao abrir ou criar o arquivo de log: %v", err)
	}

	// Configurar o logger para escrever tanto no arquivo quanto no terminal
	multiWriter := io.MultiWriter(os.Stdout, logFile)
	customLogger = log.New(multiWriter, "LOG: ", 0)
	customLogger.SetFlags(0)
	// substituir log local
	customLogger.SetOutput(log.Writer())
	customLogger.Println(time.Now().In(location).Format("2006-01-02 15:04:05"), "Logger configurado com sucesso.")
}

// Info  
func Info(message string) {
	location, _ := time.LoadLocation("America/Sao_Paulo")
	customLogger.Printf("[%s] INFO: %s\n", time.Now().In(location).Format("2006-01-02 15:04:05"), message)
}

// Error 
func Error(err error) {
	location, _ := time.LoadLocation("America/Sao_Paulo")
	customLogger.Printf("[%s] ERROR: %s\n", time.Now().In(location).Format("2006-01-02 15:04:05"), err)
}
