package servicecontext

import v1 "k8s.io/api/core/v1"

type ServiceContext struct {
	ServiceName   string
	LabelSelector map[string]string
	Ports         []v1.ServicePort
}
