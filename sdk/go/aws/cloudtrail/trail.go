// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cloudtrail

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a CloudTrail resource.
//
// > *NOTE:* For a multi-region trail, this resource must be in the home region of the trail.
//
// > *NOTE:* For an organization trail, this resource must be in the master account of the organization.
//
// ## Example Usage
// ### Basic
//
// Enable CloudTrail to capture all compatible management events in region.
// For capturing events from services like IAM, `includeGlobalServiceEvents` must be enabled.
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws"
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/cloudtrail"
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/s3"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		current, err := aws.GetCallerIdentity(ctx, nil, nil)
// 		if err != nil {
// 			return err
// 		}
// 		bucket, err := s3.NewBucket(ctx, "bucket", nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = s3.NewBucketPolicy(ctx, "bucketPolicy", &s3.BucketPolicyArgs{
// 			Bucket: bucket.ID(),
// 			Policy: pulumi.All(bucket.ID(), bucket.ID()).ApplyT(func(_args []interface{}) (string, error) {
// 				bucketId := _args[0].(string)
// 				bucketId1 := _args[1].(string)
// 				return fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v", "  {\n", "      \"Version\": \"2012-10-17\",\n", "      \"Statement\": [\n", "          {\n", "              \"Sid\": \"AWSCloudTrailAclCheck\",\n", "              \"Effect\": \"Allow\",\n", "              \"Principal\": {\n", "                \"Service\": \"cloudtrail.amazonaws.com\"\n", "              },\n", "              \"Action\": \"s3:GetBucketAcl\",\n", "              \"Resource\": \"arn:aws:s3:::", bucketId, "\"\n", "          },\n", "          {\n", "              \"Sid\": \"AWSCloudTrailWrite\",\n", "              \"Effect\": \"Allow\",\n", "              \"Principal\": {\n", "                \"Service\": \"cloudtrail.amazonaws.com\"\n", "              },\n", "              \"Action\": \"s3:PutObject\",\n", "              \"Resource\": \"arn:aws:s3:::", bucketId1, "/prefix/AWSLogs/", current.AccountId, "/*\",\n", "              \"Condition\": {\n", "                  \"StringEquals\": {\n", "                      \"s3:x-amz-acl\": \"bucket-owner-full-control\"\n", "                  }\n", "              }\n", "          }\n", "      ]\n", "  }\n"), nil
// 			}).(pulumi.StringOutput),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = cloudtrail.NewTrail(ctx, "foobar", &cloudtrail.TrailArgs{
// 			S3BucketName:               bucket.ID(),
// 			S3KeyPrefix:                pulumi.String("prefix"),
// 			IncludeGlobalServiceEvents: pulumi.Bool(false),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Data Event Logging
//
// CloudTrail can log [Data Events](https://docs.aws.amazon.com/awscloudtrail/latest/userguide/logging-data-events-with-cloudtrail.html) for certain services such as S3 bucket objects and Lambda function invocations. Additional information about data event configuration can be found in the [CloudTrail API DataResource documentation](https://docs.aws.amazon.com/awscloudtrail/latest/APIReference/API_DataResource.html).
// ### Logging All Lambda Function Invocations
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/cloudtrail"
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/s3"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		bucket, err := s3.NewBucket(ctx, "bucket", nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = cloudtrail.NewTrail(ctx, "example", &cloudtrail.TrailArgs{
// 			S3BucketName: bucket.ID(),
// 			S3KeyPrefix:  pulumi.String("prefix"),
// 			EventSelectors: cloudtrail.TrailEventSelectorArray{
// 				&cloudtrail.TrailEventSelectorArgs{
// 					ReadWriteType:           pulumi.String("All"),
// 					IncludeManagementEvents: pulumi.Bool(true),
// 					DataResources: cloudtrail.TrailEventSelectorDataResourceArray{
// 						&cloudtrail.TrailEventSelectorDataResourceArgs{
// 							Type: pulumi.String("AWS::Lambda::Function"),
// 							Values: pulumi.StringArray{
// 								pulumi.String("arn:aws:lambda"),
// 							},
// 						},
// 					},
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Logging All S3 Bucket Object Events
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/cloudtrail"
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/s3"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		bucket, err := s3.NewBucket(ctx, "bucket", nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = cloudtrail.NewTrail(ctx, "example", &cloudtrail.TrailArgs{
// 			S3BucketName: bucket.ID(),
// 			S3KeyPrefix:  pulumi.String("prefix"),
// 			EventSelectors: cloudtrail.TrailEventSelectorArray{
// 				&cloudtrail.TrailEventSelectorArgs{
// 					ReadWriteType:           pulumi.String("All"),
// 					IncludeManagementEvents: pulumi.Bool(true),
// 					DataResources: cloudtrail.TrailEventSelectorDataResourceArray{
// 						&cloudtrail.TrailEventSelectorDataResourceArgs{
// 							Type: pulumi.String("AWS::S3::Object"),
// 							Values: pulumi.StringArray{
// 								pulumi.String("arn:aws:s3:::"),
// 							},
// 						},
// 					},
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Logging Individual S3 Bucket Events
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/cloudtrail"
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/s3"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		important_bucket, err := s3.LookupBucket(ctx, &s3.LookupBucketArgs{
// 			Bucket: "important-bucket",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = cloudtrail.NewTrail(ctx, "example", &cloudtrail.TrailArgs{
// 			S3BucketName: pulumi.String(important_bucket.Id),
// 			S3KeyPrefix:  pulumi.String("prefix"),
// 			EventSelectors: cloudtrail.TrailEventSelectorArray{
// 				&cloudtrail.TrailEventSelectorArgs{
// 					ReadWriteType:           pulumi.String("All"),
// 					IncludeManagementEvents: pulumi.Bool(true),
// 					DataResources: cloudtrail.TrailEventSelectorDataResourceArray{
// 						&cloudtrail.TrailEventSelectorDataResourceArgs{
// 							Type: pulumi.String("AWS::S3::Object"),
// 							Values: pulumi.StringArray{
// 								pulumi.String(fmt.Sprintf("%v%v", important_bucket.Arn, "/")),
// 							},
// 						},
// 					},
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Sending Events to CloudWatch Logs
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws"
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/cloudtrail"
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/cloudwatch"
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/iam"
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/s3"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		current, err := aws.GetPartition(ctx, nil, nil)
// 		if err != nil {
// 			return err
// 		}
// 		exampleLogGroup, err := cloudwatch.NewLogGroup(ctx, "exampleLogGroup", nil)
// 		if err != nil {
// 			return err
// 		}
// 		testRole, err := iam.NewRole(ctx, "testRole", &iam.RoleArgs{
// 			AssumeRolePolicy: pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v", "{\n", "  \"Version\": \"2012-10-17\",\n", "  \"Statement\": [\n", "    {\n", "      \"Sid\": \"\",\n", "      \"Effect\": \"Allow\",\n", "      \"Principal\": {\n", "        \"Service\": \"cloudtrail.", current.DnsSuffix, "\"\n", "      },\n", "      \"Action\": \"sts:AssumeRole\"\n", "    }\n", "  ]\n", "}\n")),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = iam.NewRolePolicy(ctx, "testRolePolicy", &iam.RolePolicyArgs{
// 			Role:   testRole.ID(),
// 			Policy: pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v", "{\n", "  \"Version\": \"2012-10-17\",\n", "  \"Statement\": [\n", "    {\n", "      \"Sid\": \"AWSCloudTrailCreateLogStream\",\n", "      \"Effect\": \"Allow\",\n", "      \"Action\": [\n", "        \"logs:CreateLogStream\",\n", "        \"logs:PutLogEvents\"\n", "      ],\n", "      \"Resource\": \"", aws_cloudwatch_log_group.Test.Arn, ":*\"\n", "    }\n", "  ]\n", "}\n")),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = s3.NewBucket(ctx, "bucket", nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = cloudtrail.NewTrail(ctx, "exampleTrail", &cloudtrail.TrailArgs{
// 			S3BucketName:          pulumi.Any(data.Aws_s3_bucket.Important - bucket.Id),
// 			S3KeyPrefix:           pulumi.String("prefix"),
// 			CloudWatchLogsRoleArn: testRole.Arn,
// 			CloudWatchLogsGroupArn: exampleLogGroup.Arn.ApplyT(func(arn string) (string, error) {
// 				return fmt.Sprintf("%v%v", arn, ":*"), nil
// 			}).(pulumi.StringOutput),
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
// Cloudtrails can be imported using the `name`, e.g.
//
// ```sh
//  $ pulumi import aws:cloudtrail/trail:Trail sample my-sample-trail
// ```
type Trail struct {
	pulumi.CustomResourceState

	// The Amazon Resource Name of the trail.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Specifies a log group name using an Amazon Resource Name (ARN),
	// that represents the log group to which CloudTrail logs will be delivered. Note that CloudTrail requires the Log Stream wildcard.
	CloudWatchLogsGroupArn pulumi.StringPtrOutput `pulumi:"cloudWatchLogsGroupArn"`
	// Specifies the role for the CloudWatch Logs
	// endpoint to assume to write to a user’s log group.
	CloudWatchLogsRoleArn pulumi.StringPtrOutput `pulumi:"cloudWatchLogsRoleArn"`
	// Specifies whether log file integrity validation is enabled.
	// Defaults to `false`.
	EnableLogFileValidation pulumi.BoolPtrOutput `pulumi:"enableLogFileValidation"`
	// Enables logging for the trail. Defaults to `true`.
	// Setting this to `false` will pause logging.
	EnableLogging pulumi.BoolPtrOutput `pulumi:"enableLogging"`
	// Specifies an event selector for enabling data event logging. Fields documented below. Please note the [CloudTrail limits](https://docs.aws.amazon.com/awscloudtrail/latest/userguide/WhatIsCloudTrail-Limits.html) when configuring these.
	EventSelectors TrailEventSelectorArrayOutput `pulumi:"eventSelectors"`
	// The region in which the trail was created.
	HomeRegion pulumi.StringOutput `pulumi:"homeRegion"`
	// Specifies whether the trail is publishing events
	// from global services such as IAM to the log files. Defaults to `true`.
	IncludeGlobalServiceEvents pulumi.BoolPtrOutput `pulumi:"includeGlobalServiceEvents"`
	// Specifies an insight selector for identifying unusual operational activity. Fields documented below.
	InsightSelectors TrailInsightSelectorArrayOutput `pulumi:"insightSelectors"`
	// Specifies whether the trail is created in the current
	// region or in all regions. Defaults to `false`.
	IsMultiRegionTrail pulumi.BoolPtrOutput `pulumi:"isMultiRegionTrail"`
	// Specifies whether the trail is an AWS Organizations trail. Organization trails log events for the master account and all member accounts. Can only be created in the organization master account. Defaults to `false`.
	IsOrganizationTrail pulumi.BoolPtrOutput `pulumi:"isOrganizationTrail"`
	// Specifies the KMS key ARN to use to encrypt the logs delivered by CloudTrail.
	KmsKeyId pulumi.StringPtrOutput `pulumi:"kmsKeyId"`
	// Specifies the name of the trail.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the name of the S3 bucket designated for publishing log files.
	S3BucketName pulumi.StringOutput `pulumi:"s3BucketName"`
	// Specifies the S3 key prefix that follows
	// the name of the bucket you have designated for log file delivery.
	S3KeyPrefix pulumi.StringPtrOutput `pulumi:"s3KeyPrefix"`
	// Specifies the name of the Amazon SNS topic
	// defined for notification of log file delivery.
	SnsTopicName pulumi.StringPtrOutput `pulumi:"snsTopicName"`
	// A map of tags to assign to the trail. If configured with a provider [`defaultTags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
}

