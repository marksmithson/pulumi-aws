// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Provides an EventBridge Target resource.
 *
 * > **Note:** EventBridge was formerly known as CloudWatch Events. The functionality is identical.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const console = new aws.cloudwatch.EventRule("console", {
 *     description: "Capture all EC2 scaling events",
 *     eventPattern: `{
 *   "source": [
 *     "aws.autoscaling"
 *   ],
 *   "detail-type": [
 *     "EC2 Instance Launch Successful",
 *     "EC2 Instance Terminate Successful",
 *     "EC2 Instance Launch Unsuccessful",
 *     "EC2 Instance Terminate Unsuccessful"
 *   ]
 * }
 * `,
 * });
 * const testStream = new aws.kinesis.Stream("testStream", {shardCount: 1});
 * const yada = new aws.cloudwatch.EventTarget("yada", {
 *     rule: console.name,
 *     arn: testStream.arn,
 *     runCommandTargets: [
 *         {
 *             key: "tag:Name",
 *             values: ["FooBar"],
 *         },
 *         {
 *             key: "InstanceIds",
 *             values: ["i-162058cd308bffec2"],
 *         },
 *     ],
 * });
 * ```
 * ## Example RunCommand Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const stopInstancesEventRule = new aws.cloudwatch.EventRule("stopInstancesEventRule", {
 *     description: "Stop instances nightly",
 *     scheduleExpression: "cron(0 0 * * ? *)",
 * });
 * const stopInstancesEventTarget = new aws.cloudwatch.EventTarget("stopInstancesEventTarget", {
 *     arn: `arn:aws:ssm:${_var.aws_region}::document/AWS-RunShellScript`,
 *     input: "{\"commands\":[\"halt\"]}",
 *     rule: stopInstancesEventRule.name,
 *     roleArn: aws_iam_role.ssm_lifecycle.arn,
 *     runCommandTargets: [{
 *         key: "tag:Terminate",
 *         values: ["midnight"],
 *     }],
 * });
 * ```
 *
 * ## Example API Gateway target
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const exampleEventRule = new aws.cloudwatch.EventRule("exampleEventRule", {});
 * // ...
 * const exampleDeployment = new aws.apigateway.Deployment("exampleDeployment", {restApi: aws_api_gateway_rest_api.example.id});
 * // ...
 * const exampleStage = new aws.apigateway.Stage("exampleStage", {
 *     restApi: aws_api_gateway_rest_api.example.id,
 *     deployment: exampleDeployment.id,
 * });
 * // ...
 * const exampleEventTarget = new aws.cloudwatch.EventTarget("exampleEventTarget", {
 *     arn: pulumi.interpolate`${exampleStage.executionArn}/GET`,
 *     rule: exampleEventRule.id,
 *     httpTarget: {
 *         queryStringParameters: {
 *             Body: `$.detail.body`,
 *         },
 *         headerParameters: {
 *             Env: "Test",
 *         },
 *     },
 * });
 * ```
 *
 * ## Example Input Transformer Usage - JSON Object
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const exampleEventRule = new aws.cloudwatch.EventRule("exampleEventRule", {});
 * // ...
 * const exampleEventTarget = new aws.cloudwatch.EventTarget("exampleEventTarget", {
 *     arn: aws_lambda_function.example.arn,
 *     rule: exampleEventRule.id,
 *     inputTransformer: {
 *         inputPaths: {
 *             instance: `$.detail.instance`,
 *             status: `$.detail.status`,
 *         },
 *         inputTemplate: `{
 *   "instance_id": <instance>,
 *   "instance_status": <status>
 * }
 * `,
 *     },
 * });
 * ```
 *
 * ## Example Input Transformer Usage - Simple String
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const exampleEventRule = new aws.cloudwatch.EventRule("exampleEventRule", {});
 * // ...
 * const exampleEventTarget = new aws.cloudwatch.EventTarget("exampleEventTarget", {
 *     arn: aws_lambda_function.example.arn,
 *     rule: exampleEventRule.id,
 *     inputTransformer: {
 *         inputPaths: {
 *             instance: `$.detail.instance`,
 *             status: `$.detail.status`,
 *         },
 *         inputTemplate: "\"<instance> is in state <status>\"",
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * EventBridge Targets can be imported using `event_bus_name/rule-name/target-id` (if you omit `event_bus_name`, the `default` event bus will be used).
 *
 * ```sh
 *  $ pulumi import aws:cloudwatch/eventTarget:EventTarget test-event-target rule-name/target-id
 * ```
 */
export class EventTarget extends pulumi.CustomResource {
    /**
     * Get an existing EventTarget resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: EventTargetState, opts?: pulumi.CustomResourceOptions): EventTarget {
        return new EventTarget(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:cloudwatch/eventTarget:EventTarget';

    /**
     * Returns true if the given object is an instance of EventTarget.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is EventTarget {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === EventTarget.__pulumiType;
    }

    /**
     * The Amazon Resource Name (ARN) of the target.
     */
    public readonly arn!: pulumi.Output<string>;
    /**
     * Parameters used when you are using the rule to invoke an Amazon Batch Job. Documented below. A maximum of 1 are allowed.
     */
    public readonly batchTarget!: pulumi.Output<outputs.cloudwatch.EventTargetBatchTarget | undefined>;
    /**
     * Parameters used when you are providing a dead letter config. Documented below. A maximum of 1 are allowed.
     */
    public readonly deadLetterConfig!: pulumi.Output<outputs.cloudwatch.EventTargetDeadLetterConfig | undefined>;
    /**
     * Parameters used when you are using the rule to invoke Amazon ECS Task. Documented below. A maximum of 1 are allowed.
     */
    public readonly ecsTarget!: pulumi.Output<outputs.cloudwatch.EventTargetEcsTarget | undefined>;
    /**
     * The event bus to associate with the rule. If you omit this, the `default` event bus is used.
     */
    public readonly eventBusName!: pulumi.Output<string | undefined>;
    /**
     * Parameters used when you are using the rule to invoke an API Gateway REST endpoint. Documented below. A maximum of 1 is allowed.
     */
    public readonly httpTarget!: pulumi.Output<outputs.cloudwatch.EventTargetHttpTarget | undefined>;
    /**
     * Valid JSON text passed to the target. Conflicts with `inputPath` and `inputTransformer`.
     */
    public readonly input!: pulumi.Output<string | undefined>;
    /**
     * The value of the [JSONPath](http://goessner.net/articles/JsonPath/) that is used for extracting part of the matched event when passing it to the target. Conflicts with `input` and `inputTransformer`.
     */
    public readonly inputPath!: pulumi.Output<string | undefined>;
    /**
     * Parameters used when you are providing a custom input to a target based on certain event data. Conflicts with `input` and `inputPath`.
     */
    public readonly inputTransformer!: pulumi.Output<outputs.cloudwatch.EventTargetInputTransformer | undefined>;
    /**
     * Parameters used when you are using the rule to invoke an Amazon Kinesis Stream. Documented below. A maximum of 1 are allowed.
     */
    public readonly kinesisTarget!: pulumi.Output<outputs.cloudwatch.EventTargetKinesisTarget | undefined>;
    /**
     * Parameters used when you are providing retry policies. Documented below. A maximum of 1 are allowed.
     */
    public readonly retryPolicy!: pulumi.Output<outputs.cloudwatch.EventTargetRetryPolicy | undefined>;
    /**
     * The Amazon Resource Name (ARN) of the IAM role to be used for this target when the rule is triggered. Required if `ecsTarget` is used.
     */
    public readonly roleArn!: pulumi.Output<string | undefined>;
    /**
     * The name of the rule you want to add targets to.
     */
    public readonly rule!: pulumi.Output<string>;
    /**
     * Parameters used when you are using the rule to invoke Amazon EC2 Run Command. Documented below. A maximum of 5 are allowed.
     */
    public readonly runCommandTargets!: pulumi.Output<outputs.cloudwatch.EventTargetRunCommandTarget[] | undefined>;
    /**
     * Parameters used when you are using the rule to invoke an Amazon SQS Queue. Documented below. A maximum of 1 are allowed.
     */
    public readonly sqsTarget!: pulumi.Output<outputs.cloudwatch.EventTargetSqsTarget | undefined>;
    /**
     * The unique target assignment ID.  If missing, will generate a random, unique id.
     */
    public readonly targetId!: pulumi.Output<string>;

    /**
     * Create a EventTarget resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: EventTargetArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: EventTargetArgs | EventTargetState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as EventTargetState | undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["batchTarget"] = state ? state.batchTarget : undefined;
            inputs["deadLetterConfig"] = state ? state.deadLetterConfig : undefined;
            inputs["ecsTarget"] = state ? state.ecsTarget : undefined;
            inputs["eventBusName"] = state ? state.eventBusName : undefined;
            inputs["httpTarget"] = state ? state.httpTarget : undefined;
            inputs["input"] = state ? state.input : undefined;
            inputs["inputPath"] = state ? state.inputPath : undefined;
            inputs["inputTransformer"] = state ? state.inputTransformer : undefined;
            inputs["kinesisTarget"] = state ? state.kinesisTarget : undefined;
            inputs["retryPolicy"] = state ? state.retryPolicy : undefined;
            inputs["roleArn"] = state ? state.roleArn : undefined;
            inputs["rule"] = state ? state.rule : undefined;
            inputs["runCommandTargets"] = state ? state.runCommandTargets : undefined;
            inputs["sqsTarget"] = state ? state.sqsTarget : undefined;
            inputs["targetId"] = state ? state.targetId : undefined;
        } else {
            const args = argsOrState as EventTargetArgs | undefined;
            if ((!args || args.arn === undefined) && !opts.urn) {
                throw new Error("Missing required property 'arn'");
            }
            if ((!args || args.rule === undefined) && !opts.urn) {
                throw new Error("Missing required property 'rule'");
            }
            inputs["arn"] = args ? args.arn : undefined;
            inputs["batchTarget"] = args ? args.batchTarget : undefined;
            inputs["deadLetterConfig"] = args ? args.deadLetterConfig : undefined;
            inputs["ecsTarget"] = args ? args.ecsTarget : undefined;
            inputs["eventBusName"] = args ? args.eventBusName : undefined;
            inputs["httpTarget"] = args ? args.httpTarget : undefined;
            inputs["input"] = args ? args.input : undefined;
            inputs["inputPath"] = args ? args.inputPath : undefined;
            inputs["inputTransformer"] = args ? args.inputTransformer : undefined;
            inputs["kinesisTarget"] = args ? args.kinesisTarget : undefined;
            inputs["retryPolicy"] = args ? args.retryPolicy : undefined;
            inputs["roleArn"] = args ? args.roleArn : undefined;
            inputs["rule"] = args ? args.rule : undefined;
            inputs["runCommandTargets"] = args ? args.runCommandTargets : undefined;
            inputs["sqsTarget"] = args ? args.sqsTarget : undefined;
            inputs["targetId"] = args ? args.targetId : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(EventTarget.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering EventTarget resources.
 */
export interface EventTargetState {
    /**
     * The Amazon Resource Name (ARN) of the target.
     */
    readonly arn?: pulumi.Input<string>;
    /**
     * Parameters used when you are using the rule to invoke an Amazon Batch Job. Documented below. A maximum of 1 are allowed.
     */
    readonly batchTarget?: pulumi.Input<inputs.cloudwatch.EventTargetBatchTarget>;
    /**
     * Parameters used when you are providing a dead letter config. Documented below. A maximum of 1 are allowed.
     */
    readonly deadLetterConfig?: pulumi.Input<inputs.cloudwatch.EventTargetDeadLetterConfig>;
    /**
     * Parameters used when you are using the rule to invoke Amazon ECS Task. Documented below. A maximum of 1 are allowed.
     */
    readonly ecsTarget?: pulumi.Input<inputs.cloudwatch.EventTargetEcsTarget>;
    /**
     * The event bus to associate with the rule. If you omit this, the `default` event bus is used.
     */
    readonly eventBusName?: pulumi.Input<string>;
    /**
     * Parameters used when you are using the rule to invoke an API Gateway REST endpoint. Documented below. A maximum of 1 is allowed.
     */
    readonly httpTarget?: pulumi.Input<inputs.cloudwatch.EventTargetHttpTarget>;
    /**
     * Valid JSON text passed to the target. Conflicts with `inputPath` and `inputTransformer`.
     */
    readonly input?: pulumi.Input<string>;
    /**
     * The value of the [JSONPath](http://goessner.net/articles/JsonPath/) that is used for extracting part of the matched event when passing it to the target. Conflicts with `input` and `inputTransformer`.
     */
    readonly inputPath?: pulumi.Input<string>;
    /**
     * Parameters used when you are providing a custom input to a target based on certain event data. Conflicts with `input` and `inputPath`.
     */
    readonly inputTransformer?: pulumi.Input<inputs.cloudwatch.EventTargetInputTransformer>;
    /**
     * Parameters used when you are using the rule to invoke an Amazon Kinesis Stream. Documented below. A maximum of 1 are allowed.
     */
    readonly kinesisTarget?: pulumi.Input<inputs.cloudwatch.EventTargetKinesisTarget>;
    /**
     * Parameters used when you are providing retry policies. Documented below. A maximum of 1 are allowed.
     */
    readonly retryPolicy?: pulumi.Input<inputs.cloudwatch.EventTargetRetryPolicy>;
    /**
     * The Amazon Resource Name (ARN) of the IAM role to be used for this target when the rule is triggered. Required if `ecsTarget` is used.
     */
    readonly roleArn?: pulumi.Input<string>;
    /**
     * The name of the rule you want to add targets to.
     */
    readonly rule?: pulumi.Input<string>;
    /**
     * Parameters used when you are using the rule to invoke Amazon EC2 Run Command. Documented below. A maximum of 5 are allowed.
     */
    readonly runCommandTargets?: pulumi.Input<pulumi.Input<inputs.cloudwatch.EventTargetRunCommandTarget>[]>;
    /**
     * Parameters used when you are using the rule to invoke an Amazon SQS Queue. Documented below. A maximum of 1 are allowed.
     */
    readonly sqsTarget?: pulumi.Input<inputs.cloudwatch.EventTargetSqsTarget>;
    /**
     * The unique target assignment ID.  If missing, will generate a random, unique id.
     */
    readonly targetId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a EventTarget resource.
 */
export interface EventTargetArgs {
    /**
     * The Amazon Resource Name (ARN) of the target.
     */
    readonly arn: pulumi.Input<string>;
    /**
     * Parameters used when you are using the rule to invoke an Amazon Batch Job. Documented below. A maximum of 1 are allowed.
     */
    readonly batchTarget?: pulumi.Input<inputs.cloudwatch.EventTargetBatchTarget>;
    /**
     * Parameters used when you are providing a dead letter config. Documented below. A maximum of 1 are allowed.
     */
    readonly deadLetterConfig?: pulumi.Input<inputs.cloudwatch.EventTargetDeadLetterConfig>;
    /**
     * Parameters used when you are using the rule to invoke Amazon ECS Task. Documented below. A maximum of 1 are allowed.
     */
    readonly ecsTarget?: pulumi.Input<inputs.cloudwatch.EventTargetEcsTarget>;
    /**
     * The event bus to associate with the rule. If you omit this, the `default` event bus is used.
     */
    readonly eventBusName?: pulumi.Input<string>;
    /**
     * Parameters used when you are using the rule to invoke an API Gateway REST endpoint. Documented below. A maximum of 1 is allowed.
     */
    readonly httpTarget?: pulumi.Input<inputs.cloudwatch.EventTargetHttpTarget>;
    /**
     * Valid JSON text passed to the target. Conflicts with `inputPath` and `inputTransformer`.
     */
    readonly input?: pulumi.Input<string>;
    /**
     * The value of the [JSONPath](http://goessner.net/articles/JsonPath/) that is used for extracting part of the matched event when passing it to the target. Conflicts with `input` and `inputTransformer`.
     */
    readonly inputPath?: pulumi.Input<string>;
    /**
     * Parameters used when you are providing a custom input to a target based on certain event data. Conflicts with `input` and `inputPath`.
     */
    readonly inputTransformer?: pulumi.Input<inputs.cloudwatch.EventTargetInputTransformer>;
    /**
     * Parameters used when you are using the rule to invoke an Amazon Kinesis Stream. Documented below. A maximum of 1 are allowed.
     */
    readonly kinesisTarget?: pulumi.Input<inputs.cloudwatch.EventTargetKinesisTarget>;
    /**
     * Parameters used when you are providing retry policies. Documented below. A maximum of 1 are allowed.
     */
    readonly retryPolicy?: pulumi.Input<inputs.cloudwatch.EventTargetRetryPolicy>;
    /**
     * The Amazon Resource Name (ARN) of the IAM role to be used for this target when the rule is triggered. Required if `ecsTarget` is used.
     */
    readonly roleArn?: pulumi.Input<string>;
    /**
     * The name of the rule you want to add targets to.
     */
    readonly rule: pulumi.Input<string>;
    /**
     * Parameters used when you are using the rule to invoke Amazon EC2 Run Command. Documented below. A maximum of 5 are allowed.
     */
    readonly runCommandTargets?: pulumi.Input<pulumi.Input<inputs.cloudwatch.EventTargetRunCommandTarget>[]>;
    /**
     * Parameters used when you are using the rule to invoke an Amazon SQS Queue. Documented below. A maximum of 1 are allowed.
     */
    readonly sqsTarget?: pulumi.Input<inputs.cloudwatch.EventTargetSqsTarget>;
    /**
     * The unique target assignment ID.  If missing, will generate a random, unique id.
     */
    readonly targetId?: pulumi.Input<string>;
}
