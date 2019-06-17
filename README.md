# go-explore
go入门研究
#### go入门指南：https://tour.go-zh.org/list
####《go编程指南》https://github.com/Unknwon/the-way-to-go_ZH_CN/blob/master/eBook/directory.md
####《go web编程》https://github.com/astaxie/build-web-application-with-golang/blob/master/zh/preface.md

telegraf
#### handbook：https://docs.influxdata.com/telegraf/v1.11/introduction/getting-started/
#### playground：https://rootnroll.com/d/telegraf/

influxdb
#### https://docs.influxdata.com/influxdb/v1.7/introduction/getting-started/
    
Gin Web Framework & examples
https://github.com/gin-gonic/gin
    1 yaml 配置 key作Config struct，通过flag.string(filePath)读取/etc/config.yaml
      初始化yaml key（mysql、log、router run） //方法指针传递
    2 gin router init 加载log等mware、设置path路由【controllers入口】
    3 controllers
         接收ctx *gin.Context
         定义传入json解析对应的struct（json："jsonkeyName" binding："required"）
         纯go语法逻辑，调用models service api，dao通过jmoiron/sqlx 中间件，操作db
         
            
       

