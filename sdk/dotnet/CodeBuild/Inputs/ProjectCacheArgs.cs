// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.CodeBuild.Inputs
{

    public sealed class ProjectCacheArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Location of the source code from git or s3.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        [Input("modes")]
        private InputList<string>? _modes;

        /// <summary>
        /// Specifies settings that AWS CodeBuild uses to store and reuse build dependencies. Valid values:  `LOCAL_SOURCE_CACHE`, `LOCAL_DOCKER_LAYER_CACHE`, `LOCAL_CUSTOM_CACHE`.
        /// </summary>
        public InputList<string> Modes
        {
            get => _modes ?? (_modes = new InputList<string>());
            set => _modes = value;
        }

        /// <summary>
        /// Authorization type to use. The only valid value is `OAUTH`. This data type is deprecated and is no longer accurate or used. Use the `aws.codebuild.SourceCredential` resource instead.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public ProjectCacheArgs()
        {
        }
    }
}
