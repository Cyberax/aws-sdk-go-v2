// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package rds

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ModifyEventSubscriptionInput struct {
	_ struct{} `type:"structure"`

	// A value that indicates whether to activate the subscription.
	Enabled *bool `type:"boolean"`

	// A list of event categories for a SourceType that you want to subscribe to.
	// You can see a list of the categories for a given SourceType in the Events
	// (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_Events.html)
	// topic in the Amazon RDS User Guide or by using the DescribeEventCategories
	// action.
	EventCategories []string `locationNameList:"EventCategory" type:"list"`

	// The Amazon Resource Name (ARN) of the SNS topic created for event notification.
	// The ARN is created by Amazon SNS when you create a topic and subscribe to
	// it.
	SnsTopicArn *string `type:"string"`

	// The type of source that is generating the events. For example, if you want
	// to be notified of events generated by a DB instance, you would set this parameter
	// to db-instance. If this value isn't specified, all events are returned.
	//
	// Valid values: db-instance | db-parameter-group | db-security-group | db-snapshot
	SourceType *string `type:"string"`

	// The name of the RDS event notification subscription.
	//
	// SubscriptionName is a required field
	SubscriptionName *string `type:"string" required:"true"`
}

// String returns the string representation
func (s ModifyEventSubscriptionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ModifyEventSubscriptionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ModifyEventSubscriptionInput"}

	if s.SubscriptionName == nil {
		invalidParams.Add(aws.NewErrParamRequired("SubscriptionName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ModifyEventSubscriptionOutput struct {
	_ struct{} `type:"structure"`

	// Contains the results of a successful invocation of the DescribeEventSubscriptions
	// action.
	EventSubscription *EventSubscription `type:"structure"`
}

// String returns the string representation
func (s ModifyEventSubscriptionOutput) String() string {
	return awsutil.Prettify(s)
}

const opModifyEventSubscription = "ModifyEventSubscription"

// ModifyEventSubscriptionRequest returns a request value for making API operation for
// Amazon Relational Database Service.
//
// Modifies an existing RDS event notification subscription. You can't modify
// the source identifiers using this call. To change source identifiers for
// a subscription, use the AddSourceIdentifierToSubscription and RemoveSourceIdentifierFromSubscription
// calls.
//
// You can see a list of the event categories for a given SourceType in the
// Events (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_Events.html)
// topic in the Amazon RDS User Guide or by using the DescribeEventCategories
// action.
//
//    // Example sending a request using ModifyEventSubscriptionRequest.
//    req := client.ModifyEventSubscriptionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/rds-2014-10-31/ModifyEventSubscription
func (c *Client) ModifyEventSubscriptionRequest(input *ModifyEventSubscriptionInput) ModifyEventSubscriptionRequest {
	op := &aws.Operation{
		Name:       opModifyEventSubscription,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ModifyEventSubscriptionInput{}
	}

	req := c.newRequest(op, input, &ModifyEventSubscriptionOutput{})
	return ModifyEventSubscriptionRequest{Request: req, Input: input, Copy: c.ModifyEventSubscriptionRequest}
}

// ModifyEventSubscriptionRequest is the request type for the
// ModifyEventSubscription API operation.
type ModifyEventSubscriptionRequest struct {
	*aws.Request
	Input *ModifyEventSubscriptionInput
	Copy  func(*ModifyEventSubscriptionInput) ModifyEventSubscriptionRequest
}

// Send marshals and sends the ModifyEventSubscription API request.
func (r ModifyEventSubscriptionRequest) Send(ctx context.Context) (*ModifyEventSubscriptionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ModifyEventSubscriptionResponse{
		ModifyEventSubscriptionOutput: r.Request.Data.(*ModifyEventSubscriptionOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ModifyEventSubscriptionResponse is the response type for the
// ModifyEventSubscription API operation.
type ModifyEventSubscriptionResponse struct {
	*ModifyEventSubscriptionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ModifyEventSubscription request.
func (r *ModifyEventSubscriptionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
