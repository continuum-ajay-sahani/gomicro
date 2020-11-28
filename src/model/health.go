package model

import "practice/gomicro/src/entity"

// AppHealth return application health information
var AppHealth = func() entity.AppHealth {
	health := entity.AppHealth{}
	health.MemoryCache = true
	return health
}
