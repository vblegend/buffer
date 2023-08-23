package buffer

import (
	"encoding/binary"
	"errors"
	"math"
)

const (
	TRUE  = 1
	FALSE = 0
)

const (
	SIZEOF_1_BYTE = 1
	SIZEOF_2_BYTE = 2
	SIZEOF_4_BYTE = 4
	SIZEOF_8_BYTE = 8
)

// ErrorWriteOverflow
// The data write pointer is out of bounds
var ErrorWriteOverflow = errors.New("the data write pointer is out of bounds")

// ErrorReadOverflow
// The data read pointer is out of bounds
var ErrorReadOverflow = errors.New("the data read pointer is out of bounds")

// ErrorPositionOverflow
// Invalid data access pointer “offset”
var ErrorPositionOverflow = errors.New("invalid data access pointer “offset”")

// NewByteBuffer
// Create a fixed-size memory buffer
// order to determine how the data is written, should you use either binary.LittleEndian or binary.BigEndian
func NewByteBuffer(cap int, order binary.ByteOrder) *ByteBuffer {
	buf := ByteBuffer{
		buf:    make([]byte, cap),
		order:  order,
		cap:    cap,
		length: 0,
		offset: 0,
	}
	return &buf
}

// ByteBuffer
// A fixed-size buffer of read-write data
type ByteBuffer struct {
	cap    int              // buffer capacity
	buf    []byte           // buffer
	length int              // Effective data length
	offset int              // Read/write position
	order  binary.ByteOrder // Byte arrangement
}

// Length
// Gets the buffer effective data length
func (buf *ByteBuffer) Length() int {
	return buf.length
}

// Cap
// Get buffer size
func (buf *ByteBuffer) Cap() int {
	return buf.cap
}

// Position Get the read/write offset
func (buf *ByteBuffer) Position() int {
	return buf.offset
}

// SetPosition Set the read/write offset
func (buf *ByteBuffer) SetPosition(asbOffset int) error {
	if asbOffset < 0 || asbOffset > buf.length {
		return ErrorPositionOverflow
	}
	buf.offset = asbOffset
	return nil
}

// Reset
// Clear the read/write index and data in the buffer
func (buf *ByteBuffer) Reset() {
	buf.length = 0
	buf.offset = 0
}

// SetOrder
// Set the read and write mode of the buffer
// Used binary.LittleEndian or binary.BigEndian
func (buf *ByteBuffer) SetOrder(order binary.ByteOrder) {
	buf.order = order
}

// Buffer
// Gets the buffer raw byte array
func (buf *ByteBuffer) Buffer() []byte {
	return buf.buf
}

// Data
// Gets valid data in the buffer
func (buf *ByteBuffer) Data() []byte {
	return buf.buf[:buf.length]
}

func (buf *ByteBuffer) applyWrite(size int) (offset int, ok bool) {
	offset = buf.offset
	if buf.offset+size > buf.cap {
		ok = false
		return
	}
	ok = true
	buf.offset += size
	if buf.offset > buf.length {
		buf.length = buf.offset
	}
	return
}

func (buf *ByteBuffer) applyRead(size int) (offset int, ok bool) {
	offset = buf.offset
	if buf.offset+size > buf.length {
		ok = false
		return
	}
	ok = true
	buf.offset += size
	return
}

// **********************************************************************
// 写入数据
// **********************************************************************

// WriteInt8
// Writes an int8 type value to the current location
func (buf *ByteBuffer) WriteInt8(value byte) error {
	offset, ok := buf.applyWrite(SIZEOF_1_BYTE)
	if !ok {
		return ErrorWriteOverflow
	}
	buf.buf[offset] = value
	return nil
}

// WriteBoolean
// Writes a 1-byte Boolean variable at the current location
func (buf *ByteBuffer) WriteBoolean(value bool) error {
	if value {
		return buf.WriteInt8(TRUE)
	}
	return buf.WriteInt8(FALSE)
}

// WriteString
// Writes a String at the current location
func (buf *ByteBuffer) WriteString(value string) (wlen int, err error) {
	return buf.WriteBytes([]byte(value))
}

// WriteBytes
// Writes a byte array at the current location and returns the length of the data successfully written
func (buf *ByteBuffer) WriteBytes(value []byte) (int, error) {
	offset, ok := buf.applyWrite(len(value))
	if !ok {
		return 0, ErrorWriteOverflow
	}
	return copy(buf.buf[offset:], value), nil
}

// WriteInt16
// Writes an int16 type value to the current location
func (buf *ByteBuffer) WriteInt16(value int16) error {
	offset, ok := buf.applyWrite(SIZEOF_2_BYTE)
	if !ok {
		return ErrorWriteOverflow
	}
	buf.order.PutUint16(buf.buf[offset:], uint16(value))
	return nil
}

// WriteUInt16
// Writes a uint16 value at the current location
func (buf *ByteBuffer) WriteUInt16(value uint16) error {
	offset, ok := buf.applyWrite(SIZEOF_2_BYTE)
	if !ok {
		return ErrorWriteOverflow
	}
	buf.order.PutUint16(buf.buf[offset:], value)
	return nil
}

