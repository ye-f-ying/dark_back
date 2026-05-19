/*
 * @Date: 2026-05-19 15:22:25
 * @LastEditTime: 2026-05-19 15:29:30
 * @FilePath: /dark_back/models/admin_role_menu.go
 * @Description:
 */
package models

import (
	"time"

	_ "gorm.io/gorm"
)

// AdminRoleMenu 角色菜单关联表
type AdminRoleMenu struct {
	ID         uint64    `gorm:"primaryKey;autoIncrement;column:id" json:"id,string"`
	RoleID     uint64    `gorm:"type:bigint;not null;default:0;uniqueIndex:uk_role_menu;column:role_id" json:"roleId,string"` //角色ID
	MenuID     uint64    `gorm:"type:bigint;not null;default:0;uniqueIndex:uk_role_menu;column:menu_id" json:"menuId,string"` //菜单ID
	CreateTime time.Time `gorm:"autoCreateTime;column:create_time" json:"createTime"`
}

func (AdminRoleMenu) TableName() string {
	return "admin_role_menu"
}