// NewTrail registers a new resource with the given unique name, arguments, and options.
func NewTrail(ctx *pulumi.Context,
	name string, args *TrailArgs, opts ...pulumi.ResourceOption) (*Trail, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.S3BucketName == nil {
		return nil, errors.New("invalid value for required argument 'S3BucketName'")
	}
	var resource Trail
	err := ctx.RegisterResource("aws:cloudtrail/trail:Trail", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTrail gets an existing Trail resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTrail(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TrailState, opts ...pulumi.ResourceOption) (*Trail, error) {
	var resource Trail
	err := ctx.ReadResource("aws:cloudtrail/trail:Trail", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Trail resources.
type trailState struct {
	// The Amazon Resource Name of the trail.
	Arn *string `pulumi:"arn"`
	// Specifies a log group name using an Amazon Resource Name (ARN),
	// that represents the log group to which CloudTrail logs will be delivered. Note that CloudTrail requires the Log Stream wildcard.
	CloudWatchLogsGroupArn *string `pulumi:"cloudWatchLogsGroupArn"`
	// Specifies the role for the CloudWatch Logs
	// endpoint to assume to write to a user’s log group.
	CloudWatchLogsRoleArn *string `pulumi:"cloudWatchLogsRoleArn"`
	// Specifies whether log file integrity validation is enabled.
	// Defaults to `false`.
	EnableLogFileValidation *bool `pulumi:"enableLogFileValidation"`
	// Enables logging for the trail. Defaults to `true`.
	// Setting this to `false` will pause logging.
	EnableLogging *bool `pulumi:"enableLogging"`
	// Specifies an event selector for enabling data event logging. Fields documented below. Please note the [CloudTrail limits](https://docs.aws.amazon.com/awscloudtrail/latest/userguide/WhatIsCloudTrail-Limits.html) when configuring these.
	EventSelectors []TrailEventSelector `pulumi:"eventSelectors"`
	// The region in which the trail was created.
	HomeRegion *string `pulumi:"homeRegion"`
	// Specifies whether the trail is publishing events
	// from global services such as IAM to the log files. Defaults to `true`.
	IncludeGlobalServiceEvents *bool `pulumi:"includeGlobalServiceEvents"`
	// Specifies an insight selector for identifying unusual operational activity. Fields documented below.
	InsightSelectors []TrailInsightSelector `pulumi:"insightSelectors"`
	// Specifies whether the trail is created in the current
	// region or in all regions. Defaults to `false`.
	IsMultiRegionTrail *bool `pulumi:"isMultiRegionTrail"`
	// Specifies whether the trail is an AWS Organizations trail. Organization trails log events for the master account and all member accounts. Can only be created in the organization master account. Defaults to `false`.
	IsOrganizationTrail *bool `pulumi:"isOrganizationTrail"`
	// Specifies the KMS key ARN to use to encrypt the logs delivered by CloudTrail.
	KmsKeyId *string `pulumi:"kmsKeyId"`
	// Specifies the name of the trail.
	Name *string `pulumi:"name"`
	// Specifies the name of the S3 bucket designated for publishing log files.
	S3BucketName *string `pulumi:"s3BucketName"`
	// Specifies the S3 key prefix that follows
	// the name of the bucket you have designated for log file delivery.
	S3KeyPrefix *string `pulumi:"s3KeyPrefix"`
	// Specifies the name of the Amazon SNS topic
	// defined for notification of log file delivery.
	SnsTopicName *string `pulumi:"snsTopicName"`
	// A map of tags to assign to the trail. If configured with a provider [`defaultTags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll map[string]string `pulumi:"tagsAll"`
}

type TrailState struct {
	// The Amazon Resource Name of the trail.
	Arn pulumi.StringPtrInput
	// Specifies a log group name using an Amazon Resource Name (ARN),
	// that represents the log group to which CloudTrail logs will be delivered. Note that CloudTrail requires the Log Stream wildcard.
	CloudWatchLogsGroupArn pulumi.StringPtrInput
	// Specifies the role for the CloudWatch Logs
	// endpoint to assume to write to a user’s log group.
	CloudWatchLogsRoleArn pulumi.StringPtrInput
	// Specifies whether log file integrity validation is enabled.
	// Defaults to `false`.
	EnableLogFileValidation pulumi.BoolPtrInput
	// Enables logging for the trail. Defaults to `true`.
	// Setting this to `false` will pause logging.
	EnableLogging pulumi.BoolPtrInput
	// Specifies an event selector for enabling data event logging. Fields documented below. Please note the [CloudTrail limits](https://docs.aws.amazon.com/awscloudtrail/latest/userguide/WhatIsCloudTrail-Limits.html) when configuring these.
	EventSelectors TrailEventSelectorArrayInput
	// The region in which the trail was created.
	HomeRegion pulumi.StringPtrInput
	// Specifies whether the trail is publishing events
	// from global services such as IAM to the log files. Defaults to `true`.
	IncludeGlobalServiceEvents pulumi.BoolPtrInput
	// Specifies an insight selector for identifying unusual operational activity. Fields documented below.
	InsightSelectors TrailInsightSelectorArrayInput
	// Specifies whether the trail is created in the current
	// region or in all regions. Defaults to `false`.
	IsMultiRegionTrail pulumi.BoolPtrInput
	// Specifies whether the trail is an AWS Organizations trail. Organization trails log events for the master account and all member accounts. Can only be created in the organization master account. Defaults to `false`.
	IsOrganizationTrail pulumi.BoolPtrInput
	// Specifies the KMS key ARN to use to encrypt the logs delivered by CloudTrail.
	KmsKeyId pulumi.StringPtrInput
	// Specifies the name of the trail.
	Name pulumi.StringPtrInput
	// Specifies the name of the S3 bucket designated for publishing log files.
	S3BucketName pulumi.StringPtrInput
	// Specifies the S3 key prefix that follows
	// the name of the bucket you have designated for log file delivery.
	S3KeyPrefix pulumi.StringPtrInput
	// Specifies the name of the Amazon SNS topic
	// defined for notification of log file delivery.
	SnsTopicName pulumi.StringPtrInput
	// A map of tags to assign to the trail. If configured with a provider [`defaultTags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll pulumi.StringMapInput
}

func (TrailState) ElementType() reflect.Type {
	return reflect.TypeOf((*trailState)(nil)).Elem()
}

type trailArgs struct {
	// Specifies a log group name using an Amazon Resource Name (ARN),
	// that represents the log group to which CloudTrail logs will be delivered. Note that CloudTrail requires the Log Stream wildcard.
	CloudWatchLogsGroupArn *string `pulumi:"cloudWatchLogsGroupArn"`
	// Specifies the role for the CloudWatch Logs
	// endpoint to assume to write to a user’s log group.
	CloudWatchLogsRoleArn *string `pulumi:"cloudWatchLogsRoleArn"`
	// Specifies whether log file integrity validation is enabled.
	// Defaults to `false`.
	EnableLogFileValidation *bool `pulumi:"enableLogFileValidation"`
	// Enables logging for the trail. Defaults to `true`.
	// Setting this to `false` will pause logging.
	EnableLogging *bool `pulumi:"enableLogging"`
	// Specifies an event selector for enabling data event logging. Fields documented below. Please note the [CloudTrail limits](https://docs.aws.amazon.com/awscloudtrail/latest/userguide/WhatIsCloudTrail-Limits.html) when configuring these.
	EventSelectors []TrailEventSelector `pulumi:"eventSelectors"`
	// Specifies whether the trail is publishing events
	// from global services such as IAM to the log files. Defaults to `true`.
	IncludeGlobalServiceEvents *bool `pulumi:"includeGlobalServiceEvents"`
	// Specifies an insight selector for identifying unusual operational activity. Fields documented below.
	InsightSelectors []TrailInsightSelector `pulumi:"insightSelectors"`
	// Specifies whether the trail is created in the current
	// region or in all regions. Defaults to `false`.
	IsMultiRegionTrail *bool `pulumi:"isMultiRegionTrail"`
	// Specifies whether the trail is an AWS Organizations trail. Organization trails log events for the master account and all member accounts. Can only be created in the organization master account. Defaults to `false`.
	IsOrganizationTrail *bool `pulumi:"isOrganizationTrail"`
	// Specifies the KMS key ARN to use to encrypt the logs delivered by CloudTrail.
	KmsKeyId *string `pulumi:"kmsKeyId"`
	// Specifies the name of the trail.
	Name *string `pulumi:"name"`
	// Specifies the name of the S3 bucket designated for publishing log files.
	S3BucketName string `pulumi:"s3BucketName"`
	// Specifies the S3 key prefix that follows
	// the name of the bucket you have designated for log file delivery.
	S3KeyPrefix *string `pulumi:"s3KeyPrefix"`
	// Specifies the name of the Amazon SNS topic
	// defined for notification of log file delivery.
	SnsTopicName *string `pulumi:"snsTopicName"`
	// A map of tags to assign to the trail. If configured with a provider [`defaultTags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll map[string]string `pulumi:"tagsAll"`
}

// The set of arguments for constructing a Trail resource.
type TrailArgs struct {
	// Specifies a log group name using an Amazon Resource Name (ARN),
	// that represents the log group to which CloudTrail logs will be delivered. Note that CloudTrail requires the Log Stream wildcard.
	CloudWatchLogsGroupArn pulumi.StringPtrInput
	// Specifies the role for the CloudWatch Logs
	// endpoint to assume to write to a user’s log group.
	CloudWatchLogsRoleArn pulumi.StringPtrInput
	// Specifies whether log file integrity validation is enabled.
	// Defaults to `false`.
	EnableLogFileValidation pulumi.BoolPtrInput
	// Enables logging for the trail. Defaults to `true`.
	// Setting this to `false` will pause logging.
	EnableLogging pulumi.BoolPtrInput
	// Specifies an event selector for enabling data event logging. Fields documented below. Please note the [CloudTrail limits](https://docs.aws.amazon.com/awscloudtrail/latest/userguide/WhatIsCloudTrail-Limits.html) when configuring these.
	EventSelectors TrailEventSelectorArrayInput
	// Specifies whether the trail is publishing events
	// from global services such as IAM to the log files. Defaults to `true`.
	IncludeGlobalServiceEvents pulumi.BoolPtrInput
	// Specifies an insight selector for identifying unusual operational activity. Fields documented below.
	InsightSelectors TrailInsightSelectorArrayInput
	// Specifies whether the trail is created in the current
	// region or in all regions. Defaults to `false`.
	IsMultiRegionTrail pulumi.BoolPtrInput
	// Specifies whether the trail is an AWS Organizations trail. Organization trails log events for the master account and all member accounts. Can only be created in the organization master account. Defaults to `false`.
	IsOrganizationTrail pulumi.BoolPtrInput
	// Specifies the KMS key ARN to use to encrypt the logs delivered by CloudTrail.
	KmsKeyId pulumi.StringPtrInput
	// Specifies the name of the trail.
	Name pulumi.StringPtrInput
	// Specifies the name of the S3 bucket designated for publishing log files.
	S3BucketName pulumi.StringInput
	// Specifies the S3 key prefix that follows
	// the name of the bucket you have designated for log file delivery.
	S3KeyPrefix pulumi.StringPtrInput
	// Specifies the name of the Amazon SNS topic
	// defined for notification of log file delivery.
	SnsTopicName pulumi.StringPtrInput
	// A map of tags to assign to the trail. If configured with a provider [`defaultTags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll pulumi.StringMapInput
}

func (TrailArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*trailArgs)(nil)).Elem()
}

type TrailInput interface {
	pulumi.Input

	ToTrailOutput() TrailOutput
	ToTrailOutputWithContext(ctx context.Context) TrailOutput
}

func (*Trail) ElementType() reflect.Type {
	return reflect.TypeOf((*Trail)(nil))
}

func (i *Trail) ToTrailOutput() TrailOutput {
	return i.ToTrailOutputWithContext(context.Background())
}

func (i *Trail) ToTrailOutputWithContext(ctx context.Context) TrailOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TrailOutput)
}

