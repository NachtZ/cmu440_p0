// Implementation of a MultiEchoServer. Students should write their code in this file.

package p0
import (
	"bufio"
	"net"
	"strconv"
	"time"
)

const MAX = 99

type clientInfo struct {
    conn net.Conn
    msgch chan string
    live bool
}
type multiEchoServer struct {
    clients map[net.Conn]*clientInfo
    closeChan chan byte
    readChan chan string
}

// New creates and returns (but does not start) a new MultiEchoServer.
func New() MultiEchoServer {
	// TODO: implement this!
    mes := &multiEchoServer{
        clients: make(map[net.Conn]*clientInfo),
        closeChan: make(chan byte),
        readChan: make(chan string),
    }
    go func(){
        for{
            select{
                case msg := <- mes.readChan:
                    for _,cli := range mes.clients{
                        if cli.live && len(cli.msgch) < MAX{
                            cli.msgch <- string(msg)
                        }else{
                            if cli.live == false{
                                cli.conn.Close()
                                delete(mes.clients,cli.conn)
                            }
                        }
                    }
                case <-mes.closeChan:
                    return
            }
        }
    }()
	return mes
}

func (mes *multiEchoServer) Start(port int) error {
	// TODO: implement this!
    service := "localhost:" + strconv.Itoa(port)
    tcpAddr,err := net.ResolveTCPAddr("tcp",service)
    ln,err := net.ListenTCP("tcp",tcpAddr)
    if err!= nil{
        return err
    }
    go func(){
        for{
            select{
                case <- mes.closeChan:
                    ln.Close()
                    return
                default:
            }
            ln.SetDeadline(time.Now().Add(time.Millisecond))
            conn,err := ln.Accept()
            if err != nil {
                continue
            }
            client := &clientInfo{
                conn:conn,
                msgch : make(chan string, MAX),
                live: true,
            }
            mes.clients[conn] = client
            go mes.handleClientRead(client)
			go mes.handleClientWrite(client)
        }
    }()
	return nil
}

func (mes *multiEchoServer) Close() {
	// TODO: implement this!
    close(mes.closeChan)
    
    for _,client := range mes.clients{
     //   close(client.msgch)
     client.live = false
        client.conn.Close()
    }
  //  close(mes.readChan)
}

func (mes *multiEchoServer) Count() int {
	// TODO: implement this!
    ret := 0
    for _,c := range mes.clients{
        if c.live{
            ret ++
        }
    }
	return ret
}

 
func (mes *multiEchoServer) handleClientRead(client *clientInfo) {
	reader := bufio.NewReader(client.conn)
	for {
		message,err := reader.ReadBytes('\n')
		if err != nil {
			client.live = false
			return
		}
		mes.readChan <- string(message)	
	}
}

func (mes *multiEchoServer) handleClientWrite(client * clientInfo) {
	for {
		select {
        case <- mes.closeChan:
			return
		
		case message := <- client.msgch:
			if client.live {
				client.conn.Write([]byte(message))
			}
        }
	}
}

// TODO: add additional methods/functions below!
