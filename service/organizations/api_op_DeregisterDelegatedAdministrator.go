// Code generated by smithy-go-codegen DO NOT EDIT.

package organizations

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Removes the specified member AWS account as a delegated administrator for the
// specified AWS service. Deregistering a delegated administrator can have
// unintended impacts on the functionality of the enabled AWS service. See the
// documentation for the enabled service before you deregister a delegated
// administrator so that you understand any potential impacts. You can run this
// action only for AWS services that support this feature. For a current list of
// services that support it, see the column Supports Delegated Administrator in the
// table at AWS Services that you can use with AWS Organizations
// (https://docs.aws.amazon.com/organizations/latest/userguide/orgs_integrated-services-list.html)
// in the AWS Organizations User Guide. This operation can be called only from the
// organization's master account.
func (c *Client) DeregisterDelegatedAdministrator(ctx context.Context, params *DeregisterDelegatedAdministratorInput, optFns ...func(*Options)) (*DeregisterDelegatedAdministratorOutput, error) {
	stack := middleware.NewStack("DeregisterDelegatedAdministrator", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDeregisterDelegatedAdministratorMiddlewares(stack)
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	AddResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	retry.AddRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	addOpDeregisterDelegatedAdministratorValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeregisterDelegatedAdministrator(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)

	for _, fn := range options.APIOptions {
		if err := fn(stack); err != nil {
			return nil, err
		}
	}
	handler := middleware.DecorateHandler(smithyhttp.NewClientHandler(options.HTTPClient), stack)
	result, metadata, err := handler.Handle(ctx, params)
	if err != nil {
		return nil, &smithy.OperationError{
			ServiceID:     ServiceID,
			OperationName: "DeregisterDelegatedAdministrator",
			Err:           err,
		}
	}
	out := result.(*DeregisterDelegatedAdministratorOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeregisterDelegatedAdministratorInput struct {

	// The account ID number of the member account in the organization that you want to
	// deregister as a delegated administrator.
	//
	// This member is required.
	AccountId *string

	// The service principal name of an AWS service for which the account is a
	// delegated administrator. Delegated administrator privileges are revoked for only
	// the specified AWS service from the member account. If the specified service is
	// the only service for which the member account is a delegated administrator, the
	// operation also revokes Organizations read action permissions.
	//
	// This member is required.
	ServicePrincipal *string
}

type DeregisterDelegatedAdministratorOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDeregisterDelegatedAdministratorMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDeregisterDelegatedAdministrator{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeregisterDelegatedAdministrator{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeregisterDelegatedAdministrator(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "organizations",
		OperationName: "DeregisterDelegatedAdministrator",
	}
}