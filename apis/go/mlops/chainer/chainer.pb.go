/*
Copyright (c) 2024 Seldon Technologies Ltd.

Use of this software is governed BY
(1) the license included in the LICENSE file or
(2) if the license included in the LICENSE file is the Business Source License 1.1,
the Change License after the Change Date as each is defined in accordance with the LICENSE file.
*/

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.2
// source: mlops/chainer/chainer.proto

package chainer

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

type PipelineUpdateMessage_PipelineOperation int32

const (
	PipelineUpdateMessage_Unknown PipelineUpdateMessage_PipelineOperation = 0
	PipelineUpdateMessage_Create  PipelineUpdateMessage_PipelineOperation = 1
	PipelineUpdateMessage_Delete  PipelineUpdateMessage_PipelineOperation = 2
)

// Enum value maps for PipelineUpdateMessage_PipelineOperation.
var (
	PipelineUpdateMessage_PipelineOperation_name = map[int32]string{
		0: "Unknown",
		1: "Create",
		2: "Delete",
	}
	PipelineUpdateMessage_PipelineOperation_value = map[string]int32{
		"Unknown": 0,
		"Create":  1,
		"Delete":  2,
	}
)

func (x PipelineUpdateMessage_PipelineOperation) Enum() *PipelineUpdateMessage_PipelineOperation {
	p := new(PipelineUpdateMessage_PipelineOperation)
	*p = x
	return p
}

func (x PipelineUpdateMessage_PipelineOperation) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PipelineUpdateMessage_PipelineOperation) Descriptor() protoreflect.EnumDescriptor {
	return file_mlops_chainer_chainer_proto_enumTypes[0].Descriptor()
}

func (PipelineUpdateMessage_PipelineOperation) Type() protoreflect.EnumType {
	return &file_mlops_chainer_chainer_proto_enumTypes[0]
}

func (x PipelineUpdateMessage_PipelineOperation) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PipelineUpdateMessage_PipelineOperation.Descriptor instead.
func (PipelineUpdateMessage_PipelineOperation) EnumDescriptor() ([]byte, []int) {
	return file_mlops_chainer_chainer_proto_rawDescGZIP(), []int{1, 0}
}

type PipelineStepUpdate_PipelineJoinType int32

const (
	PipelineStepUpdate_Unknown PipelineStepUpdate_PipelineJoinType = 0
	PipelineStepUpdate_Inner   PipelineStepUpdate_PipelineJoinType = 1
	PipelineStepUpdate_Outer   PipelineStepUpdate_PipelineJoinType = 2
	PipelineStepUpdate_Any     PipelineStepUpdate_PipelineJoinType = 3
)

// Enum value maps for PipelineStepUpdate_PipelineJoinType.
var (
	PipelineStepUpdate_PipelineJoinType_name = map[int32]string{
		0: "Unknown",
		1: "Inner",
		2: "Outer",
		3: "Any",
	}
	PipelineStepUpdate_PipelineJoinType_value = map[string]int32{
		"Unknown": 0,
		"Inner":   1,
		"Outer":   2,
		"Any":     3,
	}
)

func (x PipelineStepUpdate_PipelineJoinType) Enum() *PipelineStepUpdate_PipelineJoinType {
	p := new(PipelineStepUpdate_PipelineJoinType)
	*p = x
	return p
}

func (x PipelineStepUpdate_PipelineJoinType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PipelineStepUpdate_PipelineJoinType) Descriptor() protoreflect.EnumDescriptor {
	return file_mlops_chainer_chainer_proto_enumTypes[1].Descriptor()
}

func (PipelineStepUpdate_PipelineJoinType) Type() protoreflect.EnumType {
	return &file_mlops_chainer_chainer_proto_enumTypes[1]
}

func (x PipelineStepUpdate_PipelineJoinType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PipelineStepUpdate_PipelineJoinType.Descriptor instead.
func (PipelineStepUpdate_PipelineJoinType) EnumDescriptor() ([]byte, []int) {
	return file_mlops_chainer_chainer_proto_rawDescGZIP(), []int{2, 0}
}

type PipelineSubscriptionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *PipelineSubscriptionRequest) Reset() {
	*x = PipelineSubscriptionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mlops_chainer_chainer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PipelineSubscriptionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PipelineSubscriptionRequest) ProtoMessage() {}

