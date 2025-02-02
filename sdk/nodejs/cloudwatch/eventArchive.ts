// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides an EventBridge event archive resource.
 *
 * > **Note:** EventBridge was formerly known as CloudWatch Events. The functionality is identical.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const orderEventBus = new aws.cloudwatch.EventBus("orderEventBus", {});
 * const orderEventArchive = new aws.cloudwatch.EventArchive("orderEventArchive", {eventSourceArn: orderEventBus.arn});
 * ```
 * ## Example all optional arguments
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const orderEventBus = new aws.cloudwatch.EventBus("orderEventBus", {});
 * const orderEventArchive = new aws.cloudwatch.EventArchive("orderEventArchive", {
 *     description: "Archived events from order service",
 *     eventSourceArn: orderEventBus.arn,
 *     retentionDays: 7,
 *     eventPattern: `{
 *   "source": ["company.team.order"]
 * }
 * `,
 * });
 * ```
 *
 * ## Import
 *
 * Event Archive can be imported using their name, for example bash
 *
 * ```sh
 *  $ pulumi import aws:cloudwatch/eventArchive:EventArchive imported_event_archive order-archive
 * ```
 */
export class EventArchive extends pulumi.CustomResource {
    /**
     * Get an existing EventArchive resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: EventArchiveState, opts?: pulumi.CustomResourceOptions): EventArchive {
        return new EventArchive(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:cloudwatch/eventArchive:EventArchive';

    /**
     * Returns true if the given object is an instance of EventArchive.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is EventArchive {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === EventArchive.__pulumiType;
    }

    /**
     * The Amazon Resource Name (ARN) of the event archive.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * The description of the new event archive.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Instructs the new event archive to only capture events matched by this pattern. By default, it attempts to archive every event received in the `eventSourceArn`.
     */
    public readonly eventPattern!: pulumi.Output<string | undefined>;
    /**
     * Event bus source ARN from where these events should be archived.
     */
    public readonly eventSourceArn!: pulumi.Output<string>;
    /**
     * The name of the new event archive. The archive name cannot exceed 48 characters.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The maximum number of days to retain events in the new event archive. By default, it archives indefinitely.
     */
    public readonly retentionDays!: pulumi.Output<number | undefined>;

    /**
     * Create a EventArchive resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: EventArchiveArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: EventArchiveArgs | EventArchiveState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as EventArchiveState | undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["eventPattern"] = state ? state.eventPattern : undefined;
            inputs["eventSourceArn"] = state ? state.eventSourceArn : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["retentionDays"] = state ? state.retentionDays : undefined;
        } else {
            const args = argsOrState as EventArchiveArgs | undefined;
            if ((!args || args.eventSourceArn === undefined) && !opts.urn) {
                throw new Error("Missing required property 'eventSourceArn'");
            }
            inputs["description"] = args ? args.description : undefined;
            inputs["eventPattern"] = args ? args.eventPattern : undefined;
            inputs["eventSourceArn"] = args ? args.eventSourceArn : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["retentionDays"] = args ? args.retentionDays : undefined;
            inputs["arn"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(EventArchive.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering EventArchive resources.
 */
export interface EventArchiveState {
    /**
     * The Amazon Resource Name (ARN) of the event archive.
     */
    readonly arn?: pulumi.Input<string>;
    /**
     * The description of the new event archive.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Instructs the new event archive to only capture events matched by this pattern. By default, it attempts to archive every event received in the `eventSourceArn`.
     */
    readonly eventPattern?: pulumi.Input<string>;
    /**
     * Event bus source ARN from where these events should be archived.
     */
    readonly eventSourceArn?: pulumi.Input<string>;
    /**
     * The name of the new event archive. The archive name cannot exceed 48 characters.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The maximum number of days to retain events in the new event archive. By default, it archives indefinitely.
     */
    readonly retentionDays?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a EventArchive resource.
 */
export interface EventArchiveArgs {
    /**
     * The description of the new event archive.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Instructs the new event archive to only capture events matched by this pattern. By default, it attempts to archive every event received in the `eventSourceArn`.
     */
    readonly eventPattern?: pulumi.Input<string>;
    /**
     * Event bus source ARN from where these events should be archived.
     */
    readonly eventSourceArn: pulumi.Input<string>;
    /**
     * The name of the new event archive. The archive name cannot exceed 48 characters.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The maximum number of days to retain events in the new event archive. By default, it archives indefinitely.
     */
    readonly retentionDays?: pulumi.Input<number>;
}
