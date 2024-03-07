// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.1
// source: agent.proto

package proto

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type InitInstanceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DevEnvNameSlug  string `protobuf:"bytes,1,opt,name=dev_env_name_slug,json=devEnvNameSlug,proto3" json:"dev_env_name_slug,omitempty"`
	GithubUserEmail string `protobuf:"bytes,2,opt,name=github_user_email,json=githubUserEmail,proto3" json:"github_user_email,omitempty"`
	UserFullName    string `protobuf:"bytes,3,opt,name=user_full_name,json=userFullName,proto3" json:"user_full_name,omitempty"`
}

func (x *InitInstanceRequest) Reset() {
	*x = InitInstanceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_agent_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InitInstanceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InitInstanceRequest) ProtoMessage() {}

func (x *InitInstanceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_agent_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InitInstanceRequest.ProtoReflect.Descriptor instead.
func (*InitInstanceRequest) Descriptor() ([]byte, []int) {
	return file_agent_proto_rawDescGZIP(), []int{0}
}

func (x *InitInstanceRequest) GetDevEnvNameSlug() string {
	if x != nil {
		return x.DevEnvNameSlug
	}
	return ""
}

func (x *InitInstanceRequest) GetGithubUserEmail() string {
	if x != nil {
		return x.GithubUserEmail
	}
	return ""
}

func (x *InitInstanceRequest) GetUserFullName() string {
	if x != nil {
		return x.UserFullName
	}
	return ""
}

type InitInstanceReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LogLineHeader             string  `protobuf:"bytes,1,opt,name=log_line_header,json=logLineHeader,proto3" json:"log_line_header,omitempty"`
	LogLine                   string  `protobuf:"bytes,2,opt,name=log_line,json=logLine,proto3" json:"log_line,omitempty"`
	GithubSshPublicKeyContent *string `protobuf:"bytes,3,opt,name=github_ssh_public_key_content,json=githubSshPublicKeyContent,proto3,oneof" json:"github_ssh_public_key_content,omitempty"`
	GithubGpgPublicKeyContent *string `protobuf:"bytes,4,opt,name=github_gpg_public_key_content,json=githubGpgPublicKeyContent,proto3,oneof" json:"github_gpg_public_key_content,omitempty"`
}

func (x *InitInstanceReply) Reset() {
	*x = InitInstanceReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_agent_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InitInstanceReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InitInstanceReply) ProtoMessage() {}

func (x *InitInstanceReply) ProtoReflect() protoreflect.Message {
	mi := &file_agent_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InitInstanceReply.ProtoReflect.Descriptor instead.
func (*InitInstanceReply) Descriptor() ([]byte, []int) {
	return file_agent_proto_rawDescGZIP(), []int{1}
}

func (x *InitInstanceReply) GetLogLineHeader() string {
	if x != nil {
		return x.LogLineHeader
	}
	return ""
}

func (x *InitInstanceReply) GetLogLine() string {
	if x != nil {
		return x.LogLine
	}
	return ""
}

func (x *InitInstanceReply) GetGithubSshPublicKeyContent() string {
	if x != nil && x.GithubSshPublicKeyContent != nil {
		return *x.GithubSshPublicKeyContent
	}
	return ""
}

func (x *InitInstanceReply) GetGithubGpgPublicKeyContent() string {
	if x != nil && x.GithubGpgPublicKeyContent != nil {
		return *x.GithubGpgPublicKeyContent
	}
	return ""
}

type BuildAndStartDevEnvRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DevEnvRepoOwner     string `protobuf:"bytes,1,opt,name=dev_env_repo_owner,json=devEnvRepoOwner,proto3" json:"dev_env_repo_owner,omitempty"`
	DevEnvRepoName      string `protobuf:"bytes,2,opt,name=dev_env_repo_name,json=devEnvRepoName,proto3" json:"dev_env_repo_name,omitempty"`
	UserConfigRepoOwner string `protobuf:"bytes,3,opt,name=user_config_repo_owner,json=userConfigRepoOwner,proto3" json:"user_config_repo_owner,omitempty"`
	UserConfigRepoName  string `protobuf:"bytes,4,opt,name=user_config_repo_name,json=userConfigRepoName,proto3" json:"user_config_repo_name,omitempty"`
}

func (x *BuildAndStartDevEnvRequest) Reset() {
	*x = BuildAndStartDevEnvRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_agent_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BuildAndStartDevEnvRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BuildAndStartDevEnvRequest) ProtoMessage() {}

