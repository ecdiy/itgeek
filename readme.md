 ```angular2html
go get github.com/ecdiy/gpa
go get github.com/gin-gonic/gin
go get github.com/cihub/seelog
go get github.com/hunterhug/go_image/graphics
go get github.com/ecdiy/itgeek

```
 
 ```angular2html
 cd ui
 npm install
```


详情介绍：[http://itgeek.top] 

QQ群：620063196

#### 目录结构
```angular2html
      gk                    后端GO代码
      db                    数据库脚本
      ui/web                pc web端
      ui/m                  h5
      tools                 工具    
      tools/post-update     编译参考，可以放在.git/hooks/ 在push代码后自动编译，有前后端的编译,ITGeek.top网站用的自动发布脚本，有需要的可以参考。
      tools/nginx-dev.conf  开发时nginx配置
      tools/post-update     开发环境nginx配置
         
```

#### 其它工具
[https://nodejs.org/en/]

#### 参数说明
BindAddr=:9000
UploadDir=./upload/  上传图片的目录
DbDriver=mysql       数据库驱动名称
DbDsn=root:root@tcp(127.0.0.1:3306)/gk?timeout=30s&charset=utf8mb4&parseTime=true  
ImgHost=http://127.0.0.1:9000

geek BindAddr=:88 UploadDir=./upload/ DbDriver=mysql DbDsn=root:root@tcp(127.0.0.1:3306)/gk?timeout=30s&charset=utf8mb4&parseTime=true ImgHost=http://127.0.0.1:9000

以上是默认值，如果与之一样可以不用配置
