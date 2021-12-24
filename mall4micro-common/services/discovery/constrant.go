package discovery

const (
	ServiceRegisterMethod   = "PUT"
	ServiceRegisterUrl      = "/agent/service/register?replace-existing-checks=true"
	ServiceDeRegisterMethod = "PUT"
	ServiceDeRegisterUrl    = "/agent/service/deregister"
)
