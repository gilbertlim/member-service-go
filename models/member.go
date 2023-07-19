package models

import (
	"github.com/mitchellh/mapstructure"
	"log"
)

type TblMember struct {
	MemberId string `json:"memberId"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
}

func CreateMember(data map[string]interface{}) (int64, error) {
	member := &TblMember{}
	err := mapstructure.Decode(data, &member)
	if err != nil {
		log.Fatal(err)
	}

	result := db.Create(member)

	return result.RowsAffected, result.Error
}

func GetMembers() ([]*TblMember, error) {
	var members []*TblMember

	results := db.Find(&members)

	return members, results.Error
}
