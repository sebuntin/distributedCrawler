## 分布式豆瓣图书数据爬虫

本项目来自于慕课网讲师ccmous的 _Google资深工程师深度讲解Go语言_ 课程，课程地址为:[here](https://coding.imooc.com/class/180.html)

### 安装依赖

* 本项目使用 Elastic Search 作为数据存储
  * Elastic Search 下载地址 [here](https://www.elastic.co/cn/downloads)

### 配置文件 

配置文件在项目主目录下的 config 目录中，内容如下

```go
const (
	// Parser names
	ParseBookTag  = "ParseBookTag"
	ParseBookList = "ParseBookList"
	ParseBookInfo = "ParseBookInfo"

	// Crawler Service Ports For Test
	CrawlerPort = 1235

	// ElasticSearch
	ElasticIndex = "douban_book"

	// Fronted
	Fronted = 8888
	// Template File Path
    	TemplateFilePaht = "{YOUR_PATH}\\distributedCrawler\\fronted\\view\\template.html"
    	View             = "{YOUR_PATH}\\distributedCrawler\\fronted\\view"

	// Rate Limiting
	Qps = 20
)
```

### 使用教程

1. 预编译好了.exe 文件可供windows系统直接使用，首先开启elastic search服务（使用默认9200端口），然后命令行进入项目主目录下的 ./bin 目录

   启动数据存储服务，--port 为数据存储服务的端口号

   ```powershell
   ./saver --port 1234
   ```
   启动成功效果如下

   ```
   ❯ .\saver --port 1234
   2020/05/26 17:50:50 elastic search client start success ...
   2020/05/26 17:50:50 Item Service start success ... PORT :1234
   ```

   

   启动爬虫服务，--port 为爬虫服务的端口号，可以开启多个服务
   
   ```shell
   ./crawler --port 9001
   ./crawler --port 9002
   ./crawler --port 9003
   ```
   
   启动成功效果如下：

   ```
   .\crawler.exe --port 9000
   2020/05/26 17:51:22 Crawler Service start success ... PORT 9000
   ```
   
   
   
   启动爬虫, --item_saver 为数据存储服务的端口号，--worker 为前面开启的多个爬虫服务的端口号，使用 ',' 分割
   
   ``` powershell
   ./douban --item_saver 1234 --worker 9001,9002,9003
   ```
   
   前端页面展示
   
   ```
   .\fronted.exe
   ```
   
   浏览器访问本地端口 8888，得到如下展示页面
   
   ![](https://github.com/sebuntin/distributedCrawler/pic/blob/master/豆瓣爬虫.png)
   
   
   
   搜索关键字，"阿加莎"，得到如下结果
   
   ![](https://github.com/sebuntin/distributedCrawler/pic/blob/master/frontedpage.png)
   
   
   
2. 如果想要自己编译，请先安装 gRPC，并开启 go module。

   ```shell
   go get -u google.golang.org/grpc
   ```

   进入项目主目录下，windows 命令行输入 ：

   ```powershell
   .\build.bat
   ```

   编译成功

   ![](https://github.com/sebuntin/distributedCrawler/pic/blob/master/buil.png)


