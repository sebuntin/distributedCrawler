cd .\rpcsupport\crawlerServer\
go build -o ..\..\bin\crawler.exe .\crawlerServer.go

cd ..\itemServer
go build -o ..\..\bin\saver.exe .\itemServer.go

cd ..\..\douban
go build -o ..\bin\douban.exe .\main.go
cd ..\fronted
go build -o ..\bin\fronted.exe .\main.go