func (i *Trail) ToTrailPtrOutput() TrailPtrOutput {
	return i.ToTrailPtrOutputWithContext(context.Background())
}

func (i *Trail) ToTrailPtrOutputWithContext(ctx context.Context) TrailPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TrailPtrOutput)
}

type TrailPtrInput interface {
	pulumi.Input

	ToTrailPtrOutput() TrailPtrOutput
	ToTrailPtrOutputWithContext(ctx context.Context) TrailPtrOutput
}

type trailPtrType TrailArgs

func (*trailPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Trail)(nil))
}

func (i *trailPtrType) ToTrailPtrOutput() TrailPtrOutput {
	return i.ToTrailPtrOutputWithContext(context.Background())
}

func (i *trailPtrType) ToTrailPtrOutputWithContext(ctx context.Context) TrailPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TrailPtrOutput)
}

// TrailArrayInput is an input type that accepts TrailArray and TrailArrayOutput values.
// You can construct a concrete instance of `TrailArrayInput` via:
//
//          TrailArray{ TrailArgs{...} }
type TrailArrayInput interface {
	pulumi.Input

	ToTrailArrayOutput() TrailArrayOutput
	ToTrailArrayOutputWithContext(context.Context) TrailArrayOutput
}

type TrailArray []TrailInput

