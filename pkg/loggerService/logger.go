package loggerService

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"path/filepath"
	"time"
)

var Log *logrus.Logger

func InitLogger(level string, outputPath string) error {
	Log = logrus.New()

	// Log formatını JSON olarak ayarlayın (isteğe bağlı)
	Log.SetFormatter(&logrus.JSONFormatter{})

	// Günlük dosyası adı ve rotasını belirleyin (örn. ./logs/YYYY-MM-DD.log)
	logDir := outputPath
	logFileName := time.Now().Format("2006-01-02") + ".log"
	logFilePath := filepath.Join(logDir, logFileName)

	// Eğer dizin yoksa, önceden oluştur
	if err := os.MkdirAll(logDir, os.ModePerm); err != nil {
		fmt.Println("Log dizinini oluştururken bir hata oluştu: ", err)
		return err
	}

	// Eğer dosya yoksa, önceden oluştur
	if _, err := os.Stat(logFilePath); os.IsNotExist(err) {
		file, err := os.Create(logFilePath)
		if err != nil {
			fmt.Println("Log dosyasını oluştururken bir hata oluştu: ", err)
			return err
		}
		defer file.Close()
	}

	file, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return err
	}

	Log.SetOutput(file)

	// Logger seviyesini ayarla
	lvl, err := logrus.ParseLevel(level)
	if err != nil {
		return err
	}
	Log.SetLevel(lvl)

	return nil
}
