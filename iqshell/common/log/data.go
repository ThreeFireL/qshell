package log

import (
	"encoding/json"
	"io"
	"os"

	"github.com/astaxie/beego/logs"
)

const (
	LevelAlert   Level = logs.LevelAlert
	LevelError   Level = logs.LevelError
	LevelWarning Level = logs.LevelWarning
	LevelInfo    Level = logs.LevelInformational
	LevelDebug   Level = logs.LevelDebug
)

type Level = int

type Config struct {
	Filename       string `json:"filename"`
	Level          int    `json:"level"`
	Daily          bool   `json:"daily"`
	MaxDays        int    `json:"maxdays"`
	StdOutColorful bool   `json:"color"`
	EnableStdout   bool   `json:"-"`
}

func (c *Config) ToJson() string {
	cfgBytes, _ := json.Marshal(c)
	return string(cfgBytes)
}

var sdtout io.Writer = os.Stdout
var sdterr io.Writer = os.Stderr

func SetStdout(o io.Writer) {
	sdtout = o
}

func SetStderr(e io.Writer) {
	sdterr = e
}
