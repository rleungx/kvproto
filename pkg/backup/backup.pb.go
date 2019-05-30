// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: backup.proto

package backup

import (
	"fmt"
	"io"
	"math"

	proto "github.com/golang/protobuf/proto"

	_ "github.com/gogo/protobuf/gogoproto"

	eraftpb "github.com/pingcap/kvproto/pkg/eraftpb"

	errorpb "github.com/pingcap/kvproto/pkg/errorpb"

	kvrpcpb "github.com/pingcap/kvproto/pkg/kvrpcpb"

	context "golang.org/x/net/context"

	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type BackupEvent_Event int32

const (
	BackupEvent_Unknown       BackupEvent_Event = 0
	BackupEvent_Snapshot      BackupEvent_Event = 1
	BackupEvent_Split         BackupEvent_Event = 2
	BackupEvent_PrepareMerge  BackupEvent_Event = 3
	BackupEvent_CommitMerge   BackupEvent_Event = 4
	BackupEvent_RollbackMerge BackupEvent_Event = 5
)

var BackupEvent_Event_name = map[int32]string{
	0: "Unknown",
	1: "Snapshot",
	2: "Split",
	3: "PrepareMerge",
	4: "CommitMerge",
	5: "RollbackMerge",
}
var BackupEvent_Event_value = map[string]int32{
	"Unknown":       0,
	"Snapshot":      1,
	"Split":         2,
	"PrepareMerge":  3,
	"CommitMerge":   4,
	"RollbackMerge": 5,
}

func (x BackupEvent_Event) String() string {
	return proto.EnumName(BackupEvent_Event_name, int32(x))
}
func (BackupEvent_Event) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_backup_07a137a2a80798d5, []int{1, 0}
}

type EntryBatch struct {
	Entries              []*eraftpb.Entry `protobuf:"bytes,1,rep,name=entries" json:"entries,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *EntryBatch) Reset()         { *m = EntryBatch{} }
func (m *EntryBatch) String() string { return proto.CompactTextString(m) }
func (*EntryBatch) ProtoMessage()    {}
func (*EntryBatch) Descriptor() ([]byte, []int) {
	return fileDescriptor_backup_07a137a2a80798d5, []int{0}
}
func (m *EntryBatch) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EntryBatch) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EntryBatch.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *EntryBatch) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EntryBatch.Merge(dst, src)
}
func (m *EntryBatch) XXX_Size() int {
	return m.Size()
}
func (m *EntryBatch) XXX_DiscardUnknown() {
	xxx_messageInfo_EntryBatch.DiscardUnknown(m)
}

var xxx_messageInfo_EntryBatch proto.InternalMessageInfo

func (m *EntryBatch) GetEntries() []*eraftpb.Entry {
	if m != nil {
		return m.Entries
	}
	return nil
}

type BackupEvent struct {
	RegionId uint64 `protobuf:"varint,1,opt,name=region_id,json=regionId,proto3" json:"region_id,omitempty"`
	// Raft log index.
	Index                uint64            `protobuf:"varint,2,opt,name=index,proto3" json:"index,omitempty"`
	RelatedRegionIds     []uint64          `protobuf:"varint,3,rep,packed,name=related_region_ids,json=relatedRegionIds" json:"related_region_ids,omitempty"`
	Event                BackupEvent_Event `protobuf:"varint,4,opt,name=event,proto3,enum=backup.BackupEvent_Event" json:"event,omitempty"`
	Dependency           uint64            `protobuf:"varint,5,opt,name=dependency,proto3" json:"dependency,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *BackupEvent) Reset()         { *m = BackupEvent{} }
func (m *BackupEvent) String() string { return proto.CompactTextString(m) }
func (*BackupEvent) ProtoMessage()    {}
func (*BackupEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_backup_07a137a2a80798d5, []int{1}
}
func (m *BackupEvent) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BackupEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BackupEvent.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *BackupEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BackupEvent.Merge(dst, src)
}
func (m *BackupEvent) XXX_Size() int {
	return m.Size()
}
func (m *BackupEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_BackupEvent.DiscardUnknown(m)
}

var xxx_messageInfo_BackupEvent proto.InternalMessageInfo

func (m *BackupEvent) GetRegionId() uint64 {
	if m != nil {
		return m.RegionId
	}
	return 0
}

