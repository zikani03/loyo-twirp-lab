// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.6.1
// source: loyo.proto

package loyo

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type EventType int32

const (
	EventType_UNSPECIFIED         EventType = 0
	EventType_PHYSICAL_PURCHASE   EventType = 1
	EventType_DIGITAL_PURCHASE    EventType = 2
	EventType_LOAN_DISBURSEMENT   EventType = 3
	EventType_OVERDUE_LOAN_CHARGE EventType = 4
	EventType_WITHDRAWAL          EventType = 5
	EventType_SERVICE_FEE         EventType = 6
	EventType_TRANSFER            EventType = 7
	EventType_REFUND              EventType = 8
	EventType_DEPOSIT             EventType = 9
	EventType_RECHARGE            EventType = 10
	EventType_COMPETITION_REWARD  EventType = 11
	EventType_UNSPECIFIED_DEBIT   EventType = 12
	EventType_UNSPECIFIED_CREDIT  EventType = 13
)

// Enum value maps for EventType.
var (
	EventType_name = map[int32]string{
		0:  "UNSPECIFIED",
		1:  "PHYSICAL_PURCHASE",
		2:  "DIGITAL_PURCHASE",
		3:  "LOAN_DISBURSEMENT",
		4:  "OVERDUE_LOAN_CHARGE",
		5:  "WITHDRAWAL",
		6:  "SERVICE_FEE",
		7:  "TRANSFER",
		8:  "REFUND",
		9:  "DEPOSIT",
		10: "RECHARGE",
		11: "COMPETITION_REWARD",
		12: "UNSPECIFIED_DEBIT",
		13: "UNSPECIFIED_CREDIT",
	}
	EventType_value = map[string]int32{
		"UNSPECIFIED":         0,
		"PHYSICAL_PURCHASE":   1,
		"DIGITAL_PURCHASE":    2,
		"LOAN_DISBURSEMENT":   3,
		"OVERDUE_LOAN_CHARGE": 4,
		"WITHDRAWAL":          5,
		"SERVICE_FEE":         6,
		"TRANSFER":            7,
		"REFUND":              8,
		"DEPOSIT":             9,
		"RECHARGE":            10,
		"COMPETITION_REWARD":  11,
		"UNSPECIFIED_DEBIT":   12,
		"UNSPECIFIED_CREDIT":  13,
	}
)

func (x EventType) Enum() *EventType {
	p := new(EventType)
	*p = x
	return p
}

func (x EventType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EventType) Descriptor() protoreflect.EnumDescriptor {
	return file_loyo_proto_enumTypes[0].Descriptor()
}

func (EventType) Type() protoreflect.EnumType {
	return &file_loyo_proto_enumTypes[0]
}

func (x EventType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EventType.Descriptor instead.
func (EventType) EnumDescriptor() ([]byte, []int) {
	return file_loyo_proto_rawDescGZIP(), []int{0}
}

type LoyoEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EventType     EventType `protobuf:"varint,1,opt,name=event_type,json=eventType,proto3,enum=EventType" json:"event_type,omitempty"`
	TransactionId string    `protobuf:"bytes,2,opt,name=transaction_id,json=transactionId,proto3" json:"transaction_id,omitempty"`
	ProductSku    string    `protobuf:"bytes,3,opt,name=product_sku,json=productSku,proto3" json:"product_sku,omitempty"`
	ServiceId     string    `protobuf:"bytes,4,opt,name=service_id,json=serviceId,proto3" json:"service_id,omitempty"`
	AccountId     string    `protobuf:"bytes,5,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	Currency      string    `protobuf:"bytes,6,opt,name=currency,proto3" json:"currency,omitempty"`
	Amount        float64   `protobuf:"fixed64,7,opt,name=amount,proto3" json:"amount,omitempty"`
	TaxApplicable float64   `protobuf:"fixed64,8,opt,name=tax_applicable,json=taxApplicable,proto3" json:"tax_applicable,omitempty"`
	TaxCharged    float64   `protobuf:"fixed64,9,opt,name=tax_charged,json=taxCharged,proto3" json:"tax_charged,omitempty"`
	// For Money transfer events
	RecipientAccountId string `protobuf:"bytes,10,opt,name=recipient_account_id,json=recipientAccountId,proto3" json:"recipient_account_id,omitempty"`
	Flagged            bool   `protobuf:"varint,11,opt,name=flagged,proto3" json:"flagged,omitempty"`
	Timestamp          int64  `protobuf:"varint,12,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	ServiceProviderId  string `protobuf:"bytes,13,opt,name=service_provider_id,json=serviceProviderId,proto3" json:"service_provider_id,omitempty"`
}

