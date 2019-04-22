// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: v1/RoomAdmin.proto

package v1

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type RoomAdminIsRoomAdminReq struct {
	// 用户id
	Uid string `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid"`
	// 主播的用户 roomid
	Roomid int64 `protobuf:"varint,2,opt,name=roomid,proto3" json:"roomid"`
	// 主播的用户 anchor_id
	AnchorId int64 `protobuf:"varint,3,opt,name=anchor_id,json=anchorId,proto3" json:"anchor_id"`
}

func (m *RoomAdminIsRoomAdminReq) Reset()         { *m = RoomAdminIsRoomAdminReq{} }
func (m *RoomAdminIsRoomAdminReq) String() string { return proto.CompactTextString(m) }
func (*RoomAdminIsRoomAdminReq) ProtoMessage()    {}
func (*RoomAdminIsRoomAdminReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_RoomAdmin_00e9d4c4f349d463, []int{0}
}
func (m *RoomAdminIsRoomAdminReq) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RoomAdminIsRoomAdminReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RoomAdminIsRoomAdminReq.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *RoomAdminIsRoomAdminReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoomAdminIsRoomAdminReq.Merge(dst, src)
}
func (m *RoomAdminIsRoomAdminReq) XXX_Size() int {
	return m.Size()
}
func (m *RoomAdminIsRoomAdminReq) XXX_DiscardUnknown() {
	xxx_messageInfo_RoomAdminIsRoomAdminReq.DiscardUnknown(m)
}

var xxx_messageInfo_RoomAdminIsRoomAdminReq proto.InternalMessageInfo

func (m *RoomAdminIsRoomAdminReq) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *RoomAdminIsRoomAdminReq) GetRoomid() int64 {
	if m != nil {
		return m.Roomid
	}
	return 0
}

func (m *RoomAdminIsRoomAdminReq) GetAnchorId() int64 {
	if m != nil {
		return m.AnchorId
	}
	return 0
}

type RoomAdminIsRoomAdminResp struct {
	// code
	Code int64 `protobuf:"varint,1,opt,name=code,proto3" json:"code"`
	// msg
	Msg string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg"`
	//
	Data *RoomAdminIsRoomAdminResp_Data `protobuf:"bytes,3,opt,name=data" json:"data"`
}

func (m *RoomAdminIsRoomAdminResp) Reset()         { *m = RoomAdminIsRoomAdminResp{} }
func (m *RoomAdminIsRoomAdminResp) String() string { return proto.CompactTextString(m) }
func (*RoomAdminIsRoomAdminResp) ProtoMessage()    {}
func (*RoomAdminIsRoomAdminResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_RoomAdmin_00e9d4c4f349d463, []int{1}
}
func (m *RoomAdminIsRoomAdminResp) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RoomAdminIsRoomAdminResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RoomAdminIsRoomAdminResp.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *RoomAdminIsRoomAdminResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoomAdminIsRoomAdminResp.Merge(dst, src)
}
func (m *RoomAdminIsRoomAdminResp) XXX_Size() int {
	return m.Size()
}
func (m *RoomAdminIsRoomAdminResp) XXX_DiscardUnknown() {
	xxx_messageInfo_RoomAdminIsRoomAdminResp.DiscardUnknown(m)
}

var xxx_messageInfo_RoomAdminIsRoomAdminResp proto.InternalMessageInfo

func (m *RoomAdminIsRoomAdminResp) GetCode() int64 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *RoomAdminIsRoomAdminResp) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *RoomAdminIsRoomAdminResp) GetData() *RoomAdminIsRoomAdminResp_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

type RoomAdminIsRoomAdminResp_UI struct {
	// 用户id
	Uid int64 `protobuf:"varint,1,opt,name=uid,proto3" json:"uid"`
	// 用户名
	Uname string `protobuf:"bytes,2,opt,name=uname,proto3" json:"uname"`
}

