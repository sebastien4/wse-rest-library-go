package application

import (
	"github.com/sebastien4/wse-rest-library-go/entity/application/helper"
	"github.com/sebastien4/wse-rest-library-go/entity/base"
)

type AdvancedSettings struct {
	base.EntityBase
	AdvancedSettings []helper.AdvancedSettingItem `json:"advancedSettings"`
}

func (a *AdvancedSettings) EntityName() string {
	return "advancedSettings"
}

func (a *AdvancedSettings) SetURI(baseURI string) {
	a.SetRestURI("")
}
