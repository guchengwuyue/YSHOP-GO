package models

import (
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	"yixiang.co/yshop/untils"
	"yixiang.co/yshop/vo/menu"
)

type SysMenu struct {
	Id     int64         `json:"id"`
	Name string          `json:"name" valid:"Required;"`
	IFrame int8          `json:"iframe"`
	Component string     `json:"component"`
	Pid int64            `json:"pid"`
	Sort int32           `json:"sort"`
	Icon string          `json:"icon"`
	Path string          `json:"path"`
	Cache int8           `json:"cache"`
	Hidden int8          `json:"hidden"`
	ComponentName string `json:"componentName"`
	Permission string    `json:"permission"`
	Type int32           `json:"type"`
	Router string        `json:"router"`
	RouterMethod  string `json:"routerMethod"`
	Children []SysMenu   `json:"children" orm:"-"`
	Label string         `orm:"-" json:"label"`
	BaseModel
}

func init() {
	orm.RegisterModel(new(SysMenu))
}


func GetAllMenus(name string) []SysMenu {
	var menus []SysMenu
	o := orm.NewOrm()
	qs := o.QueryTable("sys_menu").Filter("is_del",0)
	if name != "" {
		qs = qs.Filter("name",name)
	}

	qs.All(&menus)
	return RecursionMenuList(menus,0)
}

//递归函数
func RecursionMenuList(data []SysMenu, pid int64) []SysMenu {
	var listTree = make([]SysMenu,0)
	for _, value := range data {
		value.Label = value.Name
		if value.Pid == pid {
			value.Children = RecursionMenuList(data, value.Id)
			listTree = append(listTree, value)
		}
	}
	return listTree
}

func AddMenu(m *SysMenu) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

func UpdateByMenu(m *SysMenu) (err error) {
	o := orm.NewOrm()
	_, err = o.Update(m)
	return
}

func DelByMenu(ids []int64) (err error) {
	str := untils.ReturnQ(len(ids))
	logs.Info(str)
	o := orm.NewOrm()
	_, err = o.Raw("UPDATE sys_menu SET is_del = ? WHERE id in("+str+")", 1, ids).Exec()
	return
}

//获取权限string
func FindByRouterAndMethod(url string, method string) (permission string) {
	o := orm.NewOrm()
	var menu SysMenu
	err := o.QueryTable(new(SysMenu)).Filter("router",url).Filter("router_method",method).One(&menu)
	if err != nil {
		return ""
	}
	return menu.Permission
}

func BuildMenus(uid int64) []menu.MenuVo  {
	o := orm.NewOrm()
	var lists orm.ParamsList
	_, err := o.Raw("SELECT r.* FROM sys_role r, sys_users_roles u " +
		"WHERE r.id = u.role_id AND u.user_id = ?",uid).ValuesFlat(&lists,"id")
	if err != nil {
		logs.Error(err)
	}
	idsStr := untils.Convert(lists)
	logs.Info(idsStr)
	var menus []SysMenu
	_, e := o.Raw("select m.* from sys_menu m LEFT OUTER JOIN sys_roles_menus t on m.id= t.menu_id " +
		"LEFT OUTER JOIN sys_role r on r.id = t.role_id where m.is_del=0 and m.hidden=0 and m.type!=2 and r.id in (?)",
		idsStr).QueryRows(&menus)

	if e != nil {
		logs.Error(e)
	}

	return buildMenus(buildTree(menus))

}

func buildTree(menus []SysMenu) ([]SysMenu) {
	var trees []SysMenu
	for _, menu := range menus {
		if menu.Pid == 0 {
			trees = append(trees, menu)
		}
	}

	for k, tree := range trees {
		var child []SysMenu
		for _, it := range menus {
			if it.Pid == tree.Id {
				child = append(child,it)
			}
		}
		trees[k].Children = child
	}

	return trees

}

func buildMenus(menus []SysMenu) []menu.MenuVo {
	var list []menu.MenuVo
	for _ , menuO := range menus {
		menuList := menuO.Children
		var menuVo = new(menu.MenuVo)

		if menuO.ComponentName != "" {
			menuVo.Name = menuO.ComponentName
		}else {
			menuVo.Name = menuO.Name
		}
		if menuO.Pid == 0 {
			menuVo.Path = "/" + menuO.Path
		}else {
			menuVo.Path = menuO.Path
		}
		menuVo.Hidden = menuO.Hidden
		//判断不是外链
		if menuO.IFrame == 0 {
			if menuO.Pid == 0 {
				if menuO.Component == "" {
					menuVo.Component = "Layout"
				}else{
					menuVo.Component = menuO.Component
				}
			}else if menuO.Component != "" {
				menuVo.Component = menuO.Component
			}
		}

		menuVo.Meta = menu.MenuMetaVo{Title: menuO.Name,Icon: menuO.Icon,NoCache: !untils.IntToBool(menuO.Cache)}

		if len(menuList) > 0 {
			menuVo.AlwaysShow = true
			menuVo.Redirect = "noredirect"
			menuVo.Children = buildMenus(menuList)
		}else if menuO.Pid == 0 {
			var menuVo1 = new(menu.MenuVo)
			menuVo1.Meta = menuVo.Meta
			if menuO.IFrame == 0 {
				menuVo1.Path = "index"
				menuVo1.Name = menuVo.Name
				menuVo1.Component = menuVo.Component
			}else{
				menuVo1.Path = menuO.Path
			}
			menuVo.Name = ""
			menuVo.Meta = menu.MenuMetaVo{}
			menuVo.Component = "Layout"
			var list1 []menu.MenuVo
			list1 = append(list1,*menuVo1)
			menuVo.Children = list1
		}

		list = append(list,*menuVo)

	}

	return list
}
