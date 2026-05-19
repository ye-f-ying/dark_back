/*
 * @Date: 2026-05-19 15:21:05
 * @LastEditTime: 2026-05-19 15:28:56
 * @FilePath: /dark_back/models/admin_account_role.go
 * @Description:
 */
package models

import (
	"time"

	_ "gorm.io/gorm"
)

// AdminAccountRole 管理员角色关联表
type AdminAccountRole struct {
	ID         uint64    `gorm:"primaryKey;autoIncrement;column:id" json:"id,string"`
	AccountID  uint64    `gorm:"type:bigint;not null;default:0;uniqueIndex:uk_account_role;column:account_id" json:"accountId,string"` //管理员ID
	RoleID     uint64    `gorm:"type:bigint;not null;default:0;uniqueIndex:uk_account_role;column:role_id" json:"roleId,string"`       //角色ID
	CreateTime time.Time `gorm:"autoCreateTime;column:create_time" json:"createTime"`
}

func (AdminAccountRole) TableName() string {
	return "admin_account_role"
}
