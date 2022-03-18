package handlers

import (
	"log"
	"os"
	"process-data-volts-lambda/datavolts"
	"process-data-volts-lambda/signals"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
)

var (
	AWS_ACCESS_KEY_ID     = os.Getenv("AWS_ACCESS_KEY_ID")
	AWS_SECRET_ACCESS_KEY = os.Getenv("AWS_SECRET_ACCESS_KEY")
)

var sess *session.Session
var svc *dynamodb.DynamoDB

func init() {
	/* sess = session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	})) */

	sess = session.Must(session.NewSession(&aws.Config{
		Region:      aws.String("us-east-1"),
		Credentials: credentials.NewStaticCredentials(AWS_ACCESS_KEY_ID, AWS_SECRET_ACCESS_KEY, ""),
	}))

	svc = dynamodb.New(sess)
}

func HandleAPI() error {
	initialDate, _ := time.Parse(time.RFC3339, "2020-01-01T00:00:00.000Z")
	finalDate, _ := time.Parse(time.RFC3339, "2022-03-25T00:00:00.000Z")

	initialCondition := expression.Name("timestamp").GreaterThanEqual(expression.Value(initialDate))
	finalCondition := expression.Name("timestamp").LessThanEqual(expression.Value(finalDate))

	proj := expression.NamesList(
		expression.Name("rTensions"),
		expression.Name("sTensions"),
		expression.Name("tTensions"),
		expression.Name("rCurrents"),
		expression.Name("sCurrents"),
		expression.Name("tCurrents"),
	)

	expr, err := expression.
		NewBuilder().
		WithFilter(initialCondition).
		WithFilter(finalCondition).
		WithProjection(proj).
		Build()

	if err != nil {
		log.Println("expression problem err: ", err)
		return nil
	}

	// Build the query input parameters
	params := &dynamodb.ScanInput{
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		FilterExpression:          expr.Filter(),
		ProjectionExpression:      expr.Projection(),
		TableName:                 aws.String(datavolts.DataVoltsTableName),
	}

	// Make the DynamoDB Query API call
	result, err := svc.Scan(params)
	if err != nil {
		log.Printf("Query API call failed: %s", err)
		return nil
	}

	rTensionSignal := signals.New("rTension")
	// sTensionSignal := signals.New("sTension")
	// tTensionSignal := signals.New("tTension")
	// rCurrentSignal := signals.New("rCurrent")
	// sCurrentSignal := signals.New("sCurrent")
	// tCurrentSignal := signals.New("tCurrent")

	for _, item := range result.Items {
		dataVolts := datavolts.DataVolts{}

		err = dynamodbattribute.UnmarshalMap(item, &dataVolts)

		if err != nil {
			log.Printf("Got error unmarshalling: %s", err)
			continue
		}

		signal1 := signals.BuildSin(5, 60, 0)
		// signal2 := signals.BuildSin(5, 10, 0)
		// signal3 := signals.BuildSin(5, 30, 0)

		arr := []float64{}
		arr = append(arr, signal1.GetValues()...)
		// arr = append(arr, signal2.GetValues()...)
		// arr = append(arr, signal3.GetValues()...)

		rTensionSignal.AddValues(arr...)
		// sTensionSignal.AddValues(dataVolts.STensions...)
		// tTensionSignal.AddValues(dataVolts.TTensions...)
		// rCurrentSignal.AddValues(dataVolts.RCurrents...)
		// sCurrentSignal.AddValues(dataVolts.SCurrents...)
		// tCurrentSignal.AddValues(dataVolts.TCurrents...)
	}

	rTensionSignal.CalculateFftProperties()
	// sTensionSignal.CalculateFftProperties()
	// tTensionSignal.CalculateFftProperties()
	// rCurrentSignal.CalculateFftProperties()
	// sCurrentSignal.CalculateFftProperties()
	// tCurrentSignal.CalculateFftProperties()

	return nil
}
