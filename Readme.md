

##功能：根据字典批量查询域名是否可以注册
	   BAT check domain name regist status tool by golang.
	   Base on  RFC 945 。


###用法：
###Ugage:

	 `程序名 域名后缀 data文件夹中任一字典文件名 多线程发动间隔微秒数(如io域名建议3秒,则写3000)`,
      如window下的 sg.exe com 3py.txt 3

-----

	  `application_name  DomainTld(like'com')  dict_file_name Waittime(MicronSeconds)`
	  eg:  `./mac_speed com 3py.txt 3`

	  Please build for your platfrom  by yourself .





###各大系统使用示例：
Window 用户，有2种用法：
	a、直接执行startWindow.bat 文件
	b、用cmd进入到文件所在目录，然后输入`window.exe com 3py.txt 3 `

Linux用户：
	请从终端进入文件所在目录，然后	`./linux com 3py.txt 3 `


OSX 用户：
	请从终端进入文件所在目录，然后	`./mac com 3py.txt 3 `


### 目前只提供mac_speed 这个编译好的文件，其他的请自行编译。所有平台都已测试通过



###字典增减：可以在dict目录增加任意文本文件，一行一个字符串
			Extra dict:  pls CURD any file in dict folder, one string on line 

###域名后缀增减：直接编辑data目录下的tld.org.json 
			Extra tld:   pls modiy tld.org.json in tld folder


