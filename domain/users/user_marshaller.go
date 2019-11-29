package users

import "encoding/json"

type PublicUser struct {
	ID        uint64 `json:"id"`
	CreatedOn uint64 `json:"created_on"`
	UpdatedOn uint64 `json:"updated_on"`
	Status    string `json:"status"`
}

type PrivateUser struct {
	ID        uint64 `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	CreatedOn uint64 `json:"created_on"`
	UpdatedOn uint64 `json:"updated_on"`
	Status    string `json:"status"`
}

func (user *User) Marshall(isPublic bool) interface{} {
	if isPublic {
		return &PublicUser{
			ID:        user.ID,
			CreatedOn: user.CreatedOn,
			Status:    user.Status,
			UpdatedOn: user.UpdatedOn,
		}
	}

	userJson, _ := json.Marshal(user)
	var privateUser PrivateUser
	json.Unmarshal(userJson, &privateUser)
	return privateUser
}

func (users Users) Marshall(isPublic bool) []interface{} {
	result := make([]interface{}, len(users))
	for index, user := range users {
		result[index] = user.Marshall(isPublic)
	}
	return result
}
