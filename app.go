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
	"tld"
	//"strconv"
	"strconv"
)

var dict string = "/Users/Eric/PycharmProjects/domain/src/dict/3py.txt"
var tlds string = "to"
var tldinfo tld.Tld

var wg sync.WaitGroup
var waittime int64 =3
var (
	fileSucc *os.File
	//fileError *os.File
	//fileLog   *os.File

	wgSuccFile sync.Mutex
	//wgErrorFile sync.Mutex
	//wgLogFile sync.Mutex

)

func main() {

	args := os.Args //获取用户输入的所有参数
	if args == nil || len(args) < 3 {
		helper() //如果用户没有输入,或参数个数不够,则调用该函数提示用户
		return
	} else {
		tlds = os.Args[1]
		dict = "./dict/" + os.Args[2]
		waittime,_ = strconv.ParseInt(os.Args[3],10,64)

	}



	//var curWaitSecond int64 = 1000*1000*1000*waittime   //秒
	var curWaitSecond int64 = 1000*1000*waittime //1000*1000 为0.1秒，面对普适性的，可开N多线程。

	//var curWaitSecond int64 = 1000*1000


	//Get Tld info
	tldinfo, _ = tld.GetTld(tlds, "./data/tld.org.json")
	if tldinfo.WhoisServer == `` || tldinfo.WhoisServer == "null" {
		fmt.Println("Whois服务器地址为空，程序退出" +
			"\r\nAPP exit for the whois Server is null,please switch other tld")
		os.Exit(1)
	}

	fileSucc, _ = os.Create("./" + tldinfo.Tld + "_" + time.Stamp + "_succ.txt")
	//fileSucc.Close()
	//fileError, _ = os.Create("./" + tldinfo.Tld + "_" + time.Stamp + "_error.txt")
	//fileError.Close()
	//fileLog, _ = os.Create("./" + tldinfo.Tld + "_" + time.Stamp + "_log.txt")
	//fileLog.Close()

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
			go query(line)
			time.Sleep(time.Duration(curWaitSecond))
		}

	}

	wg.Wait()

	fmt.Println("作业任务已完成\r\n Task finish ")
	fileSucc.Close()
	os.Exit(1)

}

func query(line string) {
	wg.Add(1)

	conn, err := net.DialTimeout("tcp", tldinfo.WhoisServer+":43", 30*time.Second)
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
	domain := line + "." + tldinfo.Tld

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
		reg := regexp.MustCompile(tldinfo.Patterns.NotRegistered)
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
			wgSuccFile.Lock()
			fileSucc.WriteString(domain + "\r\n")
			wgSuccFile.Unlock()

			//wgLogFile.Lock()
		}
	} else {

		fmt.Printf(err.Error())
		//wgLogFile.Lock()
		//fileError.WriteString(domain)
		//wgLogFile.Lock()
	}

	wg.Done()

}

func helper() {
	fmt.Println("功能：根据字典批量查询域名是否可以注册\r\n\r\n" +
		"用法：`程序名 域名后缀 data文件夹中任一字典文件名 多线程发动间隔秒数(如io域名建议3秒)`, \r\n      如window下的 sg.exe com 3py.txt 3 \r\n\r\n" +
		"字典增减：可以在dict目录增加任意文本文件，一行一个字符串\r\n" +
		"域名后缀增减：直接编辑data目录下的tld.json.txt\r\n" +
		"作者:Eric.c via  http://sunorg.net  Since 2016年12月22日\r\n" +
		"授权：随意用，\r\n" +
		"\r\n\r\n" +
		"Batch check domain name resgist status\r\n" +
		"Usage: APPLACATION TLD Dict_file_name MicroSecond\r\n" +
		"Example: ./mac_seepd com 2letter.txt  3 ")
}
