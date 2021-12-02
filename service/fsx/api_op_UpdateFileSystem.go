// Code generated by smithy-go-codegen DO NOT EDIT.

package fsx

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/fsx/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Use this operation to update the configuration of an existing Amazon FSx file
// system. You can update multiple properties in a single request. For Amazon FSx
// for Windows File Server file systems, you can update the following
// properties:
//
// * AuditLogConfiguration
//
// * AutomaticBackupRetentionDays
//
// *
// DailyAutomaticBackupStartTime
//
// * SelfManagedActiveDirectoryConfiguration
//
// *
// StorageCapacity
//
// * ThroughputCapacity
//
// * WeeklyMaintenanceStartTime
//
// For FSx for
// Lustre file systems, you can update the following properties:
//
// *
// AutoImportPolicy
//
// * AutomaticBackupRetentionDays
//
// *
// DailyAutomaticBackupStartTime
//
// * DataCompressionType
//
// * StorageCapacity
//
// *
// WeeklyMaintenanceStartTime
//
// For FSx for ONTAP file systems, you can update the
// following properties:
//
// * AutomaticBackupRetentionDays
//
// *
// DailyAutomaticBackupStartTime
//
// * FsxAdminPassword
//
// *
// WeeklyMaintenanceStartTime
//
// For the Amazon FSx for OpenZFS file systems, you can
// update the following properties:
//
// * AutomaticBackupRetentionDays
//
// *
// CopyTagsToBackups
//
// * CopyTagsToVolumes
//
// * DailyAutomaticBackupStartTime
//
// *
// DiskIopsConfiguration
//
// * ThroughputCapacity
//
// * WeeklyMaintenanceStartTime
func (c *Client) UpdateFileSystem(ctx context.Context, params *UpdateFileSystemInput, optFns ...func(*Options)) (*UpdateFileSystemOutput, error) {
	if params == nil {
		params = &UpdateFileSystemInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateFileSystem", params, optFns, c.addOperationUpdateFileSystemMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateFileSystemOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The request object for the UpdateFileSystem operation.
type UpdateFileSystemInput struct {

	// The ID of the file system that you are updating.
	//
	// This member is required.
	FileSystemId *string

	// A string of up to 64 ASCII characters that Amazon FSx uses to ensure idempotent
	// updates. This string is automatically filled on your behalf when you use the
	// Command Line Interface (CLI) or an Amazon Web Services SDK.
	ClientRequestToken *string

	// The configuration object for Amazon FSx for Lustre file systems used in the
	// UpdateFileSystem operation.
	LustreConfiguration *types.UpdateFileSystemLustreConfiguration

	// The configuration updates for an Amazon FSx for NetApp ONTAP file system.
	OntapConfiguration *types.UpdateFileSystemOntapConfiguration

	// The configuration updates for an Amazon FSx for OpenZFS file system.
	OpenZFSConfiguration *types.UpdateFileSystemOpenZFSConfiguration

	// Use this parameter to increase the storage capacity of an Amazon FSx for Windows
	// File Server or Amazon FSx for Lustre file system. Specifies the storage capacity
	// target value, in GiB, to increase the storage capacity for the file system that
	// you're updating. You can't make a storage capacity increase request if there is
	// an existing storage capacity increase request in progress. For Windows file
	// systems, the storage capacity target value must be at least 10 percent greater
	// than the current storage capacity value. To increase storage capacity, the file
	// system must have at least 16 MBps of throughput capacity. For Lustre file
	// systems, the storage capacity target value can be the following:
	//
	// * For
	// SCRATCH_2 and PERSISTENT_1 SSD deployment types, valid values are in multiples
	// of 2400 GiB. The value must be greater than the current storage capacity.
	//
	// * For
	// PERSISTENT HDD file systems, valid values are multiples of 6000 GiB for 12-MBps
	// throughput per TiB file systems and multiples of 1800 GiB for 40-MBps throughput
	// per TiB file systems. The values must be greater than the current storage
	// capacity.
	//
	// * For SCRATCH_1 file systems, you can't increase the storage
	// capacity.
	//
	// For OpenZFS file systems, the input/output operations per second
	// (IOPS) automatically scale with increases to the storage capacity if IOPS is
	// configured for automatic scaling. If the storage capacity increase would result
	// in less than 3 IOPS per GiB of storage, this operation returns an error. For
	// more information, see Managing storage capacity
	// (https://docs.aws.amazon.com/fsx/latest/WindowsGuide/managing-storage-capacity.html)
	// in the Amazon FSx for Windows File Server User Guide, Managing storage and
	// throughput capacity
	// (https://docs.aws.amazon.com/fsx/latest/LustreGuide/managing-storage-capacity.html)
	// in the Amazon FSx for Lustre User Guide, and Managing storage capacity
	// (https://docs.aws.amazon.com/fsx/latest/OpenZFSGuide/managing-storage-capacity.html)
	// in the Amazon FSx for OpenZFS User Guide.
	StorageCapacity *int32

	// The configuration updates for an Amazon FSx for Windows File Server file system.
	WindowsConfiguration *types.UpdateFileSystemWindowsConfiguration

	noSmithyDocumentSerde
}

// The response object for the UpdateFileSystem operation.
type UpdateFileSystemOutput struct {

	// A description of the file system that was updated.
	FileSystem *types.FileSystem

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateFileSystemMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateFileSystem{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateFileSystem{}, middleware.After)
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
	if err = addIdempotencyToken_opUpdateFileSystemMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpUpdateFileSystemValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateFileSystem(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpUpdateFileSystem struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpUpdateFileSystem) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpUpdateFileSystem) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*UpdateFileSystemInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *UpdateFileSystemInput ")
	}

	if input.ClientRequestToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientRequestToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opUpdateFileSystemMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpUpdateFileSystem{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opUpdateFileSystem(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "fsx",
		OperationName: "UpdateFileSystem",
	}
}
