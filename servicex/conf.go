package servicex

import (
	"encoding/json"
	"fmt"
	"github.com/goools/tools/reflectx"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
)

type Initer interface {
	Init()
}

type Configuration struct {
	serviceName          string
	Feature              string
	projectRoot          string
	ShouldGenerateConfig bool
}

var config = &Configuration{}

func init() {
	config.Initialize()
}

func ProjectName() string {
	return config.ProjectName()
}

func SetServiceName(serviceName string, rootDir string) {
	config.serviceName = serviceName
	_, filename, _, _ := runtime.Caller(1)
	config.projectRoot = filepath.Join(filepath.Dir(filename), rootDir)
}

func ConfP(c interface{}) {
	tpe := reflect.TypeOf(c)
	if tpe.Kind() != reflect.Ptr {
		panic(fmt  .Errorf("ConfP pass ptr for setting value"))
	}

	_ = os.Setenv("PROJECT_NAME", config.ProjectName())

	config.MarshalFromLocal(c)

	triggerInitials(c)
}

func triggerInitials(c interface{}) {
	rv := reflectx.Indirect(reflect.ValueOf(c))
	for i := 0; i < rv.NumField(); i++ {
		value := rv.Field(i)
		if conf, ok := value.Interface().(Initer); ok {
			conf.Init()
		}
	}
}

func (conf *Configuration) ProjectName() string {
	if conf.Feature != "" {
		return conf.ServiceName() + "--" + conf.Feature
	}
	return conf.ServiceName()
}

func (conf *Configuration) ServiceName() string {
	return conf.serviceName
}

func (conf *Configuration) Prefix() string {
	return strings.ToUpper(strings.Replace(conf.serviceName, "-", "_", -1))
}

func (conf *Configuration) Initialize() {
	// TODO: DO NOTHING
}

func (conf *Configuration) MarshalFromLocal(c interface{}) {
	localConfPath := filepath.Join(conf.projectRoot, "./config/local.json")
	contents, err := ioutil.ReadFile(localConfPath)
	if err != nil {
		return
	}
	err = json.Unmarshal(contents, c)
	if err != nil {
		return
	}
}
