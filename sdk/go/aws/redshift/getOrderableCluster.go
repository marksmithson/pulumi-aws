// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package redshift

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Information about Redshift Orderable Clusters and valid parameter combinations.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/redshift"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		opt0 := "multi-node"
// 		_, err := redshift.GetOrderableCluster(ctx, &redshift.GetOrderableClusterArgs{
// 			ClusterType: &opt0,
// 			PreferredNodeTypes: []string{
// 				"dc2.large",
// 				"ds2.xlarge",
// 			},
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetOrderableCluster(ctx *pulumi.Context, args *GetOrderableClusterArgs, opts ...pulumi.InvokeOption) (*GetOrderableClusterResult, error) {
	var rv GetOrderableClusterResult
	err := ctx.Invoke("aws:redshift/getOrderableCluster:getOrderableCluster", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getOrderableCluster.
type GetOrderableClusterArgs struct {
	// Reshift Cluster type. e.g. `multi-node` or `single-node`
	ClusterType *string `pulumi:"clusterType"`
	// Redshift Cluster version. e.g. `1.0`
	ClusterVersion *string `pulumi:"clusterVersion"`
	// Redshift Cluster node type. e.g. `dc2.8xlarge`
	NodeType *string `pulumi:"nodeType"`
	// Ordered list of preferred Redshift Cluster node types. The first match in this list will be returned. If no preferred matches are found and the original search returned more than one result, an error is returned.
	PreferredNodeTypes []string `pulumi:"preferredNodeTypes"`
}

// A collection of values returned by getOrderableCluster.
type GetOrderableClusterResult struct {
	// List of Availability Zone names where the Redshit Cluster is available.
	AvailabilityZones []string `pulumi:"availabilityZones"`
	ClusterType       string   `pulumi:"clusterType"`
	ClusterVersion    string   `pulumi:"clusterVersion"`
	// The provider-assigned unique ID for this managed resource.
	Id                 string   `pulumi:"id"`
	NodeType           string   `pulumi:"nodeType"`
	PreferredNodeTypes []string `pulumi:"preferredNodeTypes"`
}
