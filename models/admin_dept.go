/*
 * @Date: 2026-05-19 15:29:44
 * @LastEditTime: 2026-05-19 15:32:08
 * @FilePath: /dark_back/models/admin_dept.go
 * @Description:
 */
package models

import (
	"time"

	"gorm.io/gorm"
)

// AdminDept 部门表
type AdminDept struct {
	ID         uint64         `gorm:"primaryKey;autoIncrement;column:id" json:"id,string"`
	DeptName   string         `gorm:"type:varchar(64);not null;default:'';uniqueIndex:uk_dept_name;column:dept_name" json:"deptName"` //部门名称
	ParentID   uint64         `gorm:"type:bigint;not null;default:0;index:idx_parent_id;column:parent_id" json:"parentId,string"`     //父部门ID
	Leader     string         `gorm:"type:varchar(64);not null;default:'';column:leader" json:"leader"`                               //负责人
	Phone      string         `gorm:"type:varchar(32);not null;default:'';column:phone" json:"phone"`                                 //联系电话
	Email      string         `gorm:"type:varchar(128);not null;default:'';column:email" json:"email"`                                //邮箱
	Status     int8           `gorm:"type:tinyint;not null;default:1;column:status" json:"status"`                                    //状态：1启用 0禁用
	Sort       int            `gorm:"type:int;not null;default:0;column:sort" json:"sort"`                                            //排序号
	CreateTime time.Time      `gorm:"autoCreateTime;column:create_time" json:"createTime"`                                            //创建时间
	UpdateTime time.Time      `gorm:"autoUpdateTime;column:update_time" json:"updateTime"`                                            //更新时间
	DeleteTime gorm.DeletedAt `gorm:"column:delete_time" json:"-"`                                                                    //删除时间（软删除）

	Children []AdminDept `gorm:"foreignKey:ParentID" json:"children,omitempty"`
}

func (AdminDept) TableName() string {
	return "admin_dept"
}
