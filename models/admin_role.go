/*
 * @Date: 2026-05-19 15:05:20
 * @LastEditTime: 2026-05-19 15:23:19
 * @FilePath: /dark_back/models/admin_role.go
 * @Description:
 */
package models

import (
	"time"

	"gorm.io/gorm"
)

// AdminRole 角色表
type AdminRole struct {
	ID          uint64         `gorm:"primaryKey;autoIncrement;column:id" json:"id,string"`
	RoleName    string         `gorm:"type:varchar(50);not null;default:'';column:role_name" json:"roleName"`                      //角色名称
	Description string         `gorm:"type:varchar(255);not null;default:'';column:description" json:"description"`                //角色描述
	Sort        int            `gorm:"type:int;not null;default:0;column:sort" json:"sort"`                                        //排序号
	Status      int8           `gorm:"type:tinyint;not null;default:1;index:idx_status;column:status" json:"status"`               //状态：1启用 0禁用
	ParentID    uint64         `gorm:"type:bigint;not null;default:0;index:idx_parent_id;column:parent_id" json:"parentId,string"` //父角色ID
	IsDefault   int8           `gorm:"type:tinyint;not null;default:0;column:is_default" json:"isDefault"`                         //是否默认角色：1是 0否
	CreateBy    uint64         `gorm:"type:bigint;not null;default:0;column:create_by" json:"createBy,string"`                     //创建人ID
	UpdateBy    uint64         `gorm:"type:bigint;not null;default:0;column:update_by" json:"updateBy,string"`                     //更新人ID
	CreateTime  time.Time      `gorm:"autoCreateTime;column:create_time" json:"createTime"`                                        //创建时间
	UpdateTime  time.Time      `gorm:"autoUpdateTime;column:update_time" json:"updateTime"`                                        //更新时间
	DeleteTime  gorm.DeletedAt `gorm:"column:delete_time" json:"-"`                                                                //删除时间（软删除）

	Menus []AdminMenu `gorm:"many2many:admin_role_menu;foreignKey:ID;joinForeignKey:RoleID;References:ID;JoinReferences:MenuID" json:"menus,omitempty"`
}

func (AdminRole) TableName() string {
	return "admin_role"
}
