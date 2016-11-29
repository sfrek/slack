package slack

import (
	"fmt"
	"log/syslog"
	"os"
	"strings"

	"github.com/sfrek/slack/rest"
	"github.com/sfrek/slack/rtm"
)

var (
	slacklog, _ = syslog.New(syslog.LOG_DAEMON, os.Args[0]+".slack")
	channel     = "#sfrektest"
	user        = "services"
)

type Slack struct {
	RestClient *rest.Client
	RtmClient  *rtm.Client
}

func NewClient(token string) *Slack {
	return &Slack{RestClient: rest.NewClient(token), RtmClient: rtm.NewClient(token)}
}

func (s *Slack) Info() string {
	return fmt.Sprintf("DUMP: RestClient |\n%s\n|\n RtmClient |\n%s\n|", s.RestClient, s.RtmClient)
}

// ToDO:
//   channel and user must be parameters
func (s *Slack) DirectMessage(message string) {
	var toSend string
	hostname, _ := os.Hostname()
	if strings.HasPrefix(message, "<!channel>") {
		toSend = message + " from " + hostname
	} else {
		toSend = hostname + ": " + message
	}
	slacklog.Info(fmt.Sprintf("gonna send message [ %s ]", toSend))
	err := s.RestClient.SendMessage(channel, toSend, user)
	if err != nil {
		slacklog.Err("imposible to send message")
		os.Exit(1)
	}
}
