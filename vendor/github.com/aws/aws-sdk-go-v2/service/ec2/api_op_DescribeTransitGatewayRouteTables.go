// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Describes one or more transit gateway route tables. By default, all transit
// gateway route tables are described. Alternatively, you can filter the results.
func (c *Client) DescribeTransitGatewayRouteTables(ctx context.Context, params *DescribeTransitGatewayRouteTablesInput, optFns ...func(*Options)) (*DescribeTransitGatewayRouteTablesOutput, error) {
	if params == nil {
		params = &DescribeTransitGatewayRouteTablesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeTransitGatewayRouteTables", params, optFns, addOperationDescribeTransitGatewayRouteTablesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeTransitGatewayRouteTablesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeTransitGatewayRouteTablesInput struct {

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun bool

	// One or more filters. The possible values are:
	//
	// * default-association-route-table
	// - Indicates whether this is the default association route table for the transit
	// gateway (true | false).
	//
	// * default-propagation-route-table - Indicates whether
	// this is the default propagation route table for the transit gateway (true |
	// false).
	//
	// * state - The state of the route table (available | deleting | deleted
	// | pending).
	//
	// * transit-gateway-id - The ID of the transit gateway.
	//
	// *
	// transit-gateway-route-table-id - The ID of the transit gateway route table.
	Filters []types.Filter

	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	MaxResults int32

	// The token for the next page of results.
	NextToken *string

	// The IDs of the transit gateway route tables.
	TransitGatewayRouteTableIds []string
}

type DescribeTransitGatewayRouteTablesOutput struct {

	// The token to use to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string

	// Information about the transit gateway route tables.
	TransitGatewayRouteTables []types.TransitGatewayRouteTable

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeTransitGatewayRouteTablesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpDescribeTransitGatewayRouteTables{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpDescribeTransitGatewayRouteTables{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeTransitGatewayRouteTables(options.Region), middleware.Before); err != nil {
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

// DescribeTransitGatewayRouteTablesAPIClient is a client that implements the
// DescribeTransitGatewayRouteTables operation.
type DescribeTransitGatewayRouteTablesAPIClient interface {
	DescribeTransitGatewayRouteTables(context.Context, *DescribeTransitGatewayRouteTablesInput, ...func(*Options)) (*DescribeTransitGatewayRouteTablesOutput, error)
}

var _ DescribeTransitGatewayRouteTablesAPIClient = (*Client)(nil)

// DescribeTransitGatewayRouteTablesPaginatorOptions is the paginator options for
// DescribeTransitGatewayRouteTables
type DescribeTransitGatewayRouteTablesPaginatorOptions struct {
	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeTransitGatewayRouteTablesPaginator is a paginator for
// DescribeTransitGatewayRouteTables
type DescribeTransitGatewayRouteTablesPaginator struct {
	options   DescribeTransitGatewayRouteTablesPaginatorOptions
	client    DescribeTransitGatewayRouteTablesAPIClient
	params    *DescribeTransitGatewayRouteTablesInput
	nextToken *string
	firstPage bool
}

// NewDescribeTransitGatewayRouteTablesPaginator returns a new
// DescribeTransitGatewayRouteTablesPaginator
func NewDescribeTransitGatewayRouteTablesPaginator(client DescribeTransitGatewayRouteTablesAPIClient, params *DescribeTransitGatewayRouteTablesInput, optFns ...func(*DescribeTransitGatewayRouteTablesPaginatorOptions)) *DescribeTransitGatewayRouteTablesPaginator {
	options := DescribeTransitGatewayRouteTablesPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	if params == nil {
		params = &DescribeTransitGatewayRouteTablesInput{}
	}

	return &DescribeTransitGatewayRouteTablesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeTransitGatewayRouteTablesPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next DescribeTransitGatewayRouteTables page.
func (p *DescribeTransitGatewayRouteTablesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeTransitGatewayRouteTablesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.DescribeTransitGatewayRouteTables(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken && prevToken != nil && p.nextToken != nil && *prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opDescribeTransitGatewayRouteTables(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DescribeTransitGatewayRouteTables",
	}
}