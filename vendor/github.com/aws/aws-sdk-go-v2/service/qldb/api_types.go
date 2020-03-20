// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package qldb

import (
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

var _ aws.Config
var _ = awsutil.Prettify

// The information about a journal export job, including the ledger name, export
// ID, when it was created, current status, and its start and end time export
// parameters.
type JournalS3ExportDescription struct {
	_ struct{} `type:"structure"`

	// The exclusive end date and time for the range of journal contents that are
	// specified in the original export request.
	//
	// ExclusiveEndTime is a required field
	ExclusiveEndTime *time.Time `type:"timestamp" required:"true"`

	// The date and time, in epoch time format, when the export job was created.
	// (Epoch time format is the number of seconds elapsed since 12:00:00 AM January
	// 1, 1970 UTC.)
	//
	// ExportCreationTime is a required field
	ExportCreationTime *time.Time `type:"timestamp" required:"true"`

	// The unique ID of the journal export job.
	//
	// ExportId is a required field
	ExportId *string `min:"22" type:"string" required:"true"`

	// The inclusive start date and time for the range of journal contents that
	// are specified in the original export request.
	//
	// InclusiveStartTime is a required field
	InclusiveStartTime *time.Time `type:"timestamp" required:"true"`

	// The name of the ledger.
	//
	// LedgerName is a required field
	LedgerName *string `min:"1" type:"string" required:"true"`

	// The Amazon Resource Name (ARN) of the IAM role that grants QLDB permissions
	// for a journal export job to do the following:
	//
	//    * Write objects into your Amazon Simple Storage Service (Amazon S3) bucket.
	//
	//    * (Optional) Use your customer master key (CMK) in AWS Key Management
	//    Service (AWS KMS) for server-side encryption of your exported data.
	//
	// RoleArn is a required field
	RoleArn *string `min:"20" type:"string" required:"true"`

	// The Amazon Simple Storage Service (Amazon S3) bucket location in which a
	// journal export job writes the journal contents.
	//
	// S3ExportConfiguration is a required field
	S3ExportConfiguration *S3ExportConfiguration `type:"structure" required:"true"`

	// The current state of the journal export job.
	//
	// Status is a required field
	Status ExportStatus `type:"string" required:"true" enum:"true"`
}

// String returns the string representation
func (s JournalS3ExportDescription) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s JournalS3ExportDescription) MarshalFields(e protocol.FieldEncoder) error {
	if s.ExclusiveEndTime != nil {
		v := *s.ExclusiveEndTime

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "ExclusiveEndTime",
			protocol.TimeValue{V: v, Format: protocol.UnixTimeFormatName, QuotedFormatTime: true}, metadata)
	}
	if s.ExportCreationTime != nil {
		v := *s.ExportCreationTime

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "ExportCreationTime",
			protocol.TimeValue{V: v, Format: protocol.UnixTimeFormatName, QuotedFormatTime: true}, metadata)
	}
	if s.ExportId != nil {
		v := *s.ExportId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "ExportId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.InclusiveStartTime != nil {
		v := *s.InclusiveStartTime

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "InclusiveStartTime",
			protocol.TimeValue{V: v, Format: protocol.UnixTimeFormatName, QuotedFormatTime: true}, metadata)
	}
	if s.LedgerName != nil {
		v := *s.LedgerName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "LedgerName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.RoleArn != nil {
		v := *s.RoleArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "RoleArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.S3ExportConfiguration != nil {
		v := s.S3ExportConfiguration

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "S3ExportConfiguration", v, metadata)
	}
	if len(s.Status) > 0 {
		v := s.Status

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Status", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	return nil
}

// Information about a ledger, including its name, state, and when it was created.
type LedgerSummary struct {
	_ struct{} `type:"structure"`

	// The date and time, in epoch time format, when the ledger was created. (Epoch
	// time format is the number of seconds elapsed since 12:00:00 AM January 1,
	// 1970 UTC.)
	CreationDateTime *time.Time `type:"timestamp"`

	// The name of the ledger.
	Name *string `min:"1" type:"string"`

	// The current status of the ledger.
	State LedgerState `type:"string" enum:"true"`
}

// String returns the string representation
func (s LedgerSummary) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s LedgerSummary) MarshalFields(e protocol.FieldEncoder) error {
	if s.CreationDateTime != nil {
		v := *s.CreationDateTime

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "CreationDateTime",
			protocol.TimeValue{V: v, Format: protocol.UnixTimeFormatName, QuotedFormatTime: true}, metadata)
	}
	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.State) > 0 {
		v := s.State

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "State", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	return nil
}