func (m *RoomAdminIsRoomAdminResp_UI) Reset()         { *m = RoomAdminIsRoomAdminResp_UI{} }
func (m *RoomAdminIsRoomAdminResp_UI) String() string { return proto.CompactTextString(m) }
func (*RoomAdminIsRoomAdminResp_UI) ProtoMessage()    {}
func (*RoomAdminIsRoomAdminResp_UI) Descriptor() ([]byte, []int) {
	return fileDescriptor_RoomAdmin_00e9d4c4f349d463, []int{1, 0}
}
func (m *RoomAdminIsRoomAdminResp_UI) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RoomAdminIsRoomAdminResp_UI) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RoomAdminIsRoomAdminResp_UI.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *RoomAdminIsRoomAdminResp_UI) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoomAdminIsRoomAdminResp_UI.Merge(dst, src)
}
func (m *RoomAdminIsRoomAdminResp_UI) XXX_Size() int {
	return m.Size()
}
func (m *RoomAdminIsRoomAdminResp_UI) XXX_DiscardUnknown() {
	xxx_messageInfo_RoomAdminIsRoomAdminResp_UI.DiscardUnknown(m)
}

var xxx_messageInfo_RoomAdminIsRoomAdminResp_UI proto.InternalMessageInfo

func (m *RoomAdminIsRoomAdminResp_UI) GetUid() int64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *RoomAdminIsRoomAdminResp_UI) GetUname() string {
	if m != nil {
		return m.Uname
	}
	return ""
}

type RoomAdminIsRoomAdminResp_Data struct {
	// banner
	Userinfo *RoomAdminIsRoomAdminResp_UI `protobuf:"bytes,1,opt,name=userinfo" json:"userinfo"`
	// 房管的用户id
	Uid int64 `protobuf:"varint,2,opt,name=uid,proto3" json:"uid"`
	// 房间号
	Roomid int64 `protobuf:"varint,3,opt,name=roomid,proto3" json:"roomid"`
	// 创建时间　"2017-07-26 17:12:51"
	Ctime string `protobuf:"bytes,4,opt,name=ctime,proto3" json:"ctime"`
}

func (m *RoomAdminIsRoomAdminResp_Data) Reset()         { *m = RoomAdminIsRoomAdminResp_Data{} }
func (m *RoomAdminIsRoomAdminResp_Data) String() string { return proto.CompactTextString(m) }
func (*RoomAdminIsRoomAdminResp_Data) ProtoMessage()    {}
func (*RoomAdminIsRoomAdminResp_Data) Descriptor() ([]byte, []int) {
	return fileDescriptor_RoomAdmin_00e9d4c4f349d463, []int{1, 1}
}
func (m *RoomAdminIsRoomAdminResp_Data) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RoomAdminIsRoomAdminResp_Data) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RoomAdminIsRoomAdminResp_Data.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *RoomAdminIsRoomAdminResp_Data) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoomAdminIsRoomAdminResp_Data.Merge(dst, src)
}
func (m *RoomAdminIsRoomAdminResp_Data) XXX_Size() int {
	return m.Size()
}
func (m *RoomAdminIsRoomAdminResp_Data) XXX_DiscardUnknown() {
	xxx_messageInfo_RoomAdminIsRoomAdminResp_Data.DiscardUnknown(m)
}

var xxx_messageInfo_RoomAdminIsRoomAdminResp_Data proto.InternalMessageInfo

func (m *RoomAdminIsRoomAdminResp_Data) GetUserinfo() *RoomAdminIsRoomAdminResp_UI {
	if m != nil {
		return m.Userinfo
	}
	return nil
}

func (m *RoomAdminIsRoomAdminResp_Data) GetUid() int64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *RoomAdminIsRoomAdminResp_Data) GetRoomid() int64 {
	if m != nil {
		return m.Roomid
	}
	return 0
}

func (m *RoomAdminIsRoomAdminResp_Data) GetCtime() string {
	if m != nil {
		return m.Ctime
	}
	return ""
}

