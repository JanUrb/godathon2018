package web

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"

	god "github.com/JanUrb/godathon2018"
	"github.com/JanUrb/godathon2018/client"
)

// Web - Struct for web specific things?
type Web struct {
	Switcher god.Switching
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:    2048,
	WriteBufferSize:   2048,
	EnableCompression: false,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// Run Websocket server to allow client registrations
func (web Web) Run() {
	log.Println("Starting webserver")

	//Serve static files
	fs := http.FileServer(http.Dir("static/voipathon"))
	http.Handle("/", fs)

	fsTestClient := http.FileServer(http.Dir("static/testclient"))
	http.Handle("/testclient/", http.StripPrefix("/testclient/", fsTestClient))

	http.HandleFunc("/ws", web.registerClient)
	log.Println("Waiting for connections")
	log.Fatal(http.ListenAndServe(":4242", nil))
}

func (web Web) registerClient(resWriter http.ResponseWriter, req *http.Request) {
	conn, err := upgrader.Upgrade(resWriter, req, nil)
	if err != nil {
		log.Println("Error while establishing websocket connection: ", err)
		return
	}
	// Spawn new client
	c := client.New(web.Switcher, conn)
	go c.Listen()
	log.Printf("New client request [IP:%s]\t", conn.RemoteAddr().String())
}
