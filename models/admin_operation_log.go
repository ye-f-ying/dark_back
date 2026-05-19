/*
 * @Date: 2026-05-19 15:34:33
 * @LastEditTime: 2026-05-19 15:36:50
 * @FilePath: /dark_back/models/admin_operation_log.go
 * @Description:
 */
package models

import "time"

// AdminOperationLog 操作日志表
type AdminOperationLog struct {
	ID            uint      `gorm:"primaryKey;autoIncrement;column:id" json:"id"`
	AccountID     uint      `gorm:"type:bigint;not null;default:0;index:idx_account_id;column:account_id" json:"accountId"`              //操作人ID
	AccountName   string    `gorm:"type:varchar(64);not null;default:'';column:account_name" json:"accountName"`                         //操作人账号
	Module        string    `gorm:"type:varchar(64);not null;default:'';index:idx_module;column:module" json:"module"`                   //操作模块
	OperationType int8      `gorm:"type:tinyint;not null;default:0;index:idx_operation_type;column:operation_type" json:"operationType"` //操作类型：1新增 2修改 3删除 4查询 5导出 6导入 9999其他
	Title         string    `gorm:"type:varchar(128);not null;default:'';column:title" json:"title"`                                     //操作标题
	URL           string    `gorm:"type:varchar(255);not null;default:'';column:url" json:"url"`                                         //请求URL
	Method        string    `gorm:"type:varchar(16);not null;default:'';column:method" json:"method"`                                    //请求方式
	RequestParams string    `gorm:"type:text;column:request_params" json:"requestParams"`                                                //请求参数
	ResponseData  string    `gorm:"type:text;column:response_data" json:"responseData"`                                                  //返回数据
	IP            string    `gorm:"type:varchar(64);not null;default:'';column:ip" json:"ip"`                                            //IP地址
	Location      string    `gorm:"type:varchar(128);not null;default:'';column:location" json:"location"`                               //IP归属地
	Duration      int       `gorm:"type:int;not null;default:0;column:duration" json:"duration"`                                         //耗时(毫秒)
	Status        int8      `gorm:"type:tinyint;not null;default:1;column:status" json:"status"`                                         //状态：1成功 0失败
	ErrorMsg      string    `gorm:"type:varchar(512);not null;default:'';column:error_msg" json:"errorMsg"`                              //错误消息
	CreateTime    time.Time `gorm:"autoCreateTime;column:create_time" json:"createTime"`                                                 //操作时间
}

func (AdminOperationLog) TableName() string {
	return "admin_operation_log"
}
