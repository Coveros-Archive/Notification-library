# Notification-library
Generic notification library for sending messages to notification providers

## Supported providers:
- [x] Slack
  - Example usage:
  ```
  import "cNotifyLib "github.com/Coveros/notification-library"
  
  func main() {
    slackNotifier := cNotifyLib.NewNotificationProvider(cNotifyLib.Slack, os.Getenv("SLACK_BOT_TOKEN"))
    err := slackNotifier.SendMsg(cNotifyLib.NotifyTemplate{
			Channel:   "SLACK_CHANNEL_ID",
			Title:     "Hello world!",
			Text:      "This is a text",
			EventType: cNotifyLib.Warning,
      Fields: map[string]string{ "field1": "value" }
		})
  }
  ```
- [ ] Discord
- [ ] Hipchat
- [ ] Microsoft teams
