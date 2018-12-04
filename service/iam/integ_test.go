// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// +build go1.10,integration

package iam_test

import (
	"context"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/awstesting/integration"
	"github.com/aws/aws-sdk-go/service/iam"
)

var _ aws.Config
var _ awserr.Error
var _ request.Request

func TestInteg_00_ListUsers(t *testing.T) {
	ctx, cancelFn := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFn()

	sess := integration.SessionWithDefaultRegion("us-east-1")
	svc := iam.New(sess)
	params := &iam.ListUsersInput{}
	_, err := svc.ListUsersWithContext(ctx, params)
	if err != nil {
		t.Errorf("expect no error, got %v", err)
	}
}
func TestInteg_01_GetUser(t *testing.T) {
	ctx, cancelFn := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFn()

	sess := integration.SessionWithDefaultRegion("us-east-1")
	svc := iam.New(sess)
	params := &iam.GetUserInput{
		UserName: aws.String("fake_user"),
	}
	_, err := svc.GetUserWithContext(ctx, params)
	if err == nil {
		t.Fatalf("expect request to fail")
	}
	aerr, ok := err.(awserr.RequestFailure)
	if !ok {
		t.Fatalf("expect awserr, was %T", err)
	}
	if v := aerr.Code(); v == request.ErrCodeSerialization {
		t.Errorf("expect API error code got serialization failure")
	}
}
