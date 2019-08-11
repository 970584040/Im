##还存在问题
1、消息没有实现本地化存储

2、页面加载有问题，没有加载没有和后台一样实现鉴权

3、登陆token只是暂时实现了一个特殊字符串，接下来准备实现单点登陆功能

4、注册功能不完善

5、忘记密码功能不完善

6、离线消息没有实现，暂时只是用户登陆才可以接收

7、聊天功能只是使用websocket功能

8、暂时只想到这么多了，后续想到了问题在更新，功能会继续完善

##接口地址
域名/user/login

其余接口地址在main函数中可以看到

数据库使用mysql，websocket使用的gorilla包下的，依赖官网下net包（需翻墙）

使用的xorm库

使用包下载地址为（不翻墙的方法，也可以到github上直接git clone）
cd $GOPATH

mkdir -p golang.org/x/net

cd golang.org/x/net

go get -u  github.com/golang/net/websocket

go get  github.com/gorilla/websocket

go get github.com/go-xorm/xorm

go get -u github.com/go-sql-driver/mysql



