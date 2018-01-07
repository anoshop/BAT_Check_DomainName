package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"regexp"
	//"sync"
	"strings"
	"sync"
	"time"
	//"strconv"
	"strconv"
)


var wgg sync.WaitGroup
var (
	fileSuccF *os.File
	//fileError *os.File
	//fileLog   *os.File

	wgSuccFileF sync.Mutex
	//wgErrorFile sync.Mutex
	//wgLogFile sync.Mutex

)

func main() {


	waittime,_ := strconv.ParseInt("10000",10,64)
	dict := "./dict/3en.txt"



	//var curWaitSecond int64 = 1000*1000*1000*waittime   //秒
	var curWaitSecond int64 = 1000*1000*waittime //1000*1000 为0.1秒，面对普适性的，可开N多线程。

	//var curWaitSecond int64 = 1000*1000



	fileSuccF, _ = os.Create("./ai_succ.txt")

	//load dict file and goto check regist status
	f, err := os.Open(dict) //打开文件
	defer f.Close()         //打开文件出错处理
	if nil == err {
		buff := bufio.NewReader(f) //读入缓存
		for {

			line, err := buff.ReadString('\n') //以'\n'为结束符读入一行
			if err != nil || io.EOF == err {
				break
			}
			 testquery(line)
			time.Sleep(time.Duration(curWaitSecond))
		}

	}

	wgg.Wait()

	fmt.Println("作业任务已完成\r\n Task finish ")
	fileSuccF.Close()
	os.Exit(1)

}

func testquery(line string) {
	wgg.Add(1)

	conn, err := net.DialTimeout("tcp","whois.ai:43", 30*time.Second)
	//conn, err := net.Dial("tcp", tldinfo.WhoisServer+":43")
	if err != nil {
		fmt.Printf("connect error :%s  AAA\n", err.Error())
		//wgLogFile.Lock()
		//fileLog.WriteString(`connect error` + err.Error())
		//wgLogFile.Lock()
	}
	if conn == nil {
		fmt.Printf("connect error")
		//wgLogFile.Lock()
		//fileLog.WriteString(`connect error`)
		//wgLogFile.Lock()
		return
	}
	defer conn.Close()

	//fmt.Printf("connect ok \n")
	//wgLogFile.Lock()
	//fileLog.WriteString(`connect ok`)
	//wgLogFile.Lock()

	line = strings.Trim(line, " ")
	line = strings.Trim(line, "\n")
	domain := line + ".ai"

	//fmt.Printf(domain + "\r\n")
	fmt.Fprintf(conn, domain+"\r\n")

	time.Sleep(time.Second)
	var buf = make([]byte, 65536)
	n, err := conn.Read(buf)

	if err == nil {
		newstr := string(buf[0 : n-1])
		//newstr = string(buf)
		//fmt.Printf(newstr)

		////正则匹配
		reg := regexp.MustCompile("/ not registred\\./")
		re := reg.FindAllString(newstr, -1)
		if re == nil {
			fmt.Printf(domain+"  has been registed\n")
			//wgLogFile.Lock()
			//fileLog.WriteString(`已被注册:` + domain)
			//
			//wgLogFile.Lock()

		} else {
			fmt.Printf(domain+" can be regist!!can be regist!!can be regist!! \n")
			//wgLogFile.Lock()
			//fileLog.WriteString(`恭喜，可以被注册:` + domain)
			wgSuccFileF.Lock()
			fileSuccF.WriteString(domain + "\r\n")
			wgSuccFileF.Unlock()

			//wgLogFile.Lock()
		}
	} else {

		fmt.Printf(err.Error())
		//wgLogFile.Lock()
		//fileError.WriteString(domain)
		//wgLogFile.Lock()
	}

	wgg.Done()

}


