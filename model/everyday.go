package model

import "gorm.io/gorm"

type Everyday struct {
	Id       int    `json:"id"`
	Prefix   string `json:"prefix,omitempty" gorm:"type:string"`
	Task     string `json:"task,omitempty" gorm:"type:string"`
	Remark   string `json:"remark,omitempty" gorm:"type:string"`
	Duration string `json:"duration,omitempty" gorm:"type:string"`
	TimeStub string `json:"time_stub,omitempty" gorm:"type:string"`
	Deleted  gorm.DeletedAt
}

func (e *Everyday) Insert() error {
	err := DB.Create(e).Error
	return err
}

func (e *Everyday) Delete() error {
	err := DB.Delete(e).Error
	return err
}

func (e *Everyday) FindAll() ([]*Everyday, error) {
	var everyday []*Everyday

	if err := DB.Find(&everyday).Error; err != nil {
		return nil, err
	}
	return everyday, nil
}
