package desafio001

import "github.com/asaskevich/govalidator"

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type User struct {
	ID    string `json:"id" valid:"uuid,notnull"`
	User  string `json:"user" valid:"notnull"`
	Email string `json:"email" valid:"notnull"`
}

func (user *User) isValid() error {
	_, err := govalidator.ValidateStruct(user)
	if err != nil {
		return err
	}
	return nil
}
