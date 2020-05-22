package onesignal

var app AppCreate

//INotification IPush
type INotification interface {
	getMessage() string
	getTitle() string
	getID() [][]string
}

//Notification Notification
type Notification struct {
	Message string
	Title   string
	ID      [][]string
}

func (a Notification) getMessage() string {
	return a.Message
}
func (a Notification) getTitle() string {
	return a.Title
}
func (a Notification) getID() [][]string {
	return a.ID
}

//ForAllUsers ForAllUsers
type ForAllUsers struct {
	Notification
}

//ForUser ForUser
type ForUser struct {
	Notification
}

func (a ForUser) newPushForUsers() INotification {
	return &ForUser{
		Notification: Notification{
			Message: a.Message,
			Title:   a.Title,
			ID:      a.ID,
		},
	}
}

func (a ForAllUsers) newPushForAllUsers() INotification {
	var all [][]string
	var s []string
	s = append(s, "All")
	all = append(all, s)
	return &ForAllUsers{
		Notification: Notification{
			Message: a.Message,
			Title:   a.Title,
			ID:      all,
		},
	}
}

//NotificationFactory NotificationFactory
func NotificationFactory(a INotification) Notification {
	return Notification{
		Message: a.getMessage(),
		Title:   a.getTitle(),
		ID:      a.getID(),
	}
}
func createNotification(msg string, title string, id ...[]string) error {
	if len(id) == 0 {
		a := ForAllUsers{Notification{Message: msg, Title: title}}
		adapte := NotificationForAllUsers{}
		notiAdapter := NotificationForAllUsersAdapter{
			users: adapte,
		}
		all := a.newPushForAllUsers()
		n := NotificationFactory(all)
		err := Push(notiAdapter, n, app)
		return err
	}
	b := ForUser{Notification{ID: id, Title: title, Message: msg}}

	user1 := b.newPushForUsers()
	ns := NotificationFactory(user1)
	den := NotificationForUser{}
	err := Push(den, ns, app)
	return err

}