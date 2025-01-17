// go-pst is a library for reading Personal Storage Table (.pst) files (written in Go/Golang).
//
// Copyright (C) 2022  Marten Mooij
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

//go:generate msgp -tests=false

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: note.proto

package properties

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

type Note struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Specifies the suggested background color of the Note object.
	NoteColor *int32 `protobuf:"varint,1,opt,name=note_color,json=noteColor,proto3,oneof" json:"note_color,omitempty" msg:"2734083"` // @gotags: msg:"2734083"
	// Specifies the height of the visible message window in pixels.
	NoteHeight *int32 `protobuf:"varint,2,opt,name=note_height,json=noteHeight,proto3,oneof" json:"note_height,omitempty" msg:"2734113"` // @gotags: msg:"2734113"
	// Specifies the width of the visible message window in pixels.
	NoteWidth *int32 `protobuf:"varint,3,opt,name=note_width,json=noteWidth,proto3,oneof" json:"note_width,omitempty" msg:"2734103"` // @gotags: msg:"2734103"
	// Specifies the distance, in pixels, from the left edge of the screen that a user interface displays a Note object.
	NoteX *int32 `protobuf:"varint,4,opt,name=note_x,json=noteX,proto3,oneof" json:"note_x,omitempty" msg:"2734123"` // @gotags: msg:"2734123"
	// Specifies the distance, in pixels, from the top edge of the screen that a user interface displays a Note object.
	NoteY *int32 `protobuf:"varint,5,opt,name=note_y,json=noteY,proto3,oneof" json:"note_y,omitempty" msg:"2734133"` // @gotags: msg:"2734133"
}

func (x *Note) Reset() {
	*x = Note{}
	if protoimpl.UnsafeEnabled {
		mi := &file_note_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Note) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Note) ProtoMessage() {}

func (x *Note) ProtoReflect() protoreflect.Message {
	mi := &file_note_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Note.ProtoReflect.Descriptor instead.
func (*Note) Descriptor() ([]byte, []int) {
	return file_note_proto_rawDescGZIP(), []int{0}
}

func (x *Note) GetNoteColor() int32 {
	if x != nil && x.NoteColor != nil {
		return *x.NoteColor
	}
	return 0
}

func (x *Note) GetNoteHeight() int32 {
	if x != nil && x.NoteHeight != nil {
		return *x.NoteHeight
	}
	return 0
}

func (x *Note) GetNoteWidth() int32 {
	if x != nil && x.NoteWidth != nil {
		return *x.NoteWidth
	}
	return 0
}

func (x *Note) GetNoteX() int32 {
	if x != nil && x.NoteX != nil {
		return *x.NoteX
	}
	return 0
}

func (x *Note) GetNoteY() int32 {
	if x != nil && x.NoteY != nil {
		return *x.NoteY
	}
	return 0
}

var File_note_proto protoreflect.FileDescriptor

var file_note_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x6e, 0x6f, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf0, 0x01, 0x0a,
	0x04, 0x4e, 0x6f, 0x74, 0x65, 0x12, 0x22, 0x0a, 0x0a, 0x6e, 0x6f, 0x74, 0x65, 0x5f, 0x63, 0x6f,
	0x6c, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x48, 0x00, 0x52, 0x09, 0x6e, 0x6f, 0x74,
	0x65, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x88, 0x01, 0x01, 0x12, 0x24, 0x0a, 0x0b, 0x6e, 0x6f, 0x74,
	0x65, 0x5f, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x48, 0x01,
	0x52, 0x0a, 0x6e, 0x6f, 0x74, 0x65, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x88, 0x01, 0x01, 0x12,
	0x22, 0x0a, 0x0a, 0x6e, 0x6f, 0x74, 0x65, 0x5f, 0x77, 0x69, 0x64, 0x74, 0x68, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x48, 0x02, 0x52, 0x09, 0x6e, 0x6f, 0x74, 0x65, 0x57, 0x69, 0x64, 0x74, 0x68,
	0x88, 0x01, 0x01, 0x12, 0x1a, 0x0a, 0x06, 0x6e, 0x6f, 0x74, 0x65, 0x5f, 0x78, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x05, 0x48, 0x03, 0x52, 0x05, 0x6e, 0x6f, 0x74, 0x65, 0x58, 0x88, 0x01, 0x01, 0x12,
	0x1a, 0x0a, 0x06, 0x6e, 0x6f, 0x74, 0x65, 0x5f, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x48,
	0x04, 0x52, 0x05, 0x6e, 0x6f, 0x74, 0x65, 0x59, 0x88, 0x01, 0x01, 0x42, 0x0d, 0x0a, 0x0b, 0x5f,
	0x6e, 0x6f, 0x74, 0x65, 0x5f, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x6e,
	0x6f, 0x74, 0x65, 0x5f, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x6e,
	0x6f, 0x74, 0x65, 0x5f, 0x77, 0x69, 0x64, 0x74, 0x68, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x6e, 0x6f,
	0x74, 0x65, 0x5f, 0x78, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x6e, 0x6f, 0x74, 0x65, 0x5f, 0x79, 0x42,
	0x28, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x6f,
	0x6f, 0x69, 0x6a, 0x74, 0x65, 0x63, 0x68, 0x2f, 0x67, 0x6f, 0x2d, 0x70, 0x73, 0x74, 0x3b, 0x70,
	0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_note_proto_rawDescOnce sync.Once
	file_note_proto_rawDescData = file_note_proto_rawDesc
)

func file_note_proto_rawDescGZIP() []byte {
	file_note_proto_rawDescOnce.Do(func() {
		file_note_proto_rawDescData = protoimpl.X.CompressGZIP(file_note_proto_rawDescData)
	})
	return file_note_proto_rawDescData
}

var file_note_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_note_proto_goTypes = []interface{}{
	(*Note)(nil), // 0: Note
}
var file_note_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_note_proto_init() }
func file_note_proto_init() {
	if File_note_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_note_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Note); i {
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
	file_note_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_note_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_note_proto_goTypes,
		DependencyIndexes: file_note_proto_depIdxs,
		MessageInfos:      file_note_proto_msgTypes,
	}.Build()
	File_note_proto = out.File
	file_note_proto_rawDesc = nil
	file_note_proto_goTypes = nil
	file_note_proto_depIdxs = nil
}
