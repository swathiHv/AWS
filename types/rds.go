package types

type Rds struct {
	Name          string
	Region        string
	Namespace     string
	Metric        string
	DimensionName string
	Stats         string
	Period        int64
	StatsVal      float64
}

func NewRds(name, region, ns, metric, dim, stats string, prd int64, statsVal float64) Rds {
	return Rds{name, region, ns, metric, dim, stats, prd, statsVal}
}
