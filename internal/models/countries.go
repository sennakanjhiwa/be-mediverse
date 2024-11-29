package models

import (
	"github.com/sev-2/raiden/pkg/db"
)

type Countries struct {
	db.ModelBase
	Id int64 `json:"id,omitempty" column:"name:id;type:bigint;primaryKey;autoIncrement;nullable:false"`
	Name string `json:"name,omitempty" column:"name:name;type:text;nullable:false"`

	// Table information
	Metadata string `json:"-" schema:"public" tableName:"countries" rlsEnable:"true" rlsForced:"false"`

	// Access control
	Acl string `json:"-" read:"" write:""`
}
