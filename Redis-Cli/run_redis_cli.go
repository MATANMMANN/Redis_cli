package Redis_Cli

import (
	"strings"
)

const address = "127.0.0.1:6379"

func Start(){
	for {
		command := GetInput(address)
		commandArray := strings.Fields(command)
		tcpData := HandleBulk(commandArray) + "\r\n"
		tcpData = HandleData(tcpData, commandArray)
		SendTcpRequest(address, tcpData)
	}
}
