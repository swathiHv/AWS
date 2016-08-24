package types

type Eip struct {
	AllocationId string
	PublicIp     string
	InstanceId   string
	Region       string
}

func NewEip(allocId, ip, instId, region string) Eip {
	return Eip{allocId, ip, instId, region}
}
