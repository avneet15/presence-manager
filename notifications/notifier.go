package notifications

import "fmt"

type Notifier interface {
	Notify(entryId string)
}

type LoggerNotifier struct {
}

func (n LoggerNotifier) Notify(entryId string) {
	fmt.Println("Entry no longer exists:",entryId)
}