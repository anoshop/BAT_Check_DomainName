package tld

//TLD 配置信息，

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"strings"
)

//!json数据对应格式,用来解析用的
type T_Tld []struct {
	Tld         string
	Description string
	WhoisServer string
	Patterns    struct {
		NotRegistered string
		WaitPeriod    string
	}
	WaitPeriod int
}

//!单个域名的信息，其本质是T_tldd的一级成员
type Tld struct {
	Tld         string
	Description string
	WhoisServer string
	Patterns    struct {
		NotRegistered string
		WaitPeriod    string
	}
	WaitPeriod int
}

var(
	Tlditem Tld
)


//获取某个tld的配置信息
func GetTld(tld string, f string) (Tld, error) {



	//一次性加载json数据，如果文件不是/开头，默认载入一个本机测试文件
	if strings.Index(f, "/") == -1 {
		f = "/Users/Eric/PycharmProjects/domain/src/tld/tld.org.json"
	}

	data, err := ioutil.ReadFile(f)
	if err != nil {

		fmt.Printf(" 打开文件失败\n")
		return Tlditem, errors.New("打开文件失败")
	}

	//!解析json数据
	resp := T_Tld{}
	err = json.Unmarshal(data, &resp)
	if err != nil {
		fmt.Println("Unmarshal: ", err.Error())
		return Tlditem, errors.New("解析json失败")
	}
	//fmt.Println(&resp)

	//遍历数据
	for _, v := range resp {
		v.Patterns.NotRegistered = strings.Trim(v.Patterns.NotRegistered, "/") //此处的修改并不会影响resp本身，仅对v当前的修改而已
		//fmt.Println(k)
		//fmt.Println(v.Tld)
		//fmt.Println(v.WhoisServer)
		//fmt.Println(v.Patterns.NotRegistered) //域名尚未注册提示文字,String类型，默认为空
		//fmt.Println(v.Patterns.WaitPeriod)    //错误提示,String类型，默认为空
		//fmt.Println(v.WaitPeriod)             //int类型，默认赋值0
		if v.Tld == tld {
			Tlditem = v
		}

	}

	//fmt.Printf("END \n")
	//fmt.Println(Tlditem )

	return Tlditem, nil

}

////测试程序
//func main() {
//	data, err := getTld("io", "11")
//	if err == nil {
//		fmt.Println(data)
//	}
//}
