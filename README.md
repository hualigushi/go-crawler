# go-crawler

go语言实现的爬虫程序

爬取珍爱网数据

### crawler-single 单机版爬虫程序

运行： go run main.go

### crawler-goroutine 并发版爬虫程序

运行：go run main.go

### crawler-elasticsearch 数据存储到elasticsearch中

1.阿里云(47.97.163.47)服务器安装docker: yum install docker

2.启动docker: service docker install

3.安装elasticsearch 镜像: docker pull docker.elastic.co/elasticsearch/elasticsearch:7.2.0

4.启动elasticsearch: docker run -p 9200:9200 -p 9300:9300 -e "discovery.type=single-node" docker.elastic.co/elasticsearch/elasticsearch:7.2.0

postman查看 47.97.163.47:9200

5.安装elasticsearch go 客户端: F:\Go\src>go get -v github.com/olivere/elastic

6.运行: go run main.go

7.运行前端: go run frontend/start.go

8.查看前端界面: localhost: 8888

### crawler_cluster 集群版爬虫程序

1-5 相同

6 运行：

 F:\Go\src\crawler>go run crawler_distributed/persist/server/itemsaver.go --port=1234

 F:\Go\src\crawler>go run crawler_distributed/worker/server/worker.go --port=9000
 F:\Go\src\crawler>go run crawler_distributed/worker/server/worker.go --port=9001
 F:\Go\src\crawler>go run crawler_distributed/worker/server/worker.go --port=9002

 F:\Go\src\crawler>go run crawler_distributed/main.go --itemsaver_host=":1234" --worker_hosts=":9000,:9001,:9002"

 F:\Go\src\crawler>go run crawler/frontend/start.go
 localhost: 8888

