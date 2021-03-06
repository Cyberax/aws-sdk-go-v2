// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iot

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
)

// The input for the DeleteTopicRule operation.
type DeleteTopicRuleInput struct {
	_ struct{} `type:"structure"`

	// The name of the rule.
	//
	// RuleName is a required field
	RuleName *string `location:"uri" locationName:"ruleName" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteTopicRuleInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteTopicRuleInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteTopicRuleInput"}

	if s.RuleName == nil {
		invalidParams.Add(aws.NewErrParamRequired("RuleName"))
	}
	if s.RuleName != nil && len(*s.RuleName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("RuleName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteTopicRuleInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.RuleName != nil {
		v := *s.RuleName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "ruleName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type DeleteTopicRuleOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteTopicRuleOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteTopicRuleOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opDeleteTopicRule = "DeleteTopicRule"

// DeleteTopicRuleRequest returns a request value for making API operation for
// AWS IoT.
//
// Deletes the rule.
//
//    // Example sending a request using DeleteTopicRuleRequest.
//    req := client.DeleteTopicRuleRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) DeleteTopicRuleRequest(input *DeleteTopicRuleInput) DeleteTopicRuleRequest {
	op := &aws.Operation{
		Name:       opDeleteTopicRule,
		HTTPMethod: "DELETE",
		HTTPPath:   "/rules/{ruleName}",
	}

	if input == nil {
		input = &DeleteTopicRuleInput{}
	}

	req := c.newRequest(op, input, &DeleteTopicRuleOutput{})
	req.Handlers.Unmarshal.Remove(restjson.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return DeleteTopicRuleRequest{Request: req, Input: input, Copy: c.DeleteTopicRuleRequest}
}

// DeleteTopicRuleRequest is the request type for the
// DeleteTopicRule API operation.
type DeleteTopicRuleRequest struct {
	*aws.Request
	Input *DeleteTopicRuleInput
	Copy  func(*DeleteTopicRuleInput) DeleteTopicRuleRequest
}

// Send marshals and sends the DeleteTopicRule API request.
func (r DeleteTopicRuleRequest) Send(ctx context.Context) (*DeleteTopicRuleResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteTopicRuleResponse{
		DeleteTopicRuleOutput: r.Request.Data.(*DeleteTopicRuleOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteTopicRuleResponse is the response type for the
// DeleteTopicRule API operation.
type DeleteTopicRuleResponse struct {
	*DeleteTopicRuleOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteTopicRule request.
func (r *DeleteTopicRuleResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