func (TrailArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*Trail)(nil))
}

func (i TrailArray) ToTrailArrayOutput() TrailArrayOutput {
	return i.ToTrailArrayOutputWithContext(context.Background())
}

func (i TrailArray) ToTrailArrayOutputWithContext(ctx context.Context) TrailArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TrailArrayOutput)
}

// TrailMapInput is an input type that accepts TrailMap and TrailMapOutput values.
// You can construct a concrete instance of `TrailMapInput` via:
//
//          TrailMap{ "key": TrailArgs{...} }
type TrailMapInput interface {
	pulumi.Input

	ToTrailMapOutput() TrailMapOutput
	ToTrailMapOutputWithContext(context.Context) TrailMapOutput
}

type TrailMap map[string]TrailInput

func (TrailMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*Trail)(nil))
}

func (i TrailMap) ToTrailMapOutput() TrailMapOutput {
	return i.ToTrailMapOutputWithContext(context.Background())
}

func (i TrailMap) ToTrailMapOutputWithContext(ctx context.Context) TrailMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TrailMapOutput)
}

type TrailOutput struct {
	*pulumi.OutputState
}

func (TrailOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Trail)(nil))
}

func (o TrailOutput) ToTrailOutput() TrailOutput {
	return o
}

func (o TrailOutput) ToTrailOutputWithContext(ctx context.Context) TrailOutput {
	return o
}

