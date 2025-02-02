// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cloudformation

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type CloudFormationTypeLoggingConfig struct {
	// Name of the CloudWatch Log Group where CloudFormation sends error logging information when invoking the type's handlers.
	LogGroupName string `pulumi:"logGroupName"`
	// Amazon Resource Name (ARN) of the IAM Role CloudFormation assumes when sending error logging information to CloudWatch Logs.
	LogRoleArn string `pulumi:"logRoleArn"`
}

// CloudFormationTypeLoggingConfigInput is an input type that accepts CloudFormationTypeLoggingConfigArgs and CloudFormationTypeLoggingConfigOutput values.
// You can construct a concrete instance of `CloudFormationTypeLoggingConfigInput` via:
//
//          CloudFormationTypeLoggingConfigArgs{...}
type CloudFormationTypeLoggingConfigInput interface {
	pulumi.Input

	ToCloudFormationTypeLoggingConfigOutput() CloudFormationTypeLoggingConfigOutput
	ToCloudFormationTypeLoggingConfigOutputWithContext(context.Context) CloudFormationTypeLoggingConfigOutput
}

type CloudFormationTypeLoggingConfigArgs struct {
	// Name of the CloudWatch Log Group where CloudFormation sends error logging information when invoking the type's handlers.
	LogGroupName pulumi.StringInput `pulumi:"logGroupName"`
	// Amazon Resource Name (ARN) of the IAM Role CloudFormation assumes when sending error logging information to CloudWatch Logs.
	LogRoleArn pulumi.StringInput `pulumi:"logRoleArn"`
}

func (CloudFormationTypeLoggingConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CloudFormationTypeLoggingConfig)(nil)).Elem()
}

func (i CloudFormationTypeLoggingConfigArgs) ToCloudFormationTypeLoggingConfigOutput() CloudFormationTypeLoggingConfigOutput {
	return i.ToCloudFormationTypeLoggingConfigOutputWithContext(context.Background())
}

func (i CloudFormationTypeLoggingConfigArgs) ToCloudFormationTypeLoggingConfigOutputWithContext(ctx context.Context) CloudFormationTypeLoggingConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudFormationTypeLoggingConfigOutput)
}

func (i CloudFormationTypeLoggingConfigArgs) ToCloudFormationTypeLoggingConfigPtrOutput() CloudFormationTypeLoggingConfigPtrOutput {
	return i.ToCloudFormationTypeLoggingConfigPtrOutputWithContext(context.Background())
}

func (i CloudFormationTypeLoggingConfigArgs) ToCloudFormationTypeLoggingConfigPtrOutputWithContext(ctx context.Context) CloudFormationTypeLoggingConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudFormationTypeLoggingConfigOutput).ToCloudFormationTypeLoggingConfigPtrOutputWithContext(ctx)
}

// CloudFormationTypeLoggingConfigPtrInput is an input type that accepts CloudFormationTypeLoggingConfigArgs, CloudFormationTypeLoggingConfigPtr and CloudFormationTypeLoggingConfigPtrOutput values.
// You can construct a concrete instance of `CloudFormationTypeLoggingConfigPtrInput` via:
//
//          CloudFormationTypeLoggingConfigArgs{...}
//
//  or:
//
//          nil
type CloudFormationTypeLoggingConfigPtrInput interface {
	pulumi.Input

	ToCloudFormationTypeLoggingConfigPtrOutput() CloudFormationTypeLoggingConfigPtrOutput
	ToCloudFormationTypeLoggingConfigPtrOutputWithContext(context.Context) CloudFormationTypeLoggingConfigPtrOutput
}

type cloudFormationTypeLoggingConfigPtrType CloudFormationTypeLoggingConfigArgs

func CloudFormationTypeLoggingConfigPtr(v *CloudFormationTypeLoggingConfigArgs) CloudFormationTypeLoggingConfigPtrInput {
	return (*cloudFormationTypeLoggingConfigPtrType)(v)
}

func (*cloudFormationTypeLoggingConfigPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudFormationTypeLoggingConfig)(nil)).Elem()
}

func (i *cloudFormationTypeLoggingConfigPtrType) ToCloudFormationTypeLoggingConfigPtrOutput() CloudFormationTypeLoggingConfigPtrOutput {
	return i.ToCloudFormationTypeLoggingConfigPtrOutputWithContext(context.Background())
}

