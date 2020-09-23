// Code generated by smithy-go-codegen DO NOT EDIT.

package storagegateway

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	"net/http"
	"time"
)

const ServiceID = "Storage Gateway"

// AWS Storage Gateway Service  <p>AWS Storage Gateway is the service that connects
// an on-premises software appliance with cloud-based storage to provide seamless
// and secure integration between an organization's on-premises IT environment and
// the AWS storage infrastructure. The service enables you to securely upload data
// to the AWS Cloud for cost effective backup and rapid disaster recovery.</p>
// <p>Use the following links to get started using the <i>AWS Storage Gateway
// Service API Reference</i>:</p> <ul> <li> <p> <a
// href="https://docs.aws.amazon.com/storagegateway/latest/userguide/AWSStorageGatewayAPI.html#AWSStorageGatewayHTTPRequestsHeaders">AWS
// Storage Gateway required request headers</a>: Describes the required headers
// that you must send with every POST request to AWS Storage Gateway.</p> </li>
// <li> <p> <a
// href="https://docs.aws.amazon.com/storagegateway/latest/userguide/AWSStorageGatewayAPI.html#AWSStorageGatewaySigningRequests">Signing
// requests</a>: AWS Storage Gateway requires that you authenticate every request
// you send; this topic describes how sign such a request.</p> </li> <li> <p> <a
// href="https://docs.aws.amazon.com/storagegateway/latest/userguide/AWSStorageGatewayAPI.html#APIErrorResponses">Error
// responses</a>: Provides reference information about AWS Storage Gateway
// errors.</p> </li> <li> <p> <a
// href="https://docs.aws.amazon.com/storagegateway/latest/APIReference/API_Operations.html">Operations
// in AWS Storage Gateway</a>: Contains detailed descriptions of all AWS Storage
// Gateway operations, their request parameters, response elements, possible
// errors, and examples of requests and responses.</p> </li> <li> <p> <a
// href="https://docs.aws.amazon.com/general/latest/gr/sg.html">AWS Storage Gateway
// endpoints and quotas:</a> Provides a list of each AWS Region and the endpoints
// available for use with AWS Storage Gateway.</p> </li> </ul> <note> <p>AWS
// Storage Gateway resource IDs are in uppercase. When you use these resource IDs
// with the Amazon EC2 API, EC2 expects resource IDs in lowercase. You must change
// your resource ID to lowercase to use it with the EC2 API. For example, in
// Storage Gateway the ID for a volume might be
// <code>vol-AA22BB012345DAF670</code>. When you use this ID with the EC2 API, you
// must change it to <code>vol-aa22bb012345daf670</code>. Otherwise, the EC2 API
// might not behave as expected.</p> </note> <important> <p>IDs for Storage Gateway
// volumes and Amazon EBS snapshots created from gateway volumes are changing to a
// longer format. Starting in December 2016, all new volumes and snapshots will be
// created with a 17-character string. Starting in April 2016, you will be able to
// use these longer IDs so you can test your systems with the new format. For more
// information, see <a href="http://aws.amazon.com/ec2/faqs/#longer-ids">Longer EC2
// and EBS resource IDs</a>.</p> <p>For example, a volume Amazon Resource Name
// (ARN) with the longer volume ID format looks like the following:</p> <p>
// <code>arn:aws:storagegateway:us-west-2:111122223333:gateway/sgw-12A3456B/volume/vol-1122AABBCCDDEEFFG</code>.</p>
// <p>A snapshot ID with the longer ID format looks like the following:
// <code>snap-78e226633445566ee</code>.</p> <p>For more information, see <a
// href="http://forums.aws.amazon.com/ann.jspa?annID=3557">Announcement: Heads-up –
// Longer AWS Storage Gateway volume and snapshot IDs coming in 2016</a>.</p>
// </important>
type Client struct {
	options Options
}

// New returns an initialized Client based on the functional options. Provide
// additional functional options to further configure the behavior of the client,
// such as changing the client's endpoint or adding custom middleware behavior.
func New(options Options, optFns ...func(*Options)) *Client {
	options = options.Copy()

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
	APIOptions []APIOptionFunc

	// The credentials object to use when signing requests.
	Credentials aws.CredentialsProvider

	// The endpoint options to be used when attempting to resolve an endpoint.
	EndpointOptions ResolverOptions

	// The service endpoint resolver.
	EndpointResolver EndpointResolver

	// Signature Version 4 (SigV4) Signer
	HTTPSignerV4 HTTPSignerV4

	// The region to send requests to. (Required)
	Region string

	// Retryer guides how HTTP requests should be retried in case of recoverable
	// failures. When nil the API client will use a default retryer.
	Retryer retry.Retryer

	// The HTTP client to invoke API calls with. Defaults to client's default HTTP
	// implementation if nil.
	HTTPClient HTTPClient
}

func (o Options) GetCredentials() aws.CredentialsProvider {
	return o.Credentials
}

func (o Options) GetEndpointOptions() ResolverOptions {
	return o.EndpointOptions
}

func (o Options) GetEndpointResolver() EndpointResolver {
	return o.EndpointResolver
}

func (o Options) GetHTTPSignerV4() HTTPSignerV4 {
	return o.HTTPSignerV4
}

func (o Options) GetRegion() string {
	return o.Region
}

func (o Options) GetRetryer() retry.Retryer {
	return o.Retryer
}

type HTTPClient interface {
	Do(*http.Request) (*http.Response, error)
}

// Copy creates a clone where the APIOptions list is deep copied.
func (o Options) Copy() Options {
	to := o
	to.APIOptions = make([]APIOptionFunc, len(o.APIOptions))
	copy(to.APIOptions, o.APIOptions)
	return to
}

type APIOptionFunc func(*middleware.Stack) error

// NewFromConfig returns a new client from the provided config.
func NewFromConfig(cfg aws.Config, optFns ...func(*Options)) *Client {
	opts := Options{
		Region:      cfg.Region,
		Retryer:     cfg.Retryer,
		HTTPClient:  cfg.HTTPClient,
		Credentials: cfg.Credentials,
	}
	return New(opts, optFns...)
}

func resolveHTTPClient(o *Options) {
	if o.HTTPClient != nil {
		return
	}
	o.HTTPClient = aws.NewBuildableHTTPClient()
}

func resolveRetryer(o *Options) {
	if o.Retryer != nil {
		return
	}
	o.Retryer = retry.NewStandard()
}

func addClientUserAgent(stack *middleware.Stack) {
	awsmiddleware.AddUserAgentKey("storagegateway")(stack)
}

func addHTTPSignerV4Middleware(stack *middleware.Stack, o Options) {
	stack.Finalize.Add(v4.NewSignHTTPRequestMiddleware(o.Credentials, o.HTTPSignerV4), middleware.After)
}

type HTTPSignerV4 interface {
	SignHTTP(ctx context.Context, credentials aws.Credentials, r *http.Request, payloadHash string, service string, region string, signingTime time.Time) error
}

func resolveHTTPSignerV4(o *Options) {
	if o.HTTPSignerV4 != nil {
		return
	}
	o.HTTPSignerV4 = v4.NewSigner()
}
