# Blog

## 项目结构图

![image-20211118134356010](README.assets/image-20211118134356010.png)

1.   `bin`  存放二进制文件
2.   `configs` 存放配置文件
3.   `docs` 存放接口文档
4.   `global` 全局变量
5.   `internal` 项目的主体部分
6.   `pkg` 为本项目设计的工具包
7.   `script` 脚本(sql语句等)
8.   `storage` 持久化层,存放日志,保存的图片等

## 数据库

![image-20211118134907626](README.assets/image-20211118134907626.png)

1.   公共model 

     ![image-20211118135023574](README.assets/image-20211118135023574.png) 

2.   `blog_article` 文章

     1.   `title` : 标题
     2.   `desc` :描述
     3.   `content` :内容
     4.   `cover_image_url` 图片链接
     5.   `state` :文章状态

3.   `blog_tag` 标签

     1.   `name` 标签名
     2.   ` state` 状态

4.   `blog_article_tag` 文章和标签的关联关系

     1.   `tag_id` 标签id
     2.   `article_id` 文章id

5.   `blog_auth` token信息

     1.   `app_key` app_key
     2.   `app_secret` app_secret

## configs 配置文件

使用viper搭配yaml文件进行配置文件的读取

1.   yaml文件

     ![image-20211118160115616](README.assets/image-20211118160115616.png)

2.   对应结构体

     ![image-20211118160204017](README.assets/image-20211118160204017.png)

3.   对应初始化viper

     ![image-20211118160409657](README.assets/image-20211118160409657.png)

4.   读取配置文件

     ![image-20211118160447875](README.assets/image-20211118160447875.png)

## swaggo 接口文档

<a href="[Go学习笔记(六) | 使用swaggo自动生成Restful API文档 | Razeen`s Blog (razeencheng.com)](https://razeencheng.com/post/go-swagger)">swaggo使用 </a>

1.   安装

     ![image-20211118162353816](README.assets/image-20211118162353816.png)

     如果出现swag未找到,请将下载的`swag.exe`放到`GOPATH/bin/`下

2.   `main`中注释

     ![image-20211118162516819](README.assets/image-20211118162516819.png)

     ![image-20211118162533424](README.assets/image-20211118162533424.png)

3.   Handle 注释

     ![image-20211118163116055](README.assets/image-20211118163116055.png)

     ![image-20211118163209660](README.assets/image-20211118163209660.png)

4.   生成文档和测试

     ![image-20211118164030705](README.assets/image-20211118164030705.png)

     ![image-20211118164055841](README.assets/image-20211118164055841.png)

     在项目运行后直接访问`IP:端口/swagger/index.html` 即可

## global 全局变量

![image-20211118164322020](README.assets/image-20211118164322020.png)

![image-20211118164335026](README.assets/image-20211118164335026.png)

用于项目中的使用

## pkg 项目包

用于将项目中的一些操作统一封装起来

![image-20211118164519089](README.assets/image-20211118164519089.png)

### app 包

![image-20211118164604475](README.assets/image-20211118164604475.png)

1.   app.go

     统一响应处理

     ![image-20211118164640456](README.assets/image-20211118164640456.png)

2.   form.go

     参数绑定 还是利用了`shouldBind`

     ![image-20211118165013420](README.assets/image-20211118165013420.png)

     举例:

     ![image-20211118165031608](README.assets/image-20211118165031608.png)

     参数绑定:

     ![image-20211118164834779](README.assets/image-20211118164834779.png)

3.   jwt.go

     API 权限访问控制

     JWT简介:

     ![image-20211118165356721](README.assets/image-20211118165356721.png)

     ![image-20211118165437883](README.assets/image-20211118165437883.png)

     ![image-20211118165448090](README.assets/image-20211118165448090.png)

     ![image-20211118165500275](README.assets/image-20211118165500275.png)

     >   第一部分指明对第三部分的加密算法以及使用的令牌类型
     >
     >   第二部分存储令牌的相关信息,如有效时间等,用于校验是否过期
     >
     >   第三部分使用一个密钥和第二部分通过第一部分的加密算法进行加密,用于校验令牌内容是否被更改过

     使用 `jwt-go` 包 进行jwt的生成和解析(判断时间和是否被修改过,令牌中无任何声明被认为是有效的)

     ![image-20211118165134898](README.assets/image-20211118165134898.png)

     对应中间件处理

     ![image-20211118172219227](README.assets/image-20211118172219227.png)

4.   分页处理

     统一对分页进行处理

     ![image-20211118170311806](README.assets/image-20211118170311806.png)

5.   time.go

     对时间的格式化处理

     ![image-20211118170330645](README.assets/image-20211118170330645.png)

### convent

1.   convent.go

     对字符转数字的封装

     ![image-20211118170420382](README.assets/image-20211118170420382.png)

### email

邮件服务

1.   email.go

     ![image-20211118170711216](README.assets/image-20211118170711216.png)

### errcode

对错误的统一标记

![image-20211118170812231](README.assets/image-20211118170812231.png)

1.   errcode.go

     对错误的操作封装

     ![image-20211118170902815](README.assets/image-20211118170902815.png)

2.   common_code.go

     通用错误码

     ![image-20211118171929667](README.assets/image-20211118171929667.png)

3.   项目模块错误码

     ![image-20211118171947000](README.assets/image-20211118171947000.png)

### limiter 限流器

使用 `ratelimit` 包实现简单高效地令牌桶

1.   基本属性

     ![image-20211118172538126](README.assets/image-20211118172538126.png)

2.   具体实现

     针对路由限流

     ![image-20211118172845751](README.assets/image-20211118172845751.png)

     对应中间件

     ![image-20211118173143944](README.assets/image-20211118173143944.png)

     使用

     对`/auth`路由进行限流

     ![image-20211118173233573](README.assets/image-20211118173233573.png)

### logger 日志管理

![image-20211118174223885](README.assets/image-20211118174223885.png)

使用 `lumberjack`进行日志的持久化保存

![image-20211118174420074](README.assets/image-20211118174420074.png)

日志内容设置

![image-20211118174534749](README.assets/image-20211118174534749.png)

日志格式化输出

![image-20211118174557586](README.assets/image-20211118174557586.png)

###  setting 热更新配置文件

使用`fsnotify`包实现

![image-20211118175031661](README.assets/image-20211118175031661.png)

`viper`则是通过 `fsnotify` 实现监听

![image-20211118175332177](README.assets/image-20211118175332177.png)

使用:

![image-20211118175446637](README.assets/image-20211118175446637.png)

初始化时加载配置文件

![image-20211118175506925](README.assets/image-20211118175506925.png)

如果文件改变则重新读取配置文件

### upload 文件相关操作

![image-20211118180951914](README.assets/image-20211118180951914.png)

### util

1.   md5.go

     ![image-20211118181038450](README.assets/image-20211118181038450.png)

## middleware中间件

### 访问日志控制

![image-20211118181431426](README.assets/image-20211118181431426.png)

将信息先保存起来,再发送

### 捕获异常

捕获异常并且短信通知

![image-20211118183104755](README.assets/image-20211118183104755.png)

## 项目内容

调用关系

router 层负责数据的整理然后与service层交互,回复数据给客户

service层负责与不同dao层沟通获取数据,返回给router

dao层负责组装数据与model层沟通,返回数据给service层

model层只负责与数据库进行操作,返回有用的数据给dao层

结论: service层的模型负责绑定客户数据和返回给客户.over