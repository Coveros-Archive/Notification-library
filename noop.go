package notification_library

type noop struct {}

func newNoop() *noop {
	return &noop{}
}

func (no *noop) SendMsg(n NotifyTemplate) error {
	return nil
}