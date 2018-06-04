// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: roachpb/io-formats.proto

package roachpb

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type IOFileFormat_FileFormat int32

const (
	IOFileFormat_Unknown      IOFileFormat_FileFormat = 0
	IOFileFormat_CSV          IOFileFormat_FileFormat = 1
	IOFileFormat_MysqlOutfile IOFileFormat_FileFormat = 2
	IOFileFormat_Mysqldump    IOFileFormat_FileFormat = 3
	IOFileFormat_PgCopy       IOFileFormat_FileFormat = 4
)

var IOFileFormat_FileFormat_name = map[int32]string{
	0: "Unknown",
	1: "CSV",
	2: "MysqlOutfile",
	3: "Mysqldump",
	4: "PgCopy",
}
var IOFileFormat_FileFormat_value = map[string]int32{
	"Unknown":      0,
	"CSV":          1,
	"MysqlOutfile": 2,
	"Mysqldump":    3,
	"PgCopy":       4,
}

func (x IOFileFormat_FileFormat) Enum() *IOFileFormat_FileFormat {
	p := new(IOFileFormat_FileFormat)
	*p = x
	return p
}
func (x IOFileFormat_FileFormat) String() string {
	return proto.EnumName(IOFileFormat_FileFormat_name, int32(x))
}
func (x *IOFileFormat_FileFormat) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(IOFileFormat_FileFormat_value, data, "IOFileFormat_FileFormat")
	if err != nil {
		return err
	}
	*x = IOFileFormat_FileFormat(value)
	return nil
}
func (IOFileFormat_FileFormat) EnumDescriptor() ([]byte, []int) {
	return fileDescriptorIoFormats, []int{0, 0}
}

type MySQLOutfileOptions_Enclose int32

const (
	MySQLOutfileOptions_Never    MySQLOutfileOptions_Enclose = 0
	MySQLOutfileOptions_Always   MySQLOutfileOptions_Enclose = 1
	MySQLOutfileOptions_Optional MySQLOutfileOptions_Enclose = 2
)

var MySQLOutfileOptions_Enclose_name = map[int32]string{
	0: "Never",
	1: "Always",
	2: "Optional",
}
var MySQLOutfileOptions_Enclose_value = map[string]int32{
	"Never":    0,
	"Always":   1,
	"Optional": 2,
}

func (x MySQLOutfileOptions_Enclose) Enum() *MySQLOutfileOptions_Enclose {
	p := new(MySQLOutfileOptions_Enclose)
	*p = x
	return p
}
func (x MySQLOutfileOptions_Enclose) String() string {
	return proto.EnumName(MySQLOutfileOptions_Enclose_name, int32(x))
}
func (x *MySQLOutfileOptions_Enclose) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(MySQLOutfileOptions_Enclose_value, data, "MySQLOutfileOptions_Enclose")
	if err != nil {
		return err
	}
	*x = MySQLOutfileOptions_Enclose(value)
	return nil
}
func (MySQLOutfileOptions_Enclose) EnumDescriptor() ([]byte, []int) {
	return fileDescriptorIoFormats, []int{2, 0}
}

type IOFileFormat struct {
	Format   IOFileFormat_FileFormat `protobuf:"varint,1,opt,name=format,enum=cockroach.roachpb.IOFileFormat_FileFormat" json:"format"`
	Csv      CSVOptions              `protobuf:"bytes,2,opt,name=csv" json:"csv"`
	MysqlOut MySQLOutfileOptions     `protobuf:"bytes,3,opt,name=mysql_out,json=mysqlOut" json:"mysql_out"`
	PgCopy   PgCopyOptions           `protobuf:"bytes,4,opt,name=pg_copy,json=pgCopy" json:"pg_copy"`
}

func (m *IOFileFormat) Reset()                    { *m = IOFileFormat{} }
func (m *IOFileFormat) String() string            { return proto.CompactTextString(m) }
func (*IOFileFormat) ProtoMessage()               {}
func (*IOFileFormat) Descriptor() ([]byte, []int) { return fileDescriptorIoFormats, []int{0} }

