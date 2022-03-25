package model

import "gorm.io/gorm"

type Yearly struct {
	Id      int    `json:"id"`
	Prefix  string `json:"prefix" gorm:"type:string"`
	Task    string `json:"task" gorm:"type:string"`   // 任务
	Cron    string `json:"cron" gorm:"type:string"`   // 执行时间
	Remark  string `json:"remark" gorm:"type:string"` // 备注
	Deleted gorm.DeletedAt
}

func (y *Yearly) Insert() error {
	err := DB.Create(y).Error
	return err
}

func (y *Yearly) Delete() error {
	err := DB.Delete(y).Error
	return err
}

func (y *Yearly) FindAll() ([]*Yearly, error) {
	var yearly []*Yearly

	if err := DB.Find(&yearly).Error; err != nil {
		return nil, err
	}
	return yearly, nil
}
