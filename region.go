package main

import (
	"errors"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

// Return all valid regions
func GetRegions() ([]string, error) {
	// The DescribeRegions call does not work without region
	svc := ec2.New(session.New(), aws.NewConfig().WithRegion("us-east-1"))

	resp, err := svc.DescribeRegions(&ec2.DescribeRegionsInput{})

	if err != nil {
		return nil, err
	}

	if resp.Regions != nil {
		regionNames := make([]string, 0, len(resp.Regions))

		for _, region := range resp.Regions {
			regionNames = append(regionNames, *region.RegionName)
		}

		return regionNames, nil
	}

	return nil, errors.New("No valid regions found")
}
