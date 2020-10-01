// Code generated by smithy-go-codegen DO NOT EDIT.

package chime

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/chime/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists the tags applied to an Amazon Chime SDK attendee resource.
func (c *Client) ListAttendeeTags(ctx context.Context, params *ListAttendeeTagsInput, optFns ...func(*Options)) (*ListAttendeeTagsOutput, error) {
	stack := middleware.NewStack("ListAttendeeTags", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpListAttendeeTagsMiddlewares(stack)
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
	addOpListAttendeeTagsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListAttendeeTags(options.Region), middleware.Before)
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
			OperationName: "ListAttendeeTags",
			Err:           err,
		}
	}
	out := result.(*ListAttendeeTagsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListAttendeeTagsInput struct {

	// The Amazon Chime SDK attendee ID.
	//
	// This member is required.
	AttendeeId *string

	// The Amazon Chime SDK meeting ID.
	//
	// This member is required.
	MeetingId *string
}

type ListAttendeeTagsOutput struct {

	// A list of tag key-value pairs.
	Tags []*types.Tag

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpListAttendeeTagsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpListAttendeeTags{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpListAttendeeTags{}, middleware.After)
}

func newServiceMetadataMiddleware_opListAttendeeTags(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "chime",
		OperationName: "ListAttendeeTags",
	}
}