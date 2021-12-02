// Code generated by smithy-go-codegen DO NOT EDIT.

package shield

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/shield/types"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
)

type validateOpAssociateDRTLogBucket struct {
}

func (*validateOpAssociateDRTLogBucket) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpAssociateDRTLogBucket) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*AssociateDRTLogBucketInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpAssociateDRTLogBucketInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpAssociateDRTRole struct {
}

func (*validateOpAssociateDRTRole) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpAssociateDRTRole) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*AssociateDRTRoleInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpAssociateDRTRoleInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpAssociateHealthCheck struct {
}

func (*validateOpAssociateHealthCheck) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpAssociateHealthCheck) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*AssociateHealthCheckInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpAssociateHealthCheckInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpAssociateProactiveEngagementDetails struct {
}

func (*validateOpAssociateProactiveEngagementDetails) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpAssociateProactiveEngagementDetails) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*AssociateProactiveEngagementDetailsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpAssociateProactiveEngagementDetailsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateProtectionGroup struct {
}

func (*validateOpCreateProtectionGroup) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateProtectionGroup) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateProtectionGroupInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateProtectionGroupInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateProtection struct {
}

func (*validateOpCreateProtection) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateProtection) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateProtectionInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateProtectionInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteProtectionGroup struct {
}

func (*validateOpDeleteProtectionGroup) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteProtectionGroup) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteProtectionGroupInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteProtectionGroupInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteProtection struct {
}

func (*validateOpDeleteProtection) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteProtection) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteProtectionInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteProtectionInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDescribeAttack struct {
}

func (*validateOpDescribeAttack) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDescribeAttack) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DescribeAttackInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDescribeAttackInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDescribeProtectionGroup struct {
}

func (*validateOpDescribeProtectionGroup) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDescribeProtectionGroup) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DescribeProtectionGroupInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDescribeProtectionGroupInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDisableApplicationLayerAutomaticResponse struct {
}

func (*validateOpDisableApplicationLayerAutomaticResponse) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDisableApplicationLayerAutomaticResponse) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DisableApplicationLayerAutomaticResponseInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDisableApplicationLayerAutomaticResponseInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDisassociateDRTLogBucket struct {
}

func (*validateOpDisassociateDRTLogBucket) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDisassociateDRTLogBucket) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DisassociateDRTLogBucketInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDisassociateDRTLogBucketInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDisassociateHealthCheck struct {
}

func (*validateOpDisassociateHealthCheck) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDisassociateHealthCheck) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DisassociateHealthCheckInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDisassociateHealthCheckInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpEnableApplicationLayerAutomaticResponse struct {
}

func (*validateOpEnableApplicationLayerAutomaticResponse) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpEnableApplicationLayerAutomaticResponse) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*EnableApplicationLayerAutomaticResponseInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpEnableApplicationLayerAutomaticResponseInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListResourcesInProtectionGroup struct {
}

func (*validateOpListResourcesInProtectionGroup) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListResourcesInProtectionGroup) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListResourcesInProtectionGroupInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListResourcesInProtectionGroupInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListTagsForResource struct {
}

func (*validateOpListTagsForResource) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListTagsForResource) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListTagsForResourceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListTagsForResourceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpTagResource struct {
}

func (*validateOpTagResource) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpTagResource) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*TagResourceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpTagResourceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUntagResource struct {
}

func (*validateOpUntagResource) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUntagResource) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UntagResourceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUntagResourceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUpdateApplicationLayerAutomaticResponse struct {
}

func (*validateOpUpdateApplicationLayerAutomaticResponse) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateApplicationLayerAutomaticResponse) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateApplicationLayerAutomaticResponseInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateApplicationLayerAutomaticResponseInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUpdateEmergencyContactSettings struct {
}

func (*validateOpUpdateEmergencyContactSettings) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateEmergencyContactSettings) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateEmergencyContactSettingsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateEmergencyContactSettingsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUpdateProtectionGroup struct {
}

func (*validateOpUpdateProtectionGroup) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateProtectionGroup) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateProtectionGroupInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateProtectionGroupInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

func addOpAssociateDRTLogBucketValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpAssociateDRTLogBucket{}, middleware.After)
}

func addOpAssociateDRTRoleValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpAssociateDRTRole{}, middleware.After)
}

func addOpAssociateHealthCheckValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpAssociateHealthCheck{}, middleware.After)
}

func addOpAssociateProactiveEngagementDetailsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpAssociateProactiveEngagementDetails{}, middleware.After)
}

func addOpCreateProtectionGroupValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateProtectionGroup{}, middleware.After)
}

func addOpCreateProtectionValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateProtection{}, middleware.After)
}

func addOpDeleteProtectionGroupValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteProtectionGroup{}, middleware.After)
}

func addOpDeleteProtectionValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteProtection{}, middleware.After)
}

func addOpDescribeAttackValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDescribeAttack{}, middleware.After)
}

func addOpDescribeProtectionGroupValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDescribeProtectionGroup{}, middleware.After)
}

func addOpDisableApplicationLayerAutomaticResponseValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDisableApplicationLayerAutomaticResponse{}, middleware.After)
}

func addOpDisassociateDRTLogBucketValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDisassociateDRTLogBucket{}, middleware.After)
}

func addOpDisassociateHealthCheckValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDisassociateHealthCheck{}, middleware.After)
}

func addOpEnableApplicationLayerAutomaticResponseValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpEnableApplicationLayerAutomaticResponse{}, middleware.After)
}

func addOpListResourcesInProtectionGroupValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListResourcesInProtectionGroup{}, middleware.After)
}

func addOpListTagsForResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListTagsForResource{}, middleware.After)
}

func addOpTagResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpTagResource{}, middleware.After)
}

func addOpUntagResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUntagResource{}, middleware.After)
}

func addOpUpdateApplicationLayerAutomaticResponseValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateApplicationLayerAutomaticResponse{}, middleware.After)
}

func addOpUpdateEmergencyContactSettingsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateEmergencyContactSettings{}, middleware.After)
}

func addOpUpdateProtectionGroupValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateProtectionGroup{}, middleware.After)
}

func validateEmergencyContact(v *types.EmergencyContact) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "EmergencyContact"}
	if v.EmailAddress == nil {
		invalidParams.Add(smithy.NewErrParamRequired("EmailAddress"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateEmergencyContactList(v []types.EmergencyContact) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "EmergencyContactList"}
	for i := range v {
		if err := validateEmergencyContact(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpAssociateDRTLogBucketInput(v *AssociateDRTLogBucketInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "AssociateDRTLogBucketInput"}
	if v.LogBucket == nil {
		invalidParams.Add(smithy.NewErrParamRequired("LogBucket"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpAssociateDRTRoleInput(v *AssociateDRTRoleInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "AssociateDRTRoleInput"}
	if v.RoleArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("RoleArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpAssociateHealthCheckInput(v *AssociateHealthCheckInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "AssociateHealthCheckInput"}
	if v.ProtectionId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ProtectionId"))
	}
	if v.HealthCheckArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("HealthCheckArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpAssociateProactiveEngagementDetailsInput(v *AssociateProactiveEngagementDetailsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "AssociateProactiveEngagementDetailsInput"}
	if v.EmergencyContactList == nil {
		invalidParams.Add(smithy.NewErrParamRequired("EmergencyContactList"))
	} else if v.EmergencyContactList != nil {
		if err := validateEmergencyContactList(v.EmergencyContactList); err != nil {
			invalidParams.AddNested("EmergencyContactList", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateProtectionGroupInput(v *CreateProtectionGroupInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateProtectionGroupInput"}
	if v.ProtectionGroupId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ProtectionGroupId"))
	}
	if len(v.Aggregation) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("Aggregation"))
	}
	if len(v.Pattern) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("Pattern"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateProtectionInput(v *CreateProtectionInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateProtectionInput"}
	if v.Name == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if v.ResourceArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteProtectionGroupInput(v *DeleteProtectionGroupInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteProtectionGroupInput"}
	if v.ProtectionGroupId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ProtectionGroupId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteProtectionInput(v *DeleteProtectionInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteProtectionInput"}
	if v.ProtectionId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ProtectionId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDescribeAttackInput(v *DescribeAttackInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DescribeAttackInput"}
	if v.AttackId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AttackId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDescribeProtectionGroupInput(v *DescribeProtectionGroupInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DescribeProtectionGroupInput"}
	if v.ProtectionGroupId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ProtectionGroupId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDisableApplicationLayerAutomaticResponseInput(v *DisableApplicationLayerAutomaticResponseInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DisableApplicationLayerAutomaticResponseInput"}
	if v.ResourceArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDisassociateDRTLogBucketInput(v *DisassociateDRTLogBucketInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DisassociateDRTLogBucketInput"}
	if v.LogBucket == nil {
		invalidParams.Add(smithy.NewErrParamRequired("LogBucket"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDisassociateHealthCheckInput(v *DisassociateHealthCheckInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DisassociateHealthCheckInput"}
	if v.ProtectionId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ProtectionId"))
	}
	if v.HealthCheckArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("HealthCheckArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpEnableApplicationLayerAutomaticResponseInput(v *EnableApplicationLayerAutomaticResponseInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "EnableApplicationLayerAutomaticResponseInput"}
	if v.ResourceArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceArn"))
	}
	if v.Action == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Action"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListResourcesInProtectionGroupInput(v *ListResourcesInProtectionGroupInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListResourcesInProtectionGroupInput"}
	if v.ProtectionGroupId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ProtectionGroupId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListTagsForResourceInput(v *ListTagsForResourceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListTagsForResourceInput"}
	if v.ResourceARN == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceARN"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpTagResourceInput(v *TagResourceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "TagResourceInput"}
	if v.ResourceARN == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceARN"))
	}
	if v.Tags == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Tags"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUntagResourceInput(v *UntagResourceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UntagResourceInput"}
	if v.ResourceARN == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceARN"))
	}
	if v.TagKeys == nil {
		invalidParams.Add(smithy.NewErrParamRequired("TagKeys"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUpdateApplicationLayerAutomaticResponseInput(v *UpdateApplicationLayerAutomaticResponseInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateApplicationLayerAutomaticResponseInput"}
	if v.ResourceArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceArn"))
	}
	if v.Action == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Action"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUpdateEmergencyContactSettingsInput(v *UpdateEmergencyContactSettingsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateEmergencyContactSettingsInput"}
	if v.EmergencyContactList != nil {
		if err := validateEmergencyContactList(v.EmergencyContactList); err != nil {
			invalidParams.AddNested("EmergencyContactList", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUpdateProtectionGroupInput(v *UpdateProtectionGroupInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateProtectionGroupInput"}
	if v.ProtectionGroupId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ProtectionGroupId"))
	}
	if len(v.Aggregation) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("Aggregation"))
	}
	if len(v.Pattern) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("Pattern"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}
