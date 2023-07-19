package service

import (
	"github.com/fatih/structs"
	"github.com/gilbertlim/member-service-go/models"
)

type Member struct {
	MemberId string
	Name     string
	Email    string
	Phone    string
	Address  string
}

func (m *Member) CreateMember() (int64, error) {
	rowsAffected, err := models.CreateMember(structs.Map(m))

	return rowsAffected, err
}

func (m *Member) GetAll() ([]*models.TblMember, error) {
	members, err := models.GetMembers()
	if err != nil {
		return nil, err
	}

	return members, nil
}
