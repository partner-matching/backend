package data

import "time"

//easyjson:json
type User struct {
	Id           int32
	UserName     string `gorm:"column:username"`
	UserAccount  string `gorm:"column:userAccount"`
	AvatarUrl    string `gorm:"column:avatarUrl"`
	Gender       int32
	UserPassword string `gorm:"column:userPassword"`
	Phone        string
	Email        string
	UserStatus   int32     `gorm:"column:userStatus"`
	CreateTime   time.Time `gorm:"column:createTime"`
	UpdateTime   time.Time `gorm:"column:updateTime"`
	IsDelete     int32     `gorm:"column:isDelete"`
	Role         int32
	Tags         string
	Profile      string
}

//easyjson:json
type Team struct {
	Id          int32
	Name        string
	Description string
	MaxNum      int32     `gorm:"column:maxNum"`
	ExpireTime  time.Time `gorm:"column:expireTime"`
	UserId      int32     `gorm:"column:userId"`
	Status      int32
	Password    string
	CreateTime  time.Time `gorm:"column:createTime"`
	UpdateTime  time.Time `gorm:"column :updateTime"`
	IsDelete    int32     `gorm:"column:isDelete"`
}
