// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package databasemigrationservice

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type StartReplicationTaskAssessmentInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the replication task.
	//
	// ReplicationTaskArn is a required field
	ReplicationTaskArn *string `type:"string" required:"true"`
}

// String returns the string representation
func (s StartReplicationTaskAssessmentInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *StartReplicationTaskAssessmentInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "StartReplicationTaskAssessmentInput"}

	if s.ReplicationTaskArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("ReplicationTaskArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type StartReplicationTaskAssessmentOutput struct {
	_ struct{} `type:"structure"`

	// The assessed replication task.
	ReplicationTask *ReplicationTask `type:"structure"`
}

// String returns the string representation
func (s StartReplicationTaskAssessmentOutput) String() string {
	return awsutil.Prettify(s)
}

const opStartReplicationTaskAssessment = "StartReplicationTaskAssessment"

// StartReplicationTaskAssessmentRequest returns a request value for making API operation for
// AWS Database Migration Service.
//
// Starts the replication task assessment for unsupported data types in the
// source database.
//
//    // Example sending a request using StartReplicationTaskAssessmentRequest.
//    req := client.StartReplicationTaskAssessmentRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/dms-2016-01-01/StartReplicationTaskAssessment
func (c *Client) StartReplicationTaskAssessmentRequest(input *StartReplicationTaskAssessmentInput) StartReplicationTaskAssessmentRequest {
	op := &aws.Operation{
		Name:       opStartReplicationTaskAssessment,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &StartReplicationTaskAssessmentInput{}
	}

	req := c.newRequest(op, input, &StartReplicationTaskAssessmentOutput{})
	return StartReplicationTaskAssessmentRequest{Request: req, Input: input, Copy: c.StartReplicationTaskAssessmentRequest}
}

// StartReplicationTaskAssessmentRequest is the request type for the
// StartReplicationTaskAssessment API operation.
type StartReplicationTaskAssessmentRequest struct {
	*aws.Request
	Input *StartReplicationTaskAssessmentInput
	Copy  func(*StartReplicationTaskAssessmentInput) StartReplicationTaskAssessmentRequest
}

// Send marshals and sends the StartReplicationTaskAssessment API request.
func (r StartReplicationTaskAssessmentRequest) Send(ctx context.Context) (*StartReplicationTaskAssessmentResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &StartReplicationTaskAssessmentResponse{
		StartReplicationTaskAssessmentOutput: r.Request.Data.(*StartReplicationTaskAssessmentOutput),
		response:                             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// StartReplicationTaskAssessmentResponse is the response type for the
// StartReplicationTaskAssessment API operation.
type StartReplicationTaskAssessmentResponse struct {
	*StartReplicationTaskAssessmentOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// StartReplicationTaskAssessment request.
func (r *StartReplicationTaskAssessmentResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
