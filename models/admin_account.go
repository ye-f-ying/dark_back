/*
 * @Date: 2026-05-19 14:59:18
 * @LastEditTime: 2026-05-19 15:16:04
 * @FilePath: /dark_back/models/admin_account.go
 * @Description:
 */
package models

import (
	"time"

	"gorm.io/gorm"
)

// AdminAccount 管理员账号表
type AdminAccount struct {
	ID            uint64    `gorm:"primaryKey;autoIncrement;column:id" json:"id,string"`
	Account       string    `gorm:"type:varchar(64);not null;default:'';uniqueIndex:uk_account;column:account" json:"account"`        //管理员账号
	Password      string    `gorm:"type:varchar(255);not null;default:'';column:password" json:"-"`                                   // 密码
	RealName      string    `gorm:"type:varchar(64);not null;default:'';column:real_name" json:"realName"`                            //真实姓名
	NickName      string    `gorm:"type:varchar(64);not null;default:'';column:nick_name" json:"nickName"`                            //昵称
	Email         string    `gorm:"type:varchar(128);not null;default:'';column:email" json:"email"`                                  //邮箱
	Phone         string    `gorm:"type:varchar(32);not null;default:'';column:phone" json:"phone"`                                   //手机号
	Avatar        string    `gorm:"type:varchar(255);not null;default:'';column:avatar" json:"avatar"`                                //头像地址
	Sex           int8      `gorm:"type:tinyint;not null;default:0;column:sex" json:"sex"`                                            //性别：0未知 1男 2女
	Status        int8      `gorm:"type:tinyint;not null;default:1;index:idx_status;column:status" json:"status"`                     //状态：1启用 0禁用
	IsAdmin       int8      `gorm:"type:tinyint;not null;default:0;column:is_admin" json:"isAdmin"`                                   //是否超级管理员：1是 0否
	LastLoginTime time.Time `gorm:"type:datetime;not null;default:'1970-01-01 00:00:00';column:last_login_time" json:"lastLoginTime"` //最后登录时间-默认注册时间
	LastLoginIP   string    `gorm:"type:varchar(64);not null;default:'';column:last_login_ip" json:"lastLoginIP"`                     //最后登录IP
	CreateBy      uint64    `gorm:"type:bigint;not null;default:0;column:create_by" json:"createBy,string"`                           //创建人ID

	UpdateBy   uint64         `gorm:"type:bigint;not null;default:0;column:update_by" json:"updateBy,string"` //更新人ID
	CreateTime time.Time      `gorm:"autoCreateTime;column:create_time" json:"createTime"`                    //创建时间
	UpdateTime time.Time      `gorm:"autoUpdateTime;column:update_time" json:"updateTime"`                    //更新时间
	DeleteTime gorm.DeletedAt `gorm:"column:delete_time" json:"-"`                                            //删除时间（软删除）

	// 关联关系
	//Roles []*AdminRole `gorm:"many2many:admin_account_role;foreignKey:ID;joinForeignKey:AccountID;References:ID;JoinReferences:RoleID" json:"roles,omitempty"`
}

func (AdminAccount) TableName() string {
	return "admin_account"
}
