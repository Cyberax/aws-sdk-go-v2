// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package imagebuilder

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type CancelImageCreationInput struct {
	_ struct{} `type:"structure"`

	// The idempotency token used to make this request idempotent.
	//
	// ClientToken is a required field
	ClientToken *string `locationName:"clientToken" min:"1" type:"string" required:"true"`

	// The Amazon Resource Name (ARN) of the image whose creation you wish to cancel.
	//
	// ImageBuildVersionArn is a required field
	ImageBuildVersionArn *string `locationName:"imageBuildVersionArn" type:"string" required:"true"`
}

// String returns the string representation
func (s CancelImageCreationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CancelImageCreationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CancelImageCreationInput"}

	if s.ClientToken == nil {
		invalidParams.Add(aws.NewErrParamRequired("ClientToken"))
	}
	if s.ClientToken != nil && len(*s.ClientToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ClientToken", 1))
	}

	if s.ImageBuildVersionArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("ImageBuildVersionArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CancelImageCreationInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.ClientToken != nil {
		v := *s.ClientToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "clientToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ImageBuildVersionArn != nil {
		v := *s.ImageBuildVersionArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "imageBuildVersionArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type CancelImageCreationOutput struct {
	_ struct{} `type:"structure"`

	// The idempotency token used to make this request idempotent.
	ClientToken *string `locationName:"clientToken" min:"1" type:"string"`

	// The Amazon Resource Name (ARN) of the image whose creation has been cancelled.
	ImageBuildVersionArn *string `locationName:"imageBuildVersionArn" type:"string"`

	// The request ID that uniquely identifies this request.
	RequestId *string `locationName:"requestId" min:"1" type:"string"`
}

// String returns the string representation
func (s CancelImageCreationOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CancelImageCreationOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.ClientToken != nil {
		v := *s.ClientToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "clientToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ImageBuildVersionArn != nil {
		v := *s.ImageBuildVersionArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "imageBuildVersionArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.RequestId != nil {
		v := *s.RequestId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "requestId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opCancelImageCreation = "CancelImageCreation"

// CancelImageCreationRequest returns a request value for making API operation for
// EC2 Image Builder.
//
// CancelImageCreation cancels the creation of Image. This operation may only
// be used on images in a non-terminal state.
//
//    // Example sending a request using CancelImageCreationRequest.
//    req := client.CancelImageCreationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/imagebuilder-2019-12-02/CancelImageCreation
func (c *Client) CancelImageCreationRequest(input *CancelImageCreationInput) CancelImageCreationRequest {
	op := &aws.Operation{
		Name:       opCancelImageCreation,
		HTTPMethod: "PUT",
		HTTPPath:   "/CancelImageCreation",
	}

	if input == nil {
		input = &CancelImageCreationInput{}
	}

	req := c.newRequest(op, input, &CancelImageCreationOutput{})
	return CancelImageCreationRequest{Request: req, Input: input, Copy: c.CancelImageCreationRequest}
}

// CancelImageCreationRequest is the request type for the
// CancelImageCreation API operation.
type CancelImageCreationRequest struct {
	*aws.Request
	Input *CancelImageCreationInput
	Copy  func(*CancelImageCreationInput) CancelImageCreationRequest
}

// Send marshals and sends the CancelImageCreation API request.
func (r CancelImageCreationRequest) Send(ctx context.Context) (*CancelImageCreationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CancelImageCreationResponse{
		CancelImageCreationOutput: r.Request.Data.(*CancelImageCreationOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CancelImageCreationResponse is the response type for the
// CancelImageCreation API operation.
type CancelImageCreationResponse struct {
	*CancelImageCreationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CancelImageCreation request.
func (r *CancelImageCreationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
