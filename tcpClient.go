package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	//建立一个链接，(Dial拨号)
	conn,err:=net.Dial("tcp","0.0.0.0:20000")
	if err != nil {
		fmt.Printf("dial failed, err%v\n",err)
		return
	}
	fmt.Println("Conn Establoshed......")

	//读入输入的信息
	reader:=bufio.NewReader(os.Stdin)
	for{
		data,err:=reader.ReadString('\n')
		if err != nil {
			fmt.Printf("read from console failed, err%v\n",err)
			break
		}

		data=strings.TrimSpace(data)
		//传输数据到服务端
		_, err =conn.Write([]byte(data))
		if err != nil {
			fmt.Println()
		}

	}

}