func (x *PipelineSubscriptionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mlops_chainer_chainer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PipelineSubscriptionRequest.ProtoReflect.Descriptor instead.
func (*PipelineSubscriptionRequest) Descriptor() ([]byte, []int) {
	return file_mlops_chainer_chainer_proto_rawDescGZIP(), []int{0}
}

func (x *PipelineSubscriptionRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type PipelineUpdateMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Op       PipelineUpdateMessage_PipelineOperation `protobuf:"varint,1,opt,name=op,proto3,enum=seldon.mlops.chainer.PipelineUpdateMessage_PipelineOperation" json:"op,omitempty"`
	Pipeline string                                  `protobuf:"bytes,2,opt,name=pipeline,proto3" json:"pipeline,omitempty"`
	Version  uint32                                  `protobuf:"varint,3,opt,name=version,proto3" json:"version,omitempty"`
	Uid      string                                  `protobuf:"bytes,4,opt,name=uid,proto3" json:"uid,omitempty"`
	Updates  []*PipelineStepUpdate                   `protobuf:"bytes,5,rep,name=updates,proto3" json:"updates,omitempty"`
}

func (x *PipelineUpdateMessage) Reset() {
	*x = PipelineUpdateMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mlops_chainer_chainer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PipelineUpdateMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PipelineUpdateMessage) ProtoMessage() {}

func (x *PipelineUpdateMessage) ProtoReflect() protoreflect.Message {
	mi := &file_mlops_chainer_chainer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PipelineUpdateMessage.ProtoReflect.Descriptor instead.
func (*PipelineUpdateMessage) Descriptor() ([]byte, []int) {
	return file_mlops_chainer_chainer_proto_rawDescGZIP(), []int{1}
}

func (x *PipelineUpdateMessage) GetOp() PipelineUpdateMessage_PipelineOperation {
	if x != nil {
		return x.Op
	}
	return PipelineUpdateMessage_Unknown
}

func (x *PipelineUpdateMessage) GetPipeline() string {
	if x != nil {
		return x.Pipeline
	}
	return ""
}

func (x *PipelineUpdateMessage) GetVersion() uint32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *PipelineUpdateMessage) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *PipelineUpdateMessage) GetUpdates() []*PipelineStepUpdate {
	if x != nil {
		return x.Updates
	}
	return nil
}

type PipelineStepUpdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// https://docs.google.com/document/d/1tX-uaOvngx1RpEyWEZ4EbEcU8D0OgYuRWVb2UAi85n4/edit
	// Pipeline Resource example, e.g. transform.outputs.traffic
	//
	//	seldon.<namespace>.<model name>.<inputs|outputs>.<tensor name>
	Sources            []*PipelineTopic                    `protobuf:"bytes,1,rep,name=sources,proto3" json:"sources,omitempty"`
	Triggers           []*PipelineTopic                    `protobuf:"bytes,2,rep,name=triggers,proto3" json:"triggers,omitempty"`
	Sink               *PipelineTopic                      `protobuf:"bytes,3,opt,name=sink,proto3" json:"sink,omitempty"`
	InputJoinTy        PipelineStepUpdate_PipelineJoinType `protobuf:"varint,4,opt,name=inputJoinTy,proto3,enum=seldon.mlops.chainer.PipelineStepUpdate_PipelineJoinType" json:"inputJoinTy,omitempty"`
	TriggersJoinTy     PipelineStepUpdate_PipelineJoinType `protobuf:"varint,5,opt,name=triggersJoinTy,proto3,enum=seldon.mlops.chainer.PipelineStepUpdate_PipelineJoinType" json:"triggersJoinTy,omitempty"`
	PassEmptyResponses bool                                `protobuf:"varint,6,opt,name=passEmptyResponses,proto3" json:"passEmptyResponses,omitempty"` // Forward empty response to following steps, default false
	JoinWindowMs       *uint32                             `protobuf:"varint,7,opt,name=joinWindowMs,proto3,oneof" json:"joinWindowMs,omitempty"`       // Join window millisecs, some nozero default (TBD)
	TensorMap          []*PipelineTensorMapping            `protobuf:"bytes,8,rep,name=tensorMap,proto3" json:"tensorMap,omitempty"`                    // optional list of tensor name mappings
	Batch              *Batch                              `protobuf:"bytes,9,opt,name=batch,proto3" json:"batch,omitempty"`                            // Batch settings
}

