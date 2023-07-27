package models

import (
	"github.com/gilbertlim/member-service-go/pkg/dto"
)

type TblMember struct {
	MemberId string `json:"memberId"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
}

func CreateMember(memberDto dto.MemberDto) (int64, error) {
	member := &TblMember{
		MemberId: memberDto.MemberId,
		Name:     memberDto.Name,
		Email:    memberDto.Email,
		Phone:    memberDto.Phone,
		Address:  memberDto.Address,
	}

	result := db.Create(member)
	return result.RowsAffected, result.Error
}

func GetMembers() ([]*TblMember, error) {
	var members []*TblMember

	results := db.Find(&members)
	return members, results.Error
}

func GetMember(memberId string) (*TblMember, error) {
	member := &TblMember{
		MemberId: memberId,
	}

	result := db.Find(&member)
	return member, result.Error
}

func DeleteMember(memberId string) (int64, error) {
	result := db.Where("member_id = ?", memberId).Delete(&TblMember{})
	return result.RowsAffected, result.Error
}