// WriteInt32
// Writes an int32 type value to the current location
func (buf *ByteBuffer) WriteInt32(value int32) error {
	offset, ok := buf.applyWrite(SIZEOF_4_BYTE)
	if !ok {
		return ErrorWriteOverflow
	}
	buf.order.PutUint32(buf.buf[offset:], uint32(value))
	return nil
}

// WriteUInt32
// Writes an UInt32 type value to the current location
func (buf *ByteBuffer) WriteUInt32(value uint32) error {
	offset, ok := buf.applyWrite(SIZEOF_4_BYTE)
	if !ok {
		return ErrorWriteOverflow
	}
	buf.order.PutUint32(buf.buf[offset:], value)
	return nil
}

// WriteInt64
// Writes an int64 type value to the current location
func (buf *ByteBuffer) WriteInt64(value int64) error {
	offset, ok := buf.applyWrite(SIZEOF_8_BYTE)
	if !ok {
		return ErrorWriteOverflow
	}
	buf.order.PutUint64(buf.buf[offset:], uint64(value))
	return nil
}

// WriteUInt64
// Writes an UInt64 type value to the current location
func (buf *ByteBuffer) WriteUInt64(value uint64) error {
	offset, ok := buf.applyWrite(SIZEOF_8_BYTE)
	if !ok {
		return ErrorWriteOverflow
	}
	buf.order.PutUint64(buf.buf[offset:], value)
	return nil
}

// WriteFloat32
// Writes an Float32 type value to the current location
func (buf *ByteBuffer) WriteFloat32(value float32) error {
	offset, ok := buf.applyWrite(SIZEOF_4_BYTE)
	if !ok {
		return ErrorWriteOverflow
	}
	buf.order.PutUint32(buf.buf[offset:], math.Float32bits(value))
	return nil
}

// WriteFloat64
// Writes an Float64 type value to the current location
func (buf *ByteBuffer) WriteFloat64(value float64) error {
	offset, ok := buf.applyWrite(SIZEOF_8_BYTE)
	if !ok {
		return ErrorWriteOverflow
	}
	buf.order.PutUint64(buf.buf[offset:], math.Float64bits(value))
	return nil
}

// **********************************************************************
// 读取数据
// **********************************************************************

// ReadInt8
// Reads a value of type int8 from the current position
func (buf *ByteBuffer) ReadInt8() (byte, error) {
	offset, ok := buf.applyRead(SIZEOF_1_BYTE)
	if !ok {
		return 0, ErrorReadOverflow
	}
	value := buf.buf[offset]
	return value, nil
}

// ReadBoolean
// Reads a 1-byte Boolean variable from the current location
func (buf *ByteBuffer) ReadBoolean() (bool, error) {
	value, ok := buf.ReadInt8()
	return value == TRUE, ok
}

// ReadString
// Reads a String from the current position
// Success is returned when length is 0, but not read
// When length is less than 0, it reads from the current position until it reaches "0" and returns the result, otherwise all subsequent data is returned
func (buf *ByteBuffer) ReadString(length int) (string, error) {
	if length == 0 {
		return "", nil
	}
	if length <= 0 {
		startIndex := buf.offset
		endIndex := buf.length
		for i := buf.offset; i < buf.length; i++ {
			if buf.buf[i] == 0 {
				endIndex = i
				break
			}
		}
		buf.offset = endIndex
		return string(buf.buf[startIndex:endIndex]), nil
	}

	if buf.offset+length > buf.length {
		// 错误的长度
		return "", ErrorReadOverflow
	}
	offset, ok := buf.applyRead(length)
	if !ok {
		return "", ErrorReadOverflow
	}
	b := buf.buf[offset : offset+length]
	return string(b), nil
}

// ReadBytes
// Success is returned when length is 0, but not read
// Reads a byte array from the current location
// When length is less than 0, the length is the length of the array "value"
// Returns the length of the data actually read
func (buf *ByteBuffer) ReadBytes(value []byte, length int) (int, error) {
	if length == 0 {
		return 0, nil
	}
	if length < 0 {
		length = len(value)
	} else if length > len(value) {
		// 错误的长度
		return 0, ErrorReadOverflow
	}
	offset, ok := buf.applyRead(length)
	if !ok {
		return 0, ErrorReadOverflow
	}
	source := buf.buf[offset : offset+length]
	return copy(value, source), nil
}

// ReadInt16
// Reads a value of type int16 from the current position
func (buf *ByteBuffer) ReadInt16() (int16, error) {
	offset, ok := buf.applyRead(SIZEOF_2_BYTE)
	if !ok {
		return 0, ErrorReadOverflow
	}
	value := buf.order.Uint16(buf.buf[offset:])
	return int16(value), nil
}

