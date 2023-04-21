package controllers

import (
	"encoding/json"
	"fmt"
	"mtv/utils"
	"net/http"
	"strings"
	"time"

	"github.com/beego/beego/v2/core/config"
	"github.com/beego/beego/v2/core/logs"
	"github.com/gorilla/websocket"
)

type Message struct {
	Id   string
	Data string
}

// websocket处理器，用于收集消息和发送消息
type Hub struct {
	//客户端列表，保存所有客户端
	clients map[string]Client
	//注册chan，客户端注册时添加到chan中
	register chan *Client
	//注销chan，客户端退出时添加到chan中，再从map中删除
	unregister chan *Client
	//消息
	send chan Message
}

// websocket连接对象，连接中包含每个连接的信息
type Client struct {
	id   string
	conn *websocket.Conn
	msg  chan Message
}

var upgrade = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool { // 跨域访问
		//如果不是get请求，返回错误
		if r.Method != "GET" {
			logs.Error("请求方式错误")
			return false
		}
		//如果路径中不包括chat，返回错误
		if r.URL.Path != "/socket" {
			logs.Error("请求路径错误")
			return false
		}

		return true
	},
}

// 初始化处理中心
var hub = &Hub{
	clients:    make(map[string]Client),
	register:   make(chan *Client),
	unregister: make(chan *Client),
	send:       make(chan Message),
}

func wsHandler(w http.ResponseWriter, r *http.Request) {

	conn, err := upgrade.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("获取连接失败:", err)
		return
	}
	id := r.URL.Query().Get("publicKey")
	//连接成功后注册客户端
	client := &Client{
		conn: conn,
		msg:  make(chan Message),
	}
	client.id = id
	hub.register <- client
	defer func() {
		hub.unregister <- client
	}()
	//得到连接后，就可以开始读写数据了
	go read(client)
	go notice(client)
	write(client)
}

func read(client *Client) {
	//从连接中循环读取信息
	for {
		_, msg, err := client.conn.ReadMessage()
		if err != nil {
			fmt.Println("客户端退出:", client.conn.RemoteAddr().String())
			hub.unregister <- client
			break
		}
		logs.Info("msg = ", string(msg))
		//将读取到的信息传入websocket处理器中的broadcast中
		var message Message
		json.Unmarshal(msg, &message)
		logs.Info("message id = ", message.Id)
		hub.send <- message
	}
}

// 好友通知
func notice(client *Client) {
	for {
		time.Sleep(5 * time.Second)
		logs.Info("notice")

		var msg []byte
		publicKey := client.id

		key := "friend_" + publicKey
		logs.Info("key = ", key)
		tmp, _ := utils.GetStr(key)
		logs.Info("public keys = ", tmp)
		if tmp != "" {
			names := strings.Split(tmp, `,`)
			name := names[0]

			if len(names) == 1 {
				tmp = ""
			} else {
				names = append(names[:1], names[2:]...)
				tmp = strings.Join(names, ",")
			}
			utils.SetStr(key, tmp, 24*time.Minute)

			msg = []byte(name)
			err := client.conn.WriteMessage(1, msg)
			if err != nil {
				fmt.Println("写入错误(notice)")
				break
			}
		}
	}
}

func write(client *Client) {
	for message := range client.msg {
		err := client.conn.WriteMessage(1, []byte(message.Data))
		if err != nil {
			fmt.Println("写入错误(write)")
			break
		}

	}
}

// 处理中心处理获取到的信息
func (h *Hub) run() {
	for {
		select {
		//从注册chan中取数据
		case client := <-h.register:
			logs.Info("id = ", client.id)
			//取到数据后将数据添加到客户端列表中
			logs.Info("register")
			h.clients[client.id] = *client
		case client := <-h.unregister:
			//从注销列表中取数据，判断客户端列表中是否存在这个客户端，存在就删掉
			logs.Info("unregister")
			if _, ok := h.clients[client.id]; ok {
				delete(h.clients, client.id)
			}
		case message := <-h.send:
			//从chan中取消息，发送到客户端的msg中
			logs.Info("receive message")
			client := h.clients[message.Id]
			select {
			case client.msg <- message:
			default:
				delete(h.clients, client.id)
				close(client.msg)
			}
		}
	}
}

func InitWebSocket() {
	logs.Info("websocket")
	port, _ := config.String("ws_port")
	go hub.run()
	addr := "127.0.0.1" + ":" + port
	http.HandleFunc("/socket", wsHandler)
	http.ListenAndServe(addr, nil)
}