func (x *PipelineStepUpdate) Reset() {
	*x = PipelineStepUpdate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mlops_chainer_chainer_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PipelineStepUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PipelineStepUpdate) ProtoMessage() {}

func (x *PipelineStepUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_mlops_chainer_chainer_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PipelineStepUpdate.ProtoReflect.Descriptor instead.
func (*PipelineStepUpdate) Descriptor() ([]byte, []int) {
	return file_mlops_chainer_chainer_proto_rawDescGZIP(), []int{2}
}

func (x *PipelineStepUpdate) GetSources() []*PipelineTopic {
	if x != nil {
		return x.Sources
	}
	return nil
}

func (x *PipelineStepUpdate) GetTriggers() []*PipelineTopic {
	if x != nil {
		return x.Triggers
	}
	return nil
}

func (x *PipelineStepUpdate) GetSink() *PipelineTopic {
	if x != nil {
		return x.Sink
	}
	return nil
}

func (x *PipelineStepUpdate) GetInputJoinTy() PipelineStepUpdate_PipelineJoinType {
	if x != nil {
		return x.InputJoinTy
	}
	return PipelineStepUpdate_Unknown
}

func (x *PipelineStepUpdate) GetTriggersJoinTy() PipelineStepUpdate_PipelineJoinType {
	if x != nil {
		return x.TriggersJoinTy
	}
	return PipelineStepUpdate_Unknown
}

func (x *PipelineStepUpdate) GetPassEmptyResponses() bool {
	if x != nil {
		return x.PassEmptyResponses
	}
	return false
}

func (x *PipelineStepUpdate) GetJoinWindowMs() uint32 {
	if x != nil && x.JoinWindowMs != nil {
		return *x.JoinWindowMs
	}
	return 0
}

func (x *PipelineStepUpdate) GetTensorMap() []*PipelineTensorMapping {
	if x != nil {
		return x.TensorMap
	}
	return nil
}

func (x *PipelineStepUpdate) GetBatch() *Batch {
	if x != nil {
		return x.Batch
	}
	return nil
}

type PipelineTensorMapping struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PipelineName   string `protobuf:"bytes,1,opt,name=pipelineName,proto3" json:"pipelineName,omitempty"`
	TopicAndTensor string `protobuf:"bytes,2,opt,name=topicAndTensor,proto3" json:"topicAndTensor,omitempty"`
	TensorName     string `protobuf:"bytes,3,opt,name=tensorName,proto3" json:"tensorName,omitempty"`
}

func (x *PipelineTensorMapping) Reset() {
	*x = PipelineTensorMapping{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mlops_chainer_chainer_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PipelineTensorMapping) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PipelineTensorMapping) ProtoMessage() {}

func (x *PipelineTensorMapping) ProtoReflect() protoreflect.Message {
	mi := &file_mlops_chainer_chainer_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PipelineTensorMapping.ProtoReflect.Descriptor instead.
func (*PipelineTensorMapping) Descriptor() ([]byte, []int) {
	return file_mlops_chainer_chainer_proto_rawDescGZIP(), []int{3}
}

func (x *PipelineTensorMapping) GetPipelineName() string {
	if x != nil {
		return x.PipelineName
	}
	return ""
}

func (x *PipelineTensorMapping) GetTopicAndTensor() string {
	if x != nil {
		return x.TopicAndTensor
	}
	return ""
}

func (x *PipelineTensorMapping) GetTensorName() string {
	if x != nil {
		return x.TensorName
	}
	return ""
}

type PipelineTopic struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PipelineName string  `protobuf:"bytes,1,opt,name=pipelineName,proto3" json:"pipelineName,omitempty"`
	TopicName    string  `protobuf:"bytes,2,opt,name=topicName,proto3" json:"topicName,omitempty"`
	Tensor       *string `protobuf:"bytes,3,opt,name=tensor,proto3,oneof" json:"tensor,omitempty"`
}