func (m *BackupEvent) GetIndex() uint64 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *BackupEvent) GetRelatedRegionIds() []uint64 {
	if m != nil {
		return m.RelatedRegionIds
	}
	return nil
}

func (m *BackupEvent) GetEvent() BackupEvent_Event {
	if m != nil {
		return m.Event
	}
	return BackupEvent_Unknown
}

func (m *BackupEvent) GetDependency() uint64 {
	if m != nil {
		return m.Dependency
	}
	return 0
}

type BackupMeta struct {
	Events                []*BackupEvent `protobuf:"bytes,1,rep,name=events" json:"events,omitempty"`
	StartBackupDependency uint64         `protobuf:"varint,2,opt,name=start_backup_dependency,json=startBackupDependency,proto3" json:"start_backup_dependency,omitempty"`
	FullBackupDependency  uint64         `protobuf:"varint,3,opt,name=full_backup_dependency,json=fullBackupDependency,proto3" json:"full_backup_dependency,omitempty"`
	IncBackupDependencies []uint64       `protobuf:"varint,4,rep,packed,name=inc_backup_dependencies,json=incBackupDependencies" json:"inc_backup_dependencies,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}       `json:"-"`
	XXX_unrecognized      []byte         `json:"-"`
	XXX_sizecache         int32          `json:"-"`
}

func (m *BackupMeta) Reset()         { *m = BackupMeta{} }
func (m *BackupMeta) String() string { return proto.CompactTextString(m) }
func (*BackupMeta) ProtoMessage()    {}
func (*BackupMeta) Descriptor() ([]byte, []int) {
	return fileDescriptor_backup_07a137a2a80798d5, []int{2}
}
func (m *BackupMeta) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BackupMeta) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BackupMeta.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *BackupMeta) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BackupMeta.Merge(dst, src)
}
func (m *BackupMeta) XXX_Size() int {
	return m.Size()
}
func (m *BackupMeta) XXX_DiscardUnknown() {
	xxx_messageInfo_BackupMeta.DiscardUnknown(m)
}

var xxx_messageInfo_BackupMeta proto.InternalMessageInfo

func (m *BackupMeta) GetEvents() []*BackupEvent {
	if m != nil {
		return m.Events
	}
	return nil
}

func (m *BackupMeta) GetStartBackupDependency() uint64 {
	if m != nil {
		return m.StartBackupDependency
	}
	return 0
}

func (m *BackupMeta) GetFullBackupDependency() uint64 {
	if m != nil {
		return m.FullBackupDependency
	}
	return 0
}

func (m *BackupMeta) GetIncBackupDependencies() []uint64 {
	if m != nil {
		return m.IncBackupDependencies
	}
	return nil
}

type BackupRequest struct {
	Context              *kvrpcpb.Context `protobuf:"bytes,1,opt,name=context" json:"context,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *BackupRequest) Reset()         { *m = BackupRequest{} }
func (m *BackupRequest) String() string { return proto.CompactTextString(m) }
func (*BackupRequest) ProtoMessage()    {}
func (*BackupRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_backup_07a137a2a80798d5, []int{3}
}
func (m *BackupRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BackupRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BackupRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *BackupRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BackupRequest.Merge(dst, src)
}
func (m *BackupRequest) XXX_Size() int {
	return m.Size()
}
func (m *BackupRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BackupRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BackupRequest proto.InternalMessageInfo

func (m *BackupRequest) GetContext() *kvrpcpb.Context {
	if m != nil {
		return m.Context
	}
	return nil
}

type BackupResponse struct {
	RegionError          *errorpb.Error `protobuf:"bytes,1,opt,name=region_error,json=regionError" json:"region_error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *BackupResponse) Reset()         { *m = BackupResponse{} }
func (m *BackupResponse) String() string { return proto.CompactTextString(m) }
func (*BackupResponse) ProtoMessage()    {}
func (*BackupResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_backup_07a137a2a80798d5, []int{4}
}
func (m *BackupResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BackupResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BackupResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *BackupResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BackupResponse.Merge(dst, src)
}
func (m *BackupResponse) XXX_Size() int {
	return m.Size()
}
func (m *BackupResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BackupResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BackupResponse proto.InternalMessageInfo

func (m *BackupResponse) GetRegionError() *errorpb.Error {
	if m != nil {
		return m.RegionError
	}
	return nil
}

func init() {
	proto.RegisterType((*EntryBatch)(nil), "backup.EntryBatch")
	proto.RegisterType((*BackupEvent)(nil), "backup.BackupEvent")
	proto.RegisterType((*BackupMeta)(nil), "backup.BackupMeta")
	proto.RegisterType((*BackupRequest)(nil), "backup.BackupRequest")
	proto.RegisterType((*BackupResponse)(nil), "backup.BackupResponse")
	proto.RegisterEnum("backup.BackupEvent_Event", BackupEvent_Event_name, BackupEvent_Event_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Backup service

type BackupClient interface {
	Backup(ctx context.Context, in *BackupRequest, opts ...grpc.CallOption) (*BackupResponse, error)
}

type backupClient struct {
	cc *grpc.ClientConn
}

func NewBackupClient(cc *grpc.ClientConn) BackupClient {
	return &backupClient{cc}
}

func (c *backupClient) Backup(ctx context.Context, in *BackupRequest, opts ...grpc.CallOption) (*BackupResponse, error) {
	out := new(BackupResponse)
	err := c.cc.Invoke(ctx, "/backup.Backup/backup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Backup service

type BackupServer interface {
	Backup(context.Context, *BackupRequest) (*BackupResponse, error)
}

func RegisterBackupServer(s *grpc.Server, srv BackupServer) {
	s.RegisterService(&_Backup_serviceDesc, srv)
}

func _Backup_Backup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BackupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackupServer).Backup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/backup.Backup/Backup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackupServer).Backup(ctx, req.(*BackupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Backup_serviceDesc = grpc.ServiceDesc{
	ServiceName: "backup.Backup",
	HandlerType: (*BackupServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "backup",
			Handler:    _Backup_Backup_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "backup.proto",
}

func (m *EntryBatch) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EntryBatch) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Entries) > 0 {
		for _, msg := range m.Entries {
			dAtA[i] = 0xa
			i++
			i = encodeVarintBackup(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *BackupEvent) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BackupEvent) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.RegionId != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintBackup(dAtA, i, uint64(m.RegionId))
	}
	if m.Index != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintBackup(dAtA, i, uint64(m.Index))
	}
	if len(m.RelatedRegionIds) > 0 {
		dAtA2 := make([]byte, len(m.RelatedRegionIds)*10)
		var j1 int
		for _, num := range m.RelatedRegionIds {
			for num >= 1<<7 {
				dAtA2[j1] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j1++
			}
			dAtA2[j1] = uint8(num)
			j1++
		}
		dAtA[i] = 0x1a
		i++
		i = encodeVarintBackup(dAtA, i, uint64(j1))
		i += copy(dAtA[i:], dAtA2[:j1])
	}
	if m.Event != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintBackup(dAtA, i, uint64(m.Event))
	}
	if m.Dependency != 0 {
		dAtA[i] = 0x28
		i++
		i = encodeVarintBackup(dAtA, i, uint64(m.Dependency))
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *BackupMeta) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BackupMeta) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Events) > 0 {
		for _, msg := range m.Events {
			dAtA[i] = 0xa
			i++
			i = encodeVarintBackup(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.StartBackupDependency != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintBackup(dAtA, i, uint64(m.StartBackupDependency))
	}
	if m.FullBackupDependency != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintBackup(dAtA, i, uint64(m.FullBackupDependency))
	}
	if len(m.IncBackupDependencies) > 0 {
		dAtA4 := make([]byte, len(m.IncBackupDependencies)*10)
		var j3 int
		for _, num := range m.IncBackupDependencies {
			for num >= 1<<7 {
				dAtA4[j3] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j3++
			}
			dAtA4[j3] = uint8(num)
			j3++
		}
		dAtA[i] = 0x22
		i++
		i = encodeVarintBackup(dAtA, i, uint64(j3))
		i += copy(dAtA[i:], dAtA4[:j3])
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *BackupRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BackupRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Context != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintBackup(dAtA, i, uint64(m.Context.Size()))
		n5, err := m.Context.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n5
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *BackupResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BackupResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.RegionError != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintBackup(dAtA, i, uint64(m.RegionError.Size()))
		n6, err := m.RegionError.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n6
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintBackup(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *EntryBatch) Size() (n int) {
	var l int
	_ = l
	if len(m.Entries) > 0 {
		for _, e := range m.Entries {
			l = e.Size()
			n += 1 + l + sovBackup(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *BackupEvent) Size() (n int) {
	var l int
	_ = l
	if m.RegionId != 0 {
		n += 1 + sovBackup(uint64(m.RegionId))
	}
	if m.Index != 0 {
		n += 1 + sovBackup(uint64(m.Index))
	}
	if len(m.RelatedRegionIds) > 0 {
		l = 0
		for _, e := range m.RelatedRegionIds {
			l += sovBackup(uint64(e))
		}
		n += 1 + sovBackup(uint64(l)) + l
	}
	if m.Event != 0 {
		n += 1 + sovBackup(uint64(m.Event))
	}
	if m.Dependency != 0 {
		n += 1 + sovBackup(uint64(m.Dependency))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *BackupMeta) Size() (n int) {
	var l int
	_ = l
	if len(m.Events) > 0 {
		for _, e := range m.Events {
			l = e.Size()
			n += 1 + l + sovBackup(uint64(l))
		}
	}
	if m.StartBackupDependency != 0 {
		n += 1 + sovBackup(uint64(m.StartBackupDependency))
	}
	if m.FullBackupDependency != 0 {
		n += 1 + sovBackup(uint64(m.FullBackupDependency))
	}
	if len(m.IncBackupDependencies) > 0 {
		l = 0
		for _, e := range m.IncBackupDependencies {
			l += sovBackup(uint64(e))
		}
		n += 1 + sovBackup(uint64(l)) + l
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *BackupRequest) Size() (n int) {
	var l int
	_ = l
	if m.Context != nil {
		l = m.Context.Size()
		n += 1 + l + sovBackup(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *BackupResponse) Size() (n int) {
	var l int
	_ = l
	if m.RegionError != nil {
		l = m.RegionError.Size()
		n += 1 + l + sovBackup(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovBackup(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozBackup(x uint64) (n int) {
	return sovBackup(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *EntryBatch) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBackup
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
			return fmt.Errorf("proto: EntryBatch: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EntryBatch: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Entries", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBackup
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
				return ErrInvalidLengthBackup
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Entries = append(m.Entries, &eraftpb.Entry{})
			if err := m.Entries[len(m.Entries)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBackup(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthBackup
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *BackupEvent) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBackup
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
			return fmt.Errorf("proto: BackupEvent: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BackupEvent: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RegionId", wireType)
			}
			m.RegionId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBackup
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RegionId |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			m.Index = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBackup
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Index |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType == 0 {
				var v uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowBackup
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.RelatedRegionIds = append(m.RelatedRegionIds, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowBackup
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= (int(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthBackup
				}
				postIndex := iNdEx + packedLen
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				for iNdEx < postIndex {
					var v uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowBackup
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= (uint64(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.RelatedRegionIds = append(m.RelatedRegionIds, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field RelatedRegionIds", wireType)
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Event", wireType)
			}
			m.Event = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBackup
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Event |= (BackupEvent_Event(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Dependency", wireType)
			}
			m.Dependency = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBackup
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Dependency |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipBackup(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthBackup
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *BackupMeta) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBackup
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
			return fmt.Errorf("proto: BackupMeta: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BackupMeta: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Events", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBackup
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
				return ErrInvalidLengthBackup
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Events = append(m.Events, &BackupEvent{})
			if err := m.Events[len(m.Events)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartBackupDependency", wireType)
			}
			m.StartBackupDependency = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBackup
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StartBackupDependency |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FullBackupDependency", wireType)
			}
			m.FullBackupDependency = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBackup
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.FullBackupDependency |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType == 0 {
				var v uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowBackup
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.IncBackupDependencies = append(m.IncBackupDependencies, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowBackup
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= (int(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthBackup
				}
				postIndex := iNdEx + packedLen
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				for iNdEx < postIndex {
					var v uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowBackup
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= (uint64(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.IncBackupDependencies = append(m.IncBackupDependencies, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field IncBackupDependencies", wireType)
			}
		default:
			iNdEx = preIndex
			skippy, err := skipBackup(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthBackup
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *BackupRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBackup
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
			return fmt.Errorf("proto: BackupRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BackupRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Context", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBackup
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
				return ErrInvalidLengthBackup
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Context == nil {
				m.Context = &kvrpcpb.Context{}
			}
			if err := m.Context.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBackup(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthBackup
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *BackupResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBackup
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
			return fmt.Errorf("proto: BackupResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BackupResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RegionError", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBackup
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
				return ErrInvalidLengthBackup
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.RegionError == nil {
				m.RegionError = &errorpb.Error{}
			}
			if err := m.RegionError.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBackup(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthBackup
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipBackup(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowBackup
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
					return 0, ErrIntOverflowBackup
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
					return 0, ErrIntOverflowBackup
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
				return 0, ErrInvalidLengthBackup
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowBackup
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
				next, err := skipBackup(dAtA[start:])
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
	ErrInvalidLengthBackup = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowBackup   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("backup.proto", fileDescriptor_backup_07a137a2a80798d5) }

var fileDescriptor_backup_07a137a2a80798d5 = []byte{
	// 534 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x53, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x8d, 0x93, 0x38, 0x49, 0xc7, 0x49, 0x6a, 0x96, 0xa4, 0x35, 0x41, 0x8a, 0x22, 0x1f, 0x50,
	0x04, 0xc8, 0x15, 0x01, 0x45, 0x42, 0xdc, 0x12, 0x72, 0xe0, 0x50, 0x09, 0x6d, 0xc5, 0x39, 0x72,
	0xec, 0x69, 0x6a, 0xd9, 0xdd, 0x35, 0xeb, 0x4d, 0x68, 0xff, 0x84, 0x0f, 0xe0, 0xc0, 0xa7, 0x70,
	0xe4, 0xc8, 0x0d, 0x14, 0x7e, 0x04, 0x79, 0xd7, 0x8e, 0x42, 0xdb, 0x8b, 0x77, 0xe7, 0xcd, 0xbc,
	0xd9, 0xb7, 0xe3, 0xb7, 0xd0, 0x5e, 0xf9, 0x41, 0xbc, 0x49, 0xbd, 0x54, 0x70, 0xc9, 0x49, 0x43,
	0x47, 0x83, 0x4e, 0xbc, 0x15, 0x69, 0x90, 0xae, 0x34, 0x3c, 0xe8, 0xa0, 0x10, 0x5c, 0x1c, 0x86,
	0xfe, 0xa5, 0xdc, 0x87, 0xbd, 0x35, 0x5f, 0x73, 0xb5, 0x3d, 0xcb, 0x77, 0x05, 0x7a, 0x2c, 0x36,
	0x99, 0x54, 0x5b, 0x0d, 0xb8, 0x53, 0x80, 0x05, 0x93, 0xe2, 0x76, 0xe6, 0xcb, 0xe0, 0x8a, 0x8c,
	0xa1, 0x89, 0x4c, 0x8a, 0x08, 0x33, 0xc7, 0x18, 0xd5, 0xc6, 0xd6, 0xa4, 0xeb, 0x95, 0x5d, 0x55,
	0x15, 0x2d, 0xd3, 0xee, 0xb7, 0x2a, 0x58, 0x33, 0x25, 0x6b, 0xb1, 0x45, 0x26, 0xc9, 0x53, 0x38,
	0x12, 0xb8, 0x8e, 0x38, 0x5b, 0x46, 0xa1, 0x63, 0x8c, 0x8c, 0x71, 0x9d, 0xb6, 0x34, 0xf0, 0x21,
	0x24, 0x3d, 0x30, 0x23, 0x16, 0xe2, 0x8d, 0x53, 0x55, 0x09, 0x1d, 0x90, 0x97, 0x40, 0x04, 0x26,
	0xbe, 0xc4, 0x70, 0xb9, 0xa7, 0x66, 0x4e, 0x6d, 0x54, 0x1b, 0xd7, 0xa9, 0x5d, 0x64, 0x68, 0xd1,
	0x22, 0x23, 0x67, 0x60, 0x62, 0x7e, 0x92, 0x53, 0x1f, 0x19, 0xe3, 0xee, 0xe4, 0x89, 0x57, 0x8c,
	0xe8, 0x40, 0x84, 0xa7, 0xbe, 0x54, 0xd7, 0x91, 0x21, 0x40, 0x88, 0x29, 0xb2, 0x10, 0x59, 0x70,
	0xeb, 0x98, 0xea, 0xe4, 0x03, 0xc4, 0x0d, 0xc0, 0xd4, 0xd2, 0x2d, 0x68, 0x7e, 0x62, 0x31, 0xe3,
	0x5f, 0x98, 0x5d, 0x21, 0x6d, 0x68, 0x5d, 0x30, 0x3f, 0xcd, 0xae, 0xb8, 0xb4, 0x0d, 0x72, 0x04,
	0xe6, 0x45, 0x9a, 0x44, 0xd2, 0xae, 0x12, 0x1b, 0xda, 0x1f, 0x05, 0xa6, 0xbe, 0xc0, 0x73, 0x14,
	0x6b, 0xb4, 0x6b, 0xe4, 0x18, 0xac, 0x39, 0xbf, 0xbe, 0x8e, 0xa4, 0x06, 0xea, 0xe4, 0x11, 0x74,
	0x28, 0x4f, 0x92, 0x5c, 0x98, 0x86, 0x4c, 0xf7, 0xb7, 0x01, 0xa0, 0x15, 0x9e, 0xa3, 0xf4, 0xc9,
	0x0b, 0x68, 0x28, 0x71, 0xe5, 0x78, 0x1f, 0x3f, 0x70, 0x0b, 0x5a, 0x94, 0x90, 0x29, 0x9c, 0x66,
	0xd2, 0x17, 0x72, 0xa9, 0x6b, 0x96, 0x07, 0xb7, 0xd1, 0x73, 0xec, 0xab, 0xb4, 0xa6, 0xbe, 0xdf,
	0x27, 0xc9, 0x1b, 0x38, 0xb9, 0xdc, 0x24, 0xc9, 0x03, 0xb4, 0x9a, 0xa2, 0xf5, 0xf2, 0xec, 0x3d,
	0xd6, 0x14, 0x4e, 0x23, 0x16, 0xdc, 0x23, 0xe5, 0x56, 0xa8, 0xab, 0x5f, 0xd2, 0x8f, 0x58, 0x70,
	0x87, 0x95, 0x1b, 0xe1, 0x1d, 0x74, 0x34, 0x4a, 0xf1, 0xf3, 0x06, 0x33, 0x49, 0x9e, 0x43, 0x33,
	0xe0, 0x4c, 0xe2, 0x8d, 0x54, 0x3e, 0xb0, 0x26, 0xb6, 0x57, 0xfa, 0x76, 0xae, 0x71, 0x5a, 0x16,
	0xb8, 0x73, 0xe8, 0x96, 0xe4, 0x2c, 0xe5, 0x2c, 0x43, 0xf2, 0x0a, 0xda, 0x85, 0x19, 0x94, 0xbb,
	0x8b, 0x16, 0xb9, 0x0d, 0xb5, 0xd7, 0x17, 0xf9, 0x4a, 0x2d, 0x5d, 0xa3, 0x82, 0xc9, 0x1c, 0x1a,
	0xba, 0x09, 0x79, 0x0b, 0xc5, 0x53, 0x21, 0xfd, 0xff, 0x07, 0x5b, 0x68, 0x1b, 0x9c, 0xdc, 0x85,
	0xf5, 0xa9, 0x6e, 0x65, 0xf6, 0xec, 0xd7, 0xf7, 0x96, 0xf1, 0x63, 0x37, 0x34, 0x7e, 0xee, 0x86,
	0xc6, 0x9f, 0xdd, 0xd0, 0xf8, 0xfa, 0x77, 0x58, 0x01, 0x9b, 0x8b, 0xb5, 0x27, 0xa3, 0x78, 0xeb,
	0xc5, 0x5b, 0xf5, 0x5e, 0x56, 0x0d, 0xb5, 0xbc, 0xfe, 0x17, 0x00, 0x00, 0xff, 0xff, 0xbb, 0xbe,
	0x5d, 0x9e, 0xa2, 0x03, 0x00, 0x00,
}
