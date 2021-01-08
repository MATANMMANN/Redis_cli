package Redis_Cli

import "strconv"

func HandleBulk(bulk []string) string{
	bulkLength:=len(bulk)
	returnArg := "*" + strconv.Itoa(bulkLength)
	return returnArg
}

func HandleData(tcpData string, command[]string) string{
	for i := 0; i < len(command); i++ {
		tcpData =tcpData + "$" + strconv.Itoa(len(command[i])) + "\r\n" + command[i] + "\r\n"
	}
	return tcpData
}
