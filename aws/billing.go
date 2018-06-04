package aws

import (
  "time"
  "fmt"

  "github.com/aws/aws-sdk-go/aws"
  "github.com/aws/aws-sdk-go/aws/session"
  "github.com/aws/aws-sdk-go/service/cloudwatch"

  "github.com/hiro-kun/AwsBillingNotifyGo/conf"
)

func GetBilling() (map[string]interface{}, error) {
  estimatedCharges := make(map[string]interface{})

	sess, err := session.NewSession(&aws.Config{Region: aws.String(conf.Region)})
	if err != nil {
		return estimatedCharges, fmt.Errorf("session create error : %s", err)
	}
	svc := cloudwatch.New(sess)

	params := &cloudwatch.GetMetricStatisticsInput{
		Namespace:  aws.String(conf.Namespace),
		MetricName: aws.String(conf.MetricName),
    // 取得する範囲(秒単位)
    Period:     aws.Int64(14400),
    // 取得開始時刻※下記だと3日前から
    StartTime:  aws.Time(time.Now().AddDate(0, 0, -3)),
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
		return estimatedCharges, fmt.Errorf("get metrics error : %s", err)
	}

  tmpMaximum := 0.0
  tmpTimeStamp := time.Now()

  // コストが一番高いデータを取得
  for _, value := range resp.Datapoints {
    if tmpMaximum < *value.Maximum {
      tmpMaximum = *value.Maximum
      tmpTimeStamp = *value.Timestamp
    }
  }

  estimatedCharges["estimatePrice"] = tmpMaximum
  estimatedCharges["timestamp"] = tmpTimeStamp

  return estimatedCharges, nil
}