// CSVOptions describe the format of csv data (delimiter, comment, etc).
type CSVOptions struct {
	// comma is an delimiter used by the CSV file; defaults to a comma.
	Comma int32 `protobuf:"varint,1,opt,name=comma" json:"comma"`
	// comment is an comment rune; zero value means comments not enabled.
	Comment int32 `protobuf:"varint,2,opt,name=comment" json:"comment"`
	// null_encoding, if not nil, is the string which identifies a NULL. Can be the empty string.
	NullEncoding *string `protobuf:"bytes,3,opt,name=null_encoding,json=nullEncoding" json:"null_encoding,omitempty"`
	// skip the first N lines of the input (e.g. to ignore column headers) when reading.
	Skip uint32 `protobuf:"varint,4,opt,name=skip" json:"skip"`
}

func (m *CSVOptions) Reset()                    { *m = CSVOptions{} }
func (m *CSVOptions) String() string            { return proto.CompactTextString(m) }
func (*CSVOptions) ProtoMessage()               {}
func (*CSVOptions) Descriptor() ([]byte, []int) { return fileDescriptorIoFormats, []int{1} }

// MySQLOutfileOptions describe the format of mysql's outfile.
type MySQLOutfileOptions struct {
	// row_separator is the delimiter between rows (mysql's --rows-terminated-by)
	RowSeparator int32 `protobuf:"varint,1,opt,name=row_separator,json=rowSeparator" json:"row_separator"`
	// field_separator is the delimiter between fields (mysql's --fields-terminated-by)
	FieldSeparator int32 `protobuf:"varint,2,opt,name=field_separator,json=fieldSeparator" json:"field_separator"`
	// enclose is the enclosing (quoting) behavior (i.e. if specified and if optional).
	Enclose MySQLOutfileOptions_Enclose `protobuf:"varint,3,opt,name=enclose,enum=cockroach.roachpb.MySQLOutfileOptions_Enclose" json:"enclose"`
	// encloser is the character used to enclose (qupte) fields (--fields-enclosed-by)
	Encloser int32 `protobuf:"varint,4,opt,name=encloser" json:"encloser"`
	// has_escape indicates that an escape character is set (mysql's default is not).
	HasEscape bool `protobuf:"varint,5,opt,name=has_escape,json=hasEscape" json:"has_escape"`
	// escape is the character used to prefix the other delimiters (--fields-escaped-by)
	Escape int32 `protobuf:"varint,6,opt,name=escape" json:"escape"`
}

func (m *MySQLOutfileOptions) Reset()                    { *m = MySQLOutfileOptions{} }
func (m *MySQLOutfileOptions) String() string            { return proto.CompactTextString(m) }
func (*MySQLOutfileOptions) ProtoMessage()               {}
func (*MySQLOutfileOptions) Descriptor() ([]byte, []int) { return fileDescriptorIoFormats, []int{2} }

// PgCopyOptions describe the format of postgresql's COPY TO STDOUT.
type PgCopyOptions struct {
	// delimiter is the delimitor between columns (DELIMITER)
	Delimiter int32 `protobuf:"varint,1,opt,name=delimiter" json:"delimiter"`
	// null is the NULL value (NULL)
	Null string `protobuf:"bytes,2,opt,name=null" json:"null"`
}

func (m *PgCopyOptions) Reset()                    { *m = PgCopyOptions{} }
func (m *PgCopyOptions) String() string            { return proto.CompactTextString(m) }
func (*PgCopyOptions) ProtoMessage()               {}
func (*PgCopyOptions) Descriptor() ([]byte, []int) { return fileDescriptorIoFormats, []int{3} }

