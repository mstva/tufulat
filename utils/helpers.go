package utils

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"strings"
)

func CreateLambdaFunction(
	scope constructs.Construct,
	handler string,
) awsapigateway.LambdaIntegration {

	_handlerName := strings.Split(handler, ".")[1]

	_functionsPath := jsii.String("functions")
	_functionName := jsii.String(AppNameUpper + _handlerName + "Function")

	_functionProps := &awslambda.FunctionProps{
		FunctionName: _functionName,
		Code:         awslambda.Code_FromAsset(_functionsPath, nil),
		Runtime:      awslambda.Runtime_GO_1_X(),
		Handler:      &handler,
	}

	_function := awslambda.NewFunction(
		scope,
		_functionName,
		_functionProps,
	)

	return awsapigateway.NewLambdaIntegration(_function, nil)

}
