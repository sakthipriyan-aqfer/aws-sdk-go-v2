// Code generated by smithy-go-codegen DO NOT EDIT.

package comprehend

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/comprehend/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Inspects text for named entities, and returns information about them. For more
// information, about named entities, see how-entities.
func (c *Client) DetectEntities(ctx context.Context, params *DetectEntitiesInput, optFns ...func(*Options)) (*DetectEntitiesOutput, error) {
	if params == nil {
		params = &DetectEntitiesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DetectEntities", params, optFns, c.addOperationDetectEntitiesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DetectEntitiesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DetectEntitiesInput struct {

	// A UTF-8 text string. Each string must contain fewer that 5,000 bytes of UTF-8
	// encoded characters.
	//
	// This member is required.
	Text *string

	// The Amazon Resource Name of an endpoint that is associated with a custom entity
	// recognition model. Provide an endpoint if you want to detect entities by using
	// your own custom model instead of the default model that is used by Amazon
	// Comprehend. If you specify an endpoint, Amazon Comprehend uses the language of
	// your custom model, and it ignores any language code that you provide in your
	// request. For information about endpoints, see Managing endpoints
	// (https://docs.aws.amazon.com/comprehend/latest/dg/manage-endpoints.html).
	EndpointArn *string

	// The language of the input documents. You can specify any of the primary
	// languages supported by Amazon Comprehend. All documents must be in the same
	// language. If your request includes the endpoint for a custom entity recognition
	// model, Amazon Comprehend uses the language of your custom model, and it ignores
	// any language code that you specify here.
	LanguageCode types.LanguageCode

	noSmithyDocumentSerde
}

type DetectEntitiesOutput struct {

	// A collection of entities identified in the input text. For each entity, the
	// response provides the entity text, entity type, where the entity text begins and
	// ends, and the level of confidence that Amazon Comprehend has in the detection.
	// If your request uses a custom entity recognition model, Amazon Comprehend
	// detects the entities that the model is trained to recognize. Otherwise, it
	// detects the default entity types. For a list of default entity types, see
	// how-entities.
	Entities []types.Entity

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDetectEntitiesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDetectEntities{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDetectEntities{}, middleware.After)
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
	if err = addOpDetectEntitiesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDetectEntities(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDetectEntities(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "comprehend",
		OperationName: "DetectEntities",
	}
}
