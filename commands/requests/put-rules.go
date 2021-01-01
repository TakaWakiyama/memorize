package main

//"https: //docs.aws.amazon.com/sdk-for-go/api/service/cloudwatchevents/#CloudWatchEvents.PutTargets"
import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatchevents"

	"fmt"
)

func main() {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	// Create the cloudwatch events client
	svc := cloudwatchevents.New(sess)

	result, err := svc.PutRule(&cloudwatchevents.PutRuleInput{
		Description:        aws.String("Created By Golang"),
		Name:               aws.String("test_golang"),
		ScheduleExpression: aws.String("cron(0 20 * * ? *)"),
		State:              aws.String("DISABLED"),
	})
	fmt.Println("Rule ARN:", result.RuleArn)
}
