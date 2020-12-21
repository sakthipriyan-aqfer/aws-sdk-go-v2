// Code generated by smithy-go-codegen DO NOT EDIT.

package s3

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	awshttp "github.com/aws/aws-sdk-go-v2/aws/transport/http"
	acceptencodingcust "github.com/aws/aws-sdk-go-v2/service/internal/accept-encoding"
	"github.com/aws/aws-sdk-go-v2/service/internal/s3shared"
	s3sharedconfig "github.com/aws/aws-sdk-go-v2/service/internal/s3shared/config"
	s3cust "github.com/aws/aws-sdk-go-v2/service/s3/internal/customizations"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/logging"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"net/http"
	"time"
)

const ServiceID = "S3"
const ServiceAPIVersion = "2006-03-01"

// Client provides the API client to make operations call for Amazon Simple Storage
// Service.
type Client struct {
	options Options
}

// New returns an initialized Client based on the functional options. Provide
// additional functional options to further configure the behavior of the client,
// such as changing the client's endpoint or adding custom middleware behavior.
func New(options Options, optFns ...func(*Options)) *Client {
	options = options.Copy()

	resolveDefaultLogger(&options)

	resolveRetryer(&options)

	resolveHTTPClient(&options)

	resolveHTTPSignerV4(&options)

	resolveDefaultEndpointConfiguration(&options)

	for _, fn := range optFns {
		fn(&options)
	}

	client := &Client{
		options: options,
	}

	return client
}

type Options struct {
	// Set of options to modify how an operation is invoked. These apply to all
	// operations invoked for this client. Use functional options on operation call to
	// modify this list for per operation behavior.
	APIOptions []func(*middleware.Stack) error

	// Configures the events that will be sent to the configured logger.
	ClientLogMode aws.ClientLogMode

	// The credentials object to use when signing requests.
	Credentials aws.CredentialsProvider

	// The endpoint options to be used when attempting to resolve an endpoint.
	EndpointOptions EndpointResolverOptions

	// The service endpoint resolver.
	EndpointResolver EndpointResolver

	// Signature Version 4 (SigV4) Signer
	HTTPSignerV4 HTTPSignerV4

	// The logger writer interface to write logging messages to.
	Logger logging.Logger

	// The region to send requests to. (Required)
	Region string

	// Retryer guides how HTTP requests should be retried in case of recoverable
	// failures. When nil the API client will use a default retryer.
	Retryer retry.Retryer

	// Allows you to enable arn region support for the service.
	UseARNRegion bool

	// Allows you to enable S3 Accelerate feature. All operations compatible with S3
	// Accelerate will use the accelerate endpoint for requests. Requests not
	// compatible will fall back to normal S3 requests. The bucket must be enabled for
	// accelerate to be used with S3 client with accelerate enabled. If the bucket is
	// not enabled for accelerate an error will be returned. The bucket name must be
	// DNS compatible to work with accelerate.
	UseAccelerate bool

	// Allows you to enable Dualstack endpoint support for the service.
	UseDualstack bool

	// Allows you to enable the client to use path-style addressing, i.e.,
	// https://s3.amazonaws.com/BUCKET/KEY. By default, the S3 client will use virtual
	// hosted bucket addressing when possible(https://BUCKET.s3.amazonaws.com/KEY).
	UsePathStyle bool

	// The HTTP client to invoke API calls with. Defaults to client's default HTTP
	// implementation if nil.
	HTTPClient HTTPClient
}

type HTTPClient interface {
	Do(*http.Request) (*http.Response, error)
}

// Copy creates a clone where the APIOptions list is deep copied.
func (o Options) Copy() Options {
	to := o
	to.APIOptions = make([]func(*middleware.Stack) error, len(o.APIOptions))
	copy(to.APIOptions, o.APIOptions)
	return to
}
func (c *Client) invokeOperation(ctx context.Context, opID string, params interface{}, optFns []func(*Options), stackFns ...func(*middleware.Stack, Options) error) (result interface{}, metadata middleware.Metadata, err error) {
	stack := middleware.NewStack(opID, smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}

	for _, fn := range stackFns {
		if err := fn(stack, options); err != nil {
			return nil, metadata, err
		}
	}

	for _, fn := range options.APIOptions {
		if err := fn(stack); err != nil {
			return nil, metadata, err
		}
	}

	handler := middleware.DecorateHandler(smithyhttp.NewClientHandler(options.HTTPClient), stack)
	result, metadata, err = handler.Handle(ctx, params)
	if err != nil {
		err = &smithy.OperationError{
			ServiceID:     ServiceID,
			OperationName: opID,
			Err:           err,
		}
	}
	return result, metadata, err
}

