package service

import (
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
	member := map[string]interface{}{
		"memberId": m.MemberId,
		"name":     m.Name,
		"email":    m.Email,
		"phone":    m.Phone,
		"address":  m.Address,
	}
	rowsAffected, err := models.CreateMember(member)

	return rowsAffected, err
}

func (m *Member) GetAll() ([]*models.TblMember, error) {
	members, err := models.GetMembers()
	if err != nil {
		return nil, err
	}

	return members, nil
}
