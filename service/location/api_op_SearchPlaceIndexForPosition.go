// Code generated by smithy-go-codegen DO NOT EDIT.

package location

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/location/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Reverse geocodes a given coordinate and returns a legible address. Allows you to
// search for Places or points of interest near a given position.
func (c *Client) SearchPlaceIndexForPosition(ctx context.Context, params *SearchPlaceIndexForPositionInput, optFns ...func(*Options)) (*SearchPlaceIndexForPositionOutput, error) {
	if params == nil {
		params = &SearchPlaceIndexForPositionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SearchPlaceIndexForPosition", params, optFns, c.addOperationSearchPlaceIndexForPositionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SearchPlaceIndexForPositionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SearchPlaceIndexForPositionInput struct {

	// The name of the place index resource you want to use for the search.
	//
	// This member is required.
	IndexName *string

	// Specifies a coordinate for the query defined by a longitude, and latitude.
	//
	// *
	// The first position is the X coordinate, or longitude.
	//
	// * The second position is
	// the Y coordinate, or latitude.
	//
	// For example,
	// position=xLongitude&position=yLatitude .
	//
	// This member is required.
	Position []float64

	// An optional paramer. The maximum number of results returned per request. Default
	// value: 50
	MaxResults int32

	noSmithyDocumentSerde
}

type SearchPlaceIndexForPositionOutput struct {

	// Returns a list of Places closest to the specified position. Each result contains
	// additional information about the Places returned.
	//
	// This member is required.
	Results []types.SearchForPositionResult

	// Contains a summary of the request.
	//
	// This member is required.
	Summary *types.SearchPlaceIndexForPositionSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationSearchPlaceIndexForPositionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpSearchPlaceIndexForPosition{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpSearchPlaceIndexForPosition{}, middleware.After)
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
	if err = addOpSearchPlaceIndexForPositionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSearchPlaceIndexForPosition(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opSearchPlaceIndexForPosition(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "geo",
		OperationName: "SearchPlaceIndexForPosition",
	}
}