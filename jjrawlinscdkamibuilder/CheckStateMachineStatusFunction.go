package jjrawlinscdkamibuilder

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/jjrawlins/cdk-ami-builder/jjrawlinscdkamibuilder/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssns"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssqs"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/jjrawlins/cdk-ami-builder/jjrawlinscdkamibuilder/internal"
)

// An AWS Lambda function which executes src/Lambdas/CheckStateMachineStatus/CheckStateMachineStatus.
type CheckStateMachineStatusFunction interface {
	awslambda.Function
	// The architecture of this Lambda Function (this is an optional attribute and defaults to X86_64).
	Architecture() awslambda.Architecture
	// Whether the addPermission() call adds any permissions.
	//
	// True for new Lambdas, false for version $LATEST and imported Lambdas
	// from different accounts.
	CanCreatePermissions() *bool
	// Access the Connections object.
	//
	// Will fail if not a VPC-enabled Lambda Function.
	Connections() awsec2.Connections
	// Returns a `lambda.Version` which represents the current version of this Lambda function. A new version will be created every time the function's configuration changes.
	//
	// You can specify options for this version using the `currentVersionOptions`
	// prop when initializing the `lambda.Function`.
	CurrentVersion() awslambda.Version
	// The DLQ (as queue) associated with this Lambda Function (this is an optional attribute).
	DeadLetterQueue() awssqs.IQueue
	// The DLQ (as topic) associated with this Lambda Function (this is an optional attribute).
	DeadLetterTopic() awssns.ITopic
	// The environment this resource belongs to.
	//
	// For resources that are created and managed by the CDK
	// (generally, those created by creating new class instances like Role, Bucket, etc.),
	// this is always the same as the environment of the stack they belong to;
	// however, for imported resources
	// (those obtained from static methods like fromRoleArn, fromBucketName, etc.),
	// that might be different than the stack they were imported into.
	Env() *awscdk.ResourceEnvironment
	// ARN of this function.
	FunctionArn() *string
	// Name of this function.
	FunctionName() *string
	// The principal this Lambda Function is running as.
	GrantPrincipal() awsiam.IPrincipal
	// Whether or not this Lambda function was bound to a VPC.
	//
	// If this is is `false`, trying to access the `connections` object will fail.
	IsBoundToVpc() *bool
	// The `$LATEST` version of this function.
	//
	// Note that this is reference to a non-specific AWS Lambda version, which
	// means the function this version refers to can return different results in
	// different invocations.
	//
	// To obtain a reference to an explicit version which references the current
	// function configuration, use `lambdaFunction.currentVersion` instead.
	LatestVersion() awslambda.IVersion
	// The LogGroup where the Lambda function's logs are made available.
	//
	// If either `logRetention` is set or this property is called, a CloudFormation custom resource is added to the stack that
	// pre-creates the log group as part of the stack deployment, if it already doesn't exist, and sets the correct log retention
	// period (never expire, by default).
	//
	// Further, if the log group already exists and the `logRetention` is not set, the custom resource will reset the log retention
	// to never expire even if it was configured with a different value.
	LogGroup() awslogs.ILogGroup
	// The tree node.
	Node() constructs.Node
	// The construct node where permissions are attached.
	PermissionsNode() constructs.Node
	// Returns a string-encoded token that resolves to the physical name that should be passed to the CloudFormation resource.
	//
	// This value will resolve to one of the following:
	// - a concrete value (e.g. `"my-awesome-bucket"`)
	// - `undefined`, when a name should be generated by CloudFormation
	// - a concrete name generated automatically during synthesis, in
	//   cross-environment scenarios.
	PhysicalName() *string
	// The ARN(s) to put into the resource field of the generated IAM policy for grantInvoke().
	ResourceArnsForGrantInvoke() *[]*string
	// Execution role associated with this function.
	Role() awsiam.IRole
	// The runtime configured for this lambda.
	Runtime() awslambda.Runtime
	// The stack in which this resource is defined.
	Stack() awscdk.Stack
	// The timeout configured for this lambda.
	Timeout() awscdk.Duration
	// Defines an alias for this function.
	//
	// The alias will automatically be updated to point to the latest version of
	// the function as it is being updated during a deployment.
	//
	// ```ts
	// declare const fn: lambda.Function;
	//
	// fn.addAlias('Live');
	//
	// // Is equivalent to
	//
	// new lambda.Alias(this, 'AliasLive', {
	//   aliasName: 'Live',
	//   version: fn.currentVersion,
	// });
	// ```.
	AddAlias(aliasName *string, options *awslambda.AliasOptions) awslambda.Alias
	// Adds an environment variable to this Lambda function.
	//
	// If this is a ref to a Lambda function, this operation results in a no-op.
	AddEnvironment(key *string, value *string, options *awslambda.EnvironmentOptions) awslambda.Function
	// Adds an event source to this function.
	//
	// Event sources are implemented in the aws-cdk-lib/aws-lambda-event-sources module.
	//
	// The following example adds an SQS Queue as an event source:
	// ```
	// import { SqsEventSource } from 'aws-cdk-lib/aws-lambda-event-sources';
	// myFunction.addEventSource(new SqsEventSource(myQueue));
	// ```.
	AddEventSource(source awslambda.IEventSource)
	// Adds an event source that maps to this AWS Lambda function.
	AddEventSourceMapping(id *string, options *awslambda.EventSourceMappingOptions) awslambda.EventSourceMapping
	// Adds a url to this lambda function.
	AddFunctionUrl(options *awslambda.FunctionUrlOptions) awslambda.FunctionUrl
	// Adds one or more Lambda Layers to this Lambda function.
	AddLayers(layers ...awslambda.ILayerVersion)
	// Adds a permission to the Lambda resource policy.
	// See: Permission for details.
	//
	AddPermission(id *string, permission *awslambda.Permission)
	// Adds a statement to the IAM role assumed by the instance.
	AddToRolePolicy(statement awsiam.PolicyStatement)
	// Apply the given removal policy to this resource.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	// Configures options for asynchronous invocation.
	ConfigureAsyncInvoke(options *awslambda.EventInvokeConfigOptions)
	// A warning will be added to functions under the following conditions: - permissions that include `lambda:InvokeFunction` are added to the unqualified function.
	//
	// - function.currentVersion is invoked before or after the permission is created.
	//
	// This applies only to permissions on Lambda functions, not versions or aliases.
	// This function is overridden as a noOp for QualifiedFunctionBase.
	ConsiderWarningOnInvokeFunctionPermissions(scope constructs.Construct, action *string)
	GeneratePhysicalName() *string
	// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
	//
	// Normally, this token will resolve to `arnAttr`, but if the resource is
	// referenced across environments, `arnComponents` will be used to synthesize
	// a concrete ARN with the resource's physical name. Make sure to reference
	// `this.physicalName` in `arnComponents`.
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
	//
	// Normally, this token will resolve to `nameAttr`, but if the resource is
	// referenced across environments, it will be resolved to `this.physicalName`,
	// which will be a concrete name.
	GetResourceNameAttribute(nameAttr *string) *string
	// Grant the given identity permissions to invoke this Lambda.
	GrantInvoke(grantee awsiam.IGrantable) awsiam.Grant
	// Grant multiple principals the ability to invoke this Lambda via CompositePrincipal.
	GrantInvokeCompositePrincipal(compositePrincipal awsiam.CompositePrincipal) *[]awsiam.Grant
	// Grant the given identity permissions to invoke the $LATEST version or unqualified version of this Lambda.
	GrantInvokeLatestVersion(grantee awsiam.IGrantable) awsiam.Grant
	// Grant the given identity permissions to invoke this Lambda Function URL.
	GrantInvokeUrl(grantee awsiam.IGrantable) awsiam.Grant
	// Grant the given identity permissions to invoke the given version of this Lambda.
	GrantInvokeVersion(grantee awsiam.IGrantable, version awslambda.IVersion) awsiam.Grant
	// Mix additional information into the hash of the Version object.
	//
	// The Lambda Function construct does its best to automatically create a new
	// Version when anything about the Function changes (its code, its layers,
	// any of the other properties).
	//
	// However, you can sometimes source information from places that the CDK cannot
	// look into, like the deploy-time values of SSM parameters. In those cases,
	// the CDK would not force the creation of a new Version object when it actually
	// should.
	//
	// This method can be used to invalidate the current Version object. Pass in
	// any string into this method, and make sure the string changes when you know
	// a new Version needs to be created.
	//
	// This method may be called more than once.
	InvalidateVersionBasedOn(x *string)
	// Return the given named metric for this Function.
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// How long execution of this Lambda takes.
	//
	// Average over 5 minutes.
	MetricDuration(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// How many invocations of this Lambda fail.
	//
	// Sum over 5 minutes.
	MetricErrors(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// How often this Lambda is invoked.
	//
	// Sum over 5 minutes.
	MetricInvocations(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// How often this Lambda is throttled.
	//
	// Sum over 5 minutes.
	MetricThrottles(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Returns a string representation of this construct.
	ToString() *string
	WarnInvokeFunctionPermissions(scope constructs.Construct)
}

// The jsii proxy struct for CheckStateMachineStatusFunction
type jsiiProxy_CheckStateMachineStatusFunction struct {
	internal.Type__awslambdaFunction
}

func (j *jsiiProxy_CheckStateMachineStatusFunction) Architecture() awslambda.Architecture {
	var returns awslambda.Architecture
	_jsii_.Get(
		j,
		"architecture",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CheckStateMachineStatusFunction) CanCreatePermissions() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"canCreatePermissions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CheckStateMachineStatusFunction) Connections() awsec2.Connections {
	var returns awsec2.Connections
	_jsii_.Get(
		j,
		"connections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CheckStateMachineStatusFunction) CurrentVersion() awslambda.Version {
	var returns awslambda.Version
	_jsii_.Get(
		j,
		"currentVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CheckStateMachineStatusFunction) DeadLetterQueue() awssqs.IQueue {
	var returns awssqs.IQueue
	_jsii_.Get(
		j,
		"deadLetterQueue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CheckStateMachineStatusFunction) DeadLetterTopic() awssns.ITopic {
	var returns awssns.ITopic
	_jsii_.Get(
		j,
		"deadLetterTopic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CheckStateMachineStatusFunction) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CheckStateMachineStatusFunction) FunctionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CheckStateMachineStatusFunction) FunctionName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CheckStateMachineStatusFunction) GrantPrincipal() awsiam.IPrincipal {
	var returns awsiam.IPrincipal
	_jsii_.Get(
		j,
		"grantPrincipal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CheckStateMachineStatusFunction) IsBoundToVpc() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isBoundToVpc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CheckStateMachineStatusFunction) LatestVersion() awslambda.IVersion {
	var returns awslambda.IVersion
	_jsii_.Get(
		j,
		"latestVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CheckStateMachineStatusFunction) LogGroup() awslogs.ILogGroup {
	var returns awslogs.ILogGroup
	_jsii_.Get(
		j,
		"logGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CheckStateMachineStatusFunction) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CheckStateMachineStatusFunction) PermissionsNode() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"permissionsNode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CheckStateMachineStatusFunction) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CheckStateMachineStatusFunction) ResourceArnsForGrantInvoke() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resourceArnsForGrantInvoke",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CheckStateMachineStatusFunction) Role() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"role",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CheckStateMachineStatusFunction) Runtime() awslambda.Runtime {
	var returns awslambda.Runtime
	_jsii_.Get(
		j,
		"runtime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CheckStateMachineStatusFunction) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CheckStateMachineStatusFunction) Timeout() awscdk.Duration {
	var returns awscdk.Duration
	_jsii_.Get(
		j,
		"timeout",
		&returns,
	)
	return returns
}


func NewCheckStateMachineStatusFunction(scope constructs.Construct, id *string, props *CheckStateMachineStatusFunctionProps) CheckStateMachineStatusFunction {
	_init_.Initialize()

	if err := validateNewCheckStateMachineStatusFunctionParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CheckStateMachineStatusFunction{}

	_jsii_.Create(
		"@jjrawlins/cdk-ami-builder.CheckStateMachineStatusFunction",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewCheckStateMachineStatusFunction_Override(c CheckStateMachineStatusFunction, scope constructs.Construct, id *string, props *CheckStateMachineStatusFunctionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@jjrawlins/cdk-ami-builder.CheckStateMachineStatusFunction",
		[]interface{}{scope, id, props},
		c,
	)
}

// Record whether specific properties in the `AWS::Lambda::Function` resource should also be associated to the Version resource.
//
// See 'currentVersion' section in the module README for more details.
func CheckStateMachineStatusFunction_ClassifyVersionProperty(propertyName *string, locked *bool) {
	_init_.Initialize()

	if err := validateCheckStateMachineStatusFunction_ClassifyVersionPropertyParameters(propertyName, locked); err != nil {
		panic(err)
	}
	_jsii_.StaticInvokeVoid(
		"@jjrawlins/cdk-ami-builder.CheckStateMachineStatusFunction",
		"classifyVersionProperty",
		[]interface{}{propertyName, locked},
	)
}

// Import a lambda function into the CDK using its ARN.
//
// For `Function.addPermissions()` to work on this imported lambda, make sure that is
// in the same account and region as the stack you are importing it into.
func CheckStateMachineStatusFunction_FromFunctionArn(scope constructs.Construct, id *string, functionArn *string) awslambda.IFunction {
	_init_.Initialize()

	if err := validateCheckStateMachineStatusFunction_FromFunctionArnParameters(scope, id, functionArn); err != nil {
		panic(err)
	}
	var returns awslambda.IFunction

	_jsii_.StaticInvoke(
		"@jjrawlins/cdk-ami-builder.CheckStateMachineStatusFunction",
		"fromFunctionArn",
		[]interface{}{scope, id, functionArn},
		&returns,
	)

	return returns
}

// Creates a Lambda function object which represents a function not defined within this stack.
//
// For `Function.addPermissions()` to work on this imported lambda, set the sameEnvironment property to true
// if this imported lambda is in the same account and region as the stack you are importing it into.
func CheckStateMachineStatusFunction_FromFunctionAttributes(scope constructs.Construct, id *string, attrs *awslambda.FunctionAttributes) awslambda.IFunction {
	_init_.Initialize()

	if err := validateCheckStateMachineStatusFunction_FromFunctionAttributesParameters(scope, id, attrs); err != nil {
		panic(err)
	}
	var returns awslambda.IFunction

	_jsii_.StaticInvoke(
		"@jjrawlins/cdk-ami-builder.CheckStateMachineStatusFunction",
		"fromFunctionAttributes",
		[]interface{}{scope, id, attrs},
		&returns,
	)

	return returns
}

// Import a lambda function into the CDK using its name.
func CheckStateMachineStatusFunction_FromFunctionName(scope constructs.Construct, id *string, functionName *string) awslambda.IFunction {
	_init_.Initialize()

	if err := validateCheckStateMachineStatusFunction_FromFunctionNameParameters(scope, id, functionName); err != nil {
		panic(err)
	}
	var returns awslambda.IFunction

	_jsii_.StaticInvoke(
		"@jjrawlins/cdk-ami-builder.CheckStateMachineStatusFunction",
		"fromFunctionName",
		[]interface{}{scope, id, functionName},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func CheckStateMachineStatusFunction_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCheckStateMachineStatusFunction_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@jjrawlins/cdk-ami-builder.CheckStateMachineStatusFunction",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns true if the construct was created by CDK, and false otherwise.
func CheckStateMachineStatusFunction_IsOwnedResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateCheckStateMachineStatusFunction_IsOwnedResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@jjrawlins/cdk-ami-builder.CheckStateMachineStatusFunction",
		"isOwnedResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func CheckStateMachineStatusFunction_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateCheckStateMachineStatusFunction_IsResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@jjrawlins/cdk-ami-builder.CheckStateMachineStatusFunction",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return the given named metric for this Lambda.
func CheckStateMachineStatusFunction_MetricAll(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	_init_.Initialize()

	if err := validateCheckStateMachineStatusFunction_MetricAllParameters(metricName, props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.StaticInvoke(
		"@jjrawlins/cdk-ami-builder.CheckStateMachineStatusFunction",
		"metricAll",
		[]interface{}{metricName, props},
		&returns,
	)

	return returns
}

// Metric for the number of concurrent executions across all Lambdas.
// Default: max over 5 minutes.
//
func CheckStateMachineStatusFunction_MetricAllConcurrentExecutions(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	_init_.Initialize()

	if err := validateCheckStateMachineStatusFunction_MetricAllConcurrentExecutionsParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.StaticInvoke(
		"@jjrawlins/cdk-ami-builder.CheckStateMachineStatusFunction",
		"metricAllConcurrentExecutions",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the Duration executing all Lambdas.
// Default: average over 5 minutes.
//
func CheckStateMachineStatusFunction_MetricAllDuration(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	_init_.Initialize()

	if err := validateCheckStateMachineStatusFunction_MetricAllDurationParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.StaticInvoke(
		"@jjrawlins/cdk-ami-builder.CheckStateMachineStatusFunction",
		"metricAllDuration",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of Errors executing all Lambdas.
// Default: sum over 5 minutes.
//
func CheckStateMachineStatusFunction_MetricAllErrors(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	_init_.Initialize()

	if err := validateCheckStateMachineStatusFunction_MetricAllErrorsParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.StaticInvoke(
		"@jjrawlins/cdk-ami-builder.CheckStateMachineStatusFunction",
		"metricAllErrors",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of invocations of all Lambdas.
// Default: sum over 5 minutes.
//
func CheckStateMachineStatusFunction_MetricAllInvocations(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	_init_.Initialize()

	if err := validateCheckStateMachineStatusFunction_MetricAllInvocationsParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.StaticInvoke(
		"@jjrawlins/cdk-ami-builder.CheckStateMachineStatusFunction",
		"metricAllInvocations",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of throttled invocations of all Lambdas.
// Default: sum over 5 minutes.
//
func CheckStateMachineStatusFunction_MetricAllThrottles(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	_init_.Initialize()

	if err := validateCheckStateMachineStatusFunction_MetricAllThrottlesParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.StaticInvoke(
		"@jjrawlins/cdk-ami-builder.CheckStateMachineStatusFunction",
		"metricAllThrottles",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of unreserved concurrent executions across all Lambdas.
// Default: max over 5 minutes.
//
func CheckStateMachineStatusFunction_MetricAllUnreservedConcurrentExecutions(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	_init_.Initialize()

	if err := validateCheckStateMachineStatusFunction_MetricAllUnreservedConcurrentExecutionsParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.StaticInvoke(
		"@jjrawlins/cdk-ami-builder.CheckStateMachineStatusFunction",
		"metricAllUnreservedConcurrentExecutions",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CheckStateMachineStatusFunction) AddAlias(aliasName *string, options *awslambda.AliasOptions) awslambda.Alias {
	if err := c.validateAddAliasParameters(aliasName, options); err != nil {
		panic(err)
	}
	var returns awslambda.Alias

	_jsii_.Invoke(
		c,
		"addAlias",
		[]interface{}{aliasName, options},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CheckStateMachineStatusFunction) AddEnvironment(key *string, value *string, options *awslambda.EnvironmentOptions) awslambda.Function {
	if err := c.validateAddEnvironmentParameters(key, value, options); err != nil {
		panic(err)
	}
	var returns awslambda.Function

	_jsii_.Invoke(
		c,
		"addEnvironment",
		[]interface{}{key, value, options},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CheckStateMachineStatusFunction) AddEventSource(source awslambda.IEventSource) {
	if err := c.validateAddEventSourceParameters(source); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addEventSource",
		[]interface{}{source},
	)
}

func (c *jsiiProxy_CheckStateMachineStatusFunction) AddEventSourceMapping(id *string, options *awslambda.EventSourceMappingOptions) awslambda.EventSourceMapping {
	if err := c.validateAddEventSourceMappingParameters(id, options); err != nil {
		panic(err)
	}
	var returns awslambda.EventSourceMapping

	_jsii_.Invoke(
		c,
		"addEventSourceMapping",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CheckStateMachineStatusFunction) AddFunctionUrl(options *awslambda.FunctionUrlOptions) awslambda.FunctionUrl {
	if err := c.validateAddFunctionUrlParameters(options); err != nil {
		panic(err)
	}
	var returns awslambda.FunctionUrl

	_jsii_.Invoke(
		c,
		"addFunctionUrl",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CheckStateMachineStatusFunction) AddLayers(layers ...awslambda.ILayerVersion) {
	args := []interface{}{}
	for _, a := range layers {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		c,
		"addLayers",
		args,
	)
}

func (c *jsiiProxy_CheckStateMachineStatusFunction) AddPermission(id *string, permission *awslambda.Permission) {
	if err := c.validateAddPermissionParameters(id, permission); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPermission",
		[]interface{}{id, permission},
	)
}

func (c *jsiiProxy_CheckStateMachineStatusFunction) AddToRolePolicy(statement awsiam.PolicyStatement) {
	if err := c.validateAddToRolePolicyParameters(statement); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addToRolePolicy",
		[]interface{}{statement},
	)
}

func (c *jsiiProxy_CheckStateMachineStatusFunction) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := c.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (c *jsiiProxy_CheckStateMachineStatusFunction) ConfigureAsyncInvoke(options *awslambda.EventInvokeConfigOptions) {
	if err := c.validateConfigureAsyncInvokeParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"configureAsyncInvoke",
		[]interface{}{options},
	)
}

func (c *jsiiProxy_CheckStateMachineStatusFunction) ConsiderWarningOnInvokeFunctionPermissions(scope constructs.Construct, action *string) {
	if err := c.validateConsiderWarningOnInvokeFunctionPermissionsParameters(scope, action); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"considerWarningOnInvokeFunctionPermissions",
		[]interface{}{scope, action},
	)
}

func (c *jsiiProxy_CheckStateMachineStatusFunction) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CheckStateMachineStatusFunction) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	if err := c.validateGetResourceArnAttributeParameters(arnAttr, arnComponents); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CheckStateMachineStatusFunction) GetResourceNameAttribute(nameAttr *string) *string {
	if err := c.validateGetResourceNameAttributeParameters(nameAttr); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CheckStateMachineStatusFunction) GrantInvoke(grantee awsiam.IGrantable) awsiam.Grant {
	if err := c.validateGrantInvokeParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		c,
		"grantInvoke",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CheckStateMachineStatusFunction) GrantInvokeCompositePrincipal(compositePrincipal awsiam.CompositePrincipal) *[]awsiam.Grant {
	if err := c.validateGrantInvokeCompositePrincipalParameters(compositePrincipal); err != nil {
		panic(err)
	}
	var returns *[]awsiam.Grant

	_jsii_.Invoke(
		c,
		"grantInvokeCompositePrincipal",
		[]interface{}{compositePrincipal},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CheckStateMachineStatusFunction) GrantInvokeLatestVersion(grantee awsiam.IGrantable) awsiam.Grant {
	if err := c.validateGrantInvokeLatestVersionParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		c,
		"grantInvokeLatestVersion",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CheckStateMachineStatusFunction) GrantInvokeUrl(grantee awsiam.IGrantable) awsiam.Grant {
	if err := c.validateGrantInvokeUrlParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		c,
		"grantInvokeUrl",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CheckStateMachineStatusFunction) GrantInvokeVersion(grantee awsiam.IGrantable, version awslambda.IVersion) awsiam.Grant {
	if err := c.validateGrantInvokeVersionParameters(grantee, version); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		c,
		"grantInvokeVersion",
		[]interface{}{grantee, version},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CheckStateMachineStatusFunction) InvalidateVersionBasedOn(x *string) {
	if err := c.validateInvalidateVersionBasedOnParameters(x); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"invalidateVersionBasedOn",
		[]interface{}{x},
	)
}

func (c *jsiiProxy_CheckStateMachineStatusFunction) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := c.validateMetricParameters(metricName, props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		c,
		"metric",
		[]interface{}{metricName, props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CheckStateMachineStatusFunction) MetricDuration(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := c.validateMetricDurationParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		c,
		"metricDuration",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CheckStateMachineStatusFunction) MetricErrors(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := c.validateMetricErrorsParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		c,
		"metricErrors",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CheckStateMachineStatusFunction) MetricInvocations(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := c.validateMetricInvocationsParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		c,
		"metricInvocations",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CheckStateMachineStatusFunction) MetricThrottles(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := c.validateMetricThrottlesParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		c,
		"metricThrottles",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CheckStateMachineStatusFunction) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CheckStateMachineStatusFunction) WarnInvokeFunctionPermissions(scope constructs.Construct) {
	if err := c.validateWarnInvokeFunctionPermissionsParameters(scope); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"warnInvokeFunctionPermissions",
		[]interface{}{scope},
	)
}

