// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package acm

import (
	"github.com/aws/aws-sdk-go/private/protocol"
)

const (

	// ErrCodeAccessDeniedException for service response error code
	// "AccessDeniedException".
	//
	// You do not have access required to perform this action.
	ErrCodeAccessDeniedException = "AccessDeniedException"

	// ErrCodeConflictException for service response error code
	// "ConflictException".
	//
	// You are trying to update a resource or configuration that is already being
	// created or updated. Wait for the previous operation to finish and try again.
	ErrCodeConflictException = "ConflictException"

	// ErrCodeInvalidArgsException for service response error code
	// "InvalidArgsException".
	//
	// One or more of of request parameters specified is not valid.
	ErrCodeInvalidArgsException = "InvalidArgsException"

	// ErrCodeInvalidArnException for service response error code
	// "InvalidArnException".
	//
	// The requested Amazon Resource Name (ARN) does not refer to an existing resource.
	ErrCodeInvalidArnException = "InvalidArnException"

	// ErrCodeInvalidDomainValidationOptionsException for service response error code
	// "InvalidDomainValidationOptionsException".
	//
	// One or more values in the DomainValidationOption structure is incorrect.
	ErrCodeInvalidDomainValidationOptionsException = "InvalidDomainValidationOptionsException"

	// ErrCodeInvalidParameterException for service response error code
	// "InvalidParameterException".
	//
	// An input parameter was invalid.
	ErrCodeInvalidParameterException = "InvalidParameterException"

	// ErrCodeInvalidStateException for service response error code
	// "InvalidStateException".
	//
	// Processing has reached an invalid state.
	ErrCodeInvalidStateException = "InvalidStateException"

	// ErrCodeInvalidTagException for service response error code
	// "InvalidTagException".
	//
	// One or both of the values that make up the key-value pair is not valid. For
	// example, you cannot specify a tag value that begins with aws:.
	ErrCodeInvalidTagException = "InvalidTagException"

	// ErrCodeLimitExceededException for service response error code
	// "LimitExceededException".
	//
	// An ACM quota has been exceeded.
	ErrCodeLimitExceededException = "LimitExceededException"

	// ErrCodeRequestInProgressException for service response error code
	// "RequestInProgressException".
	//
	// The certificate request is in process and the certificate in your account
	// has not yet been issued.
	ErrCodeRequestInProgressException = "RequestInProgressException"

	// ErrCodeResourceInUseException for service response error code
	// "ResourceInUseException".
	//
	// The certificate is in use by another Amazon Web Services service in the caller's
	// account. Remove the association and try again.
	ErrCodeResourceInUseException = "ResourceInUseException"

	// ErrCodeResourceNotFoundException for service response error code
	// "ResourceNotFoundException".
	//
	// The specified certificate cannot be found in the caller's account or the
	// caller's account cannot be found.
	ErrCodeResourceNotFoundException = "ResourceNotFoundException"

	// ErrCodeTagPolicyException for service response error code
	// "TagPolicyException".
	//
	// A specified tag did not comply with an existing tag policy and was rejected.
	ErrCodeTagPolicyException = "TagPolicyException"

	// ErrCodeThrottlingException for service response error code
	// "ThrottlingException".
	//
	// The request was denied because it exceeded a quota.
	ErrCodeThrottlingException = "ThrottlingException"

	// ErrCodeTooManyTagsException for service response error code
	// "TooManyTagsException".
	//
	// The request contains too many tags. Try the request again with fewer tags.
	ErrCodeTooManyTagsException = "TooManyTagsException"

	// ErrCodeValidationException for service response error code
	// "ValidationException".
	//
	// The supplied input failed to satisfy constraints of an Amazon Web Services
	// service.
	ErrCodeValidationException = "ValidationException"
)

var exceptionFromCode = map[string]func(protocol.ResponseMetadata) error{
	"AccessDeniedException":                   newErrorAccessDeniedException,
	"ConflictException":                       newErrorConflictException,
	"InvalidArgsException":                    newErrorInvalidArgsException,
	"InvalidArnException":                     newErrorInvalidArnException,
	"InvalidDomainValidationOptionsException": newErrorInvalidDomainValidationOptionsException,
	"InvalidParameterException":               newErrorInvalidParameterException,
	"InvalidStateException":                   newErrorInvalidStateException,
	"InvalidTagException":                     newErrorInvalidTagException,
	"LimitExceededException":                  newErrorLimitExceededException,
	"RequestInProgressException":              newErrorRequestInProgressException,
	"ResourceInUseException":                  newErrorResourceInUseException,
	"ResourceNotFoundException":               newErrorResourceNotFoundException,
	"TagPolicyException":                      newErrorTagPolicyException,
	"ThrottlingException":                     newErrorThrottlingException,
	"TooManyTagsException":                    newErrorTooManyTagsException,
	"ValidationException":                     newErrorValidationException,
}
