// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Gets a description for the specified model group.
func (c *Client) DescribeModelPackageGroup(ctx context.Context, params *DescribeModelPackageGroupInput, optFns ...func(*Options)) (*DescribeModelPackageGroupOutput, error) {
	if params == nil {
		params = &DescribeModelPackageGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeModelPackageGroup", params, optFns, c.addOperationDescribeModelPackageGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeModelPackageGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeModelPackageGroupInput struct {

	// The name of the model group to describe.
	//
	// This member is required.
	ModelPackageGroupName *string

	noSmithyDocumentSerde
}

type DescribeModelPackageGroupOutput struct {

	// Information about the user who created or modified an experiment, trial, trial
	// component, lineage group, or project.
	//
	// This member is required.
	CreatedBy *types.UserContext

	// The time that the model group was created.
	//
	// This member is required.
	CreationTime *time.Time

	// The Amazon Resource Name (ARN) of the model group.
	//
	// This member is required.
	ModelPackageGroupArn *string

	// The name of the model group.
	//
	// This member is required.
	ModelPackageGroupName *string

	// The status of the model group.
	//
	// This member is required.
	ModelPackageGroupStatus types.ModelPackageGroupStatus

	// A description of the model group.
	ModelPackageGroupDescription *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeModelPackageGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeModelPackageGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeModelPackageGroup{}, middleware.After)
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
	if err = addOpDescribeModelPackageGroupValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeModelPackageGroup(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeModelPackageGroup(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "DescribeModelPackageGroup",
	}
}
