package graylog

import (
	"encoding/json"
	"fmt"
	"net"

	"github.com/micro-in-cn/tutorials/microservice-in-micro/part6/basic/config"
	"github.com/micro/go-micro/util/log"
)

var (
	host string
	port int
)

// graylog 配置
type Graylog struct {
	Host    string `json:"host"`
	Port    int    `json:"port"`
	Appinfo interface{}
	sock    *net.UDPConn
}

func initGraylog() {
	log.Logf("[initGraylog] 初始化Graylog")

	c := config.C()

	myLog = &Graylog{}

	err := c.App("graylog", myLog)
	if err != nil {
		log.Logf("[initGraylog] %s", err)
		return
	}

	host = myLog.Host
	port = myLog.Port

	hostport := fmt.Sprintf("%s:%d", host, port)

	addr, err := net.ResolveUDPAddr("udp", hostport)
	if err != nil {
		fmt.Println("server address error. It MUST be a format like this hostname:port", port, hostport, err)
		return
	}

	// Create a udp socket and connect to server
	socket, err := net.DialUDP("udp4", nil, addr)
	if err != nil {
		fmt.Printf("connect to udpserver %v failed : %v", addr.String(), err.Error())
		return
	}
	myLog.sock = socket
	log.Logf("[initGraylog] success")

}

func (l *Graylog) Write(logs map[string]interface{}) {

	//defer l.sock.Close()
	logs["appinfo"] = l.Appinfo
	strlogs, _ := json.Marshal(logs)
	senddata := []byte(string(strlogs))
	_, err := l.sock.Write(senddata)
	if err != nil {
		fmt.Println("send data error ", err)
		return
	}

}