func resolveDefaultLogger(o *Options) {
	if o.Logger != nil {
		return
	}
	o.Logger = logging.Nop{}
}

func addSetLoggerMiddleware(stack *middleware.Stack, o Options) error {
	return middleware.AddSetLoggerMiddleware(stack, o.Logger)
}

// NewFromConfig returns a new client from the provided config.
func NewFromConfig(cfg aws.Config, optFns ...func(*Options)) *Client {
	opts := Options{
		Region:        cfg.Region,
		Retryer:       cfg.Retryer,
		HTTPClient:    cfg.HTTPClient,
		Credentials:   cfg.Credentials,
		APIOptions:    cfg.APIOptions,
		Logger:        cfg.Logger,
		ClientLogMode: cfg.ClientLogMode,
	}
	resolveAWSEndpointResolver(cfg, &opts)
	resolveClientConfig(cfg, &opts)
	return New(opts, optFns...)
}

func resolveHTTPClient(o *Options) {
	if o.HTTPClient != nil {
		return
	}
	o.HTTPClient = awshttp.NewBuildableClient()
}

func resolveRetryer(o *Options) {
	if o.Retryer != nil {
		return
	}
	o.Retryer = retry.NewStandard()
}

func resolveAWSEndpointResolver(cfg aws.Config, o *Options) {
	if cfg.EndpointResolver == nil {
		return
	}
	o.EndpointResolver = WithEndpointResolver(cfg.EndpointResolver, NewDefaultEndpointResolver())
}

func addClientUserAgent(stack *middleware.Stack) error {
	return awsmiddleware.AddUserAgentKey("s3")(stack)
}

func addHTTPSignerV4Middleware(stack *middleware.Stack, o Options) error {
	mw := v4.NewSignHTTPRequestMiddleware(v4.SignHTTPRequestMiddlewareOptions{
		CredentialsProvider: o.Credentials,
		Signer:              o.HTTPSignerV4,
		LogSigning:          o.ClientLogMode.IsSigning(),
	})
	return stack.Finalize.Add(mw, middleware.After)
}

type HTTPSignerV4 interface {
	SignHTTP(ctx context.Context, credentials aws.Credentials, r *http.Request, payloadHash string, service string, region string, signingTime time.Time, optFns ...func(*v4.SignerOptions)) error
}

func resolveHTTPSignerV4(o *Options) {
	if o.HTTPSignerV4 != nil {
		return
	}
	o.HTTPSignerV4 = newDefaultV4Signer(*o)
}

func newDefaultV4Signer(o Options) *v4.Signer {
	return v4.NewSigner(func(so *v4.SignerOptions) {
		so.Logger = o.Logger
		so.LogSigning = o.ClientLogMode.IsSigning()
		so.DisableURIPathEscaping = true
	})
}

func addRetryMiddlewares(stack *middleware.Stack, o Options) error {
	mo := retry.AddRetryMiddlewaresOptions{
		Retryer:          o.Retryer,
		LogRetryAttempts: o.ClientLogMode.IsRetries(),
	}
	return retry.AddRetryMiddlewares(stack, mo)
}

// resolves client config
func resolveClientConfig(cfg aws.Config, o *Options) error {
	if len(cfg.ConfigSources) == 0 {
		return nil
	}
	value, found, err := s3sharedconfig.ResolveUseARNRegion(context.Background(), cfg.ConfigSources)
	if err != nil {
		return err
	}
	if found {
		o.UseARNRegion = value
	}
	return nil
}

func addMetadataRetrieverMiddleware(stack *middleware.Stack) error {
	return s3shared.AddMetadataRetrieverMiddleware(stack)
}

// nopGetBucketAccessor is no-op accessor for operation that don't support bucket
// member as input
func nopGetBucketAccessor(input interface{}) (*string, bool) {
	return nil, false
}

func addResponseErrorMiddleware(stack *middleware.Stack) error {
	return s3shared.AddResponseErrorMiddleware(stack)
}

func disableAcceptEncodingGzip(stack *middleware.Stack) error {
	return acceptencodingcust.AddAcceptEncodingGzip(stack, acceptencodingcust.AddAcceptEncodingGzipOptions{})
}

// ResponseError provides the HTTP centric error type wrapping the underlying error
// with the HTTP response value and the deserialized RequestID.
type ResponseError interface {
	error

	ServiceHostID() string
	ServiceRequestID() string
}

var _ ResponseError = (*s3shared.ResponseError)(nil)

