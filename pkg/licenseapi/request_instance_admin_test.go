package licenseapi

import (
	"errors"
	"testing"
)

func TestInstancePatchInput_Validate_Valid(t *testing.T) {
	input := InstancePatchInput{
		StripeSubscriptionID: "sub_123",
	}
	if err := input.Validate(); err != nil {
		t.Fatalf("expected no error for valid input, got: %v", err)
	}
}

func TestInstancePatchInput_Validate_MissingStripeSubscriptionID(t *testing.T) {
	input := InstancePatchInput{}
	err := input.Validate()
	if err == nil {
		t.Fatal("expected error for missing StripeSubscriptionID, got nil")
	}
	if !errors.Is(err, ErrMissingStripeSubscriptionID) {
		t.Fatalf("expected ErrMissingStripeSubscriptionID, got: %v", err)
	}
}
