package models

import "time"

type BaseModel struct {
	Id int64 `json:"-" orm:"auto"`

	CreatedAt time.Time `json:"-" orm:"auto_now_add;type(datetime)"`
	UpdatedAt time.Time `json:"-" orm:"auto_now;type(datetime);null"`
}
