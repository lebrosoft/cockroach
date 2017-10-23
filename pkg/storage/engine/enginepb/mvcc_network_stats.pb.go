// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cockroach/pkg/storage/engine/enginepb/mvcc_network_stats.proto

package enginepb

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import encoding_binary "encoding/binary"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// MVCCNetworkStats is convertible to MVCCStats, but uses variable width
// encodings for most fields. This makes the encodings incompatible. Note that
// with proto3, zero-valued primitive types will not be encoded at all.
type MVCCNetworkStats struct {
	ContainsEstimates bool  `protobuf:"varint,14,opt,name=contains_estimates,json=containsEstimates,proto3" json:"contains_estimates,omitempty"`
	LastUpdateNanos   int64 `protobuf:"fixed64,1,opt,name=last_update_nanos,json=lastUpdateNanos,proto3" json:"last_update_nanos,omitempty"`
	IntentAge         int64 `protobuf:"fixed64,2,opt,name=intent_age,json=intentAge,proto3" json:"intent_age,omitempty"`
	GCBytesAge        int64 `protobuf:"fixed64,3,opt,name=gc_bytes_age,json=gcBytesAge,proto3" json:"gc_bytes_age,omitempty"`
	LiveBytes         int64 `protobuf:"zigzag64,4,opt,name=live_bytes,json=liveBytes,proto3" json:"live_bytes,omitempty"`
	LiveCount         int64 `protobuf:"zigzag64,5,opt,name=live_count,json=liveCount,proto3" json:"live_count,omitempty"`
	KeyBytes          int64 `protobuf:"zigzag64,6,opt,name=key_bytes,json=keyBytes,proto3" json:"key_bytes,omitempty"`
	KeyCount          int64 `protobuf:"zigzag64,7,opt,name=key_count,json=keyCount,proto3" json:"key_count,omitempty"`
	ValBytes          int64 `protobuf:"zigzag64,8,opt,name=val_bytes,json=valBytes,proto3" json:"val_bytes,omitempty"`
	ValCount          int64 `protobuf:"zigzag64,9,opt,name=val_count,json=valCount,proto3" json:"val_count,omitempty"`
	IntentBytes       int64 `protobuf:"zigzag64,10,opt,name=intent_bytes,json=intentBytes,proto3" json:"intent_bytes,omitempty"`
	IntentCount       int64 `protobuf:"zigzag64,11,opt,name=intent_count,json=intentCount,proto3" json:"intent_count,omitempty"`
	SysBytes          int64 `protobuf:"zigzag64,12,opt,name=sys_bytes,json=sysBytes,proto3" json:"sys_bytes,omitempty"`
	SysCount          int64 `protobuf:"zigzag64,13,opt,name=sys_count,json=sysCount,proto3" json:"sys_count,omitempty"`
}

func (m *MVCCNetworkStats) Reset()                    { *m = MVCCNetworkStats{} }
func (m *MVCCNetworkStats) String() string            { return proto.CompactTextString(m) }
func (*MVCCNetworkStats) ProtoMessage()               {}
func (*MVCCNetworkStats) Descriptor() ([]byte, []int) { return fileDescriptorMvccNetworkStats, []int{0} }