func (x *LoyoEvent) Reset() {
	*x = LoyoEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_loyo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoyoEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoyoEvent) ProtoMessage() {}

func (x *LoyoEvent) ProtoReflect() protoreflect.Message {
	mi := &file_loyo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoyoEvent.ProtoReflect.Descriptor instead.
func (*LoyoEvent) Descriptor() ([]byte, []int) {
	return file_loyo_proto_rawDescGZIP(), []int{0}
}

func (x *LoyoEvent) GetEventType() EventType {
	if x != nil {
		return x.EventType
	}
	return EventType_UNSPECIFIED
}

func (x *LoyoEvent) GetTransactionId() string {
	if x != nil {
		return x.TransactionId
	}
	return ""
}

func (x *LoyoEvent) GetProductSku() string {
	if x != nil {
		return x.ProductSku
	}
	return ""
}

func (x *LoyoEvent) GetServiceId() string {
	if x != nil {
		return x.ServiceId
	}
	return ""
}

func (x *LoyoEvent) GetAccountId() string {
	if x != nil {
		return x.AccountId
	}
	return ""
}

func (x *LoyoEvent) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

func (x *LoyoEvent) GetAmount() float64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *LoyoEvent) GetTaxApplicable() float64 {
	if x != nil {
		return x.TaxApplicable
	}
	return 0
}

func (x *LoyoEvent) GetTaxCharged() float64 {
	if x != nil {
		return x.TaxCharged
	}
	return 0
}

func (x *LoyoEvent) GetRecipientAccountId() string {
	if x != nil {
		return x.RecipientAccountId
	}
	return ""
}

func (x *LoyoEvent) GetFlagged() bool {
	if x != nil {
		return x.Flagged
	}
	return false
}

func (x *LoyoEvent) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *LoyoEvent) GetServiceProviderId() string {
	if x != nil {
		return x.ServiceProviderId
	}
	return ""
}

type LoyoIngestRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ApplicationId string       `protobuf:"bytes,1,opt,name=application_id,json=applicationId,proto3" json:"application_id,omitempty"`
	Events        []*LoyoEvent `protobuf:"bytes,2,rep,name=events,proto3" json:"events,omitempty"`
}

func (x *LoyoIngestRequest) Reset() {
	*x = LoyoIngestRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_loyo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoyoIngestRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoyoIngestRequest) ProtoMessage() {}

