// Code generated by smithy-go-codegen DO NOT EDIT.

package lakeformation

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lakeformation/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns the Lake Formation permissions for a specified table or database
// resource located at a path in Amazon S3. GetEffectivePermissionsForPath will not
// return databases and tables if the catalog is encrypted.
func (c *Client) GetEffectivePermissionsForPath(ctx context.Context, params *GetEffectivePermissionsForPathInput, optFns ...func(*Options)) (*GetEffectivePermissionsForPathOutput, error) {
	if params == nil {
		params = &GetEffectivePermissionsForPathInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetEffectivePermissionsForPath", params, optFns, c.addOperationGetEffectivePermissionsForPathMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetEffectivePermissionsForPathOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetEffectivePermissionsForPathInput struct {

	// The Amazon Resource Name (ARN) of the resource for which you want to get
	// permissions.
	//
	// This member is required.
	ResourceArn *string

	// The identifier for the Data Catalog. By default, the account ID. The Data
	// Catalog is the persistent metadata store. It contains database definitions,
	// table definitions, and other control information to manage your Lake Formation
	// environment.
	CatalogId *string

	// The maximum number of results to return.
	MaxResults *int32

	// A continuation token, if this is not the first call to retrieve this list.
	NextToken *string

	noSmithyDocumentSerde
}

type GetEffectivePermissionsForPathOutput struct {

	// A continuation token, if this is not the first call to retrieve this list.
	NextToken *string

	// A list of the permissions for the specified table or database resource located
	// at the path in Amazon S3.
	Permissions []types.PrincipalResourcePermissions

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetEffectivePermissionsForPathMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetEffectivePermissionsForPath{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetEffectivePermissionsForPath{}, middleware.After)
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
	if err = addOpGetEffectivePermissionsForPathValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetEffectivePermissionsForPath(options.Region), middleware.Before); err != nil {
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

// GetEffectivePermissionsForPathAPIClient is a client that implements the
// GetEffectivePermissionsForPath operation.
type GetEffectivePermissionsForPathAPIClient interface {
	GetEffectivePermissionsForPath(context.Context, *GetEffectivePermissionsForPathInput, ...func(*Options)) (*GetEffectivePermissionsForPathOutput, error)
}

var _ GetEffectivePermissionsForPathAPIClient = (*Client)(nil)

// GetEffectivePermissionsForPathPaginatorOptions is the paginator options for
// GetEffectivePermissionsForPath
type GetEffectivePermissionsForPathPaginatorOptions struct {
	// The maximum number of results to return.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetEffectivePermissionsForPathPaginator is a paginator for
// GetEffectivePermissionsForPath
type GetEffectivePermissionsForPathPaginator struct {
	options   GetEffectivePermissionsForPathPaginatorOptions
	client    GetEffectivePermissionsForPathAPIClient
	params    *GetEffectivePermissionsForPathInput
	nextToken *string
	firstPage bool
}

// NewGetEffectivePermissionsForPathPaginator returns a new
// GetEffectivePermissionsForPathPaginator
func NewGetEffectivePermissionsForPathPaginator(client GetEffectivePermissionsForPathAPIClient, params *GetEffectivePermissionsForPathInput, optFns ...func(*GetEffectivePermissionsForPathPaginatorOptions)) *GetEffectivePermissionsForPathPaginator {
	if params == nil {
		params = &GetEffectivePermissionsForPathInput{}
	}

	options := GetEffectivePermissionsForPathPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetEffectivePermissionsForPathPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetEffectivePermissionsForPathPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next GetEffectivePermissionsForPath page.
func (p *GetEffectivePermissionsForPathPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetEffectivePermissionsForPathOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxResults = limit

	result, err := p.client.GetEffectivePermissionsForPath(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken && prevToken != nil && p.nextToken != nil && *prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opGetEffectivePermissionsForPath(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lakeformation",
		OperationName: "GetEffectivePermissionsForPath",
	}
}
