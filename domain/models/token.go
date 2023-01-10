package models

type Token struct {
	Id     int64 `json:"id" gorm:"primaryKey"`
	Token  string `json:"token"`
	UserID int
	User   User   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
