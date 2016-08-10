package admin

import (
	"blog/controllers"
	"blog/lib"
	"blog/models"
	"net/http"
	"time"

	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/websocket"
	"github.com/sluu99/uuid"
)

var (
	ChatRooms = make(map[string]lib.ChatRoom)
)

// ChartRoomController  controller
type ChatRoomController struct {
	controllers.UspController
}

//@MenuH {"name":"系统管理","parent":"0"}
//@MenuH {"name":"系统设置","parent":"系统管理"}
//@MenuV {"name":"聊天室","parent":"系统设置"}
func (this *ChatRoomController) ChatRooms() {
	chatrooms := new(models.SysChatRoom)
	userinfo := this.GetSession(controllers.CurrentUserSession)
	opSession, _ := userinfo.(*controllers.UserInfo)
	this.Data["Rooms"] = chatrooms.GetModelsBySql(`
		select * from SysChatRoom a
		left join SysChatRoomRelation b on a.ID=b.ChatRoom
		where b.Operator =? or a.Creator=?
	`, opSession.Operator.ID, opSession.Operator.ID)
	this.Layout = this.GetTemplatetype() + "/shared/layout.tpl"
	this.TplName = this.GetTemplatetype() + "/adminPages/chatrooms.tpl"
}

func (this *ChatRoomController) Create() {
	if this.IsAjax() {
		model := new(models.SysChatRoom)
		if error := this.ParseForm(model); error != nil {
			this.Rsp(false, error.Error())
			return
		}
		model.UUID = uuid.Rand().Hex()
		model.Creator = 1
		model.CreateTime = time.Now()
		count, error := model.Add()
		if error == nil && count > 0 {
			this.Rsp(true, "Success")
		} else {
			this.Rsp(false, "")
		}
		return
	}
	this.Layout = this.GetTemplatetype() + "/shared/layout.tpl"
	this.TplName = this.GetTemplatetype() + "/adminPages/addchatroom.tpl"
}
func (this *ChatRoomController) Join() {
	userinfo := this.GetSession(controllers.CurrentUserSession)
	opSession, _ := userinfo.(*controllers.UserInfo)

	chatroomid := this.GetString("uuid")
	if chatroomid == "" {
		chatroomid = uuid.Rand().Hex()
	}
	var chatroom lib.ChatRoom
	var ChatID string
	if id, error := uuid.FromStr(chatroomid); error == nil {
		ChatID = chatroomid
		if _, ok := ChatRooms[id.Hex()]; !ok {
			chatroom = *lib.NewChatroom()
			ChatRooms[ChatID] = chatroom
		} else {
			chatroom = ChatRooms[ChatID]
		}
	}
	this.Data["ChatRoomUUID"] = ChatID
	this.Data["UserName"] = opSession.Operator.LoginName
	this.Layout = this.GetTemplatetype() + "/shared/layout.tpl"
	this.TplName = this.GetTemplatetype() + "/adminPages/chatroomjoin.tpl"
	go chatroom.Start()
}

func (this *ChatRoomController) SendMsg() {

	userinfo := this.GetSession(controllers.CurrentUserSession)
	opSession, _ := userinfo.(*controllers.UserInfo)

	chatroomid := this.GetString("uuid")

	chatroom := ChatRooms[chatroomid]

	ws, err := websocket.Upgrade(this.Ctx.ResponseWriter, this.Ctx.Request, nil, 1024, 1024)
	if _, ok := err.(websocket.HandshakeError); ok {
		http.Error(this.Ctx.ResponseWriter, "Not a websocket handshake", 400)
		return
	} else if err != nil {
		beego.Error("Cannot setup WebSocket connection:", err)
		return
	}
	chatroom.Join(opSession.Operator, ws)
	defer chatroom.Leave(opSession.Operator)
	for {
		_, p, err := ws.ReadMessage()
		if err != nil {
			return
		}
		chatroom.Publish <- chatroom.NewEvent(lib.EVENT_MESSAGE, opSession.Operator, string(p))
	}
}