func (x *PipelineTopic) Reset() {
	*x = PipelineTopic{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mlops_chainer_chainer_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PipelineTopic) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PipelineTopic) ProtoMessage() {}

func (x *PipelineTopic) ProtoReflect() protoreflect.Message {
	mi := &file_mlops_chainer_chainer_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PipelineTopic.ProtoReflect.Descriptor instead.
func (*PipelineTopic) Descriptor() ([]byte, []int) {
	return file_mlops_chainer_chainer_proto_rawDescGZIP(), []int{4}
}

func (x *PipelineTopic) GetPipelineName() string {
	if x != nil {
		return x.PipelineName
	}
	return ""
}

func (x *PipelineTopic) GetTopicName() string {
	if x != nil {
		return x.TopicName
	}
	return ""
}

func (x *PipelineTopic) GetTensor() string {
	if x != nil && x.Tensor != nil {
		return *x.Tensor
	}
	return ""
}

type Batch struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Size     *uint32 `protobuf:"varint,1,opt,name=size,proto3,oneof" json:"size,omitempty"`
	WindowMs *uint32 `protobuf:"varint,2,opt,name=windowMs,proto3,oneof" json:"windowMs,omitempty"`
	Rolling  bool    `protobuf:"varint,3,opt,name=rolling,proto3" json:"rolling,omitempty"`
}

func (x *Batch) Reset() {
	*x = Batch{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mlops_chainer_chainer_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Batch) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Batch) ProtoMessage() {}

func (x *Batch) ProtoReflect() protoreflect.Message {
	mi := &file_mlops_chainer_chainer_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Batch.ProtoReflect.Descriptor instead.
func (*Batch) Descriptor() ([]byte, []int) {
	return file_mlops_chainer_chainer_proto_rawDescGZIP(), []int{5}
}

func (x *Batch) GetSize() uint32 {
	if x != nil && x.Size != nil {
		return *x.Size
	}
	return 0
}

func (x *Batch) GetWindowMs() uint32 {
	if x != nil && x.WindowMs != nil {
		return *x.WindowMs
	}
	return 0
}

func (x *Batch) GetRolling() bool {
	if x != nil {
		return x.Rolling
	}
	return false
}

type PipelineUpdateStatusMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// TODO - include `name` to identify transformer message comes from
	Update  *PipelineUpdateMessage `protobuf:"bytes,1,opt,name=update,proto3" json:"update,omitempty"`
	Success bool                   `protobuf:"varint,2,opt,name=success,proto3" json:"success,omitempty"`
	Reason  string                 `protobuf:"bytes,3,opt,name=reason,proto3" json:"reason,omitempty"`
}

func (x *PipelineUpdateStatusMessage) Reset() {
	*x = PipelineUpdateStatusMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mlops_chainer_chainer_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PipelineUpdateStatusMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PipelineUpdateStatusMessage) ProtoMessage() {}

func (x *PipelineUpdateStatusMessage) ProtoReflect() protoreflect.Message {
	mi := &file_mlops_chainer_chainer_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PipelineUpdateStatusMessage.ProtoReflect.Descriptor instead.
func (*PipelineUpdateStatusMessage) Descriptor() ([]byte, []int) {
	return file_mlops_chainer_chainer_proto_rawDescGZIP(), []int{6}
}

func (x *PipelineUpdateStatusMessage) GetUpdate() *PipelineUpdateMessage {
	if x != nil {
		return x.Update
	}
	return nil
}

func (x *PipelineUpdateStatusMessage) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *PipelineUpdateStatusMessage) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

type PipelineUpdateStatusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PipelineUpdateStatusResponse) Reset() {
	*x = PipelineUpdateStatusResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mlops_chainer_chainer_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PipelineUpdateStatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PipelineUpdateStatusResponse) ProtoMessage() {}

