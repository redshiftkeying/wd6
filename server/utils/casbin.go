package utils

import (
	"github.com/casbin/casbin/v2/model"
)

// 初始化权限模型
func init() {
	// 支持的模型类型ACL、RBAC、多租户RBAC、ABAC
	//ACL
	//ACL with superuser
	//ACL without resources
	//ACL without users
	//RBAC
	//RBAC with resource roles
	//RBAC with domains/tenants
	//RBAC with deny-override
	//ABAC
	//ABAC with policy rule
	//RESTful (KeyMatch)
	//RESTful (KeyMatch2)
	//IP match
	//Priority
	// https://casbin.org/en/editor
	modeType := "rbac" // TODO:viper
	m := model.NewModel()
	switch modeType {
	case "acl":
		m.AddDef("r", "r", "sub, obj, act")
		m.AddDef("p", "p", "sub, obj, act")
		m.AddDef("e", "e", "some(where (p.eft == allow))")
		m.AddDef("m", "m", "r.sub == p.sub && r.obj == p.obj && r.act == p.act")
	case "rbac":
		m.AddDef("r", "r", "sub, obj, act")
		m.AddDef("p", "p", "sub, obj, act")
		m.AddDef("g", "g", "_, _")
		m.AddDef("e", "e", "some(where (p.eft == allow))")
		m.AddDef("m", "m", "g(r.sub, p.sub) && r.obj == p.obj && r.act == p.act")
	default:
		m.AddDef("r", "r", "sub, obj, act")
		m.AddDef("p", "p", "sub, obj, act")
		m.AddDef("e", "e", "some(where (p.eft == allow))")
		m.AddDef("m", "m", "r.sub == r.obj.Owner") // ABAC
	}
	//e := casbin.NewEnforcer(m, "db/role_policy.csv")
}
