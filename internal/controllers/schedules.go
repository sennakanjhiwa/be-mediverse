package controllers

import (
	"mediverse/internal/models"

	"github.com/sev-2/raiden"
)

type SchedulesDoctors struct {
	raiden.ControllerBase
	Http  string `path:"/schedules" type:"rest"`
	Model models.Schedules
}
