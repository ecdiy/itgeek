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
```angular2html
BindAddr=:9000
UploadDir=./upload/  上传图片的目录
DbDriver=mysql       数据库驱动名称
DbDsn=root:root@tcp(127.0.0.1:3306)/gk?timeout=30s&charset=utf8mb4&parseTime=true  
ImgHost=http://127.0.0.1:9000
MultiSite=0         这个参数不要设置成1，这个多站点标识，如果要设置成多站点与开发者单独联系。

geek BindAddr=:88 UploadDir=./upload/ DbDriver=mysql DbDsn=root:root@tcp(127.0.0.1:3306)/gk?timeout=30s&charset=utf8mb4&parseTime=true ImgHost=http://127.0.0.1:9000

以上是默认值，如果与之一样可以不用配置
```

#### 数据库增量升级: 
/db/geek.sql是当前版本下的数据库脚本，如果需要做增量升级的，参考这个主题
[http://www.itgeek.top/p/topic/detail,32,10021]
 
#### 使用过程有问题请到 [http://www.itgeek.top] 相应的板块反馈，截图。

#### 功能列表
|功能分类|子功能|
|-|-|
|主题|发表（-20）, 附加（-20）, 评论+-5 , 感谢+-10,收藏|
|用户|设置，上传头像，登录奖励，个人主页|
|上传图片|压缩|
|积分||
|消息||
|笔记|分类，笔记|
|分享|微博|

QQ群：620063196