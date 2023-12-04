package service

import (
	"github.com/flipped-aurora/gin-vue-admin/server/service/InfoQuire"
	"github.com/flipped-aurora/gin-vue-admin/server/service/example"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
)

type ServiceGroup struct {
	SystemServiceGroup    system.ServiceGroup
	ExampleServiceGroup   example.ServiceGroup
	InfoQuireServiceGroup InfoQuire.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