func (x *LoyoIngestRequest) ProtoReflect() protoreflect.Message {
	mi := &file_loyo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoyoIngestRequest.ProtoReflect.Descriptor instead.
func (*LoyoIngestRequest) Descriptor() ([]byte, []int) {
	return file_loyo_proto_rawDescGZIP(), []int{1}
}

func (x *LoyoIngestRequest) GetApplicationId() string {
	if x != nil {
		return x.ApplicationId
	}
	return ""
}

func (x *LoyoIngestRequest) GetEvents() []*LoyoEvent {
	if x != nil {
		return x.Events
	}
	return nil
}

type LoyoIngestResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LoyoTxnId     string `protobuf:"bytes,1,opt,name=loyo_txn_id,json=loyoTxnId,proto3" json:"loyo_txn_id,omitempty"`
	Status        string `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	ValidEvents   int32  `protobuf:"varint,3,opt,name=valid_events,json=validEvents,proto3" json:"valid_events,omitempty"`
	InvalidEvents int32  `protobuf:"varint,4,opt,name=invalid_events,json=invalidEvents,proto3" json:"invalid_events,omitempty"`
}

func (x *LoyoIngestResponse) Reset() {
	*x = LoyoIngestResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_loyo_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoyoIngestResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoyoIngestResponse) ProtoMessage() {}

func (x *LoyoIngestResponse) ProtoReflect() protoreflect.Message {
	mi := &file_loyo_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoyoIngestResponse.ProtoReflect.Descriptor instead.
func (*LoyoIngestResponse) Descriptor() ([]byte, []int) {
	return file_loyo_proto_rawDescGZIP(), []int{2}
}

func (x *LoyoIngestResponse) GetLoyoTxnId() string {
	if x != nil {
		return x.LoyoTxnId
	}
	return ""
}

func (x *LoyoIngestResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *LoyoIngestResponse) GetValidEvents() int32 {
	if x != nil {
		return x.ValidEvents
	}
	return 0
}

func (x *LoyoIngestResponse) GetInvalidEvents() int32 {
	if x != nil {
		return x.InvalidEvents
	}
	return 0
}

type ErrorResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *ErrorResponse) Reset() {
	*x = ErrorResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_loyo_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ErrorResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ErrorResponse) ProtoMessage() {}

func (x *ErrorResponse) ProtoReflect() protoreflect.Message {
	mi := &file_loyo_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ErrorResponse.ProtoReflect.Descriptor instead.
func (*ErrorResponse) Descriptor() ([]byte, []int) {
	return file_loyo_proto_rawDescGZIP(), []int{3}
}

func (x *ErrorResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *ErrorResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_loyo_proto protoreflect.FileDescriptor

var file_loyo_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x6c, 0x6f, 0x79, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd2, 0x03, 0x0a,
	0x09, 0x4c, 0x6f, 0x79, 0x6f, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x29, 0x0a, 0x0a, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0a,
	0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x09, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x73, 0x6b, 0x75, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x53, 0x6b, 0x75, 0x12, 0x1d, 0x0a,
	0x0a, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x63,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x25, 0x0a, 0x0e, 0x74, 0x61, 0x78, 0x5f, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x62, 0x6c,
	0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0d, 0x74, 0x61, 0x78, 0x41, 0x70, 0x70, 0x6c,
	0x69, 0x63, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x61, 0x78, 0x5f, 0x63, 0x68,
	0x61, 0x72, 0x67, 0x65, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0a, 0x74, 0x61, 0x78,
	0x43, 0x68, 0x61, 0x72, 0x67, 0x65, 0x64, 0x12, 0x30, 0x0a, 0x14, 0x72, 0x65, 0x63, 0x69, 0x70,
	0x69, 0x65, 0x6e, 0x74, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x72, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x66, 0x6c, 0x61,
	0x67, 0x67, 0x65, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x66, 0x6c, 0x61, 0x67,
	0x67, 0x65, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x18, 0x0c, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x12, 0x2e, 0x0a, 0x13, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x70, 0x72, 0x6f,
	0x76, 0x69, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x49,
	0x64, 0x22, 0x5e, 0x0a, 0x11, 0x4c, 0x6f, 0x79, 0x6f, 0x49, 0x6e, 0x67, 0x65, 0x73, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d,
	0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x22, 0x0a,
	0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e,
	0x4c, 0x6f, 0x79, 0x6f, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x73, 0x22, 0x96, 0x01, 0x0a, 0x12, 0x4c, 0x6f, 0x79, 0x6f, 0x49, 0x6e, 0x67, 0x65, 0x73, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x0b, 0x6c, 0x6f, 0x79, 0x6f,
	0x5f, 0x74, 0x78, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6c,
	0x6f, 0x79, 0x6f, 0x54, 0x78, 0x6e, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x21, 0x0a, 0x0c, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x69, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x5f, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x69, 0x6e, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x3d, 0x0a, 0x0d, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2a, 0x96, 0x02, 0x0a, 0x09, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45,
	0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x15, 0x0a, 0x11, 0x50, 0x48, 0x59, 0x53,
	0x49, 0x43, 0x41, 0x4c, 0x5f, 0x50, 0x55, 0x52, 0x43, 0x48, 0x41, 0x53, 0x45, 0x10, 0x01, 0x12,
	0x14, 0x0a, 0x10, 0x44, 0x49, 0x47, 0x49, 0x54, 0x41, 0x4c, 0x5f, 0x50, 0x55, 0x52, 0x43, 0x48,
	0x41, 0x53, 0x45, 0x10, 0x02, 0x12, 0x15, 0x0a, 0x11, 0x4c, 0x4f, 0x41, 0x4e, 0x5f, 0x44, 0x49,
	0x53, 0x42, 0x55, 0x52, 0x53, 0x45, 0x4d, 0x45, 0x4e, 0x54, 0x10, 0x03, 0x12, 0x17, 0x0a, 0x13,
	0x4f, 0x56, 0x45, 0x52, 0x44, 0x55, 0x45, 0x5f, 0x4c, 0x4f, 0x41, 0x4e, 0x5f, 0x43, 0x48, 0x41,
	0x52, 0x47, 0x45, 0x10, 0x04, 0x12, 0x0e, 0x0a, 0x0a, 0x57, 0x49, 0x54, 0x48, 0x44, 0x52, 0x41,
	0x57, 0x41, 0x4c, 0x10, 0x05, 0x12, 0x0f, 0x0a, 0x0b, 0x53, 0x45, 0x52, 0x56, 0x49, 0x43, 0x45,
	0x5f, 0x46, 0x45, 0x45, 0x10, 0x06, 0x12, 0x0c, 0x0a, 0x08, 0x54, 0x52, 0x41, 0x4e, 0x53, 0x46,
	0x45, 0x52, 0x10, 0x07, 0x12, 0x0a, 0x0a, 0x06, 0x52, 0x45, 0x46, 0x55, 0x4e, 0x44, 0x10, 0x08,
	0x12, 0x0b, 0x0a, 0x07, 0x44, 0x45, 0x50, 0x4f, 0x53, 0x49, 0x54, 0x10, 0x09, 0x12, 0x0c, 0x0a,
	0x08, 0x52, 0x45, 0x43, 0x48, 0x41, 0x52, 0x47, 0x45, 0x10, 0x0a, 0x12, 0x16, 0x0a, 0x12, 0x43,
	0x4f, 0x4d, 0x50, 0x45, 0x54, 0x49, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x52, 0x45, 0x57, 0x41, 0x52,
	0x44, 0x10, 0x0b, 0x12, 0x15, 0x0a, 0x11, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49,
	0x45, 0x44, 0x5f, 0x44, 0x45, 0x42, 0x49, 0x54, 0x10, 0x0c, 0x12, 0x16, 0x0a, 0x12, 0x55, 0x4e,
	0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x5f, 0x43, 0x52, 0x45, 0x44, 0x49, 0x54,
	0x10, 0x0d, 0x32, 0x4c, 0x0a, 0x11, 0x4c, 0x6f, 0x79, 0x6f, 0x49, 0x6e, 0x67, 0x65, 0x73, 0x74,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x37, 0x0a, 0x0c, 0x49, 0x6e, 0x67, 0x65, 0x73,
	0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x12, 0x2e, 0x4c, 0x6f, 0x79, 0x6f, 0x49, 0x6e,
	0x67, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x4c, 0x6f,
	0x79, 0x6f, 0x49, 0x6e, 0x67, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x42, 0x07, 0x5a, 0x05, 0x2f, 0x6c, 0x6f, 0x79, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_loyo_proto_rawDescOnce sync.Once
	file_loyo_proto_rawDescData = file_loyo_proto_rawDesc
)

func file_loyo_proto_rawDescGZIP() []byte {
	file_loyo_proto_rawDescOnce.Do(func() {
		file_loyo_proto_rawDescData = protoimpl.X.CompressGZIP(file_loyo_proto_rawDescData)
	})
	return file_loyo_proto_rawDescData
}

var file_loyo_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_loyo_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_loyo_proto_goTypes = []interface{}{
	(EventType)(0),             // 0: EventType
	(*LoyoEvent)(nil),          // 1: LoyoEvent
	(*LoyoIngestRequest)(nil),  // 2: LoyoIngestRequest
	(*LoyoIngestResponse)(nil), // 3: LoyoIngestResponse
	(*ErrorResponse)(nil),      // 4: ErrorResponse
}
var file_loyo_proto_depIdxs = []int32{
	0, // 0: LoyoEvent.event_type:type_name -> EventType
	1, // 1: LoyoIngestRequest.events:type_name -> LoyoEvent
	2, // 2: LoyoIngestService.IngestEvents:input_type -> LoyoIngestRequest
	3, // 3: LoyoIngestService.IngestEvents:output_type -> LoyoIngestResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_loyo_proto_init() }
func file_loyo_proto_init() {
	if File_loyo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_loyo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoyoEvent); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_loyo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoyoIngestRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_loyo_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoyoIngestResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_loyo_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ErrorResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_loyo_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_loyo_proto_goTypes,
		DependencyIndexes: file_loyo_proto_depIdxs,
		EnumInfos:         file_loyo_proto_enumTypes,
		MessageInfos:      file_loyo_proto_msgTypes,
	}.Build()
	File_loyo_proto = out.File
	file_loyo_proto_rawDesc = nil
	file_loyo_proto_goTypes = nil
	file_loyo_proto_depIdxs = nil
}
