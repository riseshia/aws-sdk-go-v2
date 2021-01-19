// Code generated by smithy-go-codegen DO NOT EDIT.

package appmesh

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/appmesh/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a gateway route. A gateway route is attached to a virtual gateway and
// routes traffic to an existing virtual service. If a route matches a request, it
// can distribute traffic to a target virtual service. For more information about
// gateway routes, see Gateway routes
// (https://docs.aws.amazon.com/app-mesh/latest/userguide/gateway-routes.html).
func (c *Client) CreateGatewayRoute(ctx context.Context, params *CreateGatewayRouteInput, optFns ...func(*Options)) (*CreateGatewayRouteOutput, error) {
	if params == nil {
		params = &CreateGatewayRouteInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateGatewayRoute", params, optFns, addOperationCreateGatewayRouteMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateGatewayRouteOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateGatewayRouteInput struct {

	// The name to use for the gateway route.
	//
	// This member is required.
	GatewayRouteName *string

	// The name of the service mesh to create the gateway route in.
	//
	// This member is required.
	MeshName *string

	// The gateway route specification to apply.
	//
	// This member is required.
	Spec *types.GatewayRouteSpec

	// The name of the virtual gateway to associate the gateway route with. If the
	// virtual gateway is in a shared mesh, then you must be the owner of the virtual
	// gateway resource.
	//
	// This member is required.
	VirtualGatewayName *string

	// Unique, case-sensitive identifier that you provide to ensure the idempotency of
	// the request. Up to 36 letters, numbers, hyphens, and underscores are allowed.
	ClientToken *string

	// The AWS IAM account ID of the service mesh owner. If the account ID is not your
	// own, then the account that you specify must share the mesh with your account
	// before you can create the resource in the service mesh. For more information
	// about mesh sharing, see Working with shared meshes
	// (https://docs.aws.amazon.com/app-mesh/latest/userguide/sharing.html).
	MeshOwner *string

	// Optional metadata that you can apply to the gateway route to assist with
	// categorization and organization. Each tag consists of a key and an optional
	// value, both of which you define. Tag keys can have a maximum character length of
	// 128 characters, and tag values can have a maximum length of 256 characters.
	Tags []types.TagRef
}

type CreateGatewayRouteOutput struct {

	// The full description of your gateway route following the create call.
	//
	// This member is required.
	GatewayRoute *types.GatewayRouteData

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateGatewayRouteMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateGatewayRoute{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateGatewayRoute{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpCreateGatewayRouteValidationMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}