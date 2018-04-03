package roles

import (
	"fmt"
	"net/http"
)

const (
	// Anyone is a role for any one
	Anyone = "*"
)

// 判断当前请求是否匹配某一角色
// Checker check current request match this role or not
type Checker func(req *http.Request, user interface{}) bool

// 初始化一个角色
// New initialize a new `Role`
func New() *Role {
	return &Role{}
}

// 定义角色数据结构
// Role is a struct contains all roles definitions
type Role struct {
	definitions map[string]Checker
}

// 根据条件注册角色
// Register register role with conditions
func (role *Role) Register(name string, fc Checker) {
	if role.definitions == nil {
		role.definitions = map[string]Checker{}
	}

	definition := role.definitions[name]
	if definition != nil {
		fmt.Printf("Role `%v` already defined, overwrited it!\n", name)
	}
	role.definitions[name] = fc
}

// 初始化角色的权限
// NewPermission initialize permission
func (role *Role) NewPermission() *Permission {
	return &Permission{
		Role:         role,
		AllowedRoles: map[PermissionMode][]string{},
		DeniedRoles:  map[PermissionMode][]string{},
	}
}

// 定义角色许可方法
// Allow allows permission mode for roles
func (role *Role) Allow(mode PermissionMode, roles ...string) *Permission {
	return role.NewPermission().Allow(mode, roles...)
}

// 定义角色拒绝方法
// Deny deny permission mode for roles
func (role *Role) Deny(mode PermissionMode, roles ...string) *Permission {
	return role.NewPermission().Deny(mode, roles...)
}

// 根据条件获取角色
// Get role defination
func (role *Role) Get(name string) (Checker, bool) {
	fc, ok := role.definitions[name]
	return fc, ok
}

// 删除定义的角色
// Remove role definition
func (role *Role) Remove(name string) {
	delete(role.definitions, name)
}

// 重置角色
// Reset role definitions
func (role *Role) Reset() {
	role.definitions = map[string]Checker{}
}

// 根据用户返回角色
// MatchedRoles return defined roles from user
func (role *Role) MatchedRoles(req *http.Request, user interface{}) (roles []string) {
	if definitions := role.definitions; definitions != nil {
		for name, definition := range definitions {
			if definition(req, user) {
				roles = append(roles, name)
			}
		}
	}
	return
}

// 检查当前用户的角色
// HasRole check if current user has role
func (role *Role) HasRole(req *http.Request, user interface{}, roles ...string) bool {
	if definitions := role.definitions; definitions != nil {
		for _, name := range roles {
			if definition, ok := definitions[name]; ok {
				if definition(req, user) {
					return true
				}
			}
		}
	}
	return false
}
