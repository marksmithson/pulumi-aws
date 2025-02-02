// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package sqs

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get the ARN and URL of queue in AWS Simple Queue Service (SQS).
// By using this data source, you can reference SQS queues without having to hardcode
// the ARNs as input.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/sqs"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := sqs.LookupQueue(ctx, &sqs.LookupQueueArgs{
// 			Name: "queue",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupQueue(ctx *pulumi.Context, args *LookupQueueArgs, opts ...pulumi.InvokeOption) (*LookupQueueResult, error) {
	var rv LookupQueueResult
	err := ctx.Invoke("aws:sqs/getQueue:getQueue", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getQueue.
type LookupQueueArgs struct {
	// The name of the queue to match.
	Name string `pulumi:"name"`
	// A map of tags for the resource.
	Tags map[string]string `pulumi:"tags"`
}

// A collection of values returned by getQueue.
type LookupQueueResult struct {
	// The Amazon Resource Name (ARN) of the queue.
	Arn string `pulumi:"arn"`
	// The provider-assigned unique ID for this managed resource.
	Id   string `pulumi:"id"`
	Name string `pulumi:"name"`
	// A map of tags for the resource.
	Tags map[string]string `pulumi:"tags"`
	// The URL of the queue.
	Url string `pulumi:"url"`
}
