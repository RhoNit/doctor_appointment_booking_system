package repositories

import (
	"context"

	"github.com/RhoNit/doctor_appointment_system/internal/models"
	"github.com/jackc/pgx/v5"
)

type DoctorRepository struct {
	pgxConn *pgx.Conn
}

func NewDoctorRepository(conn *pgx.Conn) *DoctorRepository {
	return &DoctorRepository{
		pgxConn: conn,
	}
}

func (r *DoctorRepository) CreateDoctor(ctx context.Context, doctor *models.Doctor) error {
	query := `INSERT INTO doctors (doctor_name, registration_number, speciality) VALUES ($1, $2, $3)`

	_, err := r.pgxConn.Exec(ctx, query, doctor.Name, doctor.RegNum, doctor.Speciality)
	if err != nil {
		return err
	}

	return nil
}

func (r *DoctorRepository) GetDoctorById(ctx context.Context, id int) (*models.Doctor, error) {
	var doctorProfile models.Doctor

	query := `SELECT doctor_id, doctor_name, registration_number, speciality FROM doctors WHERE doctor_id = $1`

	err := r.pgxConn.QueryRow(ctx, query, id).Scan(&doctorProfile.Id, &doctorProfile.Name, &doctorProfile.RegNum, &doctorProfile.Speciality)
	if err != nil {
		return nil, err
	}

	return &doctorProfile, nil
}
