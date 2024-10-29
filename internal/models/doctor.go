package models

type Doctor struct {
	Id         string `json:"doctor_id"`
	Name       string `json:"doctor_name"`
	RegNum     string `json:"registration_number"`
	Speciality string `json:"speciality"`
}

type DoctorResponse struct {
	Name       string `json:"doctor_name"`
	RegNum     string `json:"registration_number"`
	Speciality string `json:"speciality"`
}
