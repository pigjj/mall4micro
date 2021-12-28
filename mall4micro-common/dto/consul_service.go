package dto

type AddressInfo struct {
	Address string `json:"Address"`
	Port    int    `json:"Port"`
}

type Address struct {
	Lan AddressInfo `json:"Lan"`
	Wan AddressInfo `json:"Wan"`
}

type TransparentProxy struct {
	OutboundListenerPort int `json:"OutboundListenerPort"`
}

type Upstream struct {
	DestinationType string `json:"DestinationType"`
	DestinationName string `json:"DestinationName"`
	LocalBindPort   int    `json:"LocalBindPort"`
}

type Proxy struct {
	DestinationServiceName string                 `json:"DestinationServiceName"`
	DestinationServiceId   string                 `json:"DestinationServiceId"`
	LocalServiceAddress    string                 `json:"LocalServiceAddress"`
	LocalServicePort       int                    `json:"LocalServicePort"`
	LocalServiceSocketPath string                 `json:"LocalServiceSocketPath"`
	Mode                   string                 `json:"Mode"`
	TransparentProxy       TransparentProxy       `json:"TransparentProxy"`
	Config                 map[string]interface{} `json:"Config"`
	Upstreams              []Upstream             `json:"Upstreams"`
	MeshGateway            MeshGateway            `json:"MeshGateway"`
	Expose                 Expose                 `json:"Expose"`
}

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

type MeshGateway struct {
	Mode string `json:"Mode"`
}

type Path struct {
	Path          string `json:"Path"`
	LocalPathPort int    `json:"LocalPathPort"`
	ListenerPort  int    `json:"ListenerPort"`
	Protocol      string `json:"Protocol"`
}

type Expose struct {
	Check bool   `json:"Check"`
	Paths []Path `json:"Paths"`
}

type Connect struct {
	Native         bool                   `json:"Native"`
	SidecarService map[string]interface{} `json:"SidecarService"`
	Proxy          struct {
		Command []string               `json:"Command"`
		Config  map[string]interface{} `json:"Config"`
	}
}

type ConsulServiceDTO struct {
	ID                string                 `json:"ID"`
	Name              string                 `json:"Name"`
	Tags              []string               `json:"Tags"`
	Address           string                 `json:"Address"`
	ServiceMeta       map[string]interface{} `json:"Meta"`
	TaggedAddresses   Address                `json:"TaggedAddresses"`
	Port              int                    `json:"Port"`
	SocketPath        string                 `json:"SocketPath"`
	EnableTagOverride bool                   `json:"EnableTagOverride"`
	ServiceCheck      Check                  `json:"Check"`
	Kind              string                 `json:"kind"`
	ProxyDestination  string                 `json:"ProxyDestination"`
	Proxy             Proxy                  `json:"Proxy"`
	Connect           Connect                `json:"Connect"`
	ServiceWeights    Weights                `json:"Weights"`
	Token             string                 `json:"Token"`
	Namespace         string                 `json:"Namespace"`
}
