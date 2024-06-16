package mail

import (
	_ "github.com/bmerchant22/hc_hackathon/config"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var (
	user    string
	pass    string
	host    string
	port    string
	sender  string
)

func init() {
	logrus.Info("Initializing mailer")

	user = viper.GetString("MAIL.USER")
	sender = user + "@iitk.ac.in"

	pass = viper.GetString("MAIL.PASS")
	host = viper.GetString("MAIL.HOST")
	port = viper.GetString("MAIL.PORT")
}
