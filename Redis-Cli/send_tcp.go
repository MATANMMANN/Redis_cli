package Redis_Cli

import (
	"net"
	"strings"
)

func SendTcpRequest(address string,data string){
	conn, err := net.Dial("tcp",address)
	conn.Write([]byte(data))
	if err != nil {
		println("Write to server failed:", err.Error())
	}
	reply :=make([]byte,64)
	var replay_length int
	replay_length, err = conn.Read(reply)
	replay_split := strings.Fields(string(reply[:replay_length]))
	getReplay(replay_split)
	if err != nil {
		println("Write to server failed:", err.Error())
	}
	conn.Close()
}