func init() {
	proto.RegisterType((*IOFileFormat)(nil), "cockroach.roachpb.IOFileFormat")
	proto.RegisterType((*CSVOptions)(nil), "cockroach.roachpb.CSVOptions")
	proto.RegisterType((*MySQLOutfileOptions)(nil), "cockroach.roachpb.MySQLOutfileOptions")
	proto.RegisterType((*PgCopyOptions)(nil), "cockroach.roachpb.PgCopyOptions")
	proto.RegisterEnum("cockroach.roachpb.IOFileFormat_FileFormat", IOFileFormat_FileFormat_name, IOFileFormat_FileFormat_value)
	proto.RegisterEnum("cockroach.roachpb.MySQLOutfileOptions_Enclose", MySQLOutfileOptions_Enclose_name, MySQLOutfileOptions_Enclose_value)
}
func (m *IOFileFormat) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *IOFileFormat) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0x8
	i++
	i = encodeVarintIoFormats(dAtA, i, uint64(m.Format))
	dAtA[i] = 0x12
	i++
	i = encodeVarintIoFormats(dAtA, i, uint64(m.Csv.Size()))
	n1, err := m.Csv.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	dAtA[i] = 0x1a
	i++
	i = encodeVarintIoFormats(dAtA, i, uint64(m.MysqlOut.Size()))
	n2, err := m.MysqlOut.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n2
	dAtA[i] = 0x22
	i++
	i = encodeVarintIoFormats(dAtA, i, uint64(m.PgCopy.Size()))
	n3, err := m.PgCopy.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n3
	return i, nil
}

func (m *CSVOptions) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CSVOptions) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0x8
	i++
	i = encodeVarintIoFormats(dAtA, i, uint64(m.Comma))
	dAtA[i] = 0x10
	i++
	i = encodeVarintIoFormats(dAtA, i, uint64(m.Comment))
	if m.NullEncoding != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintIoFormats(dAtA, i, uint64(len(*m.NullEncoding)))
		i += copy(dAtA[i:], *m.NullEncoding)
	}
	dAtA[i] = 0x20
	i++
	i = encodeVarintIoFormats(dAtA, i, uint64(m.Skip))
	return i, nil
}

func (m *MySQLOutfileOptions) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MySQLOutfileOptions) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0x8
	i++
	i = encodeVarintIoFormats(dAtA, i, uint64(m.RowSeparator))
	dAtA[i] = 0x10
	i++
	i = encodeVarintIoFormats(dAtA, i, uint64(m.FieldSeparator))
	dAtA[i] = 0x18
	i++
	i = encodeVarintIoFormats(dAtA, i, uint64(m.Enclose))
	dAtA[i] = 0x20
	i++
	i = encodeVarintIoFormats(dAtA, i, uint64(m.Encloser))
	dAtA[i] = 0x28
	i++
	if m.HasEscape {
		dAtA[i] = 1
	} else {
		dAtA[i] = 0
	}
	i++
	dAtA[i] = 0x30
	i++
	i = encodeVarintIoFormats(dAtA, i, uint64(m.Escape))
	return i, nil
}

func (m *PgCopyOptions) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PgCopyOptions) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0x8
	i++
	i = encodeVarintIoFormats(dAtA, i, uint64(m.Delimiter))
	dAtA[i] = 0x12
	i++
	i = encodeVarintIoFormats(dAtA, i, uint64(len(m.Null)))
	i += copy(dAtA[i:], m.Null)
	return i, nil
}

func encodeVarintIoFormats(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *IOFileFormat) Size() (n int) {
	var l int
	_ = l
	n += 1 + sovIoFormats(uint64(m.Format))
	l = m.Csv.Size()
	n += 1 + l + sovIoFormats(uint64(l))
	l = m.MysqlOut.Size()
	n += 1 + l + sovIoFormats(uint64(l))
	l = m.PgCopy.Size()
	n += 1 + l + sovIoFormats(uint64(l))
	return n
}