func (o TrailOutput) ToTrailPtrOutput() TrailPtrOutput {
	return o.ToTrailPtrOutputWithContext(context.Background())
}

func (o TrailOutput) ToTrailPtrOutputWithContext(ctx context.Context) TrailPtrOutput {
	return o.ApplyT(func(v Trail) *Trail {
		return &v
	}).(TrailPtrOutput)
}

type TrailPtrOutput struct {
	*pulumi.OutputState
}

func (TrailPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Trail)(nil))
}

func (o TrailPtrOutput) ToTrailPtrOutput() TrailPtrOutput {
	return o
}

func (o TrailPtrOutput) ToTrailPtrOutputWithContext(ctx context.Context) TrailPtrOutput {
	return o
}

type TrailArrayOutput struct{ *pulumi.OutputState }

func (TrailArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Trail)(nil))
}

func (o TrailArrayOutput) ToTrailArrayOutput() TrailArrayOutput {
	return o
}

func (o TrailArrayOutput) ToTrailArrayOutputWithContext(ctx context.Context) TrailArrayOutput {
	return o
}

func (o TrailArrayOutput) Index(i pulumi.IntInput) TrailOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Trail {
		return vs[0].([]Trail)[vs[1].(int)]
	}).(TrailOutput)
}

type TrailMapOutput struct{ *pulumi.OutputState }

func (TrailMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Trail)(nil))
}

func (o TrailMapOutput) ToTrailMapOutput() TrailMapOutput {
	return o
}

func (o TrailMapOutput) ToTrailMapOutputWithContext(ctx context.Context) TrailMapOutput {
	return o
}

func (o TrailMapOutput) MapIndex(k pulumi.StringInput) TrailOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Trail {
		return vs[0].(map[string]Trail)[vs[1].(string)]
	}).(TrailOutput)
}

func init() {
	pulumi.RegisterOutputType(TrailOutput{})
	pulumi.RegisterOutputType(TrailPtrOutput{})
	pulumi.RegisterOutputType(TrailArrayOutput{})
	pulumi.RegisterOutputType(TrailMapOutput{})
}
