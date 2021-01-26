// Code generated by smithy-go-codegen DO NOT EDIT.

package iotwireless

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iotwireless/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets information about a wireless device.
func (c *Client) GetWirelessDevice(ctx context.Context, params *GetWirelessDeviceInput, optFns ...func(*Options)) (*GetWirelessDeviceOutput, error) {
	if params == nil {
		params = &GetWirelessDeviceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetWirelessDevice", params, optFns, addOperationGetWirelessDeviceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetWirelessDeviceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetWirelessDeviceInput struct {

	// The identifier of the wireless device to get.
	//
	// This member is required.
	Identifier *string

	// The type of identifier used in identifier.
	//
	// This member is required.
	IdentifierType types.WirelessDeviceIdType
}

type GetWirelessDeviceOutput struct {

	// The Amazon Resource Name of the resource.
	Arn *string

	// The description of the resource.
	Description *string

	// The name of the destination to which the device is assigned.
	DestinationName *string

	// The ID of the wireless device.
	Id *string

	// Information about the wireless device.
	LoRaWAN *types.LoRaWANDevice

	// The name of the resource.
	Name *string

	// Sidewalk device object.
	Sidewalk *types.SidewalkDevice

	// The ARN of the thing associated with the wireless device.
	ThingArn *string

	// The name of the thing associated with the wireless device. The value is empty if
	// a thing isn't associated with the device.
	ThingName *string

	// The wireless device type.
	Type types.WirelessDeviceType

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetWirelessDeviceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetWirelessDevice{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetWirelessDevice{}, middleware.After)
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
	if err = addOpGetWirelessDeviceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetWirelessDevice(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetWirelessDevice(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotwireless",
		OperationName: "GetWirelessDevice",
	}
}