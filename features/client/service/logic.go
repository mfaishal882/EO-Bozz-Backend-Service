package service

import (
	"capstone-alta1/features/client"
	"capstone-alta1/utils/helper"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"golang.org/x/crypto/bcrypt"
)

type clientService struct {
	clientRepository client.RepositoryInterface
	validate         *validator.Validate
}

func New(repo client.RepositoryInterface) client.ServiceInterface {
	return &clientService{
		clientRepository: repo,
		validate:         validator.New(),
	}
}

// Create implements user.ServiceInterface
func (service *clientService) Create(input client.Core, c echo.Context) (err error) {
	input.User.Role = "Client"
	// validasi input
	if errValidate := service.validate.Struct(input); errValidate != nil {
		return errValidate
	}

	// validasi email harus unik

	_, errFindEmail := service.clientRepository.FindUser(input.User.Email)

	if errFindEmail != nil && !strings.Contains(errFindEmail.Error(), "found") {
		return helper.ServiceErrorMsg(errFindEmail)
	}

	bytePass, errEncrypt := bcrypt.GenerateFromPassword([]byte(input.User.Password), 10)
	if errEncrypt != nil {
		log.Error(errEncrypt.Error())
		return err
	}

	input.User.Password = string(bytePass)

	errCreate := service.clientRepository.Create(input)
	if errCreate != nil {
		log.Error(errCreate.Error())
		return err
	}

	return nil
}

// GetAll implements user.ServiceInterface
func (service *clientService) GetAll(query string) (data []client.Core, err error) {
	if query == "" {
		data, err = service.clientRepository.GetAll()
	} else {
		data, err = service.clientRepository.GetAllWithSearch(query)
	}

	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	return data, err
}

func (service *clientService) GetById(id int) (data client.Core, err error) {
	data, err = service.clientRepository.GetById(id)
	if err != nil {
		log.Error(err.Error())
		return client.Core{}, err
	}

	return data, err

}

func (service *clientService) Update(input client.Core, id int, c echo.Context) error {
	if input.User.Password != "" {
		generate, _ := bcrypt.GenerateFromPassword([]byte(input.User.Password), 10)
		input.User.Password = string(generate)
	}
	err := service.clientRepository.Update(input, id)
	if err != nil {
		log.Error(err.Error())
		return err
	}

	return nil
}

func (service *clientService) Delete(id int) error {
	// proses
	err := service.clientRepository.Delete(id)
	if err != nil {
		log.Error(err.Error())
		return helper.ServiceErrorMsg(err)
	}
	return nil
}