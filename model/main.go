/*model Of User And Org*/
package model

type Org struct {
	ID          uint   `grom:"primaryKey" json:"id"`
	Name        string `json:"name"`
	Address     string `json:"address"`
	Description string `json:"desc"`
}
type User struct {
	ID      uint   `grom:"primaryKey" json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
	Age     int    `json:"age"`
	OrgId   int    `json:"orgId"`
	Email   string `json:"email" gorm:"unique"`
}
type UserWithRoleAndClass struct {
	ID      uint   `gorm:"primaryKey" json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
	Age     int    `json:"age"`
	Email   string `json:"email"`
	OrgId   int    `json:"orgId"`
	Role    string `json:"role"`
	Class   string `json:"std"`
}

type OrgPagination struct {
	Page      int   `json:"page"`
	PageSize  int   `json:"pageSize"`
	TotalPage int   `json:"totalPage"`
	Data      []Org `json:"data"`
}
type UserPagination struct {
	Page      int                    `json:"page"`
	PageSize  int                    `json:"pageSize"`
	TotalPage int                    `json:"totalPage"`
	Data      []UserWithRoleAndClass `json:"data"`
}
type Roles struct {
	ID   uint   `grom:"primaryKey" json:"id"`
	Role string `json:"role"`
}
type Classes struct {
	ID    uint   `grom:"primaryKey" json:"id"`
	Class string `json:"std"`
}
type Roleandclasses struct {
	ID      uint `grom:"primaryKey" json:"id"`
	UserId  uint `json:"userId"`
	RoleId  uint `json:"roleId"`
	ClassId uint `json:"stdId"`
}
