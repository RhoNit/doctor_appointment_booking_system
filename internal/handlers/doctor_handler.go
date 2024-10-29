package handler

import (
	"net/http"
	"strconv"

	"github.com/RhoNit/doctor_appointment_system/internal/models"
	"github.com/RhoNit/doctor_appointment_system/internal/services"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type DoctorHandler struct {
	docService *services.DoctorService
	logger     *zap.Logger
}

func NewDoctorHandler(service *services.DoctorService, logger *zap.Logger) *DoctorHandler {
	return &DoctorHandler{
		docService: service,
		logger:     logger,
	}
}

func (h *DoctorHandler) CreateDoctorProfile(c echo.Context) error {
	var doctorProfile *models.Doctor

	if err := c.Bind(&doctorProfile); err != nil {
		h.logger.Error("Failed to bind JSON data into Go type", zap.Error(err))
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "failed to bind JSON data"})
	}

	if err := h.docService.CreateDoctor(c.Request().Context(), doctorProfile); err != nil {
		h.logger.Error("Failed to create Doctor's profile", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "failed to create doctor's profile"})
	}

	doctorResp := models.DoctorResponse{
		Name:       doctorProfile.Name,
		Speciality: doctorProfile.Speciality,
		RegNum:     doctorProfile.RegNum,
	}

	return c.JSON(http.StatusCreated, doctorResp)
}

func (h *DoctorHandler) GetDoctorProfile(c echo.Context) error {
	idParam := c.Param("doctor_id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		h.logger.Error("Invalid Request", zap.Error(err))
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid request"})
	}

	docProfile, err := h.docService.GetDoctorProfileById(c.Request().Context(), id)
	if err != nil {
		h.logger.Error("Error while fetching Doctor's profile", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "couldn't fetch doctor's profile"})
	}

	return c.JSON(http.StatusOK, docProfile)
}
