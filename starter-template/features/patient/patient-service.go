package patient

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/maddalax/mhtml/starter-template/database"
	"time"
)

type Patient struct {
	Name            string
	ReasonForVisit  string
	AppointmentDate time.Time
	LocationName    string
}

type Service struct {
	ctx *fiber.Ctx
}

func NewService(ctx *fiber.Ctx) *Service {
	return &Service{}
}

type CreatePatientRequest struct {
	Name           string
	ReasonForVisit string
	LocationName   string
}

func (c *Service) Create(request CreatePatientRequest) error {
	time.Sleep(time.Second)
	database.HSet("patients", uuid.New().String(), Patient{
		Name:            request.Name,
		ReasonForVisit:  request.ReasonForVisit,
		AppointmentDate: time.Now(),
		LocationName:    "New York",
	})
	return errors.New("error creating patient")
}

func (c *Service) List() ([]*Patient, error) {
	patients, err := database.HList[Patient]("patients")
	if err != nil {
		return nil, err
	}
	return patients, nil
}
