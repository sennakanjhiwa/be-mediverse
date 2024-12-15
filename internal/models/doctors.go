package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/sev-2/raiden/pkg/db"
)

type Doctors struct {
	db.ModelBase
	Id             uuid.UUID   `json:"id,omitempty" column:"name:id;type:uuid;primaryKey;nullable:false;default:gen_random_uuid()"`
	UserId         uuid.UUID   `json:"user_id,omitempty" column:"name:user_id;type:uuid;nullable:false;unique"`
	Name           string      `json:"name,omitempty" column:"name:name;type:varchar;nullable:false"`
	Specialization string      `json:"specialization,omitempty" column:"name:specialization;type:varchar;nullable:false"`
	ContactInfo    interface{} `json:"contact_info,omitempty" column:"name:contact_info;type:json;nullable"`
	CreatedAt      *time.Time  `json:"created_at,omitempty" column:"name:created_at;type:timestampz;nullable;default:CURRENT_TIMESTAMP"`
	UpdatedAt      *time.Time  `json:"updated_at,omitempty" column:"name:updated_at;type:timestampz;nullable;default:CURRENT_TIMESTAMP"`

	// Table information
	Metadata string `json:"-" schema:"public" tableName:"doctors" rlsEnable:"true" rlsForced:"false"`

	// Access control
	Acl string `json:"-" read:"" write:""`

	// Relations
	ScheduleDoctors []*Schedules `json:"schedule_doctors,omitempty" join:"joinType:hasMany;primaryKey:id;foreignKey:doctor_id"`
}
