package stacks

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway"
	"github.com/aws/constructs-go/constructs/v10"
	"tufulat/functions"
	"tufulat/utils"
	"reflect"
	"runtime"
	"strings"
)

type ApartmentStack struct {
	Scope constructs.Construct
}

const apartmentFunctions = "apartment_functions"

func GetFunctionName(_func interface{}) string {
	_name := runtime.FuncForPC(reflect.ValueOf(_func).Pointer()).Name()
	funcName := strings.Split(_name, ".")
	return funcName[len(funcName)-1]
}

func (_stack ApartmentStack) ApartmentListFunction() awsapigateway.LambdaIntegration {
	_funcName := GetFunctionName(functions.ApartmentList)
	_handler := apartmentFunctions + "." + _funcName
	return utils.CreateLambdaFunction(_stack.Scope, _handler)
}

func (_stack ApartmentStack) ApartmentCreateFunction() awsapigateway.LambdaIntegration {
	_funcName := GetFunctionName(functions.ApartmentCreate)
	_handler := apartmentFunctions + "." + _funcName
	return utils.CreateLambdaFunction(_stack.Scope, _handler)
}

func (_stack ApartmentStack) ApartmentGetFunction() awsapigateway.LambdaIntegration {
	_funcName := GetFunctionName(functions.ApartmentGet)
	_handler := apartmentFunctions + "." + _funcName
	return utils.CreateLambdaFunction(_stack.Scope, _handler)
}

func (_stack ApartmentStack) ApartmentUpdateFunction() awsapigateway.LambdaIntegration {
	_funcName := GetFunctionName(functions.ApartmentUpdate)
	_handler := apartmentFunctions + "." + _funcName
	return utils.CreateLambdaFunction(_stack.Scope, _handler)
}

func (_stack ApartmentStack) ApartmentDeleteFunction() awsapigateway.LambdaIntegration {
	_funcName := GetFunctionName(functions.ApartmentDelete)
	_handler := apartmentFunctions + "." + _funcName
	return utils.CreateLambdaFunction(_stack.Scope, _handler)
}
