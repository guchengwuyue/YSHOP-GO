<h1 style="text-align: center">YSHOP-GO 后台管理系统</h1>


#### 项目简介
YSHOP-GO基于当前流行技术组合的前后端RBAC管理系统：Go1.15.x+Beego2.x+Jwt+Redis+Mysql8+Vue
的前后端分离系统，权限控制采用 RBAC，支持数据字典与数据权限管理，支持动态路由等

**体验地址：**  [https://go.yixiang.co](https://go.yixiang.co)

**账号密码：** `admin / 123456`

#### 项目源码

|     |   后端源码  |   前端源码  |
|---  |--- | --- |
|  码云  |  https://gitee.com/guchengwuyue/YSHOP-GO   |  https://gitee.com/guchengwuyue/YSHOP-GO-WEB   |
|  github   | https://github.com/guchengwuyue/YSHOP-GO   |  https://github.com/guchengwuyue/YSHOP-GO-WEB |


####  系统功能
- 用户管理：提供用户的相关配置 
- 角色管理：对权限与菜单进行分配，可根据部门设置角色的数据权限 
- 菜单管理：已实现菜单动态路由，后端可配置化 
- 部门管理：可配置系统组织架构，树形表格展示 
- 岗位管理：配置各个部门的职位 
- 字典管理：可维护常用一些固定的数据，如：状态，性别等 


#### 详细结构

```
- common 公共模块
    - constant.go 常量
- conf 公共配置
- controllers 控制器模块
	- admin 管理后台
	- app app模块
    - base.go 公共控制器
- dto dto模块
- jwt jwt模块
- logs 日志存放
- models 公共模型
- redis redis模块
	- pool.go 连接池
    - redis.go 常用操作
- routers 路由
- static 上传图片资源目录
- swagger swagger
- tests tests
- untils 工具模块
- vo vo模块
```
#### 启动与部署
```
1、go get 安装所需依赖
2、导入sql/yshop_go.sql,修改conf/app.conf 里数据库与redis配置
3、bee run 启动即可
4、线上部署： bee pack  -be GOOS=linux  打包然后上传服务器命令： nohup ./yshop & 
5、配置nginx 反向代理即可
```
#### 特别鸣谢

- beego:https://github.com/beego/beego/
- vue:https://github.com/vuejs/vue
- element:https://github.com/ElemeFE/element

