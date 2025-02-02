// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export * from "./portfolio";
export * from "./product";
export * from "./tagOption";

// Import resources to register:
import { Portfolio } from "./portfolio";
import { Product } from "./product";
import { TagOption } from "./tagOption";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "aws:servicecatalog/portfolio:Portfolio":
                return new Portfolio(name, <any>undefined, { urn })
            case "aws:servicecatalog/product:Product":
                return new Product(name, <any>undefined, { urn })
            case "aws:servicecatalog/tagOption:TagOption":
                return new TagOption(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("aws", "servicecatalog/portfolio", _module)
pulumi.runtime.registerResourceModule("aws", "servicecatalog/product", _module)
pulumi.runtime.registerResourceModule("aws", "servicecatalog/tagOption", _module)