func (x *PipelineUpdateStatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mlops_chainer_chainer_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PipelineUpdateStatusResponse.ProtoReflect.Descriptor instead.
func (*PipelineUpdateStatusResponse) Descriptor() ([]byte, []int) {
	return file_mlops_chainer_chainer_proto_rawDescGZIP(), []int{7}
}

var File_mlops_chainer_chainer_proto protoreflect.FileDescriptor

var file_mlops_chainer_chainer_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x6d, 0x6c, 0x6f, 0x70, 0x73, 0x2f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x2f,
	0x63, 0x68, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x73,
	0x65, 0x6c, 0x64, 0x6f, 0x6e, 0x2e, 0x6d, 0x6c, 0x6f, 0x70, 0x73, 0x2e, 0x63, 0x68, 0x61, 0x69,
	0x6e, 0x65, 0x72, 0x22, 0x31, 0x0a, 0x1b, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x53,
	0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0xac, 0x02, 0x0a, 0x15, 0x50, 0x69, 0x70, 0x65, 0x6c,
	0x69, 0x6e, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x4d, 0x0a, 0x02, 0x6f, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x3d, 0x2e, 0x73,
	0x65, 0x6c, 0x64, 0x6f, 0x6e, 0x2e, 0x6d, 0x6c, 0x6f, 0x70, 0x73, 0x2e, 0x63, 0x68, 0x61, 0x69,
	0x6e, 0x65, 0x72, 0x2e, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69,
	0x6e, 0x65, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x02, 0x6f, 0x70, 0x12,
	0x1a, 0x0a, 0x08, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x42, 0x0a, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x73, 0x65, 0x6c, 0x64, 0x6f,
	0x6e, 0x2e, 0x6d, 0x6c, 0x6f, 0x70, 0x73, 0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x2e,
	0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x53, 0x74, 0x65, 0x70, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x52, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x73, 0x22, 0x38, 0x0a, 0x11, 0x50,
	0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x0b, 0x0a, 0x07, 0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x10, 0x00, 0x12, 0x0a, 0x0a,
	0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x10, 0x02, 0x22, 0xb5, 0x05, 0x0a, 0x12, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69,
	0x6e, 0x65, 0x53, 0x74, 0x65, 0x70, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x3d, 0x0a, 0x07,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e,
	0x73, 0x65, 0x6c, 0x64, 0x6f, 0x6e, 0x2e, 0x6d, 0x6c, 0x6f, 0x70, 0x73, 0x2e, 0x63, 0x68, 0x61,
	0x69, 0x6e, 0x65, 0x72, 0x2e, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x54, 0x6f, 0x70,
	0x69, 0x63, 0x52, 0x07, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x12, 0x3f, 0x0a, 0x08, 0x74,
	0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e,
	0x73, 0x65, 0x6c, 0x64, 0x6f, 0x6e, 0x2e, 0x6d, 0x6c, 0x6f, 0x70, 0x73, 0x2e, 0x63, 0x68, 0x61,
	0x69, 0x6e, 0x65, 0x72, 0x2e, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x54, 0x6f, 0x70,
	0x69, 0x63, 0x52, 0x08, 0x74, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x73, 0x12, 0x37, 0x0a, 0x04,
	0x73, 0x69, 0x6e, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x73, 0x65, 0x6c,
	0x64, 0x6f, 0x6e, 0x2e, 0x6d, 0x6c, 0x6f, 0x70, 0x73, 0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x65,
	0x72, 0x2e, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x52,
	0x04, 0x73, 0x69, 0x6e, 0x6b, 0x12, 0x5b, 0x0a, 0x0b, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x4a, 0x6f,
	0x69, 0x6e, 0x54, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x39, 0x2e, 0x73, 0x65, 0x6c,
	0x64, 0x6f, 0x6e, 0x2e, 0x6d, 0x6c, 0x6f, 0x70, 0x73, 0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x65,
	0x72, 0x2e, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x53, 0x74, 0x65, 0x70, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x2e, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x4a, 0x6f, 0x69,
	0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0b, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x4a, 0x6f, 0x69, 0x6e,
	0x54, 0x79, 0x12, 0x61, 0x0a, 0x0e, 0x74, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x73, 0x4a, 0x6f,
	0x69, 0x6e, 0x54, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x39, 0x2e, 0x73, 0x65, 0x6c,
	0x64, 0x6f, 0x6e, 0x2e, 0x6d, 0x6c, 0x6f, 0x70, 0x73, 0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x65,
	0x72, 0x2e, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x53, 0x74, 0x65, 0x70, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x2e, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x4a, 0x6f, 0x69,
	0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0e, 0x74, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x73, 0x4a,
	0x6f, 0x69, 0x6e, 0x54, 0x79, 0x12, 0x2e, 0x0a, 0x12, 0x70, 0x61, 0x73, 0x73, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x12, 0x70, 0x61, 0x73, 0x73, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x73, 0x12, 0x27, 0x0a, 0x0c, 0x6a, 0x6f, 0x69, 0x6e, 0x57, 0x69, 0x6e,
	0x64, 0x6f, 0x77, 0x4d, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x00, 0x52, 0x0c, 0x6a,
	0x6f, 0x69, 0x6e, 0x57, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x4d, 0x73, 0x88, 0x01, 0x01, 0x12, 0x49,
	0x0a, 0x09, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x4d, 0x61, 0x70, 0x18, 0x08, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x2b, 0x2e, 0x73, 0x65, 0x6c, 0x64, 0x6f, 0x6e, 0x2e, 0x6d, 0x6c, 0x6f, 0x70, 0x73,
	0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x2e, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e,
	0x65, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x52, 0x09,
	0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x4d, 0x61, 0x70, 0x12, 0x31, 0x0a, 0x05, 0x62, 0x61, 0x74,
	0x63, 0x68, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x73, 0x65, 0x6c, 0x64, 0x6f,
	0x6e, 0x2e, 0x6d, 0x6c, 0x6f, 0x70, 0x73, 0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x2e,
	0x42, 0x61, 0x74, 0x63, 0x68, 0x52, 0x05, 0x62, 0x61, 0x74, 0x63, 0x68, 0x22, 0x3e, 0x0a, 0x10,
	0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x4a, 0x6f, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x0b, 0x0a, 0x07, 0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x10, 0x00, 0x12, 0x09, 0x0a,
	0x05, 0x49, 0x6e, 0x6e, 0x65, 0x72, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x4f, 0x75, 0x74, 0x65,
	0x72, 0x10, 0x02, 0x12, 0x07, 0x0a, 0x03, 0x41, 0x6e, 0x79, 0x10, 0x03, 0x42, 0x0f, 0x0a, 0x0d,
	0x5f, 0x6a, 0x6f, 0x69, 0x6e, 0x57, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x4d, 0x73, 0x22, 0x83, 0x01,
	0x0a, 0x15, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72,
	0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x12, 0x22, 0x0a, 0x0c, 0x70, 0x69, 0x70, 0x65, 0x6c,
	0x69, 0x6e, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70,
	0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x26, 0x0a, 0x0e, 0x74,
	0x6f, 0x70, 0x69, 0x63, 0x41, 0x6e, 0x64, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0e, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x41, 0x6e, 0x64, 0x54, 0x65, 0x6e,
	0x73, 0x6f, 0x72, 0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x4e,
	0x61, 0x6d, 0x65, 0x22, 0x79, 0x0a, 0x0d, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x54,
	0x6f, 0x70, 0x69, 0x63, 0x12, 0x22, 0x0a, 0x0c, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x69, 0x70, 0x65,
	0x6c, 0x69, 0x6e, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x6f, 0x70, 0x69,
	0x63, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x6f, 0x70,
	0x69, 0x63, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x06, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x06, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72,
	0x88, 0x01, 0x01, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x22, 0x71,
	0x0a, 0x05, 0x42, 0x61, 0x74, 0x63, 0x68, 0x12, 0x17, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x00, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x88, 0x01, 0x01,
	0x12, 0x1f, 0x0a, 0x08, 0x77, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x4d, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0d, 0x48, 0x01, 0x52, 0x08, 0x77, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x4d, 0x73, 0x88, 0x01,
	0x01, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x6f, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x07, 0x72, 0x6f, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x42, 0x07, 0x0a, 0x05, 0x5f,
	0x73, 0x69, 0x7a, 0x65, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x77, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x4d,
	0x73, 0x22, 0x94, 0x01, 0x0a, 0x1b, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x43, 0x0a, 0x06, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x2b, 0x2e, 0x73, 0x65, 0x6c, 0x64, 0x6f, 0x6e, 0x2e, 0x6d, 0x6c, 0x6f, 0x70, 0x73,
	0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x2e, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e,
	0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x06,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x22, 0x1e, 0x0a, 0x1c, 0x50, 0x69, 0x70, 0x65,
	0x6c, 0x69, 0x6e, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x89, 0x02, 0x0a, 0x07, 0x43, 0x68, 0x61,
	0x69, 0x6e, 0x65, 0x72, 0x12, 0x7e, 0x0a, 0x18, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62,
	0x65, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x73,
	0x12, 0x31, 0x2e, 0x73, 0x65, 0x6c, 0x64, 0x6f, 0x6e, 0x2e, 0x6d, 0x6c, 0x6f, 0x70, 0x73, 0x2e,
	0x63, 0x68, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x2e, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65,
	0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x73, 0x65, 0x6c, 0x64, 0x6f, 0x6e, 0x2e, 0x6d, 0x6c, 0x6f,
	0x70, 0x73, 0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x2e, 0x50, 0x69, 0x70, 0x65, 0x6c,
	0x69, 0x6e, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x22, 0x00, 0x30, 0x01, 0x12, 0x7e, 0x0a, 0x13, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x31, 0x2e, 0x73, 0x65,
	0x6c, 0x64, 0x6f, 0x6e, 0x2e, 0x6d, 0x6c, 0x6f, 0x70, 0x73, 0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e,
	0x65, 0x72, 0x2e, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x32,
	0x2e, 0x73, 0x65, 0x6c, 0x64, 0x6f, 0x6e, 0x2e, 0x6d, 0x6c, 0x6f, 0x70, 0x73, 0x2e, 0x63, 0x68,
	0x61, 0x69, 0x6e, 0x65, 0x72, 0x2e, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x42, 0x53, 0x0a, 0x17, 0x69, 0x6f, 0x2e, 0x73, 0x65, 0x6c, 0x64, 0x6f,
	0x6e, 0x2e, 0x6d, 0x6c, 0x6f, 0x70, 0x73, 0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x5a,
	0x38, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x65, 0x6c, 0x64,
	0x6f, 0x6e, 0x69, 0x6f, 0x2f, 0x73, 0x65, 0x6c, 0x64, 0x6f, 0x6e, 0x2d, 0x63, 0x6f, 0x72, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x67, 0x6f, 0x2f, 0x76, 0x32, 0x2f, 0x6d, 0x6c, 0x6f, 0x70,
	0x73, 0x2f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_mlops_chainer_chainer_proto_rawDescOnce sync.Once
	file_mlops_chainer_chainer_proto_rawDescData = file_mlops_chainer_chainer_proto_rawDesc
)

