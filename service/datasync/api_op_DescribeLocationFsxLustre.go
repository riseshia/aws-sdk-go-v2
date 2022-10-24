// Code generated by smithy-go-codegen DO NOT EDIT.

package datasync

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Provides details about how an DataSync location for an Amazon FSx for Lustre
// file system is configured.
func (c *Client) DescribeLocationFsxLustre(ctx context.Context, params *DescribeLocationFsxLustreInput, optFns ...func(*Options)) (*DescribeLocationFsxLustreOutput, error) {
	if params == nil {
		params = &DescribeLocationFsxLustreInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeLocationFsxLustre", params, optFns, c.addOperationDescribeLocationFsxLustreMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeLocationFsxLustreOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeLocationFsxLustreInput struct {

	// The Amazon Resource Name (ARN) of the FSx for Lustre location to describe.
	//
	// This member is required.
	LocationArn *string

	noSmithyDocumentSerde
}

type DescribeLocationFsxLustreOutput struct {

	// The time that the FSx for Lustre location was created.
	CreationTime *time.Time

	// The Amazon Resource Name (ARN) of the FSx for Lustre location that was
	// described.
	LocationArn *string

	// The URI of the FSx for Lustre location that was described.
	LocationUri *string

	// The Amazon Resource Names (ARNs) of the security groups that are configured for
	// the FSx for Lustre file system.
	SecurityGroupArns []string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeLocationFsxLustreMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeLocationFsxLustre{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeLocationFsxLustre{}, middleware.After)
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
	if err = addOpDescribeLocationFsxLustreValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeLocationFsxLustre(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeLocationFsxLustre(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "datasync",
		OperationName: "DescribeLocationFsxLustre",
	}
}
