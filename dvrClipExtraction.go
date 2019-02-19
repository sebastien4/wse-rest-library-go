package wserest

import (
	"strconv"
	"strings"
	"time"

	"github.com/sebastien4/wse-rest-library-go/entity/application/helper"
	"github.com/sebastien4/wse-rest-library-go/entity/base"
)

// DvrClipExtraction is DVR stores utility
type DvrClipExtraction struct {
	wowza
}

// WSEDVRStores is struct for GetAll() DVR stores
type WSEDVRStores struct {
	ServerName               string        `json:"serverName"`
	Version                  string        `json:"version"`
	DVRConverterStoreSummary []WSEDVRStore `json:"dvrconverterstoresummary"`
}

// WSEDVRStore is struct for GetAll() DVR stores
type WSEDVRStore struct {
	ID       string `json:"name"`
	Location string `json:"location"`
}

// WSEDVRConverter is struct for GetAll() DVR stores
type WSEDVRConverter struct {
	ID                string               `json:"dvrStoreName"`
	ServerName        string               `json:"serverName"`
	Version           string               `json:"version"`
	DVRConverterStore WSEDVRConverterStore `json:"DvrConverterStore"`
}

// WSEDVRConverterStore is struct for GetAll() DVR stores
type WSEDVRConverterStore struct {
	DVRStoreName        string                 `json:"dvrStoreName"`
	AudioAvailable      bool                   `json:"audioAvailable"`
	VideoAvailable      bool                   `json:"videoAvailable"`
	IsLive              bool                   `json:"isLive"`
	DVRStartTime        int64                  `json:"dvrStartTime"`
	DVREndTime          int64                  `json:"dvrEndTime"`
	Duration            int64                  `json:"duration"`
	UTCStart            int64                  `json:"utcStart"`
	UTCEnd              int64                  `json:"utcEnd"`
	OutputFilename      string                 `json:"outputFilename"`
	DVRConversionStatus WSEDVRConversionStatus `json:"conversionStatus"`
}

// WSEDVRConversionStatus is struct for GetAll() DVR stores
type WSEDVRConversionStatus struct {
	StoreName    string   `json:"storeName"`
	FileName     string   `json:"fileName"`
	State        string   `json:"state"` // The possible states are STOPPED, INITIALIZING, RUNNING, SUCCESSFUL, and ERROR.
	StatusCode   string   `json:"statusCode"`
	ErrorStrings []string `json:"errorStrings"`
	StartTime    int64    `json:"startTime"`
	EndTime      int64    `json:"endTime"`
	Duration     int64    `json:"duration"`
	CurrentChunk int64    `json:"currentChunk"`
	ChunkCount   int64    `json:"chunkCount"`
	FileSize     int64    `json:"fileSize"`
	FileDuration int64    `json:"fileDuration"`
}

// NewDvrClipExtraction creates DvrClipExtraction object
func NewDvrClipExtraction(settings *helper.Settings, appName string, appInstance string) *DvrClipExtraction {
	if appInstance == "" {
		appInstance = "_definst_"
	}

	d := new(DvrClipExtraction)
	d.init(settings)
	d.baseURI = d.host() + "/servers/" + d.serverInstance() + "/vhosts/" + d.vHostInstance() + "/applications/" + appName + "/instances/" + appInstance + "/dvrstores"

	return d
}

// Create creates a new DVR store
func (d *DvrClipExtraction) Create() (map[string]interface{}, error) {
	d.setRestURI(d.baseURI)
	response, err := d.sendRequest(d.preparePropertiesForRequest(), []base.Entity{}, POST, "")

	return response, err
}

// GetItemOld retrieves the information about a store/converter
func (d *DvrClipExtraction) GetItemOld(name string) (map[string]interface{}, error) {
	d.setRestURI(d.baseURI + "/" + name)

	return d.sendRequest(d.preparePropertiesForRequest(), []base.Entity{}, GET, "")
}

// GetItem retrieves the information about a store/converter
func (d *DvrClipExtraction) GetItem(name string) (WSEDVRConverter, error) {
	d.setRestURI(d.baseURI + "/" + name)

	var r WSEDVRConverter
	err := d.sendRequestSeb(&r, d.preparePropertiesForRequest(), []base.Entity{}, GET, "")
	return r, err
}