// ReadUInt16
// Reads a value of type UInt16 from the current position
func (buf *ByteBuffer) ReadUInt16() (uint16, error) {
	offset, ok := buf.applyRead(SIZEOF_2_BYTE)
	if !ok {
		return 0, ErrorReadOverflow
	}
	value := buf.order.Uint16(buf.buf[offset:])
	return value, nil
}

// ReadInt32
// Reads a value of type Int32 from the current position
func (buf *ByteBuffer) ReadInt32() (int32, error) {
	offset, ok := buf.applyRead(SIZEOF_4_BYTE)
	if !ok {
		return 0, ErrorReadOverflow
	}
	value := buf.order.Uint32(buf.buf[offset:])
	return int32(value), nil
}

// ReadUInt32
// Reads a value of type UInt32 from the current position
func (buf *ByteBuffer) ReadUInt32() (uint32, error) {
	offset, ok := buf.applyRead(SIZEOF_4_BYTE)
	if !ok {
		return 0, ErrorReadOverflow
	}
	value := buf.order.Uint32(buf.buf[offset:])
	return value, nil
}

// ReadInt64
// Reads a value of type Int64 from the current position
func (buf *ByteBuffer) ReadInt64() (int64, error) {
	offset, ok := buf.applyRead(SIZEOF_8_BYTE)
	if !ok {
		return 0, ErrorReadOverflow
	}
	value := buf.order.Uint64(buf.buf[offset:])
	return int64(value), nil
}

// ReadUInt64
// Reads a value of type UInt64 from the current position
func (buf *ByteBuffer) ReadUInt64() (uint64, error) {
	offset, ok := buf.applyRead(SIZEOF_8_BYTE)
	if !ok {
		return 0, ErrorReadOverflow
	}
	value := buf.order.Uint64(buf.buf[offset:])
	return value, nil
}

// ReadFloat32
// Reads a value of type Float32 from the current position
func (buf *ByteBuffer) ReadFloat32() (float32, error) {
	offset, ok := buf.applyRead(SIZEOF_4_BYTE)
	if !ok {
		return 0, ErrorReadOverflow
	}
	value := buf.order.Uint32(buf.buf[offset:])
	return math.Float32frombits(value), nil
}

// ReadFloat64
// Reads a value of type Float64 from the current position
func (buf *ByteBuffer) ReadFloat64() (float64, error) {
	offset, ok := buf.applyRead(SIZEOF_8_BYTE)
	if !ok {
		return 0, ErrorReadOverflow
	}
	value := buf.order.Uint64(buf.buf[offset:])
	return math.Float64frombits(value), nil
}

// **********************************************************************
// 不改变读/写位置情况下 写入数据
// **********************************************************************

// PutUInt16
// Fill in a uint16 value of type at the specified position offset
// This operation does not change the location of the read/write index
func (buf *ByteBuffer) PutUInt16(offset int, value uint16) error {
	if offset < 0 || offset+SIZEOF_2_BYTE > buf.length {
		return ErrorPositionOverflow
	}
	buf.order.PutUint16(buf.buf[offset:], value)
	return nil
}

// PutUInt32
// Enter a uint32 type value at the specified position offset
// This operation does not change the location of the read/write index
func (buf *ByteBuffer) PutUInt32(offset int, value uint32) error {
	if offset < 0 || (offset+SIZEOF_4_BYTE) > buf.length {
		return ErrorPositionOverflow
	}
	buf.order.PutUint32(buf.buf[offset:], value)
	return nil
}

// PutUInt64
// Enter a value of type uint64 at the specified position offset
// This operation does not change the location of the read/write index
func (buf *ByteBuffer) PutUInt64(offset int, value uint64) error {
	if offset < 0 || (offset+SIZEOF_8_BYTE) > buf.length {
		return ErrorPositionOverflow
	}
	buf.order.PutUint64(buf.buf[offset:], value)
	return nil
}

// **********************************************************************
// 不改变读/写位置情况下 读取数据
// **********************************************************************

// GetUInt16
// Reads a uint16 value from the specified position offset
// This operation does not change the location of the read/write index
func (buf *ByteBuffer) GetUInt16(offset int) (value uint16, err error) {
	if offset < 0 || (offset+SIZEOF_2_BYTE) > buf.length {
		err = ErrorPositionOverflow
		return
	}
	value = buf.order.Uint16(buf.buf[offset:])
	return
}

// GetUInt32
// Reads a uint32 type value from the specified position offset
// This operation does not change the location of the read/write index
func (buf *ByteBuffer) GetUInt32(offset int) (value uint32, err error) {
	if offset < 0 || (offset+SIZEOF_4_BYTE) > buf.length {
		err = ErrorPositionOverflow
		return
	}
	value = buf.order.Uint32(buf.buf[offset:])
	return
}

// GetUInt64
// Reads a value of type uint64 from the specified position offset
// This operation does not change the location of the read/write index
func (buf *ByteBuffer) GetUInt64(offset int) (value uint64, err error) {
	if offset < 0 || (offset+SIZEOF_8_BYTE) > buf.length {
		err = ErrorPositionOverflow
		return
	}
	value = buf.order.Uint64(buf.buf[offset:])
	return
}
