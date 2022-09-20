/*
 * @Author: kingford
 * @Date: 2022-09-20 11:49:21
 * @LastEditTime: 2022-09-20 12:00:09
 */
package entity

// User represents users table in database
type User struct {
	ID       uint64  `gorm:"primary_key:auto_increment" json:"id"`
	Name     string  `gorm:"type:varchar(255)" json:"name"`
	Email    string  `gorm:"uniqueIndex;type:varchar(255)" json:"email"`
	Password string  `gorm:"->;<-;not null" json:"-"`
	Token    string  `gorm:"-" json:"token,omitempty"`
	Books    *[]Book `json:"books,omitempty"`
}
