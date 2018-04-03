package roles

// 定义权限接口
// Permissioner permissioner interface
type Permissioner interface {
	HasPermission(mode PermissionMode, roles ...interface{}) bool
}

// 合并权限数据
// ConcatPermissioner concat permissioner
func ConcatPermissioner(ps ...Permissioner) Permissioner {
	var newPS []Permissioner
	for _, p := range ps {
		if p != nil {
			newPS = append(newPS, p)
		}
	}
	return permissioners(newPS)
}

type permissioners []Permissioner

// 检查是否具有权限
// HasPermission check has permission for permissioners or not
func (ps permissioners) HasPermission(mode PermissionMode, roles ...interface{}) bool {
	for _, p := range ps {
		if p != nil && !p.HasPermission(mode, roles) {
			return false
		}
	}

	return true
}