func (i *cloudFormationTypeLoggingConfigPtrType) ToCloudFormationTypeLoggingConfigPtrOutputWithContext(ctx context.Context) CloudFormationTypeLoggingConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudFormationTypeLoggingConfigPtrOutput)
}

type CloudFormationTypeLoggingConfigOutput struct{ *pulumi.OutputState }

func (CloudFormationTypeLoggingConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CloudFormationTypeLoggingConfig)(nil)).Elem()
}

func (o CloudFormationTypeLoggingConfigOutput) ToCloudFormationTypeLoggingConfigOutput() CloudFormationTypeLoggingConfigOutput {
	return o
}

func (o CloudFormationTypeLoggingConfigOutput) ToCloudFormationTypeLoggingConfigOutputWithContext(ctx context.Context) CloudFormationTypeLoggingConfigOutput {
	return o
}

func (o CloudFormationTypeLoggingConfigOutput) ToCloudFormationTypeLoggingConfigPtrOutput() CloudFormationTypeLoggingConfigPtrOutput {
	return o.ToCloudFormationTypeLoggingConfigPtrOutputWithContext(context.Background())
}

func (o CloudFormationTypeLoggingConfigOutput) ToCloudFormationTypeLoggingConfigPtrOutputWithContext(ctx context.Context) CloudFormationTypeLoggingConfigPtrOutput {
	return o.ApplyT(func(v CloudFormationTypeLoggingConfig) *CloudFormationTypeLoggingConfig {
		return &v
	}).(CloudFormationTypeLoggingConfigPtrOutput)
}

// Name of the CloudWatch Log Group where CloudFormation sends error logging information when invoking the type's handlers.
func (o CloudFormationTypeLoggingConfigOutput) LogGroupName() pulumi.StringOutput {
	return o.ApplyT(func(v CloudFormationTypeLoggingConfig) string { return v.LogGroupName }).(pulumi.StringOutput)
}

// Amazon Resource Name (ARN) of the IAM Role CloudFormation assumes when sending error logging information to CloudWatch Logs.
func (o CloudFormationTypeLoggingConfigOutput) LogRoleArn() pulumi.StringOutput {
	return o.ApplyT(func(v CloudFormationTypeLoggingConfig) string { return v.LogRoleArn }).(pulumi.StringOutput)
}

type CloudFormationTypeLoggingConfigPtrOutput struct{ *pulumi.OutputState }

func (CloudFormationTypeLoggingConfigPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudFormationTypeLoggingConfig)(nil)).Elem()
}

func (o CloudFormationTypeLoggingConfigPtrOutput) ToCloudFormationTypeLoggingConfigPtrOutput() CloudFormationTypeLoggingConfigPtrOutput {
	return o
}

func (o CloudFormationTypeLoggingConfigPtrOutput) ToCloudFormationTypeLoggingConfigPtrOutputWithContext(ctx context.Context) CloudFormationTypeLoggingConfigPtrOutput {
	return o
}

func (o CloudFormationTypeLoggingConfigPtrOutput) Elem() CloudFormationTypeLoggingConfigOutput {
	return o.ApplyT(func(v *CloudFormationTypeLoggingConfig) CloudFormationTypeLoggingConfig { return *v }).(CloudFormationTypeLoggingConfigOutput)
}

// Name of the CloudWatch Log Group where CloudFormation sends error logging information when invoking the type's handlers.
func (o CloudFormationTypeLoggingConfigPtrOutput) LogGroupName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CloudFormationTypeLoggingConfig) *string {
		if v == nil {
			return nil
		}
		return &v.LogGroupName
	}).(pulumi.StringPtrOutput)
}

// Amazon Resource Name (ARN) of the IAM Role CloudFormation assumes when sending error logging information to CloudWatch Logs.
func (o CloudFormationTypeLoggingConfigPtrOutput) LogRoleArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CloudFormationTypeLoggingConfig) *string {
		if v == nil {
			return nil
		}
		return &v.LogRoleArn
	}).(pulumi.StringPtrOutput)
}

type StackSetAutoDeployment struct {
	// Whether or not auto-deployment is enabled.
	Enabled *bool `pulumi:"enabled"`
	// Whether or not to retain stacks when the account is removed.
	RetainStacksOnAccountRemoval *bool `pulumi:"retainStacksOnAccountRemoval"`
}

