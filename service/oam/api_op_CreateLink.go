// Code generated by smithy-go-codegen DO NOT EDIT.

package oam

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/oam/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a link between a source account and a sink that you have created in a
// monitoring account. Before you create a link, you must create a sink in the
// monitoring account and create a sink policy in that account. The sink policy
// must permit the source account to link to it. You can grant permission to source
// accounts by granting permission to an entire organization or to individual
// accounts. For more information, see CreateSink
// (https://docs.aws.amazon.com/OAM/latest/APIReference/API_CreateSink.html) and
// PutSinkPolicy
// (https://docs.aws.amazon.com/OAM/latest/APIReference/API_PutSinkPolicy.html).
// Each monitoring account can be linked to as many as 100,000 source accounts.
// Each source account can be linked to as many as five monitoring accounts.
func (c *Client) CreateLink(ctx context.Context, params *CreateLinkInput, optFns ...func(*Options)) (*CreateLinkOutput, error) {
	if params == nil {
		params = &CreateLinkInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateLink", params, optFns, c.addOperationCreateLinkMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateLinkOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateLinkInput struct {

	// Specify a friendly human-readable name to use to identify this source account
	// when you are viewing data from it in the monitoring account. You can use a
	// custom label or use the following variables:
	//
	// * $AccountName is the name of the
	// account
	//
	// * $AccountEmail is the globally unique email address of the account
	//
	// *
	// $AccountEmailNoDomain is the email address of the account without the domain
	// name
	//
	// This member is required.
	LabelTemplate *string

	// An array of strings that define which types of data that the source account
	// shares with the monitoring account.
	//
	// This member is required.
	ResourceTypes []types.ResourceType

	// The ARN of the sink to use to create this link. You can use ListSinks
	// (https://docs.aws.amazon.com/OAM/latest/APIReference/API_ListSinks.html) to find
	// the ARNs of sinks. For more information about sinks, see CreateSink
	// (https://docs.aws.amazon.com/OAM/latest/APIReference/API_CreateSink.html).
	//
	// This member is required.
	SinkIdentifier *string

	// Assigns one or more tags (key-value pairs) to the link. Tags can help you
	// organize and categorize your resources. You can also use them to scope user
	// permissions by granting a user permission to access or change only resources
	// with certain tag values. For more information about using tags to control
	// access, see Controlling access to Amazon Web Services resources using tags
	// (https://docs.aws.amazon.com/IAM/latest/UserGuide/access_tags.html).
	Tags map[string]string

	noSmithyDocumentSerde
}

type CreateLinkOutput struct {

	// The ARN of the link that is newly created.
	Arn *string

	// The random ID string that Amazon Web Services generated as part of the link ARN.
	Id *string

	// The label that you assigned to this link. If the labelTemplate includes
	// variables, this field displays the variables resolved to their actual values.
	Label *string

	// The exact label template that you specified, with the variables not resolved.
	LabelTemplate *string

	// The resource types supported by this link.
	ResourceTypes []string

	// The ARN of the sink that is used for this link.
	SinkArn *string

	// The tags assigned to the link.
	Tags map[string]string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateLinkMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateLink{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateLink{}, middleware.After)
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
	if err = addOpCreateLinkValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateLink(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateLink(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "oam",
		OperationName: "CreateLink",
	}
}