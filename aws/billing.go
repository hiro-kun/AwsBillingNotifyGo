package aws

import (
  "time"

  "github.com/aws/aws-sdk-go/aws"
  "github.com/aws/aws-sdk-go/aws/session"
  "github.com/aws/aws-sdk-go/service/cloudwatch"

  "github.com/hiro-kun/AwsBillingNotifyGo/conf"
)

func GetBilling() float64 {
	sess, err := session.NewSession(&aws.Config{Region: aws.String(conf.Region)})
	if err != nil {
		panic(err)
	}

	svc := cloudwatch.New(sess)

	params := &cloudwatch.GetMetricStatisticsInput{
		Namespace:  aws.String(conf.Namespace),
		MetricName: aws.String(conf.MetricName),
		Period:     aws.Int64(21600),
		StartTime:  aws.Time(time.Now().Add(time.Duration(21600) * time.Second * -1)),
		EndTime:    aws.Time(time.Now()),
		Statistics: []*string{
			aws.String(cloudwatch.StatisticMaximum),
		},
		Dimensions: []*cloudwatch.Dimension{
			{
				Name:  aws.String(conf.DimensionName),
				Value: aws.String(conf.DimensionValue),
			},
		},
		Unit: aws.String(cloudwatch.StandardUnitNone),
	}

	resp, err := svc.GetMetricStatistics(params)
	if err != nil {
		panic(err)
	}

	return float64(*resp.Datapoints[0].Maximum)
}
