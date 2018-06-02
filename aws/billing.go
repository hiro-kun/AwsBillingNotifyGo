package aws

import (
  "time"
  "fmt"

  "github.com/aws/aws-sdk-go/aws"
  "github.com/aws/aws-sdk-go/aws/session"
  "github.com/aws/aws-sdk-go/service/cloudwatch"

  "github.com/hiro-kun/AwsBillingNotifyGo/conf"
)

func GetBilling() (float64, int, error) {
  var metrics float64

	sess, err := session.NewSession(&aws.Config{Region: aws.String(conf.Region)})
	if err != nil {
		return metrics, conf.ExitCodeError, fmt.Errorf("session create error : %s", err)
	}

	svc := cloudwatch.New(sess)

	params := &cloudwatch.GetMetricStatisticsInput{
		Namespace:  aws.String(conf.Namespace),
		MetricName: aws.String(conf.MetricName),
    Period:     aws.Int64(21600),
    StartTime:  aws.Time(time.Now().Add(time.Duration(43200) * time.Second * -1)),
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
		return metrics, conf.ExitCodeError, fmt.Errorf("get metrics error : %s", err)
	}

	return float64(*resp.Datapoints[0].Maximum), conf.ExitCodeOk, nil
}
