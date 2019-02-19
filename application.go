package wserest

import (
	"github.com/sebastien4/wse-rest-library-go/entity/application"
	"github.com/sebastien4/wse-rest-library-go/entity/application/helper"
	"github.com/sebastien4/wse-rest-library-go/entity/base"
)

// Application Operations
type Application struct {
	wowza
}

// WSEApps is struct for GetAll() applications
type WSEApps struct {
	ServerName   string   `json:"serverName"`
	Applications []WSEApp `json:"applications"`
}

// WSEApp is struct for GetAll() applications
type WSEApp struct {
	ID                   string `json:"id"`
	AppType              string `json:"appType"`
	HREF                 string `json:"href"`
	DRMEnabled           bool   `json:"drmEnabled"`
	DVREnabled           bool   `json:"dvrEnabled"`
	StreamTargetsEnabled bool   `json:"streamTargetsEnabled"`
	TranscoderEnabled    bool   `json:"transcoderEnabled"`
}

// NewApplication create Application object
func NewApplication(
	settings *helper.Settings,
	name string,
	appType string,
	readAccess string,
	writeAccess string,
	description string) *Application {
	if name == "" {
		name = "live"
	}
	if appType == "" {
		appType = "Live"
	}
	if readAccess == "" {
		readAccess = "*"
	}
	if writeAccess == "" {
		writeAccess = "*"
	}
	if description == "" {
		description = "*"
	}

	a := new(Application)
	a.init(settings)
	a.props["name"] = name
	a.props["appType"] = appType
	a.props["clientStreamReadAccess"] = readAccess
	a.props["clientStreamWriteAccess"] = writeAccess
	a.props["description"] = description
	a.baseURI = a.host() + "/servers/" + a.serverInstance() + "/vhosts/" + a.vHostInstance() + "/applications/" + name
	return a
}

func (a *Application) setParameters() {
	a.AddSkipParameter("name")
	a.AddSkipParameter("clientStreamReadAccess")
	a.AddSkipParameter("appType")
	a.AddSkipParameter("clientStreamWriteAccess")
	a.AddSkipParameter("description")
}

// Get retrieves the specified Application configuration
func (a *Application) Get() (map[string]interface{}, error) {
	a.setParameters()

	a.setRestURI(a.baseURI)

	return a.sendRequest(a.preparePropertiesForRequest(), []base.Entity{}, GET, "")
}

// GetAdvanced retrieves the specified advanced Application configuration
func (a *Application) GetAdvanced() (map[string]interface{}, error) {
	a.setParameters()

	a.setRestURI(a.baseURI + "/adv")

	return a.sendRequest(a.preparePropertiesForRequest(), []base.Entity{}, GET, "")
}

// GetAllOld retrieves the list of Applications
func (a *Application) GetAllOld() (map[string]interface{}, error) {
	a.setParameters()

	a.setRestURI(a.host() + "/servers/" + a.serverInstance() + "/vhosts/" + a.vHostInstance() + "/applications")

	return a.sendRequest(a.preparePropertiesForRequest(), []base.Entity{}, GET, "")
}

// GetAll retrieves the list of Applications
func (a *Application) GetAll() (WSEApps, error) {
	a.setParameters()

	a.setRestURI(a.host() + "/servers/" + a.serverInstance() + "/vhosts/" + a.vHostInstance() + "/applications")

	var r WSEApps
	err := a.sendRequestSeb(&r, a.preparePropertiesForRequest(), []base.Entity{}, GET, "")
	return r, err
}

// Create adds the specified Application configuration
func (a *Application) Create(
	streamConfig *application.StreamConfig,
	securityConfig *application.SecurityConfig,
	modules *application.Modules,
	dvrConfig *application.DvrConfig,
	transConfig *application.TranscoderConfig,
	drmConfig *application.DrmConfig,
) (map[string]interface{}, error) {
	a.setRestURI(a.baseURI)

	args := []base.Entity{}
	if streamConfig != nil {
		args = append(args, streamConfig)
	}
	if securityConfig != nil {
		args = append(args, securityConfig)
	}
	if modules != nil {
		args = append(args, modules)
	}
	if dvrConfig != nil {
		args = append(args, dvrConfig)
	}
	if transConfig != nil {
		args = append(args, transConfig)
	}
	if drmConfig != nil {
		args = append(args, drmConfig)
	}
	entities := a.getEntities(args, a.baseURI)

	return a.sendRequest(a.preparePropertiesForRequest(), entities, POST, "")
}

// Update updates the specified Application configuration
func (a *Application) Update(
	streamConfig *application.StreamConfig,
	securityConfig *application.SecurityConfig,
	modules *application.Modules,
	dvrConfig *application.DvrConfig,
	transConfig *application.TranscoderConfig,
	drmConfig *application.DrmConfig,
) (map[string]interface{}, error) {
	a.setRestURI(a.baseURI)

	args := []base.Entity{}
	if streamConfig != nil {
		args = append(args, streamConfig)
	}
	if securityConfig != nil {
		args = append(args, securityConfig)
	}
	if modules != nil {
		args = append(args, modules)
	}
	if dvrConfig != nil {
		args = append(args, dvrConfig)
	}
	if transConfig != nil {
		args = append(args, transConfig)
	}
	if drmConfig != nil {
		args = append(args, drmConfig)
	}
	entities := a.getEntities(args, a.baseURI)

	return a.sendRequest(a.preparePropertiesForRequest(), entities, PUT, "")
}

// UpdateAdvanced updates the specified advanced Application configuration
func (a *Application) UpdateAdvanced(advancedSettings *application.AdvancedSettings, modules *application.Modules) (map[string]interface{}, error) {
	entities := a.getEntities(nil, a.baseURI)
	props := make(map[string]interface{})
	props["advancedSettings"] = advancedSettings.AdvancedSettings
	props["modules"] = modules.ModuleList
	props["restURI"] = a.baseURI + "/adv"

	return a.sendRequest(props, entities, PUT, "")
}

// Remove deletes the specified Application configuration
func (a *Application) Remove() (map[string]interface{}, error) {
	a.setRestURI(a.baseURI)
	return a.sendRequest(a.preparePropertiesForRequest(), []base.Entity{}, DELETE, "")
}

// Name return name property
func (a *Application) Name() string {
	return a.props["name"].(string)
}
