// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: v1/grpc/booking_service.proto

package grpc

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

type BookingQuery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Movie        string `protobuf:"bytes,2,opt,name=movie,proto3" json:"movie,omitempty"`
	BookingName  string `protobuf:"bytes,3,opt,name=booking_name,json=bookingName,proto3" json:"booking_name,omitempty"`
	BookingEmail string `protobuf:"bytes,4,opt,name=booking_email,json=bookingEmail,proto3" json:"booking_email,omitempty"`
	Location     string `protobuf:"bytes,5,opt,name=location,proto3" json:"location,omitempty"`
	Date         string `protobuf:"bytes,6,opt,name=date,proto3" json:"date,omitempty"`
}

func (x *BookingQuery) Reset() {
	*x = BookingQuery{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_grpc_booking_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BookingQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BookingQuery) ProtoMessage() {}

func (x *BookingQuery) ProtoReflect() protoreflect.Message {
	mi := &file_v1_grpc_booking_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BookingQuery.ProtoReflect.Descriptor instead.
func (*BookingQuery) Descriptor() ([]byte, []int) {
	return file_v1_grpc_booking_service_proto_rawDescGZIP(), []int{0}
}

func (x *BookingQuery) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *BookingQuery) GetMovie() string {
	if x != nil {
		return x.Movie
	}
	return ""
}

func (x *BookingQuery) GetBookingName() string {
	if x != nil {
		return x.BookingName
	}
	return ""
}

func (x *BookingQuery) GetBookingEmail() string {
	if x != nil {
		return x.BookingEmail
	}
	return ""
}

func (x *BookingQuery) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *BookingQuery) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

type BookingDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Movie        string `protobuf:"bytes,2,opt,name=movie,proto3" json:"movie,omitempty"`
	BookingName  string `protobuf:"bytes,3,opt,name=booking_name,json=bookingName,proto3" json:"booking_name,omitempty"`
	BookingEmail string `protobuf:"bytes,4,opt,name=booking_email,json=bookingEmail,proto3" json:"booking_email,omitempty"`
	Date         string `protobuf:"bytes,5,opt,name=date,proto3" json:"date,omitempty"`
	Location     string `protobuf:"bytes,6,opt,name=location,proto3" json:"location,omitempty"`
	Screen       string `protobuf:"bytes,7,opt,name=screen,proto3" json:"screen,omitempty"`
	SeatNumber   int32  `protobuf:"varint,8,opt,name=seat_number,json=seatNumber,proto3" json:"seat_number,omitempty"`
	Status       string `protobuf:"bytes,9,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *BookingDetails) Reset() {
	*x = BookingDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_grpc_booking_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BookingDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BookingDetails) ProtoMessage() {}

func (x *BookingDetails) ProtoReflect() protoreflect.Message {
	mi := &file_v1_grpc_booking_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BookingDetails.ProtoReflect.Descriptor instead.
func (*BookingDetails) Descriptor() ([]byte, []int) {
	return file_v1_grpc_booking_service_proto_rawDescGZIP(), []int{1}
}

func (x *BookingDetails) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *BookingDetails) GetMovie() string {
	if x != nil {
		return x.Movie
	}
	return ""
}

func (x *BookingDetails) GetBookingName() string {
	if x != nil {
		return x.BookingName
	}
	return ""
}

func (x *BookingDetails) GetBookingEmail() string {
	if x != nil {
		return x.BookingEmail
	}
	return ""
}

func (x *BookingDetails) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

func (x *BookingDetails) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *BookingDetails) GetScreen() string {
	if x != nil {
		return x.Screen
	}
	return ""
}

func (x *BookingDetails) GetSeatNumber() int32 {
	if x != nil {
		return x.SeatNumber
	}
	return 0
}

func (x *BookingDetails) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

var File_v1_grpc_booking_service_proto protoreflect.FileDescriptor

var file_v1_grpc_booking_service_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x76, 0x31, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e,
	0x67, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x22, 0xac, 0x01, 0x0a, 0x0c, 0x42, 0x6f, 0x6f,
	0x6b, 0x69, 0x6e, 0x67, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x6f, 0x76,
	0x69, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x12,
	0x21, 0x0a, 0x0c, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x62, 0x6f, 0x6f, 0x6b, 0x69,
	0x6e, 0x67, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x22, 0xff, 0x01, 0x0a, 0x0e, 0x42, 0x6f, 0x6f, 0x6b,
	0x69, 0x6e, 0x67, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x6f,
	0x76, 0x69, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6d, 0x6f, 0x76, 0x69, 0x65,
	0x12, 0x21, 0x0a, 0x0c, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x62, 0x6f, 0x6f, 0x6b,
	0x69, 0x6e, 0x67, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x63, 0x72, 0x65,
	0x65, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x63, 0x72, 0x65, 0x65, 0x6e,
	0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x65, 0x61, 0x74, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x73, 0x65, 0x61, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x32, 0x86, 0x03, 0x0a, 0x13, 0x4d, 0x6f,
	0x76, 0x69, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x3f, 0x0a, 0x0b, 0x6d, 0x61, 0x6b, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67,
	0x12, 0x15, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x69,
	0x6e, 0x67, 0x51, 0x75, 0x65, 0x72, 0x79, 0x1a, 0x17, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e,
	0x67, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73,
	0x22, 0x00, 0x12, 0x45, 0x0a, 0x11, 0x67, 0x65, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67,
	0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x15, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e,
	0x67, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x51, 0x75, 0x65, 0x72, 0x79, 0x1a, 0x17,
	0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67,
	0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x22, 0x00, 0x12, 0x4e, 0x0a, 0x18, 0x67, 0x65, 0x74,
	0x41, 0x6c, 0x6c, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x73, 0x42, 0x79, 0x4c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x15, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e,
	0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x51, 0x75, 0x65, 0x72, 0x79, 0x1a, 0x17, 0x2e, 0x62,
	0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x73, 0x22, 0x00, 0x30, 0x01, 0x12, 0x4a, 0x0a, 0x14, 0x67, 0x65, 0x74,
	0x41, 0x6c, 0x6c, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x73, 0x42, 0x79, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x15, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x42, 0x6f, 0x6f, 0x6b,
	0x69, 0x6e, 0x67, 0x51, 0x75, 0x65, 0x72, 0x79, 0x1a, 0x17, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69,
	0x6e, 0x67, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x73, 0x22, 0x00, 0x30, 0x01, 0x12, 0x4b, 0x0a, 0x15, 0x67, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x42,
	0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x73, 0x42, 0x79, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x12, 0x15,
	0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x1a, 0x17, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e,
	0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x22, 0x00,
	0x30, 0x01, 0x42, 0x7b, 0x0a, 0x23, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x62, 0x61, 0x69, 0x74, 0x6d,
	0x61, 0x6e, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e,
	0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x42, 0x11, 0x4d, 0x6f, 0x76, 0x69, 0x65,
	0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x39,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x62, 0x61, 0x69, 0x74,
	0x6d, 0x61, 0x6e, 0x67, 0x61, 0x6c, 0x6b, 0x61, 0x72, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e,
	0x67, 0x2d, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x61, 0x70,
	0x69, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x72, 0x70, 0x63, 0xa2, 0x02, 0x03, 0x52, 0x54, 0x47, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_grpc_booking_service_proto_rawDescOnce sync.Once
	file_v1_grpc_booking_service_proto_rawDescData = file_v1_grpc_booking_service_proto_rawDesc
)

func file_v1_grpc_booking_service_proto_rawDescGZIP() []byte {
	file_v1_grpc_booking_service_proto_rawDescOnce.Do(func() {
		file_v1_grpc_booking_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_grpc_booking_service_proto_rawDescData)
	})
	return file_v1_grpc_booking_service_proto_rawDescData
}

var file_v1_grpc_booking_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_v1_grpc_booking_service_proto_goTypes = []interface{}{
	(*BookingQuery)(nil),   // 0: booking.BookingQuery
	(*BookingDetails)(nil), // 1: booking.BookingDetails
}
var file_v1_grpc_booking_service_proto_depIdxs = []int32{
	0, // 0: booking.MovieBookingService.makeBooking:input_type -> booking.BookingQuery
	0, // 1: booking.MovieBookingService.getBookingDetails:input_type -> booking.BookingQuery
	0, // 2: booking.MovieBookingService.getAllBookingsByLocation:input_type -> booking.BookingQuery
	0, // 3: booking.MovieBookingService.getAllBookingsByName:input_type -> booking.BookingQuery
	0, // 4: booking.MovieBookingService.getAllBookingsByMovie:input_type -> booking.BookingQuery
	1, // 5: booking.MovieBookingService.makeBooking:output_type -> booking.BookingDetails
	1, // 6: booking.MovieBookingService.getBookingDetails:output_type -> booking.BookingDetails
	1, // 7: booking.MovieBookingService.getAllBookingsByLocation:output_type -> booking.BookingDetails
	1, // 8: booking.MovieBookingService.getAllBookingsByName:output_type -> booking.BookingDetails
	1, // 9: booking.MovieBookingService.getAllBookingsByMovie:output_type -> booking.BookingDetails
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_v1_grpc_booking_service_proto_init() }
func file_v1_grpc_booking_service_proto_init() {
	if File_v1_grpc_booking_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v1_grpc_booking_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BookingQuery); i {
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
		file_v1_grpc_booking_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BookingDetails); i {
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
			RawDescriptor: file_v1_grpc_booking_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_v1_grpc_booking_service_proto_goTypes,
		DependencyIndexes: file_v1_grpc_booking_service_proto_depIdxs,
		MessageInfos:      file_v1_grpc_booking_service_proto_msgTypes,
	}.Build()
	File_v1_grpc_booking_service_proto = out.File
	file_v1_grpc_booking_service_proto_rawDesc = nil
	file_v1_grpc_booking_service_proto_goTypes = nil
	file_v1_grpc_booking_service_proto_depIdxs = nil
}
