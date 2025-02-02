# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

# Export this package's modules as members:
from .custom_data_identifier import *
from .findings_filter import *
from .member_account_association import *
from .s3_bucket_association import *
from ._inputs import *
from . import outputs

def _register_module():
    import pulumi
    from .. import _utilities


    class Module(pulumi.runtime.ResourceModule):
        _version = _utilities.get_semver_version()

        def version(self):
            return Module._version

        def construct(self, name: str, typ: str, urn: str) -> pulumi.Resource:
            if typ == "aws:macie/customDataIdentifier:CustomDataIdentifier":
                return CustomDataIdentifier(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:macie/findingsFilter:FindingsFilter":
                return FindingsFilter(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:macie/memberAccountAssociation:MemberAccountAssociation":
                return MemberAccountAssociation(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:macie/s3BucketAssociation:S3BucketAssociation":
                return S3BucketAssociation(name, pulumi.ResourceOptions(urn=urn))
            else:
                raise Exception(f"unknown resource type {typ}")


    _module_instance = Module()
    pulumi.runtime.register_resource_module("aws", "macie/customDataIdentifier", _module_instance)
    pulumi.runtime.register_resource_module("aws", "macie/findingsFilter", _module_instance)
    pulumi.runtime.register_resource_module("aws", "macie/memberAccountAssociation", _module_instance)
    pulumi.runtime.register_resource_module("aws", "macie/s3BucketAssociation", _module_instance)

_register_module()
