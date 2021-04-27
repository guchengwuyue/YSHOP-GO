package models

import (
	"context"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	"strconv"
	"yixiang.co/yshop/dto"
	"yixiang.co/yshop/untils"
)

type SysUser struct {
	Id     int64 `json:"id"`
	Avatar string `json:"avatar"`
	Email  string `json:"email"`
	Enabled int8 `json:"enabled"`
	Password string `json:"password"`
	Username string `json:"username" valid:"Required;"`
	//DeptId int32
	Phone string `json:"phone"`
	//JobId int32
	NickName string      `json:"nickName"`
	Sex string           `json:"sex"`
	Roles []*SysRole        `json:"roles" orm:"rel(m2m);rel_through(yixiang.co/yshop/models.SysUsersRoles)"`
	Jobs *SysJob         `json:"job" orm:"column(job_id);bigint;rel(one)""`
	Depts *SysDept          `json:"dept" orm:"column(dept_id);bigint;rel(one)""`
	Permissions []string `orm:"-"`
	RoleIds []int64      `json:"roleIds" orm:"-"`
	BaseModel
}

type RoleId struct {
	Id int64 `json:"id"`
}


func init() {
	orm.RegisterModel(new(SysUser))
}


func FindByUserId(id int64) ([]string,error)  {
	o := orm.NewOrm()
	var roles []SysRole
	_, err := o.Raw("SELECT r.* FROM sys_role r, sys_users_roles u " +
		"WHERE r.id = u.role_id AND u.user_id = ?", id).QueryRows(&roles)
	for k, _ := range roles {
		_, err = o.LoadRelated(&roles[k], "Menus")
	}

	var permissions []string

	for _, v := range roles {
		menus := v.Menus
		for _, m := range menus {
			if m.Permission == "" {
				continue
			}
			permissions = append(permissions, m.Permission)
		}
	}

	return permissions, err
}

//根据用户名返回
func GetUserByUsername(name string) (v *SysUser, err error)  {
	o := orm.NewOrm()
	user := &SysUser{}
	err = o.QueryTable(new(SysUser)).Filter("username",name).RelatedSel().One(user)
	if _, err = o.LoadRelated(user, "Roles");err != nil{
		return nil, err
	}
	if err == nil {
		permissions, _ := FindByUserId(user.Id)
		user.Permissions = permissions
		return user, nil
	}

	return nil, err
}

// GetUserById retrieves User by Id. Returns error if
// Id doesn't exist
func GetUserById(id int64) (v *SysUser, err error) {
	//var userlist []User
	o := orm.NewOrm()
	v = &SysUser{Id: id}

	err = o.QueryTable(new(SysUser)).Filter("Id", id).RelatedSel().One(v)


	if _, err = o.LoadRelated(v, "Roles");err!=nil{
		return nil, err
	}
	if err == nil {
		return v, nil
	}



	return nil, err
}

// get all
func GetAllUser(base dto.BasePage,query ...interface{}) (int,[]SysUser)  {
	var (
		tableName = "sys_user"
		users []SysUser
		condition = ""
	)
	if base.Blurry != "" {
		condition = " and username= '" + base.Blurry + "'"
	}
	if len(query) > 0 {
		deptId := query[0].(int64)
		enabled := query[1].(int64)
		if deptId >= 0 {
			condition += " and dept_id=" + strconv.FormatInt(deptId,10)
		}
		if enabled >= 0 {
			condition += " and enabled=" + strconv.FormatInt(enabled,10)
		}
	}
	total, _, rs := GetPagesInfo(tableName, base.Page, base.Size, condition)
	rs.QueryRows(&users)

	o := orm.NewOrm()
	for k, _ := range users {
		_, _ = o.LoadRelated(&users[k], "Jobs")
		_, _ = o.LoadRelated(&users[k], "Depts")
		_, _ = o.LoadRelated(&users[k], "Roles")
	}

	return total,users
}

func UpdateCurrentUser(m *SysUser) ( err error)  {
	o := orm.NewOrm()
	_, err = o.Update(m)
	return
}


func AddUser(m *SysUser) (id int64, err error) {
	o := orm.NewOrm()
	//transaction
	err = o.DoTx(func(ctx context.Context, txOrm orm.TxOrmer) error {
		// data
		m.Password = untils.HashAndSalt([]byte("123456"))
		id, e := txOrm.Insert(m)
		if e != nil {
			return e
		}
		var ee error
		// add user_role
		for _, roleId := range m.RoleIds {
			_, ee = txOrm.Raw("INSERT INTO sys_users_roles (user_id,role_id) VALUES (?,?)", id, roleId).Exec()
		}
		return ee
	})
	return 0,err
}

func UpdateByUser(m *SysUser) (err error) {
	o := orm.NewOrm()
	//transaction
	err = o.DoTx(func(ctx context.Context, txOrm orm.TxOrmer) error {
		// data
		_, e := txOrm.Update(m)
		if e != nil {
			return e
		}
		//先删除
		_, eee := txOrm.Raw("delete from sys_users_roles WHERE user_id = ?",  m.Id).Exec()
		if eee != nil {
			return eee
		}

		var ee error
		for _, roleId := range m.RoleIds {
			_, ee = txOrm.Raw("INSERT INTO sys_users_roles (user_id,role_id) VALUES (?,?)", m.Id, roleId).Exec()
		}

		return ee
	})

	logs.Error(err)

	return
}

func DelByUser(ids []int64) (err error) {
	str := untils.ReturnQ(len(ids))
	o := orm.NewOrm()
	err = o.DoTx(func(ctx context.Context, txOrm orm.TxOrmer) error {
		_, e1 := txOrm.Raw("UPDATE sys_user SET is_del = ? WHERE id in("+str+")", 1, ids).Exec()
		_, e2 := txOrm.Raw("delete from sys_users_roles WHERE user_id in("+str+")",  ids).Exec()

		if e1 != nil {
			return e1
		}
		return e2
	})
	return
}



