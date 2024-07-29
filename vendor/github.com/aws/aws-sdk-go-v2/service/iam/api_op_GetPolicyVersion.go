// Code generated by smithy-go-codegen DO NOT EDIT.

package iam

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves information about the specified version of the specified managed
// policy, including the policy document.
//
// Policies returned by this operation are URL-encoded compliant with [RFC 3986]. You can
// use a URL decoding method to convert the policy back to plain JSON text. For
// example, if you use Java, you can use the decode method of the
// java.net.URLDecoder utility class in the Java SDK. Other languages and SDKs
// provide similar functionality.
//
// To list the available versions for a policy, use ListPolicyVersions.
//
// This operation retrieves information about managed policies. To retrieve
// information about an inline policy that is embedded in a user, group, or role,
// use GetUserPolicy, GetGroupPolicy, or GetRolePolicy.
//
// For more information about the types of policies, see [Managed policies and inline policies] in the IAM User Guide.
//
// For more information about managed policy versions, see [Versioning for managed policies] in the IAM User Guide.
//
// [RFC 3986]: https://tools.ietf.org/html/rfc3986
// [Versioning for managed policies]: https://docs.aws.amazon.com/IAM/latest/UserGuide/policies-managed-versions.html
// [Managed policies and inline policies]: https://docs.aws.amazon.com/IAM/latest/UserGuide/policies-managed-vs-inline.html
func (c *Client) GetPolicyVersion(ctx context.Context, params *GetPolicyVersionInput, optFns ...func(*Options)) (*GetPolicyVersionOutput, error) {
	if params == nil {
		params = &GetPolicyVersionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetPolicyVersion", params, optFns, c.addOperationGetPolicyVersionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetPolicyVersionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetPolicyVersionInput struct {

	// The Amazon Resource Name (ARN) of the managed policy that you want information
	// about.
	//
	// For more information about ARNs, see [Amazon Resource Names (ARNs)] in the Amazon Web Services General
	// Reference.
	//
	// [Amazon Resource Names (ARNs)]: https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html
	//
	// This member is required.
	PolicyArn *string

	// Identifies the policy version to retrieve.
	//
	// This parameter allows (through its [regex pattern]) a string of characters that consists of
	// the lowercase letter 'v' followed by one or two digits, and optionally followed
	// by a period '.' and a string of letters and digits.
	//
	// [regex pattern]: http://wikipedia.org/wiki/regex
	//
	// This member is required.
	VersionId *string

	noSmithyDocumentSerde
}

// Contains the response to a successful GetPolicyVersion request.
type GetPolicyVersionOutput struct {

	// A structure containing details about the policy version.
	PolicyVersion *types.PolicyVersion

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetPolicyVersionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpGetPolicyVersion{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpGetPolicyVersion{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetPolicyVersion"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = addClientRequestID(stack); err != nil {
		return err
	}
	if err = addComputeContentLength(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addComputePayloadSHA256(stack); err != nil {
		return err
	}
	if err = addRetry(stack, options); err != nil {
		return err
	}
	if err = addRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = addRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addTimeOffsetBuild(stack, c); err != nil {
		return err
	}
	if err = addUserAgentRetryMode(stack, options); err != nil {
		return err
	}
	if err = addOpGetPolicyVersionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetPolicyVersion(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRecursionDetection(stack); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opGetPolicyVersion(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetPolicyVersion",
	}
}
