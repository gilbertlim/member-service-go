package member_service

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

func (m *Member) GetAll() ([]*models.TblMember, error) {
	members, err := models.GetMembers()
	if err != nil {
		return nil, err
	}

	return members, nil
}
