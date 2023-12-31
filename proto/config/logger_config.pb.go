// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.12.4
// source: proto/config/logger_config.proto

package config

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

type LoggerConfig_LogLevel int32

const (
	LoggerConfig_LOG_LEVEL_DEBUG   LoggerConfig_LogLevel = 1
	LoggerConfig_LOG_LEVEL_INFO    LoggerConfig_LogLevel = 2
	LoggerConfig_LOG_LEVEL_WARNING LoggerConfig_LogLevel = 3
	LoggerConfig_LOG_LEVEL_ERROR   LoggerConfig_LogLevel = 4
	LoggerConfig_LOG_LEVEL_FATAL   LoggerConfig_LogLevel = 5
)

// Enum value maps for LoggerConfig_LogLevel.
var (
	LoggerConfig_LogLevel_name = map[int32]string{
		1: "LOG_LEVEL_DEBUG",
		2: "LOG_LEVEL_INFO",
		3: "LOG_LEVEL_WARNING",
		4: "LOG_LEVEL_ERROR",
		5: "LOG_LEVEL_FATAL",
	}
	LoggerConfig_LogLevel_value = map[string]int32{
		"LOG_LEVEL_DEBUG":   1,
		"LOG_LEVEL_INFO":    2,
		"LOG_LEVEL_WARNING": 3,
		"LOG_LEVEL_ERROR":   4,
		"LOG_LEVEL_FATAL":   5,
	}
)

func (x LoggerConfig_LogLevel) Enum() *LoggerConfig_LogLevel {
	p := new(LoggerConfig_LogLevel)
	*p = x
	return p
}

func (x LoggerConfig_LogLevel) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LoggerConfig_LogLevel) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_config_logger_config_proto_enumTypes[0].Descriptor()
}

func (LoggerConfig_LogLevel) Type() protoreflect.EnumType {
	return &file_proto_config_logger_config_proto_enumTypes[0]
}

func (x LoggerConfig_LogLevel) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *LoggerConfig_LogLevel) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = LoggerConfig_LogLevel(num)
	return nil
}

// Deprecated: Use LoggerConfig_LogLevel.Descriptor instead.
func (LoggerConfig_LogLevel) EnumDescriptor() ([]byte, []int) {
	return file_proto_config_logger_config_proto_rawDescGZIP(), []int{0, 0}
}

type LoggerConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileWriterConfig    *LoggerConfig_FileWriterConfig    `protobuf:"bytes,1,opt,name=file_writer_config,json=fileWriterConfig" json:"file_writer_config,omitempty"`
	ConsoleWriterConfig *LoggerConfig_ConsoleWriterConfig `protobuf:"bytes,2,opt,name=console_writer_config,json=consoleWriterConfig" json:"console_writer_config,omitempty"`
}

func (x *LoggerConfig) Reset() {
	*x = LoggerConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_config_logger_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoggerConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoggerConfig) ProtoMessage() {}

func (x *LoggerConfig) ProtoReflect() protoreflect.Message {
	mi := &file_proto_config_logger_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoggerConfig.ProtoReflect.Descriptor instead.
func (*LoggerConfig) Descriptor() ([]byte, []int) {
	return file_proto_config_logger_config_proto_rawDescGZIP(), []int{0}
}

func (x *LoggerConfig) GetFileWriterConfig() *LoggerConfig_FileWriterConfig {
	if x != nil {
		return x.FileWriterConfig
	}
	return nil
}

func (x *LoggerConfig) GetConsoleWriterConfig() *LoggerConfig_ConsoleWriterConfig {
	if x != nil {
		return x.ConsoleWriterConfig
	}
	return nil
}

type LoggerConfig_FileWriterConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 是否将日志写入文件
	Enable *bool `protobuf:"varint,1,opt,name=enable,def=0" json:"enable,omitempty"`
	// 日志级别, 只写入 >= LogLevel 的日志
	LogLevel *LoggerConfig_LogLevel `protobuf:"varint,2,opt,name=log_level,json=logLevel,enum=ums.config.proto.LoggerConfig_LogLevel,def=1" json:"log_level,omitempty"`
	// Info 日志存放路径
	InfoLogPath *string `protobuf:"bytes,3,opt,name=info_log_path,json=infoLogPath,def=logs/log.info" json:"info_log_path,omitempty"`
	// Wf 日志存放路径 (大于 Warning 级别的日志)
	WfLogPath *string `protobuf:"bytes,4,opt,name=wf_log_path,json=wfLogPath,def=logs/log.wf" json:"wf_log_path,omitempty"`
	// 保存小时数, -1 表示不进行日志切割
	RetainHours *int32 `protobuf:"varint,5,opt,name=retain_hours,json=retainHours,def=48" json:"retain_hours,omitempty"`
}

