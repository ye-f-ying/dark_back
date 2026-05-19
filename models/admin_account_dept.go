/*
 * @Date: 2026-05-19 15:32:37
 * @LastEditTime: 2026-05-19 15:33:43
 * @FilePath: /dark_back/models/admin_account_dept.go
 * @Description:
 */
package models

import "time"

// AdminAccountDept 管理员部门关联表
type AdminAccountDept struct {
	ID         uint64    `gorm:"primaryKey;autoIncrement;column:id" json:"id,string"`
	AccountID  uint64    `gorm:"type:bigint;not null;default:0;uniqueIndex:uk_account_dept;column:account_id" json:"accountId,string"` //管理员ID
	DeptID     uint64    `gorm:"type:bigint;not null;default:0;uniqueIndex:uk_account_dept;column:dept_id" json:"deptId,string"`       //部门ID
	IsMain     int8      `gorm:"type:tinyint;not null;default:0;column:is_main" json:"isMain"`                                         //是否主部门：1是 0否
	CreateTime time.Time `gorm:"autoCreateTime;column:create_time" json:"createTime"`                                                  //创建时间
}

func (AdminAccountDept) TableName() string {
	return "admin_account_dept"
}