func file_mlops_chainer_chainer_proto_rawDescGZIP() []byte {
	file_mlops_chainer_chainer_proto_rawDescOnce.Do(func() {
		file_mlops_chainer_chainer_proto_rawDescData = protoimpl.X.CompressGZIP(file_mlops_chainer_chainer_proto_rawDescData)
	})
	return file_mlops_chainer_chainer_proto_rawDescData
}

var file_mlops_chainer_chainer_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_mlops_chainer_chainer_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_mlops_chainer_chainer_proto_goTypes = []any{
	(PipelineUpdateMessage_PipelineOperation)(0), // 0: seldon.mlops.chainer.PipelineUpdateMessage.PipelineOperation
	(PipelineStepUpdate_PipelineJoinType)(0),     // 1: seldon.mlops.chainer.PipelineStepUpdate.PipelineJoinType
	(*PipelineSubscriptionRequest)(nil),          // 2: seldon.mlops.chainer.PipelineSubscriptionRequest
	(*PipelineUpdateMessage)(nil),                // 3: seldon.mlops.chainer.PipelineUpdateMessage
	(*PipelineStepUpdate)(nil),                   // 4: seldon.mlops.chainer.PipelineStepUpdate
	(*PipelineTensorMapping)(nil),                // 5: seldon.mlops.chainer.PipelineTensorMapping
	(*PipelineTopic)(nil),                        // 6: seldon.mlops.chainer.PipelineTopic
	(*Batch)(nil),                                // 7: seldon.mlops.chainer.Batch
	(*PipelineUpdateStatusMessage)(nil),          // 8: seldon.mlops.chainer.PipelineUpdateStatusMessage
	(*PipelineUpdateStatusResponse)(nil),         // 9: seldon.mlops.chainer.PipelineUpdateStatusResponse
}
var file_mlops_chainer_chainer_proto_depIdxs = []int32{
	0,  // 0: seldon.mlops.chainer.PipelineUpdateMessage.op:type_name -> seldon.mlops.chainer.PipelineUpdateMessage.PipelineOperation
	4,  // 1: seldon.mlops.chainer.PipelineUpdateMessage.updates:type_name -> seldon.mlops.chainer.PipelineStepUpdate
	6,  // 2: seldon.mlops.chainer.PipelineStepUpdate.sources:type_name -> seldon.mlops.chainer.PipelineTopic
	6,  // 3: seldon.mlops.chainer.PipelineStepUpdate.triggers:type_name -> seldon.mlops.chainer.PipelineTopic
	6,  // 4: seldon.mlops.chainer.PipelineStepUpdate.sink:type_name -> seldon.mlops.chainer.PipelineTopic
	1,  // 5: seldon.mlops.chainer.PipelineStepUpdate.inputJoinTy:type_name -> seldon.mlops.chainer.PipelineStepUpdate.PipelineJoinType
	1,  // 6: seldon.mlops.chainer.PipelineStepUpdate.triggersJoinTy:type_name -> seldon.mlops.chainer.PipelineStepUpdate.PipelineJoinType
	5,  // 7: seldon.mlops.chainer.PipelineStepUpdate.tensorMap:type_name -> seldon.mlops.chainer.PipelineTensorMapping
	7,  // 8: seldon.mlops.chainer.PipelineStepUpdate.batch:type_name -> seldon.mlops.chainer.Batch
	3,  // 9: seldon.mlops.chainer.PipelineUpdateStatusMessage.update:type_name -> seldon.mlops.chainer.PipelineUpdateMessage
	2,  // 10: seldon.mlops.chainer.Chainer.SubscribePipelineUpdates:input_type -> seldon.mlops.chainer.PipelineSubscriptionRequest
	8,  // 11: seldon.mlops.chainer.Chainer.PipelineUpdateEvent:input_type -> seldon.mlops.chainer.PipelineUpdateStatusMessage
	3,  // 12: seldon.mlops.chainer.Chainer.SubscribePipelineUpdates:output_type -> seldon.mlops.chainer.PipelineUpdateMessage
	9,  // 13: seldon.mlops.chainer.Chainer.PipelineUpdateEvent:output_type -> seldon.mlops.chainer.PipelineUpdateStatusResponse
	12, // [12:14] is the sub-list for method output_type
	10, // [10:12] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_mlops_chainer_chainer_proto_init() }
func file_mlops_chainer_chainer_proto_init() {
	if File_mlops_chainer_chainer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mlops_chainer_chainer_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*PipelineSubscriptionRequest); i {
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
		file_mlops_chainer_chainer_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*PipelineUpdateMessage); i {
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
		file_mlops_chainer_chainer_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*PipelineStepUpdate); i {
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
		file_mlops_chainer_chainer_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*PipelineTensorMapping); i {
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
		file_mlops_chainer_chainer_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*PipelineTopic); i {
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
		file_mlops_chainer_chainer_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*Batch); i {
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
		file_mlops_chainer_chainer_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*PipelineUpdateStatusMessage); i {
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
		file_mlops_chainer_chainer_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*PipelineUpdateStatusResponse); i {
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
	file_mlops_chainer_chainer_proto_msgTypes[2].OneofWrappers = []any{}
	file_mlops_chainer_chainer_proto_msgTypes[4].OneofWrappers = []any{}
	file_mlops_chainer_chainer_proto_msgTypes[5].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_mlops_chainer_chainer_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_mlops_chainer_chainer_proto_goTypes,
		DependencyIndexes: file_mlops_chainer_chainer_proto_depIdxs,
		EnumInfos:         file_mlops_chainer_chainer_proto_enumTypes,
		MessageInfos:      file_mlops_chainer_chainer_proto_msgTypes,
	}.Build()
	File_mlops_chainer_chainer_proto = out.File
	file_mlops_chainer_chainer_proto_rawDesc = nil
	file_mlops_chainer_chainer_proto_goTypes = nil
	file_mlops_chainer_chainer_proto_depIdxs = nil
}
