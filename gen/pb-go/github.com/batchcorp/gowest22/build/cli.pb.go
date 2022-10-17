// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: cli.proto

package build

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

type SubCommand struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: kong:"help='Sub action to perform',default='bar'"
	Action string `protobuf:"bytes,1,opt,name=action,proto3" json:"action,omitempty"`
	// @gotags: kong:"help='Optional arguments for sub action',optional"
	Args []string `protobuf:"bytes,2,rep,name=args,proto3" json:"args,omitempty"`
	// @gotags: kong:"help='Number of times to execute',default=1"
	Count int64 `protobuf:"varint,3,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *SubCommand) Reset() {
	*x = SubCommand{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cli_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubCommand) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubCommand) ProtoMessage() {}

func (x *SubCommand) ProtoReflect() protoreflect.Message {
	mi := &file_cli_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubCommand.ProtoReflect.Descriptor instead.
func (*SubCommand) Descriptor() ([]byte, []int) {
	return file_cli_proto_rawDescGZIP(), []int{0}
}

func (x *SubCommand) GetAction() string {
	if x != nil {
		return x.Action
	}
	return ""
}

func (x *SubCommand) GetArgs() []string {
	if x != nil {
		return x.Args
	}
	return nil
}

func (x *SubCommand) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

type Base struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: kong:"help='Action to perform',default='foo'"
	Action string `protobuf:"bytes,1,opt,name=action,proto3" json:"action,omitempty"`
	// @gotags: kong:"cmd,help='Execute subcommand'"
	Sub *SubCommand `protobuf:"bytes,2,opt,name=sub,proto3" json:"sub,omitempty"`
	// @gotags: kong:"help='Enable debug output',short=d"
	Debug bool `protobuf:"varint,3,opt,name=debug,proto3" json:"debug,omitempty"`
	// @gotags: kong:"help='Display plumber version'"
	Version string `protobuf:"bytes,4,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *Base) Reset() {
	*x = Base{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cli_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Base) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Base) ProtoMessage() {}

func (x *Base) ProtoReflect() protoreflect.Message {
	mi := &file_cli_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Base.ProtoReflect.Descriptor instead.
func (*Base) Descriptor() ([]byte, []int) {
	return file_cli_proto_rawDescGZIP(), []int{1}
}

func (x *Base) GetAction() string {
	if x != nil {
		return x.Action
	}
	return ""
}

func (x *Base) GetSub() *SubCommand {
	if x != nil {
		return x.Sub
	}
	return nil
}

func (x *Base) GetDebug() bool {
	if x != nil {
		return x.Debug
	}
	return false
}

func (x *Base) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

var File_cli_proto protoreflect.FileDescriptor

var file_cli_proto_rawDesc = []byte{
	0x0a, 0x09, 0x63, 0x6c, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x66, 0x6c, 0x61,
	0x67, 0x73, 0x22, 0x4e, 0x0a, 0x0a, 0x53, 0x75, 0x62, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64,
	0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x72, 0x67, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x61, 0x72, 0x67, 0x73, 0x12, 0x14, 0x0a, 0x05,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x22, 0x73, 0x0a, 0x04, 0x42, 0x61, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x23, 0x0a, 0x03, 0x73, 0x75, 0x62, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x11, 0x2e, 0x66, 0x6c, 0x61, 0x67, 0x73, 0x2e, 0x53, 0x75, 0x62, 0x43, 0x6f, 0x6d, 0x6d, 0x61,
	0x6e, 0x64, 0x52, 0x03, 0x73, 0x75, 0x62, 0x12, 0x14, 0x0a, 0x05, 0x64, 0x65, 0x62, 0x75, 0x67,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x64, 0x65, 0x62, 0x75, 0x67, 0x12, 0x18, 0x0a,
	0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x42, 0x25, 0x5a, 0x23, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x61, 0x74, 0x63, 0x68, 0x63, 0x6f, 0x72, 0x70, 0x2f,
	0x67, 0x6f, 0x77, 0x65, 0x73, 0x74, 0x32, 0x32, 0x2f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cli_proto_rawDescOnce sync.Once
	file_cli_proto_rawDescData = file_cli_proto_rawDesc
)

func file_cli_proto_rawDescGZIP() []byte {
	file_cli_proto_rawDescOnce.Do(func() {
		file_cli_proto_rawDescData = protoimpl.X.CompressGZIP(file_cli_proto_rawDescData)
	})
	return file_cli_proto_rawDescData
}

var file_cli_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_cli_proto_goTypes = []interface{}{
	(*SubCommand)(nil), // 0: flags.SubCommand
	(*Base)(nil),       // 1: flags.Base
}
var file_cli_proto_depIdxs = []int32{
	0, // 0: flags.Base.sub:type_name -> flags.SubCommand
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_cli_proto_init() }
func file_cli_proto_init() {
	if File_cli_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cli_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubCommand); i {
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
		file_cli_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Base); i {
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
			RawDescriptor: file_cli_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cli_proto_goTypes,
		DependencyIndexes: file_cli_proto_depIdxs,
		MessageInfos:      file_cli_proto_msgTypes,
	}.Build()
	File_cli_proto = out.File
	file_cli_proto_rawDesc = nil
	file_cli_proto_goTypes = nil
	file_cli_proto_depIdxs = nil
}
