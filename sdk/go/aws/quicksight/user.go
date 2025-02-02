// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package quicksight

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource for managing QuickSight User
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/quicksight"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := quicksight.NewUser(ctx, "example", &quicksight.UserArgs{
// 			Email:        pulumi.String("author@example.com"),
// 			IdentityType: pulumi.String("IAM"),
// 			UserName:     pulumi.String("an-author"),
// 			UserRole:     pulumi.String("AUTHOR"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// Importing is currently not supported on this resource.
type User struct {
	pulumi.CustomResourceState

	// Amazon Resource Name (ARN) of the user
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The ID for the AWS account that the user is in. Currently, you use the ID for the AWS account that contains your Amazon QuickSight account.
	AwsAccountId pulumi.StringOutput `pulumi:"awsAccountId"`
	// The email address of the user that you want to register.
	Email pulumi.StringOutput `pulumi:"email"`
	// The ARN of the IAM user or role that you are registering with Amazon QuickSight.
	IamArn pulumi.StringPtrOutput `pulumi:"iamArn"`
	// Amazon QuickSight supports several ways of managing the identity of users. This parameter accepts either  `IAM` or `QUICKSIGHT`.
	IdentityType pulumi.StringOutput `pulumi:"identityType"`
	// The namespace. Currently, you should set this to `default`.
	Namespace pulumi.StringPtrOutput `pulumi:"namespace"`
	// The name of the IAM session to use when assuming roles that can embed QuickSight dashboards.
	SessionName pulumi.StringPtrOutput `pulumi:"sessionName"`
	// The Amazon QuickSight user name that you want to create for the user you are registering.
	UserName pulumi.StringPtrOutput `pulumi:"userName"`
	// The Amazon QuickSight role of the user. The user role can be one of the following: `READER`, `AUTHOR`, or `ADMIN`
	UserRole pulumi.StringOutput `pulumi:"userRole"`
}