func init() {
	proto.RegisterType((*RoomAdminIsRoomAdminReq)(nil), "live_user.v1.RoomAdminIsRoomAdminReq")
	proto.RegisterType((*RoomAdminIsRoomAdminResp)(nil), "live_user.v1.RoomAdminIsRoomAdminResp")
	proto.RegisterType((*RoomAdminIsRoomAdminResp_UI)(nil), "live_user.v1.RoomAdminIsRoomAdminResp.UI")
	proto.RegisterType((*RoomAdminIsRoomAdminResp_Data)(nil), "live_user.v1.RoomAdminIsRoomAdminResp.Data")
}
func (m *RoomAdminIsRoomAdminReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RoomAdminIsRoomAdminReq) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Uid) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintRoomAdmin(dAtA, i, uint64(len(m.Uid)))
		i += copy(dAtA[i:], m.Uid)
	}
	if m.Roomid != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintRoomAdmin(dAtA, i, uint64(m.Roomid))
	}
	if m.AnchorId != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintRoomAdmin(dAtA, i, uint64(m.AnchorId))
	}
	return i, nil
}

func (m *RoomAdminIsRoomAdminResp) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RoomAdminIsRoomAdminResp) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Code != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintRoomAdmin(dAtA, i, uint64(m.Code))
	}
	if len(m.Msg) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintRoomAdmin(dAtA, i, uint64(len(m.Msg)))
		i += copy(dAtA[i:], m.Msg)
	}
	if m.Data != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintRoomAdmin(dAtA, i, uint64(m.Data.Size()))
		n1, err := m.Data.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	return i, nil
}

func (m *RoomAdminIsRoomAdminResp_UI) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RoomAdminIsRoomAdminResp_UI) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Uid != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintRoomAdmin(dAtA, i, uint64(m.Uid))
	}
	if len(m.Uname) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintRoomAdmin(dAtA, i, uint64(len(m.Uname)))
		i += copy(dAtA[i:], m.Uname)
	}
	return i, nil
}

func (m *RoomAdminIsRoomAdminResp_Data) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RoomAdminIsRoomAdminResp_Data) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Userinfo != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintRoomAdmin(dAtA, i, uint64(m.Userinfo.Size()))
		n2, err := m.Userinfo.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if m.Uid != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintRoomAdmin(dAtA, i, uint64(m.Uid))
	}
	if m.Roomid != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintRoomAdmin(dAtA, i, uint64(m.Roomid))
	}
	if len(m.Ctime) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintRoomAdmin(dAtA, i, uint64(len(m.Ctime)))
		i += copy(dAtA[i:], m.Ctime)
	}
	return i, nil
}

func encodeVarintRoomAdmin(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *RoomAdminIsRoomAdminReq) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Uid)
	if l > 0 {
		n += 1 + l + sovRoomAdmin(uint64(l))
	}
	if m.Roomid != 0 {
		n += 1 + sovRoomAdmin(uint64(m.Roomid))
	}
	if m.AnchorId != 0 {
		n += 1 + sovRoomAdmin(uint64(m.AnchorId))
	}
	return n
}

func (m *RoomAdminIsRoomAdminResp) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Code != 0 {
		n += 1 + sovRoomAdmin(uint64(m.Code))
	}
	l = len(m.Msg)
	if l > 0 {
		n += 1 + l + sovRoomAdmin(uint64(l))
	}
	if m.Data != nil {
		l = m.Data.Size()
		n += 1 + l + sovRoomAdmin(uint64(l))
	}
	return n
}

func (m *RoomAdminIsRoomAdminResp_UI) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Uid != 0 {
		n += 1 + sovRoomAdmin(uint64(m.Uid))
	}
	l = len(m.Uname)
	if l > 0 {
		n += 1 + l + sovRoomAdmin(uint64(l))
	}
	return n
}

func (m *RoomAdminIsRoomAdminResp_Data) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Userinfo != nil {
		l = m.Userinfo.Size()
		n += 1 + l + sovRoomAdmin(uint64(l))
	}
	if m.Uid != 0 {
		n += 1 + sovRoomAdmin(uint64(m.Uid))
	}
	if m.Roomid != 0 {
		n += 1 + sovRoomAdmin(uint64(m.Roomid))
	}
	l = len(m.Ctime)
	if l > 0 {
		n += 1 + l + sovRoomAdmin(uint64(l))
	}
	return n
}