// StackSetAutoDeploymentInput is an input type that accepts StackSetAutoDeploymentArgs and StackSetAutoDeploymentOutput values.
// You can construct a concrete instance of `StackSetAutoDeploymentInput` via:
//
//          StackSetAutoDeploymentArgs{...}
type StackSetAutoDeploymentInput interface {
	pulumi.Input

	ToStackSetAutoDeploymentOutput() StackSetAutoDeploymentOutput
	ToStackSetAutoDeploymentOutputWithContext(context.Context) StackSetAutoDeploymentOutput
}

type StackSetAutoDeploymentArgs struct {
	// Whether or not auto-deployment is enabled.
	Enabled pulumi.BoolPtrInput `pulumi:"enabled"`
	// Whether or not to retain stacks when the account is removed.
	RetainStacksOnAccountRemoval pulumi.BoolPtrInput `pulumi:"retainStacksOnAccountRemoval"`
}

func (StackSetAutoDeploymentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StackSetAutoDeployment)(nil)).Elem()
}

func (i StackSetAutoDeploymentArgs) ToStackSetAutoDeploymentOutput() StackSetAutoDeploymentOutput {
	return i.ToStackSetAutoDeploymentOutputWithContext(context.Background())
}

func (i StackSetAutoDeploymentArgs) ToStackSetAutoDeploymentOutputWithContext(ctx context.Context) StackSetAutoDeploymentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StackSetAutoDeploymentOutput)
}

func (i StackSetAutoDeploymentArgs) ToStackSetAutoDeploymentPtrOutput() StackSetAutoDeploymentPtrOutput {
	return i.ToStackSetAutoDeploymentPtrOutputWithContext(context.Background())
}

func (i StackSetAutoDeploymentArgs) ToStackSetAutoDeploymentPtrOutputWithContext(ctx context.Context) StackSetAutoDeploymentPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StackSetAutoDeploymentOutput).ToStackSetAutoDeploymentPtrOutputWithContext(ctx)
}

// StackSetAutoDeploymentPtrInput is an input type that accepts StackSetAutoDeploymentArgs, StackSetAutoDeploymentPtr and StackSetAutoDeploymentPtrOutput values.
// You can construct a concrete instance of `StackSetAutoDeploymentPtrInput` via:
//
//          StackSetAutoDeploymentArgs{...}
//
//  or:
//
//          nil
type StackSetAutoDeploymentPtrInput interface {
	pulumi.Input

	ToStackSetAutoDeploymentPtrOutput() StackSetAutoDeploymentPtrOutput
	ToStackSetAutoDeploymentPtrOutputWithContext(context.Context) StackSetAutoDeploymentPtrOutput
}

type stackSetAutoDeploymentPtrType StackSetAutoDeploymentArgs

func StackSetAutoDeploymentPtr(v *StackSetAutoDeploymentArgs) StackSetAutoDeploymentPtrInput {
	return (*stackSetAutoDeploymentPtrType)(v)
}

func (*stackSetAutoDeploymentPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**StackSetAutoDeployment)(nil)).Elem()
}

func (i *stackSetAutoDeploymentPtrType) ToStackSetAutoDeploymentPtrOutput() StackSetAutoDeploymentPtrOutput {
	return i.ToStackSetAutoDeploymentPtrOutputWithContext(context.Background())
}

func (i *stackSetAutoDeploymentPtrType) ToStackSetAutoDeploymentPtrOutputWithContext(ctx context.Context) StackSetAutoDeploymentPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StackSetAutoDeploymentPtrOutput)
}

type StackSetAutoDeploymentOutput struct{ *pulumi.OutputState }

func (StackSetAutoDeploymentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StackSetAutoDeployment)(nil)).Elem()
}

func (o StackSetAutoDeploymentOutput) ToStackSetAutoDeploymentOutput() StackSetAutoDeploymentOutput {
	return o
}

func (o StackSetAutoDeploymentOutput) ToStackSetAutoDeploymentOutputWithContext(ctx context.Context) StackSetAutoDeploymentOutput {
	return o
}

func (o StackSetAutoDeploymentOutput) ToStackSetAutoDeploymentPtrOutput() StackSetAutoDeploymentPtrOutput {
	return o.ToStackSetAutoDeploymentPtrOutputWithContext(context.Background())
}

func (o StackSetAutoDeploymentOutput) ToStackSetAutoDeploymentPtrOutputWithContext(ctx context.Context) StackSetAutoDeploymentPtrOutput {
	return o.ApplyT(func(v StackSetAutoDeployment) *StackSetAutoDeployment {
		return &v
	}).(StackSetAutoDeploymentPtrOutput)
}

