package kodeec_api

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

func NewKodeecApiStack(
	scope constructs.Construct,
	id string,
	props *awscdk.StackProps,
) awscdk.Stack {

	stack := awscdk.NewStack(scope, &id, props)

	return stack
}

func main() {
	defer jsii.Close()

	app := awscdk.NewApp(nil)

	NewKodeecApiStack(app, "KodeecApiStack", nil)

	app.Synth(nil)
}
