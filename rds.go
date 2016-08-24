package main

import (
	"AWS/types"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatch"
)

func getDay(month time.Month) int {

	switch month {

	case time.January, time.March, time.May, time.July, time.August, time.October, time.December:
		return 31
	case time.February:
		return 28
	default:
		return 30
	}
}

//MEtrics for the last 2 weeks
func getTimes() (time.Time, time.Time) {
	end := time.Now()
	st := time.Now()
	year, month, day := end.Date()
	if (day - 14) > 0 {
		st = time.Date(year, month, day-14, 12, 0, 0, 0, time.UTC)
	} else if (day - 14) == 0 {
		day1 := getDay(month - 1)
		st = time.Date(year, month-1, day1, 12, 0, 0, 0, time.UTC)
	} else {
		day1 := getDay(month - 1)
		st = time.Date(year, month-1, day1+(day-14), 12, 0, 0, 0, time.UTC)
	}

	return end, st
}

func GetRdss(rds types.Rds) ([]types.Rds, error) {

	regions, err := GetRegions()
	if err != nil {
		return []types.Rds{}, err
	}
	var rdss []types.Rds
	for _, region := range regions {
		svc := cloudwatch.New(session.New(), aws.NewConfig().WithRegion(region))
		rds.Region = region
		//metricName := "DatabaseConnections"
		//name := "DBInstanceIdentifier"
		dim := cloudwatch.DimensionFilter{Name: &rds.DimensionName}
		dims := []*cloudwatch.DimensionFilter{}
		dims = append(dims, &dim)
		//namespace := "AWS/RDS"
		input := cloudwatch.ListMetricsInput{MetricName: &rds.Metric, Namespace: &rds.Namespace, Dimensions: dims}
		output, err := svc.ListMetrics(&input)
		if err != nil {
			fmt.Println("Error occured : ", err.Error())
			return []types.Rds{}, err
		}
		for _, metric := range output.Metrics {
			//stats := "Average"
			//var period int64
			//period = 1209600
			end, st := getTimes()
			arrayIn := []*string{}
			arrayIn = append(arrayIn, &rds.Stats)
			in := cloudwatch.GetMetricStatisticsInput{Dimensions: metric.Dimensions, MetricName: &rds.Metric, Namespace: &rds.Namespace, Statistics: arrayIn, Period: &rds.Period, StartTime: &st, EndTime: &end}
			out, err1 := svc.GetMetricStatistics(&in)
			if err1 != nil {
				fmt.Println(err1.Error())
				return []types.Rds{}, err1
			}
			if len(out.Datapoints) == 0 {

				rds.Name = *metric.Dimensions[0].Value
				rdss = append(rdss, rds)
				continue
			}
			rds.Name = *metric.Dimensions[0].Value
			rds.StatsVal = *out.Datapoints[0].Average
			rdss = append(rdss, rds)
		}
	}
	return rdss, nil
}
