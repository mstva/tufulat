package main

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"tufulat/routes"
	"tufulat/stacks"
	"tufulat/utils"
)

func NewTufulatStack(
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

	NewTufulatStack(app, "TufulatStack", nil)

	app.Synth(nil)
}
