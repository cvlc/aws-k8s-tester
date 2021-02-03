// Code generated by smithy-go-codegen DO NOT EDIT.

package ecr

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves the pre-signed Amazon S3 download URL corresponding to an image layer.
// You can only get URLs for image layers that are referenced in an image. When an
// image is pulled, the GetDownloadUrlForLayer API is called once per image layer
// that is not already cached. This operation is used by the Amazon ECR proxy and
// is not generally used by customers for pulling and pushing images. In most
// cases, you should use the docker CLI to pull, tag, and push images.
func (c *Client) GetDownloadUrlForLayer(ctx context.Context, params *GetDownloadUrlForLayerInput, optFns ...func(*Options)) (*GetDownloadUrlForLayerOutput, error) {
	if params == nil {
		params = &GetDownloadUrlForLayerInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetDownloadUrlForLayer", params, optFns, addOperationGetDownloadUrlForLayerMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetDownloadUrlForLayerOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetDownloadUrlForLayerInput struct {

	// The digest of the image layer to download.
	//
	// This member is required.
	LayerDigest *string

	// The name of the repository that is associated with the image layer to download.
	//
	// This member is required.
	RepositoryName *string

	// The AWS account ID associated with the registry that contains the image layer to
	// download. If you do not specify a registry, the default registry is assumed.
	RegistryId *string
}

type GetDownloadUrlForLayerOutput struct {

	// The pre-signed Amazon S3 download URL for the requested layer.
	DownloadUrl *string

	// The digest of the image layer to download.
	LayerDigest *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetDownloadUrlForLayerMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetDownloadUrlForLayer{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetDownloadUrlForLayer{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpGetDownloadUrlForLayerValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetDownloadUrlForLayer(options.Region), middleware.Before); err != nil {
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
	return nil
}

func newServiceMetadataMiddleware_opGetDownloadUrlForLayer(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ecr",
		OperationName: "GetDownloadUrlForLayer",
	}
}