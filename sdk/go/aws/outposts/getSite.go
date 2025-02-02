// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package outposts

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides details about an Outposts Site.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/outposts"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		opt0 := "example"
// 		_, err := outposts.GetSite(ctx, &outposts.GetSiteArgs{
// 			Name: &opt0,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetSite(ctx *pulumi.Context, args *GetSiteArgs, opts ...pulumi.InvokeOption) (*GetSiteResult, error) {
	var rv GetSiteResult
	err := ctx.Invoke("aws:outposts/getSite:getSite", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSite.
type GetSiteArgs struct {
	// Identifier of the Site.
	Id *string `pulumi:"id"`
	// Name of the Site.
	Name *string `pulumi:"name"`
}

// A collection of values returned by getSite.
type GetSiteResult struct {
	// AWS Account identifier.
	AccountId string `pulumi:"accountId"`
	// Description.
	Description string `pulumi:"description"`
	Id          string `pulumi:"id"`
	Name        string `pulumi:"name"`
}