func (x *BuildAndStartDevEnvRequest) ProtoReflect() protoreflect.Message {
	mi := &file_agent_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BuildAndStartDevEnvRequest.ProtoReflect.Descriptor instead.
func (*BuildAndStartDevEnvRequest) Descriptor() ([]byte, []int) {
	return file_agent_proto_rawDescGZIP(), []int{2}
}

func (x *BuildAndStartDevEnvRequest) GetDevEnvRepoOwner() string {
	if x != nil {
		return x.DevEnvRepoOwner
	}
	return ""
}

func (x *BuildAndStartDevEnvRequest) GetDevEnvRepoName() string {
	if x != nil {
		return x.DevEnvRepoName
	}
	return ""
}

func (x *BuildAndStartDevEnvRequest) GetUserConfigRepoOwner() string {
	if x != nil {
		return x.UserConfigRepoOwner
	}
	return ""
}

func (x *BuildAndStartDevEnvRequest) GetUserConfigRepoName() string {
	if x != nil {
		return x.UserConfigRepoName
	}
	return ""
}

type BuildAndStartDevEnvReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LogLineHeader string `protobuf:"bytes,1,opt,name=log_line_header,json=logLineHeader,proto3" json:"log_line_header,omitempty"`
	LogLine       string `protobuf:"bytes,2,opt,name=log_line,json=logLine,proto3" json:"log_line,omitempty"`
}

func (x *BuildAndStartDevEnvReply) Reset() {
	*x = BuildAndStartDevEnvReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_agent_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BuildAndStartDevEnvReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BuildAndStartDevEnvReply) ProtoMessage() {}

func (x *BuildAndStartDevEnvReply) ProtoReflect() protoreflect.Message {
	mi := &file_agent_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BuildAndStartDevEnvReply.ProtoReflect.Descriptor instead.
func (*BuildAndStartDevEnvReply) Descriptor() ([]byte, []int) {
	return file_agent_proto_rawDescGZIP(), []int{3}
}

func (x *BuildAndStartDevEnvReply) GetLogLineHeader() string {
	if x != nil {
		return x.LogLineHeader
	}
	return ""
}

func (x *BuildAndStartDevEnvReply) GetLogLine() string {
	if x != nil {
		return x.LogLine
	}
	return ""
}

var File_agent_proto protoreflect.FileDescriptor

var file_agent_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x61,
	0x67, 0x65, 0x6e, 0x74, 0x22, 0x92, 0x01, 0x0a, 0x13, 0x49, 0x6e, 0x69, 0x74, 0x49, 0x6e, 0x73,
	0x74, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x29, 0x0a, 0x11,
	0x64, 0x65, 0x76, 0x5f, 0x65, 0x6e, 0x76, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x5f, 0x73, 0x6c, 0x75,
	0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x64, 0x65, 0x76, 0x45, 0x6e, 0x76, 0x4e,
	0x61, 0x6d, 0x65, 0x53, 0x6c, 0x75, 0x67, 0x12, 0x2a, 0x0a, 0x11, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x55, 0x73, 0x65, 0x72, 0x45, 0x6d,
	0x61, 0x69, 0x6c, 0x12, 0x24, 0x0a, 0x0e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x66, 0x75, 0x6c, 0x6c,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x75, 0x73, 0x65,
	0x72, 0x46, 0x75, 0x6c, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0xa8, 0x02, 0x0a, 0x11, 0x49, 0x6e,
	0x69, 0x74, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12,
	0x26, 0x0a, 0x0f, 0x6c, 0x6f, 0x67, 0x5f, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x68, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6c, 0x6f, 0x67, 0x4c, 0x69, 0x6e,
	0x65, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x19, 0x0a, 0x08, 0x6c, 0x6f, 0x67, 0x5f, 0x6c,
	0x69, 0x6e, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6c, 0x6f, 0x67, 0x4c, 0x69,
	0x6e, 0x65, 0x12, 0x45, 0x0a, 0x1d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x5f, 0x73, 0x73, 0x68,
	0x5f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x6b, 0x65, 0x79, 0x5f, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x19, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x53, 0x73, 0x68, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x43,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x88, 0x01, 0x01, 0x12, 0x45, 0x0a, 0x1d, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x5f, 0x67, 0x70, 0x67, 0x5f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x6b,
	0x65, 0x79, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x01, 0x52, 0x19, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x47, 0x70, 0x67, 0x50, 0x75, 0x62,
	0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x88, 0x01, 0x01,
	0x42, 0x20, 0x0a, 0x1e, 0x5f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x5f, 0x73, 0x73, 0x68, 0x5f,
	0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x6b, 0x65, 0x79, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x42, 0x20, 0x0a, 0x1e, 0x5f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x5f, 0x67, 0x70,
	0x67, 0x5f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x6b, 0x65, 0x79, 0x5f, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x22, 0xdc, 0x01, 0x0a, 0x1a, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x41, 0x6e,
	0x64, 0x53, 0x74, 0x61, 0x72, 0x74, 0x44, 0x65, 0x76, 0x45, 0x6e, 0x76, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x2b, 0x0a, 0x12, 0x64, 0x65, 0x76, 0x5f, 0x65, 0x6e, 0x76, 0x5f, 0x72,
	0x65, 0x70, 0x6f, 0x5f, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0f, 0x64, 0x65, 0x76, 0x45, 0x6e, 0x76, 0x52, 0x65, 0x70, 0x6f, 0x4f, 0x77, 0x6e, 0x65, 0x72,
	0x12, 0x29, 0x0a, 0x11, 0x64, 0x65, 0x76, 0x5f, 0x65, 0x6e, 0x76, 0x5f, 0x72, 0x65, 0x70, 0x6f,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x64, 0x65, 0x76,
	0x45, 0x6e, 0x76, 0x52, 0x65, 0x70, 0x6f, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x33, 0x0a, 0x16, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x72, 0x65, 0x70, 0x6f, 0x5f,
	0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x75, 0x73, 0x65,
	0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x70, 0x6f, 0x4f, 0x77, 0x6e, 0x65, 0x72,
	0x12, 0x31, 0x0a, 0x15, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f,
	0x72, 0x65, 0x70, 0x6f, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x12, 0x75, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x70, 0x6f, 0x4e,
	0x61, 0x6d, 0x65, 0x22, 0x5d, 0x0a, 0x18, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x41, 0x6e, 0x64, 0x53,
	0x74, 0x61, 0x72, 0x74, 0x44, 0x65, 0x76, 0x45, 0x6e, 0x76, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12,
	0x26, 0x0a, 0x0f, 0x6c, 0x6f, 0x67, 0x5f, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x68, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6c, 0x6f, 0x67, 0x4c, 0x69, 0x6e,
	0x65, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x19, 0x0a, 0x08, 0x6c, 0x6f, 0x67, 0x5f, 0x6c,
	0x69, 0x6e, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6c, 0x6f, 0x67, 0x4c, 0x69,
	0x6e, 0x65, 0x32, 0xb0, 0x01, 0x0a, 0x05, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x12, 0x48, 0x0a, 0x0c,
	0x49, 0x6e, 0x69, 0x74, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x1a, 0x2e, 0x61,
	0x67, 0x65, 0x6e, 0x74, 0x2e, 0x49, 0x6e, 0x69, 0x74, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74,
	0x2e, 0x49, 0x6e, 0x69, 0x74, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x22, 0x00, 0x30, 0x01, 0x12, 0x5d, 0x0a, 0x13, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x41,
	0x6e, 0x64, 0x53, 0x74, 0x61, 0x72, 0x74, 0x44, 0x65, 0x76, 0x45, 0x6e, 0x76, 0x12, 0x21, 0x2e,
	0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x41, 0x6e, 0x64, 0x53, 0x74,
	0x61, 0x72, 0x74, 0x44, 0x65, 0x76, 0x45, 0x6e, 0x76, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1f, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x41, 0x6e,
	0x64, 0x53, 0x74, 0x61, 0x72, 0x74, 0x44, 0x65, 0x76, 0x45, 0x6e, 0x76, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x22, 0x00, 0x30, 0x01, 0x42, 0x22, 0x5a, 0x20, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x2d, 0x73, 0x68, 0x2f, 0x61, 0x67,
	0x65, 0x6e, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_agent_proto_rawDescOnce sync.Once
	file_agent_proto_rawDescData = file_agent_proto_rawDesc
)