// NewUser registers a new resource with the given unique name, arguments, and options.
func NewUser(ctx *pulumi.Context,
	name string, args *UserArgs, opts ...pulumi.ResourceOption) (*User, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Email == nil {
		return nil, errors.New("invalid value for required argument 'Email'")
	}
	if args.IdentityType == nil {
		return nil, errors.New("invalid value for required argument 'IdentityType'")
	}
	if args.UserRole == nil {
		return nil, errors.New("invalid value for required argument 'UserRole'")
	}
	var resource User
	err := ctx.RegisterResource("aws:quicksight/user:User", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetUser gets an existing User resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetUser(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *UserState, opts ...pulumi.ResourceOption) (*User, error) {
	var resource User
	err := ctx.ReadResource("aws:quicksight/user:User", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering User resources.
type userState struct {
	// Amazon Resource Name (ARN) of the user
	Arn *string `pulumi:"arn"`
	// The ID for the AWS account that the user is in. Currently, you use the ID for the AWS account that contains your Amazon QuickSight account.
	AwsAccountId *string `pulumi:"awsAccountId"`
	// The email address of the user that you want to register.
	Email *string `pulumi:"email"`
	// The ARN of the IAM user or role that you are registering with Amazon QuickSight.
	IamArn *string `pulumi:"iamArn"`
	// Amazon QuickSight supports several ways of managing the identity of users. This parameter accepts either  `IAM` or `QUICKSIGHT`.
	IdentityType *string `pulumi:"identityType"`
	// The namespace. Currently, you should set this to `default`.
	Namespace *string `pulumi:"namespace"`
	// The name of the IAM session to use when assuming roles that can embed QuickSight dashboards.
	SessionName *string `pulumi:"sessionName"`
	// The Amazon QuickSight user name that you want to create for the user you are registering.
	UserName *string `pulumi:"userName"`
	// The Amazon QuickSight role of the user. The user role can be one of the following: `READER`, `AUTHOR`, or `ADMIN`
	UserRole *string `pulumi:"userRole"`
}

type UserState struct {
	// Amazon Resource Name (ARN) of the user
	Arn pulumi.StringPtrInput
	// The ID for the AWS account that the user is in. Currently, you use the ID for the AWS account that contains your Amazon QuickSight account.
	AwsAccountId pulumi.StringPtrInput
	// The email address of the user that you want to register.
	Email pulumi.StringPtrInput
	// The ARN of the IAM user or role that you are registering with Amazon QuickSight.
	IamArn pulumi.StringPtrInput
	// Amazon QuickSight supports several ways of managing the identity of users. This parameter accepts either  `IAM` or `QUICKSIGHT`.
	IdentityType pulumi.StringPtrInput
	// The namespace. Currently, you should set this to `default`.
	Namespace pulumi.StringPtrInput
	// The name of the IAM session to use when assuming roles that can embed QuickSight dashboards.
	SessionName pulumi.StringPtrInput
	// The Amazon QuickSight user name that you want to create for the user you are registering.
	UserName pulumi.StringPtrInput
	// The Amazon QuickSight role of the user. The user role can be one of the following: `READER`, `AUTHOR`, or `ADMIN`
	UserRole pulumi.StringPtrInput
}

func (UserState) ElementType() reflect.Type {
	return reflect.TypeOf((*userState)(nil)).Elem()
}

type userArgs struct {
	// The ID for the AWS account that the user is in. Currently, you use the ID for the AWS account that contains your Amazon QuickSight account.
	AwsAccountId *string `pulumi:"awsAccountId"`
	// The email address of the user that you want to register.
	Email string `pulumi:"email"`
	// The ARN of the IAM user or role that you are registering with Amazon QuickSight.
	IamArn *string `pulumi:"iamArn"`
	// Amazon QuickSight supports several ways of managing the identity of users. This parameter accepts either  `IAM` or `QUICKSIGHT`.
	IdentityType string `pulumi:"identityType"`
	// The namespace. Currently, you should set this to `default`.
	Namespace *string `pulumi:"namespace"`
	// The name of the IAM session to use when assuming roles that can embed QuickSight dashboards.
	SessionName *string `pulumi:"sessionName"`
	// The Amazon QuickSight user name that you want to create for the user you are registering.
	UserName *string `pulumi:"userName"`
	// The Amazon QuickSight role of the user. The user role can be one of the following: `READER`, `AUTHOR`, or `ADMIN`
	UserRole string `pulumi:"userRole"`
}

// The set of arguments for constructing a User resource.
type UserArgs struct {
	// The ID for the AWS account that the user is in. Currently, you use the ID for the AWS account that contains your Amazon QuickSight account.
	AwsAccountId pulumi.StringPtrInput
	// The email address of the user that you want to register.
	Email pulumi.StringInput
	// The ARN of the IAM user or role that you are registering with Amazon QuickSight.
	IamArn pulumi.StringPtrInput
	// Amazon QuickSight supports several ways of managing the identity of users. This parameter accepts either  `IAM` or `QUICKSIGHT`.
	IdentityType pulumi.StringInput
	// The namespace. Currently, you should set this to `default`.
	Namespace pulumi.StringPtrInput
	// The name of the IAM session to use when assuming roles that can embed QuickSight dashboards.
	SessionName pulumi.StringPtrInput
	// The Amazon QuickSight user name that you want to create for the user you are registering.
	UserName pulumi.StringPtrInput
	// The Amazon QuickSight role of the user. The user role can be one of the following: `READER`, `AUTHOR`, or `ADMIN`
	UserRole pulumi.StringInput
}

func (UserArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*userArgs)(nil)).Elem()
}

type UserInput interface {
	pulumi.Input

	ToUserOutput() UserOutput
	ToUserOutputWithContext(ctx context.Context) UserOutput
}

func (*User) ElementType() reflect.Type {
	return reflect.TypeOf((*User)(nil))
}

func (i *User) ToUserOutput() UserOutput {
	return i.ToUserOutputWithContext(context.Background())
}

func (i *User) ToUserOutputWithContext(ctx context.Context) UserOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserOutput)
}

func (i *User) ToUserPtrOutput() UserPtrOutput {
	return i.ToUserPtrOutputWithContext(context.Background())
}

func (i *User) ToUserPtrOutputWithContext(ctx context.Context) UserPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserPtrOutput)
}