// ConvertGroup convert group
func (d *DvrClipExtraction) ConvertGroup(nameArr []string) (map[string]interface{}, error) {
	d.setNoParams()
	d.setRestURI(d.baseURI + "/actions/convert?dvrConverterStoreList=" + strings.Join(nameArr, ","))

	return d.sendRequest(d.preparePropertiesForRequest(), []base.Entity{}, PUT, "")
}

// Convert converts
func (d *DvrClipExtraction) Convert(name string, startTime int64, endTime int64, outputFolder, outputFileName string, debugEnabled bool) (map[string]interface{}, error) {
	d.setNoParams()
	query := ""

	if startTime != 0 {
		query += "dvrConverterStartTime=" + strconv.FormatInt(startTime, 10)
	}

	if endTime != 0 {
		if query != "" {
			query += "&"
		}
		query += "dvrConverterEndTime=" + strconv.FormatInt(endTime, 10)
	}

	if outputFolder != "" {
		if query != "" {
			query += "&"
		}
		query += "dvrConverterDefaultFileDestination=" + outputFolder
	}

	if outputFileName != "" {
		if query != "" {
			query += "&"
		}
		query += "dvrConverterOutputFilename=" + outputFileName
	}

	if query != "" {
		query += "&"
	}
	query += "dvrConverterDebugConversions=" + strconv.FormatBool(debugEnabled)

	if len(query) > 0 {
		query = "?" + query
	}

	d.setRestURI(d.baseURI + "/" + name + "/actions/convert" + query)

	return d.sendRequest(d.preparePropertiesForRequest(), []base.Entity{}, PUT, "")
}

// ClearCache clear cache
func (d *DvrClipExtraction) ClearCache() (map[string]interface{}, error) {
	d.setRestURI(d.baseURI + "/actions/expire")

	return d.sendRequest(d.preparePropertiesForRequest(), []base.Entity{}, PUT, "")
}

// DebugConversions converts
func (d *DvrClipExtraction) DebugConversions(name string) (map[string]interface{}, error) {
	d.setRestURI(d.baseURI + "/" + name + "/actions/convert?dvrConverterDebugConversions=true")

	return d.sendRequest(d.preparePropertiesForRequest(), []base.Entity{}, PUT, "")
}

// ConvertByDurationWithStartTime conver by duration with start time
func (d *DvrClipExtraction) ConvertByDurationWithStartTime(name string, startTime *time.Time, duration *time.Duration, outputFileName string) (map[string]interface{}, error) {
	d.setNoParams()
	query := ""
	if startTime != nil {
		query += "dvrConverterStartTime=" + strconv.FormatInt(startTime.Unix(), 10)
	}
	if duration != nil {
		if query != "" {
			query += "&"
		}
		query += "dvrConverterDuration=" + strconv.FormatInt(int64(*duration/time.Millisecond), 10)
	}
	if outputFileName != "" {
		if query != "" {
			query += "&"
		}
		query += "dvrConverterOutputFilename=" + outputFileName
	}
	if len(query) > 0 {
		query = "?" + query
	}

	d.setRestURI(d.baseURI + "/" + name + "/actions/convert" + query)

	return d.sendRequest(d.preparePropertiesForRequest(), []base.Entity{}, PUT, "")
}

// ConvertByDurationWithStartTimeSeb converts by duration with start time
func (d *DvrClipExtraction) ConvertByDurationWithStartTimeSeb(name string, startTime int64, duration int64, outputFileName string, debugEnabled bool) (map[string]interface{}, error) {
	d.setNoParams()
	query := ""

	if startTime != 0 {
		query += "dvrConverterStartTime=" + strconv.FormatInt(startTime, 10)
	}

	if duration != 0 {
		if query != "" {
			query += "&"
		}
		query += "dvrConverterDuration=" + strconv.FormatInt(duration, 10)
	}

	if outputFileName != "" {
		if query != "" {
			query += "&"
		}
		query += "dvrConverterOutputFilename=" + outputFileName
	}

	if query != "" {
		query += "&"
	}
	query += "dvrConverterDebugConversions=" + strconv.FormatBool(debugEnabled)

	if len(query) > 0 {
		query = "?" + query
	}

	d.setRestURI(d.baseURI + "/" + name + "/actions/convert" + query)

	return d.sendRequest(d.preparePropertiesForRequest(), []base.Entity{}, PUT, "")
}

