package lib

import (
	"container/list"
	"encoding/json"
	"time"

	"blog/models"

	"github.com/astaxie/beego"
	"github.com/gorilla/websocket"
)

//用户
type Subscriber struct {
	User models.SysOperator
	Conn *websocket.Conn // Only for WebSocket users; otherwise nil.
}

//房间
type ChatRoom struct {
	// Channel for new join users.
	Subscribe chan Subscriber
	// Channel for exit users.
	Unsubscribe chan models.SysOperator
	// Send events here to publish them.
	Publish chan Event
	// Long polling waiting list.
	WaitingList *list.List
	Subscribers *list.List
}

func NewChatroom() *ChatRoom {
	chatroom := &ChatRoom{
		Subscribe:   make(chan Subscriber, 10),
		Unsubscribe: make(chan models.SysOperator, 10),
		Publish:     make(chan Event, 10),
		WaitingList: list.New(),
		Subscribers: list.New(),
	}
	return chatroom
}
func (this *ChatRoom) isUserExist(subscribers *list.List, user models.SysOperator) bool {
	for sub := this.Subscribers.Front(); sub != nil; sub = sub.Next() {
		if sub.Value.(Subscriber).User == user {
			return true
		}
	}
	return false
}
func (this *ChatRoom) NewEvent(ep EventType, user models.SysOperator, msg string) Event {
	return Event{ep, user, int(time.Now().Unix()), msg}
}

func (this *ChatRoom) Join(user models.SysOperator, ws *websocket.Conn) {
	this.Subscribe <- Subscriber{User: user, Conn: ws}
}

func (this *ChatRoom) Leave(user models.SysOperator) {
	this.Unsubscribe <- user
}

func (this *ChatRoom) Start() {
	for {
		select {
		case sub := <-this.Subscribe:
			if !this.isUserExist(this.Subscribers, sub.User) {
				this.Subscribers.PushBack(sub) // Add user to the end of list.
				// Publish a JOIN event.
				this.Publish <- this.NewEvent(EVENT_JOIN, sub.User, "")
				beego.Info("New user:", sub.User.RealName, ";WebSocket:", sub.Conn != nil)
			} else {
				beego.Info("Old user:", sub.User.RealName, ";WebSocket:", sub.Conn != nil)
			}
		case event := <-this.Publish:
			// Notify waiting list.
			for ch := this.WaitingList.Back(); ch != nil; ch = ch.Prev() {
				ch.Value.(chan bool) <- true
				this.WaitingList.Remove(ch)
			}

			this.BroadcastWebSocket(event)
			NewArchive(event)

			if event.Type == EVENT_MESSAGE {
				beego.Info("Message from", event.User, ";Content:", event.Content)
			}
		case unsub := <-this.Unsubscribe:
			for sub := this.Subscribers.Front(); sub != nil; sub = sub.Next() {
				if sub.Value.(Subscriber).User == unsub {
					this.Subscribers.Remove(sub)
					// Clone connection.
					ws := sub.Value.(Subscriber).Conn
					if ws != nil {
						ws.Close()
						beego.Error("WebSocket closed:", unsub)
					}
					this.Publish <- this.NewEvent(EVENT_LEAVE, unsub, "") // Publish a LEAVE event.
					break
				}
			}
		}
	}
}

// broadcastWebSocket broadcasts messages to WebSocket users.
func (this *ChatRoom) BroadcastWebSocket(event Event) {
	data, err := json.Marshal(event)
	if err != nil {
		beego.Error("Fail to marshal event:", err)
		return
	}

	for sub := this.Subscribers.Front(); sub != nil; sub = sub.Next() {
		// Immediately send event to WebSocket users.
		ws := sub.Value.(Subscriber).Conn
		if ws != nil {
			if ws.WriteMessage(websocket.TextMessage, data) != nil {
				// User disconnected.
				this.Unsubscribe <- sub.Value.(Subscriber).User
			}
		}
	}
}
