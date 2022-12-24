package routes

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway"
	"github.com/aws/jsii-runtime-go"
	"tufulat/stacks"
	"tufulat/utils"
)

func ApartmentRoutes(api awsapigateway.RestApi, apartmentStack stacks.ApartmentStack) {

	apartments := api.Root().AddResource(jsii.String("apartments"), nil)
	apartment := apartments.AddResource(jsii.String("{apartment_id}"), nil)

	// api/apartments
	apartments.AddMethod(jsii.String(utils.GET), apartmentStack.ApartmentListFunction(), nil)
	apartments.AddMethod(jsii.String(utils.POST), apartmentStack.ApartmentCreateFunction(), nil)
	// api/apartments/{apartment_id}
	apartment.AddMethod(jsii.String(utils.GET), apartmentStack.ApartmentGetFunction(), nil)
	apartment.AddMethod(jsii.String(utils.DELETE), apartmentStack.ApartmentDeleteFunction(), nil)
	apartment.AddMethod(jsii.String(utils.PUT), apartmentStack.ApartmentUpdateFunction(), nil)

}
