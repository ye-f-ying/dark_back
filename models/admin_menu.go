/*
 * @Date: 2026-05-19 15:06:27
 * @LastEditTime: 2026-05-19 15:23:08
 * @FilePath: /dark_back/models/admin_menu.go
 * @Description:
 */
package models

import (
	"time"

	_ "gorm.io/gorm"
)

// AdminMenu 菜单权限表
type AdminMenu struct {
	ID             uint64    `gorm:"primaryKey;autoIncrement;column:id" json:"id,string"`
	MenuName       string    `gorm:"type:varchar(64);not null;default:'';column:menu_name" json:"menuName"`                                             //菜单名称
	MenuType       int8      `gorm:"type:tinyint;not null;default:1;index:idx_menu_type;column:menu_type" json:"menuType"`                              //类型：1目录 2菜单 3按钮
	PermissionCode string    `gorm:"type:varchar(128);not null;default:'';uniqueIndex:uk_permission_code;column:permission_code" json:"permissionCode"` //权限标识（如 system:user:list）
	ParentID       uint      `gorm:"type:bigint;not null;default:0;index:idx_parent_id;column:parent_id" json:"parentId"`                               //父菜单ID
	Icon           string    `gorm:"type:varchar(128);not null;default:'';column:icon" json:"icon"`                                                     //图标
	RoutePath      string    `gorm:"type:varchar(255);not null;default:'';column:route_path" json:"routePath"`                                          //路由路径
	RouteName      string    `gorm:"type:varchar(128);not null;default:'';column:route_name" json:"routeName"`                                          //路由名称
	Component      string    `gorm:"type:varchar(255);not null;default:'';column:component" json:"component"`                                           //组件路径
	QueryParams    string    `gorm:"type:varchar(255);not null;default:'';column:query_params" json:"queryParams"`                                      //路由参数
	IsFrame        int8      `gorm:"type:tinyint;not null;default:1;column:is_frame" json:"isFrame"`                                                    //是否内嵌：1是 0否
	IsCache        int8      `gorm:"type:tinyint;not null;default:1;column:is_cache" json:"isCache"`                                                    //是否缓存：1缓存 0不缓存
	Visible        int8      `gorm:"type:tinyint;not null;default:1;column:visible" json:"visible"`                                                     //是否显示：1显示 0隐藏
	Status         int8      `gorm:"type:tinyint;not null;default:1;index:idx_status;column:status" json:"status"`                                      //状态：1启用 0禁用
	Sort           int       `gorm:"type:int;not null;default:0;column:sort" json:"sort"`                                                               //排序号
	Remark         string    `gorm:"type:varchar(255);not null;default:'';column:remark" json:"remark"`                                                 //备注
	CreateBy       uint64    `gorm:"type:bigint;not null;default:0;column:create_by" json:"createBy,string"`                                            //创建人ID
	UpdateBy       uint64    `gorm:"type:bigint;not null;default:0;column:update_by" json:"updateBy,string"`                                            //更新人ID
	CreateTime     time.Time `gorm:"autoCreateTime;column:create_time" json:"createTime"`                                                               //创建时间
	UpdateTime     time.Time `gorm:"autoUpdateTime;column:update_time" json:"updateTime"`                                                               //更新时间

	Children []AdminMenu `gorm:"foreignKey:ParentID" json:"children,omitempty"`
}

func (AdminMenu) TableName() string {
	return "admin_menu"
}
