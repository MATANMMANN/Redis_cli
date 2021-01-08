package Redis_Cli

import (
	"bufio"
	"fmt"
	"os"
)

func GetInput (address string) string{
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("redis " +address+"> ")
	text, _ := reader.ReadString('\n')
	return text
}
