package config

import (
	"os"
	"reflect"
	"regexp"
	
	"gopkg.in/yaml.v3"
)

var Cfg = cfg{}
var reVar = regexp.MustCompile(`^\${(\w+)}$`)

type cfg struct {
	Debug    bool     `yaml:"debug"`
	Port     string   `yaml:"port"`
	Context  ctx      `yaml:"context"`
	Timezone string   `yaml:"timezone"`
	Services services `yaml:"services"`
	Database database `yaml:"database"`
	Jwt      jwt      `yaml:"jwt"`
}

type ctx struct {
	Timeout int `yaml:"timeout"`
}

type jwt struct {
	PrivateKey          string `yaml:"private-key"`
	PublicKey           string `yaml:"public-key"`
	AccessPeriodMobile  int64  `yaml:"access-token-expire-period-mobile"`
	RefreshPeriodMobile int64  `yaml:"refresh-token-expire-period-mobile"`
	AccessPeriodBo      int64  `yaml:"access-token-expire-period-mobile"`
	RefreshPeriodBo     int64  `yaml:"refresh-token-expire-period-backoffice"`
}

type database struct {
	Provider string `yaml:"provider"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Name     string `yaml:"name"`
}

type services struct {
	KlikPromo klikPromo `yaml:"klik-promo"`
}

type klikPromo struct {
	Url                    string `yaml:"url"`
	ProcessScheduledStatus string `yaml:"process-scheduled-status"`
	ProcessExpiredStatus   string `yaml:"process-expired-status"`
}

func LoadEnv() {
	var file []byte
	
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	
	if os.Getenv("ENV") == "production" {
		file, err = os.ReadFile(wd + "/application.yaml")
		if err != nil {
			panic(err)
		}
	} else if os.Getenv("ENV") == "staging" {
		file, err = os.ReadFile(wd + "/application-staging.yaml")
		if err != nil {
			panic(err)
		}
	} else {
		file, err = os.ReadFile(wd + "/application-dev.yaml")
		if err != nil {
			panic(err)
		}
	}
	
	err = yaml.Unmarshal(file, &Cfg)
	if err != nil {
		panic(err)
	}
	
	_fromenv(&Cfg)
}

func _fromenv(v interface{}) {
	_reflectEnv(reflect.ValueOf(v).Elem())
}

func _reflectEnv(rv reflect.Value) {
	for i := 0; i < rv.NumField(); i++ {
		fv := rv.Field(i)
		
		if fv.Kind() == reflect.Ptr {
			fv = fv.Elem()
		}
		
		if fv.Kind() == reflect.Struct {
			_reflectEnv(fv)
			continue
		}
		
		if fv.Kind() == reflect.String {
			match := reVar.FindStringSubmatch(fv.String())
			if len(match) > 1 {
				fv.SetString(os.Getenv(match[1]))
			}
		}
	}
}
