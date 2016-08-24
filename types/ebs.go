package types

type Tag struct {
	Key   string
	Value string
}

type Ebs struct {
	VolumeId string
	AZ       string
	Tags     []Tag
}

func NewEbs(volId, az string, tags []Tag) Ebs {
	return Ebs{volId, az, tags}
}

func NewTag(key, value string) Tag {
	return Tag{key, value}
}
