package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/sev-2/raiden/pkg/db"
)

type Schedules struct {
	db.ModelBase
	Id        uuid.UUID  `json:"id,omitempty" column:"name:id;type:uuid;primaryKey;nullable:false;default:gen_random_uuid()"`
	DoctorId  uuid.UUID  `json:"doctor_id,omitempty" column:"name:doctor_id;type:uuid;nullable:false"`
	Day       string     `json:"day,omitempty" column:"name:day;type:text;nullable:false"`
	StartTime time.Time  `json:"start_time,omitempty" column:"name:start_time;type:timestampz;nullable:false"`
	EndTime   time.Time  `json:"end_time,omitempty" column:"name:end_time;type:timestampz;nullable:false"`
	CreatedAt *time.Time `json:"created_at,omitempty" column:"name:created_at;type:timestampz;nullable;default:CURRENT_TIMESTAMP"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" column:"name:updated_at;type:timestampz;nullable;default:CURRENT_TIMESTAMP"`

	// Table information
	Metadata string `json:"-" schema:"public" tableName:"schedules" rlsEnable:"true" rlsForced:"false"`

	// Access control
	Acl string `json:"-" read:"" write:""`

	// Relations
	Doctor *Doctors `json:"doctor,omitempty" join:"joinType:hasOne;primaryKey:id;foreignKey:doctor_id"`
}