// ConvertByDurationWithEndTime convert by duration with end time
func (d *DvrClipExtraction) ConvertByDurationWithEndTime(name string, endTime *time.Time, duration *time.Duration, outputFileName string) (map[string]interface{}, error) {
	d.setNoParams()
	query := ""
	if endTime != nil {
		query += "dvrConverterEndTime=" + strconv.FormatInt(endTime.Unix(), 10)
	}
	if duration != nil {
		if query != "" {
			query += "&"
		}
		query += "dvrConverterDuration=" + strconv.FormatInt(int64(*duration/time.Millisecond), 10)
	}
	if outputFileName != "" {
		if query != "" {
			query += "&"
		}
		query += "dvrConverterOutputFilename=" + outputFileName
	}
	if len(query) > 0 {
		query = "?" + query
	}

	d.setRestURI(d.baseURI + "/" + name + "/actions/convert" + query)

	return d.sendRequest(d.preparePropertiesForRequest(), []base.Entity{}, PUT, "")
}

// ConvertOld converts
func (d *DvrClipExtraction) ConvertOld(name string, startTime *time.Time, endTime *time.Time, outputFileName string) (map[string]interface{}, error) {
	d.setNoParams()
	query := ""
	if startTime != nil {
		query += "dvrConverterStartTime=" + strconv.FormatInt(startTime.Unix(), 10)
	}
	if endTime != nil {
		if query != "" {
			query += "&"
		}
		query += "dvrConverterEndTime=" + strconv.FormatInt(endTime.Unix(), 10)
	}
	if outputFileName != "" {
		if query != "" {
			query += "&"
		}
		query += "dvrConverterOutputFilename=" + outputFileName
	}
	if len(query) > 0 {
		query = "?" + query
	}

	d.setRestURI(d.baseURI + "/" + name + "/actions/convert" + query)

	return d.sendRequest(d.preparePropertiesForRequest(), []base.Entity{}, PUT, "")
}

// ConvertByDurationWithEndTimeSeb convert by duration with end time
func (d *DvrClipExtraction) ConvertByDurationWithEndTimeSeb(name string, endTime int64, duration int64, outputFileName string, debugEnabled bool) (map[string]interface{}, error) {
	d.setNoParams()
	query := ""

	if endTime != 0 {
		query += "dvrConverterEndTime=" + strconv.FormatInt(endTime, 10)
	}

	if duration != 0 {
		if query != "" {
			query += "&"
		}
		query += "dvrConverterDuration=" + strconv.FormatInt(duration, 10)
	}

	if outputFileName != "" {
		if query != "" {
			query += "&"
		}
		query += "dvrConverterOutputFilename=" + outputFileName
	}

	if query != "" {
		query += "&"
	}
	query += "dvrConverterDebugConversions=" + strconv.FormatBool(debugEnabled)

	if len(query) > 0 {
		query = "?" + query
	}

	d.setRestURI(d.baseURI + "/" + name + "/actions/convert" + query)

	return d.sendRequest(d.preparePropertiesForRequest(), []base.Entity{}, PUT, "")
}

// GetAllOld retrieves the list of DVR stores associated with this application instance
func (d *DvrClipExtraction) GetAllOld() (map[string]interface{}, error) {
	d.setNoParams()

	d.setRestURI(d.baseURI)

	return d.sendRequest(d.preparePropertiesForRequest(), []base.Entity{}, GET, "")
}

// GetAll retrieves the list of Applications
func (d *DvrClipExtraction) GetAll() (WSEDVRStores, error) {
	d.setNoParams()

	d.setRestURI(d.baseURI)

	var r WSEDVRStores
	err := d.sendRequestSeb(&r, d.preparePropertiesForRequest(), []base.Entity{}, GET, "")
	return r, err
}

func (d *DvrClipExtraction) setNoParams() {

}

// Remove delete DVR store
func (d *DvrClipExtraction) Remove(fileName string) (map[string]interface{}, error) {
	d.setNoParams()
	d.setRestURI(d.baseURI + "/" + fileName)

	return d.sendRequest(d.preparePropertiesForRequest(), []base.Entity{}, DELETE, "")
}
