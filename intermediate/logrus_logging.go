package intermediate

import "github.com/sirupsen/logrus"

func main() {

	log := logrus.New()

	// set log level
	log.SetLevel(logrus.InfoLevel)

	// set log format
	log.SetFormatter(&logrus.JSONFormatter{})

	// logging examples
	log.Info("This is an info message")
	log.Warn("This is a warning message")
	log.Error("This is an error message")

	log.WithFields(logrus.Fields{
		"username": "johndoe",
		"method": "GET",
	}).Info("User logged in.")

}