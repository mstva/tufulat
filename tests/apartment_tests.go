package tests

import (
	"testing"

	_ "github.com/aws/aws-cdk-go/awscdk/v2"
	_ "github.com/aws/aws-cdk-go/awscdk/v2/assertions"
	_ "github.com/aws/jsii-runtime-go"
)

func TestTufulatStack(t *testing.T) {
	// GIVEN
	//app := awscdk.NewApp(nil)

	//// WHEN
	//stack := stacks.NewTufulatStack(app, "MyStack", nil)
	//
	//// THEN
	//template := assertions.Template_FromStack(stack)
	//
	//template.HasResourceProperties(jsii.String("AWS::SQS::Queue"), map[string]interface{}{
	//	"VisibilityTimeout": 300,
	//})
	//template.ResourceCountIs(jsii.String("AWS::SNS::Topic"), jsii.Number(1))
}