// Whether or not auto-deployment is enabled.
func (o StackSetAutoDeploymentOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v StackSetAutoDeployment) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

// Whether or not to retain stacks when the account is removed.
func (o StackSetAutoDeploymentOutput) RetainStacksOnAccountRemoval() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v StackSetAutoDeployment) *bool { return v.RetainStacksOnAccountRemoval }).(pulumi.BoolPtrOutput)
}

type StackSetAutoDeploymentPtrOutput struct{ *pulumi.OutputState }

func (StackSetAutoDeploymentPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StackSetAutoDeployment)(nil)).Elem()
}

func (o StackSetAutoDeploymentPtrOutput) ToStackSetAutoDeploymentPtrOutput() StackSetAutoDeploymentPtrOutput {
	return o
}

func (o StackSetAutoDeploymentPtrOutput) ToStackSetAutoDeploymentPtrOutputWithContext(ctx context.Context) StackSetAutoDeploymentPtrOutput {
	return o
}

func (o StackSetAutoDeploymentPtrOutput) Elem() StackSetAutoDeploymentOutput {
	return o.ApplyT(func(v *StackSetAutoDeployment) StackSetAutoDeployment { return *v }).(StackSetAutoDeploymentOutput)
}

// Whether or not auto-deployment is enabled.
func (o StackSetAutoDeploymentPtrOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *StackSetAutoDeployment) *bool {
		if v == nil {
			return nil
		}
		return v.Enabled
	}).(pulumi.BoolPtrOutput)
}

// Whether or not to retain stacks when the account is removed.
func (o StackSetAutoDeploymentPtrOutput) RetainStacksOnAccountRemoval() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *StackSetAutoDeployment) *bool {
		if v == nil {
			return nil
		}
		return v.RetainStacksOnAccountRemoval
	}).(pulumi.BoolPtrOutput)
}

type GetCloudFormationTypeLoggingConfig struct {
	// Name of the CloudWatch Log Group where CloudFormation sends error logging information when invoking the type's handlers.
	LogGroupName string `pulumi:"logGroupName"`
	// Amazon Resource Name (ARN) of the IAM Role CloudFormation assumes when sending error logging information to CloudWatch Logs.
	LogRoleArn string `pulumi:"logRoleArn"`
}

// GetCloudFormationTypeLoggingConfigInput is an input type that accepts GetCloudFormationTypeLoggingConfigArgs and GetCloudFormationTypeLoggingConfigOutput values.
// You can construct a concrete instance of `GetCloudFormationTypeLoggingConfigInput` via:
//
//          GetCloudFormationTypeLoggingConfigArgs{...}
type GetCloudFormationTypeLoggingConfigInput interface {
	pulumi.Input

	ToGetCloudFormationTypeLoggingConfigOutput() GetCloudFormationTypeLoggingConfigOutput
	ToGetCloudFormationTypeLoggingConfigOutputWithContext(context.Context) GetCloudFormationTypeLoggingConfigOutput
}

type GetCloudFormationTypeLoggingConfigArgs struct {
	// Name of the CloudWatch Log Group where CloudFormation sends error logging information when invoking the type's handlers.
	LogGroupName pulumi.StringInput `pulumi:"logGroupName"`
	// Amazon Resource Name (ARN) of the IAM Role CloudFormation assumes when sending error logging information to CloudWatch Logs.
	LogRoleArn pulumi.StringInput `pulumi:"logRoleArn"`
}

func (GetCloudFormationTypeLoggingConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCloudFormationTypeLoggingConfig)(nil)).Elem()
}

func (i GetCloudFormationTypeLoggingConfigArgs) ToGetCloudFormationTypeLoggingConfigOutput() GetCloudFormationTypeLoggingConfigOutput {
	return i.ToGetCloudFormationTypeLoggingConfigOutputWithContext(context.Background())
}

func (i GetCloudFormationTypeLoggingConfigArgs) ToGetCloudFormationTypeLoggingConfigOutputWithContext(ctx context.Context) GetCloudFormationTypeLoggingConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetCloudFormationTypeLoggingConfigOutput)
}

