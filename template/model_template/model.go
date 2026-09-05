/*
 * @Date: 2026-06-16 13:47:14
 * @LastEditTime: 2026-06-16 14:37:32
 * @FilePath: /dark_back/template/model_template/model.go
 * @Description:
 */
package model_template

import (
	"fmt"

	"github.com/ye-f-ying/dark_pkg/pkg/db"

	"gorm.io/gorm"
)

/**
 * @description: 获取模型
 * @return {*}
 */
func GetModel() *gorm.DB {
	return db.GetGormDBMYSQL()
}

/**
 * @description: 获取单条
 * @param {int64} id
 * @return {*}
 */
func GetById[T any](id int64) (*T, error) {
	var obj T

	err := GetModel().
		First(&obj, id).
		Error

	if err != nil {
		return nil, err
	}

	return &obj, nil
}

/**
 * @description: 创建
 * @param {any} obj
 * @return {*}
 */
func Create[T any](obj *T) error {
	if obj == nil {
		return fmt.Errorf("model is nil")
	}

	return GetModel().
		Create(obj).
		Error
}

/**
 * @description: 根据ID更新
 * @param {int64} int64
 * @param {*any} obj
 * @return {*}
 */
func UpdateById[T any](id int64, obj *T) error {
	if obj == nil {
		return fmt.Errorf("model is nil")
	}

	return GetModel().
		Model(new(T)).
		Where("id = ?", id).
		Omit("id").
		Updates(obj).
		Error
}

/**
 * @description: 保存
 * @param {*any} obj
 * @return {*}
 */
func Save[T any](obj *T) error {
	if obj == nil {
		return fmt.Errorf("model is nil")
	}

	return GetModel().
		Save(obj).
		Error
}

/**
 * @description: 删除
 * @param {int64} id
 * @return {*}
 */
func DeleteById[T any](id int64) error {
	var obj T

	return GetModel().
		Where("id = ?", id).
		Delete(&obj).
		Error
}

/**
 * @description: 统计
 * @return {*}
 */
func Count[T any](where map[string]interface{}) (int64, error) {
	var obj T
	var total int64

	db := GetModel().Model(&obj)

	if len(where) > 0 {
		db = db.Where(where)
	}

	err := db.Count(&total).Error

	return total, err
}

/**
 * @description: 列表查询
 * @return {*}
 */
func List[T any](
	where map[string]interface{},
	page int,
	pageSize int,
	order string,
) ([]T, int64, error) {

	var list []T
	var total int64
	var obj T

	db := GetModel().Model(&obj)

	if len(where) > 0 {
		db = db.Where(where)
	}

	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if order != "" {
		db = db.Order(order)
	}

	if page > 0 && pageSize > 0 {
		db = db.
			Offset((page - 1) * pageSize).
			Limit(pageSize)
	}

	if err := db.Find(&list).Error; err != nil {
		return nil, 0, err
	}

	return list, total, nil
}
