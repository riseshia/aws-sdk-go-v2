// Code generated by smithy-go-codegen DO NOT EDIT.

package iotfleetwise

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iotfleetwise/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a vehicle model (model manifest) that specifies signals (attributes,
// branches, sensors, and actuators). For more information, see Vehicle models
// (https://docs.aws.amazon.com/iot-fleetwise/latest/developerguide/vehicle-models.html)
// in the Amazon Web Services IoT FleetWise Developer Guide.
func (c *Client) CreateModelManifest(ctx context.Context, params *CreateModelManifestInput, optFns ...func(*Options)) (*CreateModelManifestOutput, error) {
	if params == nil {
		params = &CreateModelManifestInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateModelManifest", params, optFns, c.addOperationCreateModelManifestMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateModelManifestOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateModelManifestInput struct {

	// The name of the vehicle model to create.
	//
	// This member is required.
	Name *string

	// A list of nodes, which are a general abstraction of signals.
	//
	// This member is required.
	Nodes []string

	// The Amazon Resource Name (ARN) of a signal catalog.
	//
	// This member is required.
	SignalCatalogArn *string

	// A brief description of the vehicle model.
	Description *string

	// Metadata that can be used to manage the vehicle model.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateModelManifestOutput struct {

	// The ARN of the created vehicle model.
	//
	// This member is required.
	Arn *string

	// The name of the created vehicle model.
	//
	// This member is required.
	Name *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateModelManifestMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpCreateModelManifest{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpCreateModelManifest{}, middleware.After)
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
	if err = addOpCreateModelManifestValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateModelManifest(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateModelManifest(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotfleetwise",
		OperationName: "CreateModelManifest",
	}
}