type UserPtrInput interface {
	pulumi.Input

	ToUserPtrOutput() UserPtrOutput
	ToUserPtrOutputWithContext(ctx context.Context) UserPtrOutput
}

type userPtrType UserArgs

func (*userPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**User)(nil))
}

func (i *userPtrType) ToUserPtrOutput() UserPtrOutput {
	return i.ToUserPtrOutputWithContext(context.Background())
}

func (i *userPtrType) ToUserPtrOutputWithContext(ctx context.Context) UserPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserPtrOutput)
}

// UserArrayInput is an input type that accepts UserArray and UserArrayOutput values.
// You can construct a concrete instance of `UserArrayInput` via:
//
//          UserArray{ UserArgs{...} }
type UserArrayInput interface {
	pulumi.Input

	ToUserArrayOutput() UserArrayOutput
	ToUserArrayOutputWithContext(context.Context) UserArrayOutput
}

type UserArray []UserInput

func (UserArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*User)(nil))
}

func (i UserArray) ToUserArrayOutput() UserArrayOutput {
	return i.ToUserArrayOutputWithContext(context.Background())
}

func (i UserArray) ToUserArrayOutputWithContext(ctx context.Context) UserArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserArrayOutput)
}

// UserMapInput is an input type that accepts UserMap and UserMapOutput values.
// You can construct a concrete instance of `UserMapInput` via:
//
//          UserMap{ "key": UserArgs{...} }
type UserMapInput interface {
	pulumi.Input

	ToUserMapOutput() UserMapOutput
	ToUserMapOutputWithContext(context.Context) UserMapOutput
}

type UserMap map[string]UserInput

func (UserMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*User)(nil))
}

func (i UserMap) ToUserMapOutput() UserMapOutput {
	return i.ToUserMapOutputWithContext(context.Background())
}

func (i UserMap) ToUserMapOutputWithContext(ctx context.Context) UserMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserMapOutput)
}

type UserOutput struct {
	*pulumi.OutputState
}

func (UserOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*User)(nil))
}

func (o UserOutput) ToUserOutput() UserOutput {
	return o
}

func (o UserOutput) ToUserOutputWithContext(ctx context.Context) UserOutput {
	return o
}

func (o UserOutput) ToUserPtrOutput() UserPtrOutput {
	return o.ToUserPtrOutputWithContext(context.Background())
}

func (o UserOutput) ToUserPtrOutputWithContext(ctx context.Context) UserPtrOutput {
	return o.ApplyT(func(v User) *User {
		return &v
	}).(UserPtrOutput)
}

type UserPtrOutput struct {
	*pulumi.OutputState
}

func (UserPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**User)(nil))
}

func (o UserPtrOutput) ToUserPtrOutput() UserPtrOutput {
	return o
}

func (o UserPtrOutput) ToUserPtrOutputWithContext(ctx context.Context) UserPtrOutput {
	return o
}

type UserArrayOutput struct{ *pulumi.OutputState }

func (UserArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]User)(nil))
}

func (o UserArrayOutput) ToUserArrayOutput() UserArrayOutput {
	return o
}

func (o UserArrayOutput) ToUserArrayOutputWithContext(ctx context.Context) UserArrayOutput {
	return o
}

func (o UserArrayOutput) Index(i pulumi.IntInput) UserOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) User {
		return vs[0].([]User)[vs[1].(int)]
	}).(UserOutput)
}

type UserMapOutput struct{ *pulumi.OutputState }

func (UserMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]User)(nil))
}

func (o UserMapOutput) ToUserMapOutput() UserMapOutput {
	return o
}

func (o UserMapOutput) ToUserMapOutputWithContext(ctx context.Context) UserMapOutput {
	return o
}

func (o UserMapOutput) MapIndex(k pulumi.StringInput) UserOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) User {
		return vs[0].(map[string]User)[vs[1].(string)]
	}).(UserOutput)
}

func init() {
	pulumi.RegisterOutputType(UserOutput{})
	pulumi.RegisterOutputType(UserPtrOutput{})
	pulumi.RegisterOutputType(UserArrayOutput{})
	pulumi.RegisterOutputType(UserMapOutput{})
}
