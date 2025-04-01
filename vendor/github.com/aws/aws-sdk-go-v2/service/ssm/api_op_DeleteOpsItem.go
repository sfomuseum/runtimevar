// Code generated by smithy-go-codegen DO NOT EDIT.

package ssm

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Delete an OpsItem. You must have permission in Identity and Access Management
// (IAM) to delete an OpsItem.
//
// Note the following important information about this operation.
//
//   - Deleting an OpsItem is irreversible. You can't restore a deleted OpsItem.
//
//   - This operation uses an eventual consistency model, which means the system
//     can take a few minutes to complete this operation. If you delete an OpsItem and
//     immediately call, for example, GetOpsItem, the deleted OpsItem might still appear in
//     the response.
//
//   - This operation is idempotent. The system doesn't throw an exception if you
//     repeatedly call this operation for the same OpsItem. If the first call is
//     successful, all additional calls return the same successful response as the
//     first call.
//
//   - This operation doesn't support cross-account calls. A delegated
//     administrator or management account can't delete OpsItems in other accounts,
//     even if OpsCenter has been set up for cross-account administration. For more
//     information about cross-account administration, see [Setting up OpsCenter to centrally manage OpsItems across accounts]in the Systems Manager
//     User Guide.
//
// [Setting up OpsCenter to centrally manage OpsItems across accounts]: https://docs.aws.amazon.com/systems-manager/latest/userguide/OpsCenter-setting-up-cross-account.html
func (c *Client) DeleteOpsItem(ctx context.Context, params *DeleteOpsItemInput, optFns ...func(*Options)) (*DeleteOpsItemOutput, error) {
	if params == nil {
		params = &DeleteOpsItemInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteOpsItem", params, optFns, c.addOperationDeleteOpsItemMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteOpsItemOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteOpsItemInput struct {

	// The ID of the OpsItem that you want to delete.
	//
	// This member is required.
	OpsItemId *string

	noSmithyDocumentSerde
}

type DeleteOpsItemOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteOpsItemMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteOpsItem{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteOpsItem{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DeleteOpsItem"); err != nil {
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
	if err = addSpanRetryLoop(stack, options); err != nil {
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
	if err = addCredentialSource(stack, options); err != nil {
		return err
	}
	if err = addOpDeleteOpsItemValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteOpsItem(options.Region), middleware.Before); err != nil {
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
	if err = addSpanInitializeStart(stack); err != nil {
		return err
	}
	if err = addSpanInitializeEnd(stack); err != nil {
		return err
	}
	if err = addSpanBuildRequestStart(stack); err != nil {
		return err
	}
	if err = addSpanBuildRequestEnd(stack); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opDeleteOpsItem(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DeleteOpsItem",
	}
}
