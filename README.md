## cloudgo-io
#### 一个简单的web应用程序，功能包括：

+ 支持静态文件服务
+ 支持简单js访问
+ 提交表单，并输出一个表格
+ 对 /unknown 给出开发中的提示，返回码 501

ps:通上次使用了martini框架，本次作业也使用martini框架。

### 测试

1.启动服务器

`go run main.go`

![启动服务器:](https://github.com/wuxuemin/cloudgo-io/blob/master/image/1.png)

2.支持静态文件服务

`curl -v http://localhost:9090`

##### curl测试结果
![curl测试:](https://github.com/wuxuemin/cloudgo-io/blob/master/image/2-2.png)
![curl测试:](https://github.com/wuxuemin/cloudgo-io/blob/master/image/2-3.png)

##### 通过浏览器更加直观展示：
![浏览器:](https://github.com/wuxuemin/cloudgo-io/blob/master/image/2-1.png)

3.支持简单js访问

`curl -v http://localhost:9090/json`
##### curl测试结果
![curl测试:](https://github.com/wuxuemin/cloudgo-io/blob/master/image/3-2.png)

![curl测试:](https://github.com/wuxuemin/cloudgo-io/blob/master/image/3-1.png)

##### 通过浏览器更加直观展示：
![浏览器:](https://github.com/wuxuemin/cloudgo-io/blob/master/image/3-3.png)


4.提交表单，并输出一个表格

`curl -v http://localhost:9090/login`

##### curl测试结果
![curl测试:](https://github.com/wuxuemin/cloudgo-io/blob/master/image/4-2.png)
![curl测试:](https://github.com/wuxuemin/cloudgo-io/blob/master/image/4-1.png)

##### 通过浏览器更加直观展示：
![浏览器:](https://github.com/wuxuemin/cloudgo-io/blob/master/image/4-3.png)

输入登录信息登录：
![浏览器:](https://github.com/wuxuemin/cloudgo-io/blob/master/image/4-4.png)


5.对 /unknown 给出开发中的提示，返回码 501

`curl -v http://localhost:9090/unknown`
##### curl测试结果
![curl测试:](https://github.com/wuxuemin/cloudgo-io/blob/master/image/5-2.png)

![curl测试:](https://github.com/wuxuemin/cloudgo-io/blob/master/image/5-1.png)

##### 通过浏览器更加直观展示：
![浏览器:](https://github.com/wuxuemin/cloudgo-io/blob/master/image/5-3.png)