func sovRoomAdmin(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozRoomAdmin(x uint64) (n int) {
	return sovRoomAdmin(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *RoomAdminIsRoomAdminReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRoomAdmin
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: RoomAdminIsRoomAdminReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RoomAdminIsRoomAdminReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Uid", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRoomAdmin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRoomAdmin
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Uid = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Roomid", wireType)
			}
			m.Roomid = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRoomAdmin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Roomid |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AnchorId", wireType)
			}
			m.AnchorId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRoomAdmin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AnchorId |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipRoomAdmin(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRoomAdmin
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *RoomAdminIsRoomAdminResp) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRoomAdmin
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: RoomAdminIsRoomAdminResp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RoomAdminIsRoomAdminResp: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Code", wireType)
			}
			m.Code = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRoomAdmin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Code |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Msg", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRoomAdmin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRoomAdmin
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Msg = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRoomAdmin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRoomAdmin
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Data == nil {
				m.Data = &RoomAdminIsRoomAdminResp_Data{}
			}
			if err := m.Data.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRoomAdmin(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRoomAdmin
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *RoomAdminIsRoomAdminResp_UI) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRoomAdmin
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: UI: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UI: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Uid", wireType)
			}
			m.Uid = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRoomAdmin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Uid |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Uname", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRoomAdmin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRoomAdmin
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Uname = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRoomAdmin(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRoomAdmin
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *RoomAdminIsRoomAdminResp_Data) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRoomAdmin
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Data: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Data: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Userinfo", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRoomAdmin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRoomAdmin
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Userinfo == nil {
				m.Userinfo = &RoomAdminIsRoomAdminResp_UI{}
			}
			if err := m.Userinfo.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Uid", wireType)
			}
			m.Uid = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRoomAdmin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Uid |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Roomid", wireType)
			}
			m.Roomid = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRoomAdmin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Roomid |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ctime", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRoomAdmin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRoomAdmin
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Ctime = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRoomAdmin(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRoomAdmin
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipRoomAdmin(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRoomAdmin
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowRoomAdmin
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowRoomAdmin
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthRoomAdmin
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowRoomAdmin
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipRoomAdmin(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthRoomAdmin = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRoomAdmin   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("v1/RoomAdmin.proto", fileDescriptor_RoomAdmin_00e9d4c4f349d463) }

var fileDescriptor_RoomAdmin_00e9d4c4f349d463 = []byte{
	// 410 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x4d, 0x8e, 0xd3, 0x30,
	0x14, 0xc7, 0xeb, 0x24, 0x33, 0x34, 0x6f, 0x66, 0x36, 0xde, 0x10, 0xa2, 0x51, 0x3c, 0xaa, 0x04,
	0x1a, 0x40, 0xa4, 0x6a, 0xb9, 0x00, 0x44, 0x6c, 0xb2, 0x35, 0x9a, 0x0d, 0x0b, 0xa2, 0x34, 0xc9,
	0xa4, 0x16, 0x38, 0x6e, 0xf3, 0x75, 0x02, 0x0e, 0xc0, 0x09, 0xb8, 0x08, 0x17, 0x60, 0xd9, 0x25,
	0xab, 0x08, 0xb5, 0xbb, 0x9c, 0x02, 0xd9, 0x09, 0x69, 0x85, 0x5a, 0xa9, 0x9b, 0xf7, 0xf1, 0x97,
	0xdf, 0x7b, 0x3f, 0x3f, 0x1b, 0x70, 0x3d, 0x9b, 0x52, 0x21, 0xf8, 0xfb, 0x98, 0xb3, 0xcc, 0x5d,
	0xe5, 0xa2, 0x14, 0xf8, 0xfa, 0x2b, 0xab, 0x93, 0xa0, 0x2a, 0x92, 0xdc, 0xad, 0x67, 0xf6, 0x9b,
	0x94, 0x95, 0xcb, 0x6a, 0xe1, 0x46, 0x82, 0x4f, 0x53, 0x91, 0x8a, 0xa9, 0x3a, 0xb4, 0xa8, 0x1e,
	0x55, 0xa6, 0x12, 0x15, 0x75, 0xc5, 0x93, 0x6f, 0x08, 0x9e, 0x0e, 0x0d, 0xfd, 0x62, 0x08, 0x69,
	0xb2, 0xc6, 0xcf, 0x40, 0xaf, 0x58, 0x6c, 0xa1, 0x3b, 0x74, 0x6f, 0x7a, 0x4f, 0xda, 0x86, 0xc8,
	0x94, 0x4a, 0x83, 0x27, 0x70, 0x99, 0x0b, 0xc1, 0x59, 0x6c, 0x69, 0x77, 0xe8, 0x5e, 0xf7, 0xa0,
	0x6d, 0x48, 0xaf, 0xd0, 0xde, 0xe3, 0x57, 0x60, 0x86, 0x59, 0xb4, 0x14, 0x79, 0xc0, 0x62, 0x4b,
	0x57, 0xc7, 0x6e, 0xda, 0x86, 0xec, 0x45, 0x3a, 0xee, 0x42, 0x3f, 0x9e, 0xfc, 0xd0, 0xc1, 0x3a,
	0x8e, 0x51, 0xac, 0xf0, 0x2d, 0x18, 0x91, 0x88, 0x13, 0x05, 0xa2, 0x7b, 0xe3, 0xb6, 0x21, 0x2a,
	0xa7, 0xca, 0x4a, 0x4a, 0x5e, 0xa4, 0x8a, 0xa3, 0xa7, 0xe4, 0x45, 0x4a, 0xa5, 0xc1, 0x3e, 0x18,
	0x71, 0x58, 0x86, 0x6a, 0xf8, 0xd5, 0xfc, 0xb5, 0x7b, 0xb8, 0x28, 0xf7, 0xd4, 0x38, 0xf7, 0x43,
	0x58, 0x86, 0xdd, 0x14, 0x59, 0x4c, 0x95, 0xb5, 0xdf, 0x81, 0xf6, 0xe0, 0x1f, 0x6e, 0x44, 0xff,
	0x6f, 0x23, 0x04, 0x2e, 0xaa, 0x2c, 0xe4, 0x49, 0x0f, 0x62, 0xb6, 0x0d, 0xe9, 0x04, 0xda, 0x39,
	0xfb, 0x27, 0x02, 0x43, 0xb6, 0xc6, 0x1f, 0x61, 0x2c, 0x19, 0x58, 0xf6, 0x28, 0x54, 0xa7, 0xab,
	0xf9, 0xcb, 0x33, 0xc9, 0x1e, 0x7c, 0xef, 0xba, 0x6d, 0xc8, 0x50, 0x4e, 0x87, 0xe8, 0x1f, 0x99,
	0x76, 0x84, 0x6c, 0xff, 0x56, 0xfa, 0xc9, 0xb7, 0x22, 0x70, 0x11, 0x95, 0x8c, 0x27, 0x96, 0xb1,
	0xa7, 0x57, 0x02, 0xed, 0xdc, 0xfc, 0x0b, 0x98, 0x03, 0x0b, 0xfe, 0x0c, 0x37, 0xac, 0x08, 0x64,
	0x69, 0x10, 0x2a, 0xe1, 0xf9, 0x39, 0x17, 0x58, 0xdb, 0x2f, 0xce, 0xbb, 0xa7, 0x77, 0xfb, 0x6b,
	0xeb, 0xa0, 0xcd, 0xd6, 0x41, 0x7f, 0xb6, 0x0e, 0xfa, 0xbe, 0x73, 0x46, 0x9b, 0x9d, 0x33, 0xfa,
	0xbd, 0x73, 0x46, 0x9f, 0xb4, 0x7a, 0xb6, 0xb8, 0x54, 0x3f, 0xf7, 0xed, 0xdf, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x8f, 0x71, 0xc5, 0x66, 0x0c, 0x03, 0x00, 0x00,
}