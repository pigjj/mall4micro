package dto

type Check struct {
	DeregisterCriticalServiceAfter string   `json:"DeregisterCriticalServiceAfter"`
	Args                           []string `json:"Args"`
	Interval                       string   `json:"Interval"`
	Timeout                        string   `json:"Timeout"`
}

type Weights struct {
	Passing int `json:"Passing"`
	Warning int `json:"Warning"`
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
