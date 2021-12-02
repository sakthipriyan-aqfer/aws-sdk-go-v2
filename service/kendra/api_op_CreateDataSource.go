// Code generated by smithy-go-codegen DO NOT EDIT.

package kendra

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/kendra/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a data source that you want to use with an Amazon Kendra index. You
// specify a name, data source connector type and description for your data source.
// You also specify configuration information for the data source connector.
// CreateDataSource is a synchronous operation. The operation returns 200 if the
// data source was successfully created. Otherwise, an exception is raised. Amazon
// S3 and custom
// (https://docs.aws.amazon.com/kendra/latest/dg/data-source-custom.html) data
// sources are the only supported data sources in the Amazon Web Services GovCloud
// (US-West) region.
func (c *Client) CreateDataSource(ctx context.Context, params *CreateDataSourceInput, optFns ...func(*Options)) (*CreateDataSourceOutput, error) {
	if params == nil {
		params = &CreateDataSourceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateDataSource", params, optFns, c.addOperationCreateDataSourceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateDataSourceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateDataSourceInput struct {

	// The identifier of the index that should be associated with this data source.
	//
	// This member is required.
	IndexId *string

	// A unique name for the data source. A data source name can't be changed without
	// deleting and recreating the data source.
	//
	// This member is required.
	Name *string

	// The type of repository that contains the data source.
	//
	// This member is required.
	Type types.DataSourceType

	// A token that you provide to identify the request to create a data source.
	// Multiple calls to the CreateDataSource operation with the same client token will
	// create only one data source.
	ClientToken *string

	// The connector configuration information that is required to access the
	// repository. You can't specify the Configuration parameter when the Type
	// parameter is set to CUSTOM. If you do, you receive a ValidationException
	// exception. The Configuration parameter is required for all other data sources.
	Configuration *types.DataSourceConfiguration

	// Configuration information for altering document metadata and content during the
	// document ingestion process when you create a data source. For more information
	// on how to create, modify and delete document metadata, or make other content
	// alterations when you ingest documents into Amazon Kendra, see Customizing
	// document metadata during the ingestion process
	// (https://docs.aws.amazon.com/kendra/latest/dg/custom-document-enrichment.html).
	CustomDocumentEnrichmentConfiguration *types.CustomDocumentEnrichmentConfiguration

	// A description for the data source.
	Description *string

	// The code for a language. This allows you to support a language for all documents
	// when creating the data source. English is supported by default. For more
	// information on supported languages, including their codes, see Adding documents
	// in languages other than English
	// (https://docs.aws.amazon.com/kendra/latest/dg/in-adding-languages.html).
	LanguageCode *string

	// The Amazon Resource Name (ARN) of a role with permission to access the data
	// source. For more information, see IAM Roles for Amazon Kendra
	// (https://docs.aws.amazon.com/kendra/latest/dg/iam-roles.html). You can't specify
	// the RoleArn parameter when the Type parameter is set to CUSTOM. If you do, you
	// receive a ValidationException exception. The RoleArn parameter is required for
	// all other data sources.
	RoleArn *string

	// Sets the frequency that Amazon Kendra will check the documents in your
	// repository and update the index. If you don't set a schedule Amazon Kendra will
	// not periodically update the index. You can call the StartDataSourceSyncJob
	// operation to update the index. You can't specify the Schedule parameter when the
	// Type parameter is set to CUSTOM. If you do, you receive a ValidationException
	// exception.
	Schedule *string

	// A list of key-value pairs that identify the data source. You can use the tags to
	// identify and organize your resources and to control access to resources.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateDataSourceOutput struct {

	// A unique identifier for the data source.
	//
	// This member is required.
	Id *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateDataSourceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateDataSource{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateDataSource{}, middleware.After)
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
	if err = addIdempotencyToken_opCreateDataSourceMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateDataSourceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateDataSource(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateDataSource struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateDataSource) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateDataSource) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateDataSourceInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateDataSourceInput ")
	}

	if input.ClientToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreateDataSourceMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateDataSource{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateDataSource(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kendra",
		OperationName: "CreateDataSource",
	}
}
