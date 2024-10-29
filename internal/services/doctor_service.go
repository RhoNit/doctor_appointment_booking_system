package services

import (
	"context"

	"github.com/RhoNit/doctor_appointment_system/internal/models"
	"github.com/RhoNit/doctor_appointment_system/internal/repositories"
)

type DoctorService struct {
	docRepo *repositories.DoctorRepository
}

func NewDoctorService(repo *repositories.DoctorRepository) *DoctorService {
	return &DoctorService{
		docRepo: repo,
	}
}

func (s *DoctorService) CreateDoctor(ctx context.Context, doctor *models.Doctor) error {
	return s.docRepo.CreateDoctor(ctx, doctor)
}

func (s *DoctorService) GetDoctorProfileById(ctx context.Context, id int) (*models.Doctor, error) {
	return s.docRepo.GetDoctorById(ctx, id)
}
