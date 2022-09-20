/*
 * @Author: kingford
 * @Date: 2022-09-20 11:49:21
 * @LastEditTime: 2022-09-20 12:00:11
 */
package entity

// Book struct represents books table in database
type Book struct {
	ID          uint64 `gorm:"primary_key:auto_increment" json:"id"`
	Title       string `gorm:"type:varchar(255)" json:"title"`
	Description string `gorm:"type:text" json:"description"`
	UserID      uint64 `gorm:"not null" json:"-"`
	User        User   `gorm:"foreignkey:UserID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"user"`
}
