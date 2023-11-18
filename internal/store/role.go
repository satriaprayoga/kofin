package store

import "time"

type KRole struct {
	RoleID      uint   `json:"role_id" gorm:"primary_key;auto_increment:true"`
	Name        string `gorm:"size:50;not null;unique" json:"name"`
	Description string `gorm:"size:255;not null" json:"description"`

	UserInput string    `json:"user_input" gorm:"type:varchar(20)"`
	UserEdit  string    `json:"user_edit" gorm:"type:varchar(20)"`
	TimeInput time.Time `json:"time_input" gorm:"type:timestamp(0) without time zone;default:now()"`
	TimeEdit  time.Time `json:"time_edit" gorm:"type:timestamp(0) without time zone;default:now()"`
}