// Default values for LoggerConfig_FileWriterConfig fields.
const (
	Default_LoggerConfig_FileWriterConfig_Enable      = bool(false)
	Default_LoggerConfig_FileWriterConfig_LogLevel    = LoggerConfig_LOG_LEVEL_DEBUG
	Default_LoggerConfig_FileWriterConfig_InfoLogPath = string("logs/log.info")
	Default_LoggerConfig_FileWriterConfig_WfLogPath   = string("logs/log.wf")
	Default_LoggerConfig_FileWriterConfig_RetainHours = int32(48)
)

func (x *LoggerConfig_FileWriterConfig) Reset() {
	*x = LoggerConfig_FileWriterConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_config_logger_config_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoggerConfig_FileWriterConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoggerConfig_FileWriterConfig) ProtoMessage() {}

func (x *LoggerConfig_FileWriterConfig) ProtoReflect() protoreflect.Message {
	mi := &file_proto_config_logger_config_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoggerConfig_FileWriterConfig.ProtoReflect.Descriptor instead.
func (*LoggerConfig_FileWriterConfig) Descriptor() ([]byte, []int) {
	return file_proto_config_logger_config_proto_rawDescGZIP(), []int{0, 0}
}

func (x *LoggerConfig_FileWriterConfig) GetEnable() bool {
	if x != nil && x.Enable != nil {
		return *x.Enable
	}
	return Default_LoggerConfig_FileWriterConfig_Enable
}

func (x *LoggerConfig_FileWriterConfig) GetLogLevel() LoggerConfig_LogLevel {
	if x != nil && x.LogLevel != nil {
		return *x.LogLevel
	}
	return Default_LoggerConfig_FileWriterConfig_LogLevel
}

func (x *LoggerConfig_FileWriterConfig) GetInfoLogPath() string {
	if x != nil && x.InfoLogPath != nil {
		return *x.InfoLogPath
	}
	return Default_LoggerConfig_FileWriterConfig_InfoLogPath
}

func (x *LoggerConfig_FileWriterConfig) GetWfLogPath() string {
	if x != nil && x.WfLogPath != nil {
		return *x.WfLogPath
	}
	return Default_LoggerConfig_FileWriterConfig_WfLogPath
}

func (x *LoggerConfig_FileWriterConfig) GetRetainHours() int32 {
	if x != nil && x.RetainHours != nil {
		return *x.RetainHours
	}
	return Default_LoggerConfig_FileWriterConfig_RetainHours
}

type LoggerConfig_ConsoleWriterConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 是否将日志输出到控制台
	Enable *bool `protobuf:"varint,1,opt,name=enable,def=1" json:"enable,omitempty"`
	// 日志级别, 只写入 >= LogLevel 的日志
	LogLevel *LoggerConfig_LogLevel `protobuf:"varint,2,opt,name=log_level,json=logLevel,enum=ums.config.proto.LoggerConfig_LogLevel,def=2" json:"log_level,omitempty"`
	// 是否开启彩色输出
	EnableColor *bool `protobuf:"varint,3,opt,name=enable_color,json=enableColor,def=1" json:"enable_color,omitempty"`
}

// Default values for LoggerConfig_ConsoleWriterConfig fields.
const (
	Default_LoggerConfig_ConsoleWriterConfig_Enable      = bool(true)
	Default_LoggerConfig_ConsoleWriterConfig_LogLevel    = LoggerConfig_LOG_LEVEL_INFO
	Default_LoggerConfig_ConsoleWriterConfig_EnableColor = bool(true)
)

func (x *LoggerConfig_ConsoleWriterConfig) Reset() {
	*x = LoggerConfig_ConsoleWriterConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_config_logger_config_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoggerConfig_ConsoleWriterConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoggerConfig_ConsoleWriterConfig) ProtoMessage() {}

