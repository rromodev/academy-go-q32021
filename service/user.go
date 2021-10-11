package service

import (
	"github.com/rromodev/academy-go-q32021/model"
)

type CSVUserGetter interface {
	GetRecordById(id int, rowType string) ([]string, error)
	WriteRecord(record []string) (string, error)
}

type ExternalApiSetter interface {
	GetFromUrl() ([]string, error)
}

type UserService struct {
	userData    CSVUserGetter
	externalApi ExternalApiSetter
}

func NewUserService(userData CSVUserGetter, externalApi ExternalApiSetter) UserService {
	return UserService{userData: userData, externalApi: externalApi}
}

const USER string = "user"

func (us UserService) GetUserById(id int) (*model.User, error) {

	record, err := us.userData.GetRecordById(id, USER)

	if err != nil {
		return nil, err
	}
	user := model.User{
		ID:       id,
		Name:     record[2],
		LastName: record[3],
		Email:    record[4],
		Avatar:   record[5],
	}

	return &user, nil
}

func (us UserService) StoreNewInfo() (string, error) {
	user, err := us.externalApi.GetFromUrl()
	if err != nil {
		return "", err
	}

	us.userData.WriteRecord(user)
	return "ok", nil
}
