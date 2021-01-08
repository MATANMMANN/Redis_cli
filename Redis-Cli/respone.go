package Redis_Cli

import (
	"fmt"
	"strings"
)

func getReplay(replay []string){
	firstChar:=string(replay[0][0])
	if firstChar=="-" {
		replay[0] = strings.Replace(replay[0],"-","(error) ",1)
		replayString:= strings.Join(replay," ")
		fmt.Println(replayString)
	} else if firstChar=="+" {
		fmt.Println("OK")
	}else if replay[0]=="$-1"{
		fmt.Println("(nil)")
	} else {
		fmt.Println("\"" + replay[1] + "\"")
	}
}