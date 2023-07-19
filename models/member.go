package models

type TblMember struct {
	MemberId string `json:"member_id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
}

func CreateMember(data map[string]interface{}) (int64, error) {
	member := TblMember{
		MemberId: data["memberId"].(string),
		Name:     data["name"].(string),
		Email:    data["email"].(string),
		Phone:    data["phone"].(string),
		Address:  data["address"].(string),
	}
	result := db.Create(&member)

	return result.RowsAffected, result.Error
}

func GetMembers() ([]*TblMember, error) {
	var members []*TblMember

	results := db.Find(&members)

	return members, results.Error
}
