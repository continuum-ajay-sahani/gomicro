package model

import (
	"practice/gomicro/src/entity"
	"practice/gomicro/src/persistent"
)

// AdminCreate create parking lot information into the system
var AdminCreate = func(ad entity.AdminCreate) error {
	return persistent.CreateLot(ad)
}
