package notification_library

type NotificationProvider string
type NotifyEventType string
var (
	Failure NotifyEventType = "failure"
	Success NotifyEventType = "success"
	Warning NotifyEventType = "warning"
	Slack NotificationProvider = "slack"
	Noop NotificationProvider = "noop"
	//TODO: add more as needed
)


type NotifyTemplate struct {
	Channel string
	Title string
	Text string
	EventType NotifyEventType
	Fields map[string]string
}

type Notify interface {
	SendMsg(n NotifyTemplate) error
}

func NewNotificationProvider(provider NotificationProvider, token string) Notify {
	switch provider {
	case Slack:
		return newSlack(token)
	default:
		return newNoop()
	}
	return nil
}
