package roles

import "net/http"

// 定义全局角色变量
// Global global role instance
var Global = &Role{}

// 注册角色
// Register register role with conditions
func Register(name string, fc Checker) {
	Global.Register(name, fc)
}

// 允许授权（模式）给角色
// Allow allows permission mode for roles
func Allow(mode PermissionMode, roles ...string) *Permission {
	return Global.Allow(mode, roles...)
}

// 拒绝授权（模式）给角色
// Deny deny permission mode for roles
func Deny(mode PermissionMode, roles ...string) *Permission {
	return Global.Deny(mode, roles...)
}

// 获取角色定义
// Get role defination
func Get(name string) (Checker, bool) {
	return Global.Get(name)
}

// 删除定义的全局角色实例
// Remove role definition from global role instance
func Remove(name string) {
	Global.Remove(name)
}

// 重置定义的全局角色实例
// Reset role definitions from global role instance
func Reset() {
	Global.Reset()
}

// 匹配定义的全局角色实例
// MatchedRoles return defined roles from user
func MatchedRoles(req *http.Request, user interface{}) []string {
	return Global.MatchedRoles(req, user)
}

// 检查当前的用户角色
// HasRole check if current user has role
func HasRole(req *http.Request, user interface{}, roles ...string) bool {
	return Global.HasRole(req, user)
}

// 初始化一个默认的角色及权限
// NewPermission initialize a new permission for default role
func NewPermission() *Permission {
	return Global.NewPermission()
}