func (x *LoggerConfig_ConsoleWriterConfig) ProtoReflect() protoreflect.Message {
	mi := &file_proto_config_logger_config_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoggerConfig_ConsoleWriterConfig.ProtoReflect.Descriptor instead.
func (*LoggerConfig_ConsoleWriterConfig) Descriptor() ([]byte, []int) {
	return file_proto_config_logger_config_proto_rawDescGZIP(), []int{0, 1}
}

func (x *LoggerConfig_ConsoleWriterConfig) GetEnable() bool {
	if x != nil && x.Enable != nil {
		return *x.Enable
	}
	return Default_LoggerConfig_ConsoleWriterConfig_Enable
}

func (x *LoggerConfig_ConsoleWriterConfig) GetLogLevel() LoggerConfig_LogLevel {
	if x != nil && x.LogLevel != nil {
		return *x.LogLevel
	}
	return Default_LoggerConfig_ConsoleWriterConfig_LogLevel
}

func (x *LoggerConfig_ConsoleWriterConfig) GetEnableColor() bool {
	if x != nil && x.EnableColor != nil {
		return *x.EnableColor
	}
	return Default_LoggerConfig_ConsoleWriterConfig_EnableColor
}

var File_proto_config_logger_config_proto protoreflect.FileDescriptor

var file_proto_config_logger_config_proto_rawDesc = []byte{
	0x0a, 0x20, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x6c,
	0x6f, 0x67, 0x67, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x10, 0x75, 0x6d, 0x73, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x92, 0x06, 0x0a, 0x0c, 0x4c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x5d, 0x0a, 0x12, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x77, 0x72,
	0x69, 0x74, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x2f, 0x2e, 0x75, 0x6d, 0x73, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x57, 0x72, 0x69, 0x74, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x52, 0x10, 0x66, 0x69, 0x6c, 0x65, 0x57, 0x72, 0x69, 0x74, 0x65, 0x72, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x12, 0x66, 0x0a, 0x15, 0x63, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x5f,
	0x77, 0x72, 0x69, 0x74, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x75, 0x6d, 0x73, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x57, 0x72, 0x69, 0x74, 0x65,
	0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x13, 0x63, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65,
	0x57, 0x72, 0x69, 0x74, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x1a, 0x8f, 0x02, 0x0a,
	0x10, 0x46, 0x69, 0x6c, 0x65, 0x57, 0x72, 0x69, 0x74, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x12, 0x1d, 0x0a, 0x06, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x3a, 0x05, 0x66, 0x61, 0x6c, 0x73, 0x65, 0x52, 0x06, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65,
	0x12, 0x55, 0x0a, 0x09, 0x6c, 0x6f, 0x67, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x27, 0x2e, 0x75, 0x6d, 0x73, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x2e, 0x4c, 0x6f, 0x67, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x3a, 0x0f, 0x4c, 0x4f,
	0x47, 0x5f, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f, 0x44, 0x45, 0x42, 0x55, 0x47, 0x52, 0x08, 0x6c,
	0x6f, 0x67, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x31, 0x0a, 0x0d, 0x69, 0x6e, 0x66, 0x6f, 0x5f,
	0x6c, 0x6f, 0x67, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x3a, 0x0d,
	0x6c, 0x6f, 0x67, 0x73, 0x2f, 0x6c, 0x6f, 0x67, 0x2e, 0x69, 0x6e, 0x66, 0x6f, 0x52, 0x0b, 0x69,
	0x6e, 0x66, 0x6f, 0x4c, 0x6f, 0x67, 0x50, 0x61, 0x74, 0x68, 0x12, 0x2b, 0x0a, 0x0b, 0x77, 0x66,
	0x5f, 0x6c, 0x6f, 0x67, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x3a,
	0x0b, 0x6c, 0x6f, 0x67, 0x73, 0x2f, 0x6c, 0x6f, 0x67, 0x2e, 0x77, 0x66, 0x52, 0x09, 0x77, 0x66,
	0x4c, 0x6f, 0x67, 0x50, 0x61, 0x74, 0x68, 0x12, 0x25, 0x0a, 0x0c, 0x72, 0x65, 0x74, 0x61, 0x69,
	0x6e, 0x5f, 0x68, 0x6f, 0x75, 0x72, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x3a, 0x02, 0x34,
	0x38, 0x52, 0x0b, 0x72, 0x65, 0x74, 0x61, 0x69, 0x6e, 0x48, 0x6f, 0x75, 0x72, 0x73, 0x1a, 0xb2,
	0x01, 0x0a, 0x13, 0x43, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x57, 0x72, 0x69, 0x74, 0x65, 0x72,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x1c, 0x0a, 0x06, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x3a, 0x04, 0x74, 0x72, 0x75, 0x65, 0x52, 0x06, 0x65, 0x6e,
	0x61, 0x62, 0x6c, 0x65, 0x12, 0x54, 0x0a, 0x09, 0x6c, 0x6f, 0x67, 0x5f, 0x6c, 0x65, 0x76, 0x65,
	0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x27, 0x2e, 0x75, 0x6d, 0x73, 0x2e, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x6f, 0x67, 0x67, 0x65,
	0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x4c, 0x6f, 0x67, 0x4c, 0x65, 0x76, 0x65, 0x6c,
	0x3a, 0x0e, 0x4c, 0x4f, 0x47, 0x5f, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f, 0x49, 0x4e, 0x46, 0x4f,
	0x52, 0x08, 0x6c, 0x6f, 0x67, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x27, 0x0a, 0x0c, 0x65, 0x6e,
	0x61, 0x62, 0x6c, 0x65, 0x5f, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08,
	0x3a, 0x04, 0x74, 0x72, 0x75, 0x65, 0x52, 0x0b, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x43, 0x6f,
	0x6c, 0x6f, 0x72, 0x22, 0x74, 0x0a, 0x08, 0x4c, 0x6f, 0x67, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12,
	0x13, 0x0a, 0x0f, 0x4c, 0x4f, 0x47, 0x5f, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f, 0x44, 0x45, 0x42,
	0x55, 0x47, 0x10, 0x01, 0x12, 0x12, 0x0a, 0x0e, 0x4c, 0x4f, 0x47, 0x5f, 0x4c, 0x45, 0x56, 0x45,
	0x4c, 0x5f, 0x49, 0x4e, 0x46, 0x4f, 0x10, 0x02, 0x12, 0x15, 0x0a, 0x11, 0x4c, 0x4f, 0x47, 0x5f,
	0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f, 0x57, 0x41, 0x52, 0x4e, 0x49, 0x4e, 0x47, 0x10, 0x03, 0x12,
	0x13, 0x0a, 0x0f, 0x4c, 0x4f, 0x47, 0x5f, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f, 0x45, 0x52, 0x52,
	0x4f, 0x52, 0x10, 0x04, 0x12, 0x13, 0x0a, 0x0f, 0x4c, 0x4f, 0x47, 0x5f, 0x4c, 0x45, 0x56, 0x45,
	0x4c, 0x5f, 0x46, 0x41, 0x54, 0x41, 0x4c, 0x10, 0x05, 0x42, 0x0e, 0x5a, 0x0c, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
}

var (
	file_proto_config_logger_config_proto_rawDescOnce sync.Once
	file_proto_config_logger_config_proto_rawDescData = file_proto_config_logger_config_proto_rawDesc
)

func file_proto_config_logger_config_proto_rawDescGZIP() []byte {
	file_proto_config_logger_config_proto_rawDescOnce.Do(func() {
		file_proto_config_logger_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_config_logger_config_proto_rawDescData)
	})
	return file_proto_config_logger_config_proto_rawDescData
}