// GetCloudFormationTypeLoggingConfigArrayInput is an input type that accepts GetCloudFormationTypeLoggingConfigArray and GetCloudFormationTypeLoggingConfigArrayOutput values.
// You can construct a concrete instance of `GetCloudFormationTypeLoggingConfigArrayInput` via:
//
//          GetCloudFormationTypeLoggingConfigArray{ GetCloudFormationTypeLoggingConfigArgs{...} }
type GetCloudFormationTypeLoggingConfigArrayInput interface {
	pulumi.Input

	ToGetCloudFormationTypeLoggingConfigArrayOutput() GetCloudFormationTypeLoggingConfigArrayOutput
	ToGetCloudFormationTypeLoggingConfigArrayOutputWithContext(context.Context) GetCloudFormationTypeLoggingConfigArrayOutput
}

type GetCloudFormationTypeLoggingConfigArray []GetCloudFormationTypeLoggingConfigInput

func (GetCloudFormationTypeLoggingConfigArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetCloudFormationTypeLoggingConfig)(nil)).Elem()
}

func (i GetCloudFormationTypeLoggingConfigArray) ToGetCloudFormationTypeLoggingConfigArrayOutput() GetCloudFormationTypeLoggingConfigArrayOutput {
	return i.ToGetCloudFormationTypeLoggingConfigArrayOutputWithContext(context.Background())
}

func (i GetCloudFormationTypeLoggingConfigArray) ToGetCloudFormationTypeLoggingConfigArrayOutputWithContext(ctx context.Context) GetCloudFormationTypeLoggingConfigArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetCloudFormationTypeLoggingConfigArrayOutput)
}

type GetCloudFormationTypeLoggingConfigOutput struct{ *pulumi.OutputState }

func (GetCloudFormationTypeLoggingConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCloudFormationTypeLoggingConfig)(nil)).Elem()
}

func (o GetCloudFormationTypeLoggingConfigOutput) ToGetCloudFormationTypeLoggingConfigOutput() GetCloudFormationTypeLoggingConfigOutput {
	return o
}

func (o GetCloudFormationTypeLoggingConfigOutput) ToGetCloudFormationTypeLoggingConfigOutputWithContext(ctx context.Context) GetCloudFormationTypeLoggingConfigOutput {
	return o
}

// Name of the CloudWatch Log Group where CloudFormation sends error logging information when invoking the type's handlers.
func (o GetCloudFormationTypeLoggingConfigOutput) LogGroupName() pulumi.StringOutput {
	return o.ApplyT(func(v GetCloudFormationTypeLoggingConfig) string { return v.LogGroupName }).(pulumi.StringOutput)
}

// Amazon Resource Name (ARN) of the IAM Role CloudFormation assumes when sending error logging information to CloudWatch Logs.
func (o GetCloudFormationTypeLoggingConfigOutput) LogRoleArn() pulumi.StringOutput {
	return o.ApplyT(func(v GetCloudFormationTypeLoggingConfig) string { return v.LogRoleArn }).(pulumi.StringOutput)
}

type GetCloudFormationTypeLoggingConfigArrayOutput struct{ *pulumi.OutputState }

func (GetCloudFormationTypeLoggingConfigArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetCloudFormationTypeLoggingConfig)(nil)).Elem()
}

func (o GetCloudFormationTypeLoggingConfigArrayOutput) ToGetCloudFormationTypeLoggingConfigArrayOutput() GetCloudFormationTypeLoggingConfigArrayOutput {
	return o
}

func (o GetCloudFormationTypeLoggingConfigArrayOutput) ToGetCloudFormationTypeLoggingConfigArrayOutputWithContext(ctx context.Context) GetCloudFormationTypeLoggingConfigArrayOutput {
	return o
}

func (o GetCloudFormationTypeLoggingConfigArrayOutput) Index(i pulumi.IntInput) GetCloudFormationTypeLoggingConfigOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetCloudFormationTypeLoggingConfig {
		return vs[0].([]GetCloudFormationTypeLoggingConfig)[vs[1].(int)]
	}).(GetCloudFormationTypeLoggingConfigOutput)
}

func init() {
	pulumi.RegisterOutputType(CloudFormationTypeLoggingConfigOutput{})
	pulumi.RegisterOutputType(CloudFormationTypeLoggingConfigPtrOutput{})
	pulumi.RegisterOutputType(StackSetAutoDeploymentOutput{})
	pulumi.RegisterOutputType(StackSetAutoDeploymentPtrOutput{})
	pulumi.RegisterOutputType(GetCloudFormationTypeLoggingConfigOutput{})
	pulumi.RegisterOutputType(GetCloudFormationTypeLoggingConfigArrayOutput{})
}
