<h1 style="text-align: center">beego-vue-admin 后台管理系统</h1>


#### 项目简介
beego-vue-admin(目前版本v1.2)基于当前流行技术组合的前后端RBAC管理系统：Go1.15.x+Beego2.x+Jwt+Redis+Mysql8+Vue
的前后端分离系统，权限控制采用 RBAC，支持数据字典与数据权限管理，支持动态路由等

**体验地址：**  [https://go.yixiang.co](https://go.yixiang.co)

**账号密码：** `admin / 123456`

#### 项目源码

|     |   后端源码  |   前端源码  |
|---  |--- | --- |
|  码云  |  https://gitee.com/guchengwuyue/beego-vue-admin   |  https://gitee.com/guchengwuyue/beego-vue-admin-web   |
|  github   | https://github.com/guchengwuyue/beego-vue-admin  |  https://github.com/guchengwuyue/beego-vue-admin-web |


####  系统功能
- 用户管理：提供用户的相关配置 
- 角色管理：对权限与菜单进行分配，可根据部门设置角色的数据权限 
- 菜单管理：已实现菜单动态路由，后端可配置化 
- 部门管理：可配置系统组织架构，树形表格展示 
- 岗位管理：配置各个部门的职位 
- 字典管理：可维护常用一些固定的数据，如：状态，性别等 
- 日志管理：用户操日志记录 


#### 详细结构

```
- common 公共模块
    - jwt jwt模块
    - redis redis模块
    - untils 工具模块
    - constant.go 常量
- conf 公共配置
- controllers 控制器模块
	- admin 管理后台
	- app app模块
    - base.go 公共控制器
- fiters 过滤器
- initialize 初始化
- logs 日志存放
- models 公共模型
	- dto dto模块
	- vo vo模块
    - xxx.go 公共model
- routers 路由
- static 上传图片资源目录
- swagger swagger
- tests tests
```
#### 启动与部署
```
1、go get 安装所需依赖
2、导入sql/yshop_go.sql,修改conf/app.conf 里数据库与redis配置
3、bee run 启动即可
4、线上部署： bee pack  -be GOOS=linux  打包然后上传服务器命令： nohup ./yshop & 
5、配置nginx 反向代理即可
```
### 效果图
|  ![输入图片说明](https://images.gitee.com/uploads/images/2021/0324/091742_84ebf16e_477893.png "01.png") | ![输入图片说明](https://images.gitee.com/uploads/images/2021/0324/091938_e3408f92_477893.png "02.png")  |
|---|---|
|  ![输入图片说明](https://images.gitee.com/uploads/images/2021/0324/092327_601b2411_477893.png "03.png") |  ![输入图片说明](https://images.gitee.com/uploads/images/2021/0324/092343_9cadff96_477893.png "04.png") |

#### 特别鸣谢

- beego:https://github.com/beego/beego/
- vue:https://github.com/vuejs/vue
- element:https://github.com/ElemeFE/element