// GetHostIDMetadata retrieves the host id from middleware metadata returns host id
// as string along with a boolean indicating presence of hostId on middleware
// metadata.
func GetHostIDMetadata(metadata middleware.Metadata) (string, bool) {
	return s3shared.GetHostIDMetadata(metadata)
}

func addRequestResponseLogging(stack *middleware.Stack, o Options) error {
	return stack.Deserialize.Add(&smithyhttp.RequestResponseLogger{
		LogRequest:          o.ClientLogMode.IsRequest(),
		LogRequestWithBody:  o.ClientLogMode.IsRequestWithBody(),
		LogResponse:         o.ClientLogMode.IsResponse(),
		LogResponseWithBody: o.ClientLogMode.IsResponseWithBody(),
	}, middleware.After)
}

// HTTPPresignerV4 represents presigner interface used by presign url client
type HTTPPresignerV4 interface {
	PresignHTTP(
		ctx context.Context, credentials aws.Credentials, r *http.Request,
		payloadHash string, service string, region string, signingTime time.Time,
		optFns ...func(*v4.SignerOptions),
	) (url string, signedHeader http.Header, err error)
}

// PresignOptions represents the presign client options
type PresignOptions struct {

	// ClientOptions are list of functional options to mutate client options used by
	// the presign client.
	ClientOptions []func(*Options)

	// Presigner is the presigner used by the presign url client
	Presigner HTTPPresignerV4

	// Expires sets the expiration duration for the generated presign url. This should
	// be the duration in seconds the presigned URL should be considered valid for. If
	// not set or set to zero, presign url would default to expire after 900 seconds.
	Expires time.Duration
}

func (o PresignOptions) copy() PresignOptions {
	clientOptions := make([]func(*Options), len(o.ClientOptions))
	copy(clientOptions, o.ClientOptions)
	o.ClientOptions = clientOptions
	return o
}

// WithPresignClientFromClientOptions is a helper utility to retrieve a function
// that takes PresignOption as input
func WithPresignClientFromClientOptions(optFns ...func(*Options)) func(*PresignOptions) {
	return withPresignClientFromClientOptions(optFns).options
}

type withPresignClientFromClientOptions []func(*Options)

func (w withPresignClientFromClientOptions) options(o *PresignOptions) {
	o.ClientOptions = append(o.ClientOptions, w...)
}

// WithPresignExpires is a helper utility to append Expires value on presign
// options optional function
func WithPresignExpires(dur time.Duration) func(*PresignOptions) {
	return withPresignExpires(dur).options
}

type withPresignExpires time.Duration

func (w withPresignExpires) options(o *PresignOptions) {
	o.Expires = time.Duration(w)
}

// PresignClient represents the presign url client
type PresignClient struct {
	client  *Client
	options PresignOptions
}

// NewPresignClient generates a presign client using provided API Client and
// presign options
func NewPresignClient(c *Client, optFns ...func(*PresignOptions)) *PresignClient {
	var options PresignOptions
	for _, fn := range optFns {
		fn(&options)
	}
	if len(options.ClientOptions) != 0 {
		c = New(c.options, options.ClientOptions...)
	}

	if options.Presigner == nil {
		options.Presigner = newDefaultV4Signer(c.options)
	}

	return &PresignClient{
		client:  c,
		options: options,
	}
}

func withNopHTTPClientAPIOption(o *Options) {
	o.HTTPClient = smithyhttp.NopClient{}
}

type presignConverter PresignOptions

func (c presignConverter) convertToPresignMiddleware(stack *middleware.Stack, options Options) (err error) {
	stack.Finalize.Clear()
	stack.Deserialize.Clear()
	stack.Build.Remove((*awsmiddleware.ClientRequestID)(nil).ID())
	pmw := v4.NewPresignHTTPRequestMiddleware(v4.PresignHTTPRequestMiddlewareOptions{
		CredentialsProvider: options.Credentials,
		Presigner:           c.Presigner,
		LogSigning:          options.ClientLogMode.IsSigning(),
	})
	err = stack.Finalize.Add(pmw, middleware.After)
	if err != nil {
		return err
	}
	if c.Expires < 0 {
		return fmt.Errorf("presign URL duration must be 0 or greater, %v", c.Expires)
	}
	// add middleware to set expiration for s3 presigned url, if expiration is set to
	// 0, this middleware sets a default expiration of 900 seconds
	err = stack.Build.Add(&s3cust.AddExpiresOnPresignedURL{Expires: c.Expires}, middleware.After)
	if err != nil {
		return err
	}
	return nil
}