func file_agent_proto_rawDescGZIP() []byte {
	file_agent_proto_rawDescOnce.Do(func() {
		file_agent_proto_rawDescData = protoimpl.X.CompressGZIP(file_agent_proto_rawDescData)
	})
	return file_agent_proto_rawDescData
}

var file_agent_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_agent_proto_goTypes = []interface{}{
	(*InitInstanceRequest)(nil),        // 0: agent.InitInstanceRequest
	(*InitInstanceReply)(nil),          // 1: agent.InitInstanceReply
	(*BuildAndStartDevEnvRequest)(nil), // 2: agent.BuildAndStartDevEnvRequest
	(*BuildAndStartDevEnvReply)(nil),   // 3: agent.BuildAndStartDevEnvReply
}
var file_agent_proto_depIdxs = []int32{
	0, // 0: agent.Agent.InitInstance:input_type -> agent.InitInstanceRequest
	2, // 1: agent.Agent.BuildAndStartDevEnv:input_type -> agent.BuildAndStartDevEnvRequest
	1, // 2: agent.Agent.InitInstance:output_type -> agent.InitInstanceReply
	3, // 3: agent.Agent.BuildAndStartDevEnv:output_type -> agent.BuildAndStartDevEnvReply
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_agent_proto_init() }
func file_agent_proto_init() {
	if File_agent_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_agent_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InitInstanceRequest); i {
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
		file_agent_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InitInstanceReply); i {
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
		file_agent_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BuildAndStartDevEnvRequest); i {
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
		file_agent_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BuildAndStartDevEnvReply); i {
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
	file_agent_proto_msgTypes[1].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_agent_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_agent_proto_goTypes,
		DependencyIndexes: file_agent_proto_depIdxs,
		MessageInfos:      file_agent_proto_msgTypes,
	}.Build()
	File_agent_proto = out.File
	file_agent_proto_rawDesc = nil
	file_agent_proto_goTypes = nil
	file_agent_proto_depIdxs = nil
}