// The encryption settings that are used by a journal export job to write data
// in an Amazon Simple Storage Service (Amazon S3) bucket.
type S3EncryptionConfiguration struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) for a customer master key (CMK) in AWS Key
	// Management Service (AWS KMS).
	//
	// You must provide a KmsKeyArn if you specify SSE_KMS as the ObjectEncryptionType.
	//
	// KmsKeyArn is not required if you specify SSE_S3 as the ObjectEncryptionType.
	KmsKeyArn *string `min:"20" type:"string"`

	// The Amazon S3 object encryption type.
	//
	// To learn more about server-side encryption options in Amazon S3, see Protecting
	// Data Using Server-Side Encryption (https://docs.aws.amazon.com/AmazonS3/latest/dev/serv-side-encryption.html)
	// in the Amazon S3 Developer Guide.
	//
	// ObjectEncryptionType is a required field
	ObjectEncryptionType S3ObjectEncryptionType `type:"string" required:"true" enum:"true"`
}

// String returns the string representation
func (s S3EncryptionConfiguration) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *S3EncryptionConfiguration) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "S3EncryptionConfiguration"}
	if s.KmsKeyArn != nil && len(*s.KmsKeyArn) < 20 {
		invalidParams.Add(aws.NewErrParamMinLen("KmsKeyArn", 20))
	}
	if len(s.ObjectEncryptionType) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("ObjectEncryptionType"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s S3EncryptionConfiguration) MarshalFields(e protocol.FieldEncoder) error {
	if s.KmsKeyArn != nil {
		v := *s.KmsKeyArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "KmsKeyArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.ObjectEncryptionType) > 0 {
		v := s.ObjectEncryptionType

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "ObjectEncryptionType", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	return nil
}

// The Amazon Simple Storage Service (Amazon S3) bucket location in which a
// journal export job writes the journal contents.
type S3ExportConfiguration struct {
	_ struct{} `type:"structure"`

	// The Amazon S3 bucket name in which a journal export job writes the journal
	// contents.
	//
	// The bucket name must comply with the Amazon S3 bucket naming conventions.
	// For more information, see Bucket Restrictions and Limitations (https://docs.aws.amazon.com/AmazonS3/latest/dev/BucketRestrictions.html)
	// in the Amazon S3 Developer Guide.
	//
	// Bucket is a required field
	Bucket *string `min:"3" type:"string" required:"true"`

	// The encryption settings that are used by a journal export job to write data
	// in an Amazon S3 bucket.
	//
	// EncryptionConfiguration is a required field
	EncryptionConfiguration *S3EncryptionConfiguration `type:"structure" required:"true"`

	// The prefix for the Amazon S3 bucket in which a journal export job writes
	// the journal contents.
	//
	// The prefix must comply with Amazon S3 key naming rules and restrictions.
	// For more information, see Object Key and Metadata (https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingMetadata.html)
	// in the Amazon S3 Developer Guide.
	//
	// The following are examples of valid Prefix values:
	//
	//    * JournalExports-ForMyLedger/Testing/
	//
	//    * JournalExports
	//
	//    * My:Tests/
	//
	// Prefix is a required field
	Prefix *string `type:"string" required:"true"`
}

// String returns the string representation
func (s S3ExportConfiguration) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *S3ExportConfiguration) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "S3ExportConfiguration"}

	if s.Bucket == nil {
		invalidParams.Add(aws.NewErrParamRequired("Bucket"))
	}
	if s.Bucket != nil && len(*s.Bucket) < 3 {
		invalidParams.Add(aws.NewErrParamMinLen("Bucket", 3))
	}

	if s.EncryptionConfiguration == nil {
		invalidParams.Add(aws.NewErrParamRequired("EncryptionConfiguration"))
	}

	if s.Prefix == nil {
		invalidParams.Add(aws.NewErrParamRequired("Prefix"))
	}
	if s.EncryptionConfiguration != nil {
		if err := s.EncryptionConfiguration.Validate(); err != nil {
			invalidParams.AddNested("EncryptionConfiguration", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s S3ExportConfiguration) MarshalFields(e protocol.FieldEncoder) error {
	if s.Bucket != nil {
		v := *s.Bucket

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Bucket", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.EncryptionConfiguration != nil {
		v := s.EncryptionConfiguration

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "EncryptionConfiguration", v, metadata)
	}
	if s.Prefix != nil {
		v := *s.Prefix

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Prefix", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// A structure that can contain an Amazon Ion value in multiple encoding formats.
type ValueHolder struct {
	_ struct{} `type:"structure" sensitive:"true"`

	// An Amazon Ion plaintext value contained in a ValueHolder structure.
	IonText *string `min:"1" type:"string" sensitive:"true"`
}

// String returns the string representation
func (s ValueHolder) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ValueHolder) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ValueHolder"}
	if s.IonText != nil && len(*s.IonText) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("IonText", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ValueHolder) MarshalFields(e protocol.FieldEncoder) error {
	if s.IonText != nil {
		v := *s.IonText

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "IonText", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}