func (m *CSVOptions) Size() (n int) {
	var l int
	_ = l
	n += 1 + sovIoFormats(uint64(m.Comma))
	n += 1 + sovIoFormats(uint64(m.Comment))
	if m.NullEncoding != nil {
		l = len(*m.NullEncoding)
		n += 1 + l + sovIoFormats(uint64(l))
	}
	n += 1 + sovIoFormats(uint64(m.Skip))
	return n
}

func (m *MySQLOutfileOptions) Size() (n int) {
	var l int
	_ = l
	n += 1 + sovIoFormats(uint64(m.RowSeparator))
	n += 1 + sovIoFormats(uint64(m.FieldSeparator))
	n += 1 + sovIoFormats(uint64(m.Enclose))
	n += 1 + sovIoFormats(uint64(m.Encloser))
	n += 2
	n += 1 + sovIoFormats(uint64(m.Escape))
	return n
}

func (m *PgCopyOptions) Size() (n int) {
	var l int
	_ = l
	n += 1 + sovIoFormats(uint64(m.Delimiter))
	l = len(m.Null)
	n += 1 + l + sovIoFormats(uint64(l))
	return n
}

func sovIoFormats(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozIoFormats(x uint64) (n int) {
	return sovIoFormats(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *IOFileFormat) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIoFormats
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
			return fmt.Errorf("proto: IOFileFormat: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IOFileFormat: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Format", wireType)
			}
			m.Format = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIoFormats
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Format |= (IOFileFormat_FileFormat(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Csv", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIoFormats
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
				return ErrInvalidLengthIoFormats
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Csv.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MysqlOut", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIoFormats
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
				return ErrInvalidLengthIoFormats
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MysqlOut.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PgCopy", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIoFormats
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
				return ErrInvalidLengthIoFormats
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.PgCopy.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipIoFormats(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthIoFormats
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
func (m *CSVOptions) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIoFormats
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
			return fmt.Errorf("proto: CSVOptions: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CSVOptions: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Comma", wireType)
			}
			m.Comma = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIoFormats
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Comma |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Comment", wireType)
			}
			m.Comment = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIoFormats
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Comment |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NullEncoding", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIoFormats
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
				return ErrInvalidLengthIoFormats
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := string(dAtA[iNdEx:postIndex])
			m.NullEncoding = &s
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Skip", wireType)
			}
			m.Skip = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIoFormats
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Skip |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipIoFormats(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthIoFormats
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
func (m *MySQLOutfileOptions) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIoFormats
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
			return fmt.Errorf("proto: MySQLOutfileOptions: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MySQLOutfileOptions: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RowSeparator", wireType)
			}
			m.RowSeparator = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIoFormats
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RowSeparator |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FieldSeparator", wireType)
			}
			m.FieldSeparator = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIoFormats
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.FieldSeparator |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Enclose", wireType)
			}
			m.Enclose = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIoFormats
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Enclose |= (MySQLOutfileOptions_Enclose(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Encloser", wireType)
			}
			m.Encloser = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIoFormats
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Encloser |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field HasEscape", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIoFormats
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
			m.HasEscape = bool(v != 0)
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Escape", wireType)
			}
			m.Escape = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIoFormats
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Escape |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipIoFormats(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthIoFormats
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
func (m *PgCopyOptions) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIoFormats
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
			return fmt.Errorf("proto: PgCopyOptions: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PgCopyOptions: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Delimiter", wireType)
			}
			m.Delimiter = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIoFormats
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Delimiter |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Null", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIoFormats
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
				return ErrInvalidLengthIoFormats
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Null = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipIoFormats(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthIoFormats
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
func skipIoFormats(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowIoFormats
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
					return 0, ErrIntOverflowIoFormats
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
					return 0, ErrIntOverflowIoFormats
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
				return 0, ErrInvalidLengthIoFormats
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowIoFormats
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
				next, err := skipIoFormats(dAtA[start:])
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
	ErrInvalidLengthIoFormats = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowIoFormats   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("roachpb/io-formats.proto", fileDescriptorIoFormats) }

var fileDescriptorIoFormats = []byte{
	// 553 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0x4f, 0x8f, 0xd2, 0x40,
	0x14, 0x67, 0xf8, 0x57, 0xfa, 0x16, 0xb0, 0x3e, 0x3d, 0x34, 0x1b, 0xad, 0x58, 0x13, 0x83, 0x26,
	0xdb, 0x4d, 0x48, 0x3c, 0x1b, 0x97, 0xb0, 0x71, 0x13, 0x59, 0x74, 0x89, 0x7b, 0xf0, 0x42, 0x6a,
	0x19, 0xa0, 0xa1, 0xed, 0x8c, 0x9d, 0xb2, 0x84, 0x0f, 0x61, 0xe2, 0xc7, 0xe2, 0xe8, 0xc1, 0x83,
	0x27, 0xa3, 0x78, 0xf5, 0x43, 0x98, 0x4e, 0xa7, 0x02, 0x91, 0xc3, 0xde, 0x26, 0xbf, 0x7f, 0xef,
	0xf5, 0xf7, 0x0a, 0x66, 0xcc, 0x5c, 0x6f, 0xc6, 0x3f, 0x9e, 0xfa, 0xec, 0x64, 0xc2, 0xe2, 0xd0,
	0x4d, 0x84, 0xc3, 0x63, 0x96, 0x30, 0xbc, 0xeb, 0x31, 0x6f, 0x2e, 0x59, 0x47, 0x69, 0x8e, 0xef,
	0x4f, 0xd9, 0x94, 0x49, 0xf6, 0x34, 0x7d, 0x65, 0x42, 0xfb, 0x4f, 0x11, 0xea, 0x17, 0x83, 0x73,
	0x3f, 0xa0, 0xe7, 0x32, 0x00, 0x5f, 0x43, 0x35, 0x8b, 0x32, 0x49, 0x8b, 0xb4, 0x9b, 0x9d, 0xe7,
	0xce, 0x7f, 0x51, 0xce, 0xae, 0xc1, 0xd9, 0x3e, 0xcf, 0xca, 0xeb, 0x1f, 0x8f, 0x0a, 0x57, 0xca,
	0x8f, 0x2f, 0xa0, 0xe4, 0x89, 0x1b, 0xb3, 0xd8, 0x22, 0xed, 0xa3, 0xce, 0xc3, 0x03, 0x31, 0xdd,
	0xe1, 0xf5, 0x80, 0x27, 0x3e, 0x8b, 0x84, 0x72, 0xa6, 0x7a, 0xbc, 0x00, 0x3d, 0x5c, 0x89, 0x4f,
	0xc1, 0x88, 0x2d, 0x12, 0xb3, 0x24, 0xcd, 0x4f, 0x0f, 0x98, 0xfb, 0xab, 0xe1, 0xbb, 0x37, 0x83,
	0x45, 0x32, 0xf1, 0x03, 0xba, 0x9f, 0x52, 0x93, 0xf6, 0xc1, 0x22, 0xc1, 0x97, 0xa0, 0xf1, 0xe9,
	0xc8, 0x63, 0x7c, 0x65, 0x96, 0x65, 0x50, 0xeb, 0x40, 0xd0, 0xdb, 0x69, 0x97, 0xf1, 0xd5, 0x7e,
	0x44, 0x95, 0x4b, 0xd0, 0x1e, 0x00, 0xec, 0x54, 0x73, 0x04, 0xda, 0xfb, 0x68, 0x1e, 0xb1, 0x65,
	0x64, 0x14, 0x50, 0x83, 0x52, 0x77, 0x78, 0x6d, 0x10, 0x34, 0xa0, 0xde, 0x57, 0x03, 0xd3, 0x5d,
	0x8c, 0x22, 0x36, 0x40, 0x97, 0xc8, 0x78, 0x11, 0x72, 0xa3, 0x84, 0x00, 0xd5, 0x6c, 0x86, 0x51,
	0xb6, 0x3f, 0x13, 0x80, 0xed, 0x67, 0xe3, 0x31, 0x54, 0x3c, 0x16, 0x86, 0xae, 0xec, 0xba, 0xa2,
	0x86, 0x67, 0x10, 0x5a, 0xa0, 0xa5, 0x0f, 0x1a, 0x25, 0xb2, 0xc2, 0x9c, 0xcd, 0x41, 0x7c, 0x06,
	0x8d, 0x68, 0x11, 0x04, 0x23, 0x1a, 0x79, 0x6c, 0xec, 0x47, 0x53, 0xd9, 0x95, 0x2e, 0x55, 0xe4,
	0xaa, 0x9e, 0x52, 0x3d, 0xc5, 0xa0, 0x09, 0x65, 0x31, 0xf7, 0xb9, 0x2c, 0xa1, 0xa1, 0x72, 0x24,
	0x62, 0x7f, 0x2b, 0xc2, 0xbd, 0x03, 0x4d, 0xa6, 0xe1, 0x31, 0x5b, 0x8e, 0x04, 0xe5, 0x6e, 0xec,
	0x26, 0x2c, 0xde, 0x5b, 0xb0, 0x1e, 0xb3, 0xe5, 0x30, 0x67, 0xf0, 0x04, 0xee, 0x4c, 0x7c, 0x1a,
	0x8c, 0x77, 0xc4, 0xbb, 0xfb, 0x36, 0x25, 0xb9, 0x95, 0x5f, 0x82, 0x46, 0x23, 0x2f, 0x60, 0x82,
	0xca, 0x85, 0x9b, 0x1d, 0xe7, 0x76, 0xc7, 0x75, 0x7a, 0x99, 0x2b, 0xaf, 0x41, 0x85, 0x60, 0x0b,
	0x6a, 0xea, 0x19, 0xcb, 0xef, 0xcb, 0xe7, 0xfe, 0x43, 0xf1, 0x09, 0xc0, 0xcc, 0x15, 0x23, 0x2a,
	0x3c, 0x97, 0x53, 0xb3, 0xd2, 0x22, 0xed, 0x9a, 0xd2, 0xe8, 0x33, 0x57, 0xf4, 0x24, 0x8c, 0x0f,
	0xa0, 0xaa, 0x04, 0xd5, 0x9d, 0x10, 0x85, 0xd9, 0x0e, 0x68, 0x6a, 0x3c, 0xea, 0x50, 0xb9, 0xa4,
	0x37, 0x34, 0x36, 0x0a, 0xe9, 0x61, 0x5f, 0x05, 0x4b, 0x77, 0x25, 0x0c, 0x82, 0x75, 0xa8, 0x65,
	0x8b, 0xba, 0x81, 0x51, 0xb4, 0xfb, 0xd0, 0xd8, 0xfb, 0xad, 0xd0, 0x06, 0x7d, 0x4c, 0x03, 0x3f,
	0xf4, 0x13, 0xba, 0xdf, 0xe5, 0x16, 0x4e, 0xaf, 0x94, 0x5e, 0x4d, 0xb6, 0xa7, 0xe7, 0x57, 0x4a,
	0x91, 0xb3, 0xc7, 0xeb, 0x5f, 0x56, 0x61, 0xbd, 0xb1, 0xc8, 0xd7, 0x8d, 0x45, 0xbe, 0x6f, 0x2c,
	0xf2, 0x73, 0x63, 0x91, 0x2f, 0xbf, 0xad, 0xc2, 0x07, 0x4d, 0x35, 0xf6, 0x37, 0x00, 0x00, 0xff,
	0xff, 0x41, 0x56, 0x6d, 0xab, 0x0b, 0x04, 0x00, 0x00,
}
