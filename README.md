

# beego_element_cms    [![Go Report Card](https://goreportcard.com/badge/github.com/beego/beego)](https://github.com/beego/beego/actions/workflows/test.yml) <img src="https://img.shields.io/badge/golang-1.16-blue"/> <img src="https://img.shields.io/badge/element--plus-2.0.1-green"/>

## 技术选项
- 前端： 用基于 [Vue](https://vuejs.org) 的 [Element](https://github.com/ElemeFE/element) 构建基础页面。
- 后端：用[beego v2](https://github.com/beego/beego) 高性能 快速开发的 web框架
- 数据库：采用`MySql`(8)版本，使用 beego orm 模块 实现对数据库的基本操作。
- API文档：使用`Swagger`构建自动化文档。
- 官网界面： 使用beego html/template 视图模板引擎（官网前台考虑到seo优化相关因素，所以用模板引擎输出，而不采用前后端分离）

## server 项目

## 导入数据库

进入到doc 文件夹  导入db_kaiyu.sql 到mysql数据

## 配置数据库链接

进入到server 文件夹 conf 下 app.conf 修改配置文件

```
appname = kaiyu-web
httpport = 8080
runmode = dev
copyrequestbody  = true
sessionon = true
[mysql]
conn = root:root@tcp(127.0.0.1:3306)/db_kaiyu
```
> 本项目支持 fist code 模式，除了初始数据表数据外其他的数据表如果不存在会在启动时候自动新增 ，修改某字段的时候只要在models文件夹修改相应的struct
数据库会相应的自动隐射新增修改


### 安装 bee 工具
```
  go get -u github.com/beego/bee/v2
```
> 安装完之后，bee 可执行文件默认存放在 $GOPATH/bin 里面，所以你需要把 $GOPATH/bin 添加到你的环境变量中，才可以进行下一步。

使用 Goland 等编辑工具，打开server目录，不可以打开 beego_elment_cms

```
  # 克隆项目
  git clone  https://github.com/Echosong/beego_element_cms.get
  # 进入server文件夹
  cd server
  # 使用 go mod 并安装go依赖包
  go generate
  
  # 通过bee 运行
  bee run 
  
  # 编译
  go build
```


### swagger自动化API文档
```
  bee gao docs
```
> 执行上面的命令后，server目录下会出现docs文件夹里的 docs.go, swagger.json, swagger.yaml 三个文件更新，启动go服务之后, 在浏览器输入 http://localhost:8080/swagger/index.html 即可查看swagger文档

## web  项目
```
# 进入web文件夹
cd web

# 安装依赖
cnpm install || npm install

# 启动web项目
npm run serve

# 打包编译
npm run build

```

> 注意前端的菜单，对应这后端的菜单权限管理，后端没有专门做菜单关联添加相关功能，可以通过数据库 表 hz_permission 进行修改添加

## 演示

官网前台界面： http://localhost:8080

![img](https://raw.githubusercontent.com/Echosong/beego_element_cms/master/doc/1.png)

后台：  http://localhost:8081

账号 admin  密码 654321

![img](https://raw.githubusercontent.com/Echosong/beego_element_cms/master/doc/2.png)

```
proxy: {
			"/api":{
				target: "http://localhost:8080/api",
				changeOrigin: true,
				pathRewrite:{
					'^/api':''
				}
			},
			'/upload':{
				target: "http://localhost:8080/upload",
				changeOrigin: true,
				pathRewrite: {
					'^/upload':''
				}
			}
			
		}
```
> 发布时候 nginx 运行请参考上面做相应的反向代理



## 商用

> 任意个人 或者公司使用， 遵循  MIT 最宽泛的开源协议

## 联系我们

如果觉得对你有帮助，请点个小星星


qq:313690636

qq群： 571627871

![img](https://github.com/Echosong/beego_element_cms/blob/main/doc/wx.png?raw=true)