var file_proto_config_logger_config_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_config_logger_config_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_config_logger_config_proto_goTypes = []interface{}{
	(LoggerConfig_LogLevel)(0),               // 0: ums.config.proto.LoggerConfig.LogLevel
	(*LoggerConfig)(nil),                     // 1: ums.config.proto.LoggerConfig
	(*LoggerConfig_FileWriterConfig)(nil),    // 2: ums.config.proto.LoggerConfig.FileWriterConfig
	(*LoggerConfig_ConsoleWriterConfig)(nil), // 3: ums.config.proto.LoggerConfig.ConsoleWriterConfig
}
var file_proto_config_logger_config_proto_depIdxs = []int32{
	2, // 0: ums.config.proto.LoggerConfig.file_writer_config:type_name -> ums.config.proto.LoggerConfig.FileWriterConfig
	3, // 1: ums.config.proto.LoggerConfig.console_writer_config:type_name -> ums.config.proto.LoggerConfig.ConsoleWriterConfig
	0, // 2: ums.config.proto.LoggerConfig.FileWriterConfig.log_level:type_name -> ums.config.proto.LoggerConfig.LogLevel
	0, // 3: ums.config.proto.LoggerConfig.ConsoleWriterConfig.log_level:type_name -> ums.config.proto.LoggerConfig.LogLevel
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_proto_config_logger_config_proto_init() }
func file_proto_config_logger_config_proto_init() {
	if File_proto_config_logger_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_config_logger_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoggerConfig); i {
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
		file_proto_config_logger_config_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoggerConfig_FileWriterConfig); i {
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
		file_proto_config_logger_config_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoggerConfig_ConsoleWriterConfig); i {
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
			RawDescriptor: file_proto_config_logger_config_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_config_logger_config_proto_goTypes,
		DependencyIndexes: file_proto_config_logger_config_proto_depIdxs,
		EnumInfos:         file_proto_config_logger_config_proto_enumTypes,
		MessageInfos:      file_proto_config_logger_config_proto_msgTypes,
	}.Build()
	File_proto_config_logger_config_proto = out.File
	file_proto_config_logger_config_proto_rawDesc = nil
	file_proto_config_logger_config_proto_goTypes = nil
	file_proto_config_logger_config_proto_depIdxs = nil
}