func init() {
	proto.RegisterType((*MVCCNetworkStats)(nil), "cockroach.storage.engine.enginepb.MVCCNetworkStats")
}
func (this *MVCCNetworkStats) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*MVCCNetworkStats)
	if !ok {
		that2, ok := that.(MVCCNetworkStats)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if this.ContainsEstimates != that1.ContainsEstimates {
		return false
	}
	if this.LastUpdateNanos != that1.LastUpdateNanos {
		return false
	}
	if this.IntentAge != that1.IntentAge {
		return false
	}
	if this.GCBytesAge != that1.GCBytesAge {
		return false
	}
	if this.LiveBytes != that1.LiveBytes {
		return false
	}
	if this.LiveCount != that1.LiveCount {
		return false
	}
	if this.KeyBytes != that1.KeyBytes {
		return false
	}
	if this.KeyCount != that1.KeyCount {
		return false
	}
	if this.ValBytes != that1.ValBytes {
		return false
	}
	if this.ValCount != that1.ValCount {
		return false
	}
	if this.IntentBytes != that1.IntentBytes {
		return false
	}
	if this.IntentCount != that1.IntentCount {
		return false
	}
	if this.SysBytes != that1.SysBytes {
		return false
	}
	if this.SysCount != that1.SysCount {
		return false
	}
	return true
}
func (m *MVCCNetworkStats) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MVCCNetworkStats) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.LastUpdateNanos != 0 {
		dAtA[i] = 0x9
		i++
		encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(m.LastUpdateNanos))
		i += 8
	}
	if m.IntentAge != 0 {
		dAtA[i] = 0x11
		i++
		encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(m.IntentAge))
		i += 8
	}
	if m.GCBytesAge != 0 {
		dAtA[i] = 0x19
		i++
		encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(m.GCBytesAge))
		i += 8
	}
	if m.LiveBytes != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintMvccNetworkStats(dAtA, i, uint64((uint64(m.LiveBytes)<<1)^uint64((m.LiveBytes>>63))))
	}
	if m.LiveCount != 0 {
		dAtA[i] = 0x28
		i++
		i = encodeVarintMvccNetworkStats(dAtA, i, uint64((uint64(m.LiveCount)<<1)^uint64((m.LiveCount>>63))))
	}
	if m.KeyBytes != 0 {
		dAtA[i] = 0x30
		i++
		i = encodeVarintMvccNetworkStats(dAtA, i, uint64((uint64(m.KeyBytes)<<1)^uint64((m.KeyBytes>>63))))
	}
	if m.KeyCount != 0 {
		dAtA[i] = 0x38
		i++
		i = encodeVarintMvccNetworkStats(dAtA, i, uint64((uint64(m.KeyCount)<<1)^uint64((m.KeyCount>>63))))
	}
	if m.ValBytes != 0 {
		dAtA[i] = 0x40
		i++
		i = encodeVarintMvccNetworkStats(dAtA, i, uint64((uint64(m.ValBytes)<<1)^uint64((m.ValBytes>>63))))
	}
	if m.ValCount != 0 {
		dAtA[i] = 0x48
		i++
		i = encodeVarintMvccNetworkStats(dAtA, i, uint64((uint64(m.ValCount)<<1)^uint64((m.ValCount>>63))))
	}
	if m.IntentBytes != 0 {
		dAtA[i] = 0x50
		i++
		i = encodeVarintMvccNetworkStats(dAtA, i, uint64((uint64(m.IntentBytes)<<1)^uint64((m.IntentBytes>>63))))
	}
	if m.IntentCount != 0 {
		dAtA[i] = 0x58
		i++
		i = encodeVarintMvccNetworkStats(dAtA, i, uint64((uint64(m.IntentCount)<<1)^uint64((m.IntentCount>>63))))
	}
	if m.SysBytes != 0 {
		dAtA[i] = 0x60
		i++
		i = encodeVarintMvccNetworkStats(dAtA, i, uint64((uint64(m.SysBytes)<<1)^uint64((m.SysBytes>>63))))
	}
	if m.SysCount != 0 {
		dAtA[i] = 0x68
		i++
		i = encodeVarintMvccNetworkStats(dAtA, i, uint64((uint64(m.SysCount)<<1)^uint64((m.SysCount>>63))))
	}
	if m.ContainsEstimates {
		dAtA[i] = 0x70
		i++
		if m.ContainsEstimates {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	return i, nil
}

func encodeVarintMvccNetworkStats(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *MVCCNetworkStats) Size() (n int) {
	var l int
	_ = l
	if m.LastUpdateNanos != 0 {
		n += 9
	}
	if m.IntentAge != 0 {
		n += 9
	}
	if m.GCBytesAge != 0 {
		n += 9
	}
	if m.LiveBytes != 0 {
		n += 1 + sozMvccNetworkStats(uint64(m.LiveBytes))
	}
	if m.LiveCount != 0 {
		n += 1 + sozMvccNetworkStats(uint64(m.LiveCount))
	}
	if m.KeyBytes != 0 {
		n += 1 + sozMvccNetworkStats(uint64(m.KeyBytes))
	}
	if m.KeyCount != 0 {
		n += 1 + sozMvccNetworkStats(uint64(m.KeyCount))
	}
	if m.ValBytes != 0 {
		n += 1 + sozMvccNetworkStats(uint64(m.ValBytes))
	}
	if m.ValCount != 0 {
		n += 1 + sozMvccNetworkStats(uint64(m.ValCount))
	}
	if m.IntentBytes != 0 {
		n += 1 + sozMvccNetworkStats(uint64(m.IntentBytes))
	}
	if m.IntentCount != 0 {
		n += 1 + sozMvccNetworkStats(uint64(m.IntentCount))
	}
	if m.SysBytes != 0 {
		n += 1 + sozMvccNetworkStats(uint64(m.SysBytes))
	}
	if m.SysCount != 0 {
		n += 1 + sozMvccNetworkStats(uint64(m.SysCount))
	}
	if m.ContainsEstimates {
		n += 2
	}
	return n
}

func sovMvccNetworkStats(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozMvccNetworkStats(x uint64) (n int) {
	return sovMvccNetworkStats(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MVCCNetworkStats) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMvccNetworkStats
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
			return fmt.Errorf("proto: MVCCNetworkStats: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MVCCNetworkStats: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastUpdateNanos", wireType)
			}
			m.LastUpdateNanos = 0
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			m.LastUpdateNanos = int64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
		case 2:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field IntentAge", wireType)
			}
			m.IntentAge = 0
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			m.IntentAge = int64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
		case 3:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field GCBytesAge", wireType)
			}
			m.GCBytesAge = 0
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			m.GCBytesAge = int64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LiveBytes", wireType)
			}
			var v uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMvccNetworkStats
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
			v = (v >> 1) ^ uint64((int64(v&1)<<63)>>63)
			m.LiveBytes = int64(v)
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LiveCount", wireType)
			}
			var v uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMvccNetworkStats
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
			v = (v >> 1) ^ uint64((int64(v&1)<<63)>>63)
			m.LiveCount = int64(v)
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field KeyBytes", wireType)
			}
			var v uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMvccNetworkStats
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
			v = (v >> 1) ^ uint64((int64(v&1)<<63)>>63)
			m.KeyBytes = int64(v)
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field KeyCount", wireType)
			}
			var v uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMvccNetworkStats
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
			v = (v >> 1) ^ uint64((int64(v&1)<<63)>>63)
			m.KeyCount = int64(v)
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValBytes", wireType)
			}
			var v uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMvccNetworkStats
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
			v = (v >> 1) ^ uint64((int64(v&1)<<63)>>63)
			m.ValBytes = int64(v)
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValCount", wireType)
			}
			var v uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMvccNetworkStats
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
			v = (v >> 1) ^ uint64((int64(v&1)<<63)>>63)
			m.ValCount = int64(v)
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IntentBytes", wireType)
			}
			var v uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMvccNetworkStats
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
			v = (v >> 1) ^ uint64((int64(v&1)<<63)>>63)
			m.IntentBytes = int64(v)
		case 11:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IntentCount", wireType)
			}
			var v uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMvccNetworkStats
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
			v = (v >> 1) ^ uint64((int64(v&1)<<63)>>63)
			m.IntentCount = int64(v)
		case 12:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SysBytes", wireType)
			}
			var v uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMvccNetworkStats
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
			v = (v >> 1) ^ uint64((int64(v&1)<<63)>>63)
			m.SysBytes = int64(v)
		case 13:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SysCount", wireType)
			}
			var v uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMvccNetworkStats
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
			v = (v >> 1) ^ uint64((int64(v&1)<<63)>>63)
			m.SysCount = int64(v)
		case 14:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContainsEstimates", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMvccNetworkStats
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.ContainsEstimates = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipMvccNetworkStats(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMvccNetworkStats
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
func skipMvccNetworkStats(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMvccNetworkStats
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
					return 0, ErrIntOverflowMvccNetworkStats
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
					return 0, ErrIntOverflowMvccNetworkStats
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
				return 0, ErrInvalidLengthMvccNetworkStats
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowMvccNetworkStats
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
				next, err := skipMvccNetworkStats(dAtA[start:])
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
	ErrInvalidLengthMvccNetworkStats = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMvccNetworkStats   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("cockroach/pkg/storage/engine/enginepb/mvcc_network_stats.proto", fileDescriptorMvccNetworkStats)
}

var fileDescriptorMvccNetworkStats = []byte{
	// 393 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x92, 0xb1, 0xae, 0xd3, 0x30,
	0x18, 0x85, 0xaf, 0xb9, 0x97, 0xd2, 0xb8, 0xa5, 0xb4, 0x11, 0x43, 0x04, 0x22, 0x6d, 0x99, 0x2a,
	0x24, 0x12, 0x24, 0x36, 0x06, 0xa4, 0x36, 0x42, 0x4c, 0x74, 0x08, 0x82, 0x35, 0x72, 0x8d, 0x65,
	0xa2, 0xa4, 0x76, 0x14, 0xbb, 0x41, 0x79, 0x0b, 0x1e, 0x81, 0xc7, 0xe9, 0xc8, 0xc8, 0x84, 0x20,
	0x2c, 0x3c, 0x01, 0x33, 0xb2, 0xff, 0xc4, 0xbd, 0x53, 0xad, 0xf3, 0x9d, 0xef, 0xfc, 0x52, 0x15,
	0xfc, 0x9a, 0x4a, 0x5a, 0xd4, 0x92, 0xd0, 0xcf, 0x71, 0x55, 0xf0, 0x58, 0x69, 0x59, 0x13, 0xce,
	0x62, 0x26, 0x78, 0x2e, 0x86, 0x9f, 0xea, 0x10, 0x1f, 0x1b, 0x4a, 0x33, 0xc1, 0xf4, 0x17, 0x59,
	0x17, 0x99, 0xd2, 0x44, 0xab, 0xa8, 0xaa, 0xa5, 0x96, 0xfe, 0xda, 0xf9, 0x51, 0xef, 0x46, 0x20,
	0x45, 0x83, 0xfb, 0xe8, 0x21, 0x97, 0x5c, 0xda, 0x76, 0x6c, 0x5e, 0x20, 0x3e, 0xfd, 0x77, 0x8d,
	0xe7, 0xef, 0x3e, 0x26, 0xc9, 0x1e, 0x46, 0xdf, 0x9b, 0x4d, 0xff, 0x19, 0x5e, 0x94, 0x44, 0xe9,
	0xec, 0x54, 0x7d, 0x22, 0x9a, 0x65, 0x82, 0x08, 0xa9, 0x02, 0xb4, 0x42, 0x9b, 0x79, 0xfa, 0xc0,
	0x80, 0x0f, 0x36, 0xdf, 0x9b, 0xd8, 0x7f, 0x82, 0x71, 0x2e, 0x34, 0x13, 0x3a, 0x23, 0x9c, 0x05,
	0x77, 0x6c, 0xc9, 0x83, 0x64, 0xcb, 0x99, 0xff, 0x02, 0x4f, 0x39, 0xcd, 0x0e, 0xad, 0x66, 0xca,
	0x16, 0xae, 0x4d, 0x61, 0x37, 0xeb, 0x7e, 0x2e, 0xf1, 0xdb, 0x64, 0x67, 0xe2, 0x2d, 0x67, 0x29,
	0xe6, 0x74, 0x78, 0x9b, 0xc1, 0x32, 0x6f, 0x18, 0x38, 0xc1, 0xcd, 0x0a, 0x6d, 0xfc, 0xd4, 0x33,
	0x89, 0x6d, 0x38, 0x4c, 0xe5, 0x49, 0xe8, 0xe0, 0xee, 0x05, 0x27, 0x26, 0xf0, 0x1f, 0x63, 0xaf,
	0x60, 0x6d, 0x2f, 0x8f, 0x2c, 0x1d, 0x17, 0xac, 0x05, 0xb7, 0x87, 0xa0, 0xde, 0x73, 0xd0, 0x99,
	0x0d, 0x29, 0x7b, 0x73, 0x0c, 0xb0, 0x21, 0xa5, 0x33, 0x0d, 0x04, 0xd3, 0x73, 0x10, 0xcc, 0x35,
	0x9e, 0xf6, 0x7f, 0x01, 0xc8, 0xd8, 0xf2, 0x09, 0x64, 0xe0, 0x5f, 0x2a, 0x30, 0x31, 0xb9, 0x5d,
	0x71, 0xf7, 0x55, 0xab, 0xfa, 0x89, 0x29, 0x9c, 0x50, 0xad, 0x72, 0xf7, 0x0d, 0x04, 0xf9, 0xbe,
	0x83, 0x60, 0x3e, 0xc7, 0x3e, 0x95, 0x42, 0x93, 0x5c, 0xa8, 0x8c, 0x29, 0x9d, 0x1f, 0x89, 0x99,
	0x98, 0xad, 0xd0, 0x66, 0x9c, 0x2e, 0x06, 0xf2, 0x66, 0x00, 0xaf, 0x6e, 0xfe, 0x7e, 0x5b, 0xa2,
	0x5d, 0x70, 0xfe, 0x1d, 0x5e, 0x9d, 0xbb, 0x10, 0x7d, 0xef, 0x42, 0xf4, 0xa3, 0x0b, 0xd1, 0xaf,
	0x2e, 0x44, 0x5f, 0xff, 0x84, 0x57, 0x87, 0x91, 0xfd, 0x32, 0x5e, 0xfe, 0x0f, 0x00, 0x00, 0xff,
	0xff, 0x34, 0x29, 0x90, 0x45, 0x94, 0x02, 0x00, 0x00,
}