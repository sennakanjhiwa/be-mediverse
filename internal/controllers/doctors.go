package controllers

import (
	"mediverse/internal/models"

	"github.com/sev-2/raiden"
)

type DoctorsController struct {
	raiden.ControllerBase
	Http  string `path:"/doctors" type:"rest"`
	Model models.Doctors
}
