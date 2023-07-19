package models

import (
	"errors"
	"github.com/jinzhu/gorm"
)

type TblMember struct {
	MemberId string `json:"member_id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
}

func GetMembers() ([]*TblMember, error) {
	var members []*TblMember

	results := db.Find(&members)

	errors.Is(results.Error, gorm.ErrRecordNotFound)

	return members, nil
}
