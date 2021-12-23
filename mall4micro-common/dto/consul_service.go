package dto

type Check struct {
	DeregisterCriticalServiceAfter string   `json:"DeregisterCriticalServiceAfter" yaml:"deregister_critical_service_after"`
	Args                           []string `json:"Args" yaml:"args"`
	Interval                       string   `json:"Interval" yaml:"interval"`
	Timeout                        string   `json:"Timeout" yaml:"timeout"`
}

type Weights struct {
	Passing int `json:"Passing" yaml:"passing"`
	Warning int `json:"Warning" yaml:"warning"`
}

type ConsulServiceDTO struct {
	ID                string                 `json:"ID"`
	Name              string                 `json:"Name"`
	Tags              []string               `json:"Tags"`
	Address           string                 `json:"Address"`
	Port              int                    `json:"Port"`
	ServiceMeta       map[string]interface{} `json:"Meta"`
	EnableTagOverride bool                   `json:"EnableTagOverride"`
	ServiceCheck      Check                  `json:"Check"`
	ServiceWeights    Weights                `json:"Weights"`
}
