package main


import (
	//"fmt"
	//"net"
	//"time"
)
import (
	"os"
	//"time"
	"strconv"
)

func main() {

	fileSucc, _ := os.Create("./1000.txt")

	for i := 1;i < 1000 ; i=i+1{
		fileSucc.WriteString(strconv.Itoa(i)+"\n")
		//println(i);
		//println("\r\n");



	}

	//conn, err := net.DialTimeout("tcp", "whois.tn:43", 30*time.Second)
	////conn, err := net.Dial("tcp", tldinfo.WhoisServer+":43")
	//if err != nil {
	//	fmt.Printf("connect error :%s  AAA\n", err.Error())
	//	//wgLogFile.Lock()
	//	//fileLog.WriteString(`connect error` + err.Error())
	//	//wgLogFile.Lock()
	//}
	//if conn == nil {
	//	fmt.Printf("connect error")
	//	//wgLogFile.Lock()
	//	//fileLog.WriteString(`connect error`)
	//	//wgLogFile.Lock()
	//	return
	//}
	//defer conn.Close()
	//
	////fmt.Printf("connect ok \n")
	////wgLogFile.Lock()
	////fileLog.WriteString(`connect ok`)
	////wgLogFile.Lock()
	//
	////fmt.Printf(domain + "\r\n")
	//fmt.Fprintf(conn, "4444444444.tn\r\n")
	//
	//time.Sleep(time.Second)
	//var buf = make([]byte, 65536)
	//n, err := conn.Read(buf)
	//
	//if err == nil {
	//	newstr := string(buf[0 : n-1])
	//
	//	//newstr = string(buf)
	//	fmt.Printf(newstr)
	//}
}
