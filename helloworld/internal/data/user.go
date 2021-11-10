package main

import "time"

type User struct {
	ID          int64     `gorm:"column:id" db:"column:id" json:"id" form:"id"`
	Account     string    `gorm:"column:account" db:"column:account" json:"account" form:"account"`
	Password    string    `gorm:"column:password" db:"column:password" json:"password" form:"password"`
	UserName    string    `gorm:"column:user_name" db:"column:user_name" json:"user_name" form:"user_name"`
	NikeName    string    `gorm:"column:nike_name" db:"column:nike_name" json:"nike_name" form:"nike_name"`
	Phone       string    `gorm:"column:phone" db:"column:phone" json:"phone" form:"phone"`
	Avatar      string    `gorm:"column:avatar" db:"column:avatar" json:"avatar" form:"avatar"`
	QrCode      string    `gorm:"column:qr_code" db:"column:qr_code" json:"qr_code" form:"qr_code"`
	IdCard      string    `gorm:"column:id_card" db:"column:id_card" json:"id_card" form:"id_card"`
	TrueName    string    `gorm:"column:true_name" db:"column:true_name" json:"true_name" form:"true_name"`
	TrueImg     string    `gorm:"column:true_img" db:"column:true_img" json:"true_img" form:"true_img"`
	Fingerprint string    `gorm:"column:fingerprint" db:"column:fingerprint" json:"fingerprint" form:"fingerprint"`
	Sex         int64     `gorm:"column:sex" db:"column:sex" json:"sex" form:"sex"`
	Email       string    `gorm:"column:email" db:"column:email" json:"email" form:"email"`
	Birthday    string    `gorm:"column:birthday" db:"column:birthday" json:"birthday" form:"birthday"`
	Region      string    `gorm:"column:region" db:"column:region" json:"region" form:"region"`
	Integral    float64   `gorm:"column:integral" db:"column:integral" json:"integral" form:"integral"`
	WxUnionId   string    `gorm:"column:wx_union_id" db:"column:wx_union_id" json:"wx_union_id" form:"wx_union_id"`
	WxOpenId    string    `gorm:"column:wx_open_id" db:"column:wx_open_id" json:"wx_open_id" form:"wx_open_id"`
	OpenIdType  int64     `gorm:"column:open_id_type" db:"column:open_id_type" json:"open_id_type" form:"open_id_type"`
	Ip          string    `gorm:"column:ip" db:"column:ip" json:"ip" form:"ip"`
	Mac         string    `gorm:"column:mac" db:"column:mac" json:"mac" form:"mac"`
	LevelId     int64     `gorm:"column:level_id" db:"column:level_id" json:"level_id" form:"level_id"`
	ParentId    int64     `gorm:"column:parent_id" db:"column:parent_id" json:"parent_id" form:"parent_id"`
	IsDisabled  int64     `gorm:"column:is_disabled" db:"column:is_disabled" json:"is_disabled" form:"is_disabled"`
	IsDeleted   int64     `gorm:"column:is_deleted" db:"column:is_deleted" json:"is_deleted" form:"is_deleted"`
	IsCertified int64     `gorm:"column:is_certified" db:"column:is_certified" json:"is_certified" form:"is_certified"`
	CreatedAt   time.Time `gorm:"column:created_at" db:"column:created_at" json:"created_at" form:"created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at" db:"column:updated_at" json:"updated_at" form:"updated_at"`
	DeletedAt   time.Time `gorm:"column:deleted_at" db:"column:deleted_at" json:"deleted_at" form:"deleted_at"`
}

type UserRepo struct {
}

func (u UserRepo) CreateUser() {

}
