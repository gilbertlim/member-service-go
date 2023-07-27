package service

import (
	"github.com/gilbertlim/member-service-go/models"
	"github.com/gilbertlim/member-service-go/pkg/dto"
)

func CreateMember(memberDto dto.MemberDto) (int64, error) {
	rowsAffected, err := models.CreateMember(memberDto)

	return rowsAffected, err
}

func GetMembers() ([]*models.TblMember, error) {
	members, err := models.GetMembers()
	if err != nil {
		return nil, err
	}

	return members, nil
}

func GetMember(memberId string) (*models.TblMember, error) {
	member, err := models.GetMember(memberId)
	if err != nil {
		return nil, err
	}

	return member, nil
}

func DeleteMember(memberId string) (int64, error) {
	rowsAffected, err := models.DeleteMember(memberId)
	return rowsAffected, err
}
