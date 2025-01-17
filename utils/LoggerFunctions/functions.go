package LoggerFunctions

import (
	"os"

	"github.com/sirupsen/logrus"
)

var Log *logrus.Logger

func init() {
	Log = logrus.New()

	// Setting up the log file
	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		logrus.Fatalf("Failed to log to file, using default stderr: %v", err)
	}
	Log.Out = file
	Log.Formatter = &logrus.JSONFormatter{} // Log in JSON format for easier parsing
}
