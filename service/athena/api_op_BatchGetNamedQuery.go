// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package athena

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type BatchGetNamedQueryInput struct {
	_ struct{} `type:"structure"`

	// An array of query IDs.
	//
	// NamedQueryIds is a required field
	NamedQueryIds []string `min:"1" type:"list" required:"true"`
}

// String returns the string representation
func (s BatchGetNamedQueryInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *BatchGetNamedQueryInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "BatchGetNamedQueryInput"}

	if s.NamedQueryIds == nil {
		invalidParams.Add(aws.NewErrParamRequired("NamedQueryIds"))
	}
	if s.NamedQueryIds != nil && len(s.NamedQueryIds) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NamedQueryIds", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type BatchGetNamedQueryOutput struct {
	_ struct{} `type:"structure"`

	// Information about the named query IDs submitted.
	NamedQueries []NamedQuery `type:"list"`

	// Information about provided query IDs.
	UnprocessedNamedQueryIds []UnprocessedNamedQueryId `type:"list"`
}

// String returns the string representation
func (s BatchGetNamedQueryOutput) String() string {
	return awsutil.Prettify(s)
}

const opBatchGetNamedQuery = "BatchGetNamedQuery"

// BatchGetNamedQueryRequest returns a request value for making API operation for
// Amazon Athena.
//
// Returns the details of a single named query or a list of up to 50 queries,
// which you provide as an array of query ID strings. Requires you to have access
// to the workgroup in which the queries were saved. Use ListNamedQueriesInput
// to get the list of named query IDs in the specified workgroup. If information
// could not be retrieved for a submitted query ID, information about the query
// ID submitted is listed under UnprocessedNamedQueryId. Named queries differ
// from executed queries. Use BatchGetQueryExecutionInput to get details about
// each unique query execution, and ListQueryExecutionsInput to get a list of
// query execution IDs.
//
//    // Example sending a request using BatchGetNamedQueryRequest.
//    req := client.BatchGetNamedQueryRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/athena-2017-05-18/BatchGetNamedQuery
func (c *Client) BatchGetNamedQueryRequest(input *BatchGetNamedQueryInput) BatchGetNamedQueryRequest {
	op := &aws.Operation{
		Name:       opBatchGetNamedQuery,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &BatchGetNamedQueryInput{}
	}

	req := c.newRequest(op, input, &BatchGetNamedQueryOutput{})
	return BatchGetNamedQueryRequest{Request: req, Input: input, Copy: c.BatchGetNamedQueryRequest}
}

// BatchGetNamedQueryRequest is the request type for the
// BatchGetNamedQuery API operation.
type BatchGetNamedQueryRequest struct {
	*aws.Request
	Input *BatchGetNamedQueryInput
	Copy  func(*BatchGetNamedQueryInput) BatchGetNamedQueryRequest
}

// Send marshals and sends the BatchGetNamedQuery API request.
func (r BatchGetNamedQueryRequest) Send(ctx context.Context) (*BatchGetNamedQueryResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &BatchGetNamedQueryResponse{
		BatchGetNamedQueryOutput: r.Request.Data.(*BatchGetNamedQueryOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// BatchGetNamedQueryResponse is the response type for the
// BatchGetNamedQuery API operation.
type BatchGetNamedQueryResponse struct {
	*BatchGetNamedQueryOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// BatchGetNamedQuery request.
func (r *BatchGetNamedQueryResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
