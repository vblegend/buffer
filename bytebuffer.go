package buffer

import (
	"encoding/binary"
	"errors"
	"math"
)

// ErrorWriteOverflow 写入指针溢出
var ErrorWriteOverflow = errors.New("Data written to the index exceeds the limit")

// ErrorReadOverflow 读取指针溢出
var ErrorReadOverflow = errors.New("Data Read to the index exceeds the limit")

// ErrorPositionOverflow 读取指针溢出
var ErrorPositionOverflow = errors.New("Data Read to the index exceeds the limit")

// NewByteBuffer 创建一个内存缓冲区
// order 决定数据以什么方式写入，应使用 binary.LittleEndian 或 binary.BigEndian 其中之一
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

// ByteBuffer 固定大小的可读写数据缓冲区
type ByteBuffer struct {
	cap    int              // 缓冲区容量
	buf    []byte           // 缓冲区数据
	length int              // 有效数据长度
	offset int              // 读/写 位置
	order  binary.ByteOrder //字节排列方式
}

// Length 获取缓存有效数据长度
func (buf *ByteBuffer) Length() int {
	return buf.length
}

// Cap 获取缓存大小
func (buf *ByteBuffer) Cap() int {
	return buf.cap
}

// Position 获取 读取/写入位置
func (buf *ByteBuffer) Position() int {
	return buf.offset
}

// SetPosition 设置 读取/写入位置
func (buf *ByteBuffer) SetPosition(asbOffset int) error {
	if asbOffset < 0 || asbOffset > buf.length {
		return ErrorPositionOverflow
	}
	buf.offset = asbOffset
	return nil
}

// Reset 清除缓冲区内数据并使读写索引归0
func (buf *ByteBuffer) Reset() {
	buf.length = 0
	buf.offset = 0
}

// SetOrder 设置缓冲区的大小端读写方式
// Used binary.LittleEndian or binary.BigEndian
func (buf *ByteBuffer) SetOrder(order binary.ByteOrder) {
	buf.order = order
}

// Buffer 缓冲区
func (buf *ByteBuffer) Buffer() []byte {
	return buf.buf
}

// Data 有效数据
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

// WriteInt8 在当前位置写入一个int8类型数值
func (buf *ByteBuffer) WriteInt8(value byte) error {
	offset, ok := buf.applyWrite(1)
	if !ok {
		return ErrorWriteOverflow
	}
	buf.buf[offset] = value
	return nil
}

// WriteBoolean 在当前位置写入一个1字节的布尔变量
func (buf *ByteBuffer) WriteBoolean(value bool) error {
	if value {
		return buf.WriteInt8(1)
	}
	return buf.WriteInt8(0)
}

// WriteString 在当前位置写入一个String字符串
func (buf *ByteBuffer) WriteString(value string) (wlen int, err error) {
	wlen = len(value)
	offset, ok := buf.applyWrite(wlen)
	if !ok {
		wlen = 0
		err = ErrorWriteOverflow
		return
	}
	copy(buf.buf[offset:], value)
	return
}

// WriteBytes 在当前位置写入一个byte数组
func (buf *ByteBuffer) WriteBytes(value []byte) error {
	offset, ok := buf.applyWrite(len(value))
	if !ok {
		return ErrorWriteOverflow
	}
	copy(buf.buf[offset:], value)
	return nil
}

// WriteInt16 在当前位置写入一个int16类型数值
func (buf *ByteBuffer) WriteInt16(value int16) error {
	offset, ok := buf.applyWrite(2)
	if !ok {
		return ErrorWriteOverflow
	}
	buf.order.PutUint16(buf.buf[offset:], uint16(value))
	return nil
}

// WriteUInt16 在当前位置写入一个uint16类型数值
func (buf *ByteBuffer) WriteUInt16(value uint16) error {
	offset, ok := buf.applyWrite(2)
	if !ok {
		return ErrorWriteOverflow
	}
	buf.order.PutUint16(buf.buf[offset:], value)
	return nil
}

// WriteInt32 在当前位置写入一个int32类型数值
func (buf *ByteBuffer) WriteInt32(value int32) error {
	offset, ok := buf.applyWrite(4)
	if !ok {
		return ErrorWriteOverflow
	}
	buf.order.PutUint32(buf.buf[offset:], uint32(value))
	return nil
}

// WriteUInt32 在当前位置写入一个uint32类型数值
func (buf *ByteBuffer) WriteUInt32(value uint32) error {
	offset, ok := buf.applyWrite(4)
	if !ok {
		return ErrorWriteOverflow
	}
	buf.order.PutUint32(buf.buf[offset:], value)
	return nil
}

// WriteInt64 在当前位置写入一个int64类型数值
func (buf *ByteBuffer) WriteInt64(value int64) error {
	offset, ok := buf.applyWrite(8)
	if !ok {
		return ErrorWriteOverflow
	}
	buf.order.PutUint64(buf.buf[offset:], uint64(value))
	return nil
}

// WriteUInt64 在当前位置写入一个uint64类型数值
func (buf *ByteBuffer) WriteUInt64(value uint64) error {
	offset, ok := buf.applyWrite(8)
	if !ok {
		return ErrorWriteOverflow
	}
	buf.order.PutUint64(buf.buf[offset:], value)
	return nil
}

// WriteFloat32 在当前位置写入一个float32类型数值
func (buf *ByteBuffer) WriteFloat32(value float32) error {
	offset, ok := buf.applyWrite(4)
	if !ok {
		return ErrorWriteOverflow
	}
	buf.order.PutUint32(buf.buf[offset:], math.Float32bits(value))
	return nil
}

// WriteFloat64 在当前位置写入一个float64类型数值
func (buf *ByteBuffer) WriteFloat64(value float64) error {
	offset, ok := buf.applyWrite(8)
	if !ok {
		return ErrorWriteOverflow
	}
	buf.order.PutUint64(buf.buf[offset:], math.Float64bits(value))
	return nil
}

// **********************************************************************
// 读取数据
// **********************************************************************

// ReadInt8 从当前位置读取一个int8类型数值
func (buf *ByteBuffer) ReadInt8() (byte, error) {
	offset, ok := buf.applyRead(1)
	if !ok {
		return 0, ErrorReadOverflow
	}
	value := buf.buf[offset]
	return value, nil
}

// ReadBoolean 从当前位置读取一个1字节的布尔变量
func (buf *ByteBuffer) ReadBoolean() (bool, error) {
	value, ok := buf.ReadInt8()
	return value == 1, ok
}

// ReadString 从当前位置读取一个String字符串
func (buf *ByteBuffer) ReadString(length int) (string, error) {
	offset, ok := buf.applyRead(length)
	if !ok {
		return "", ErrorReadOverflow
	}
	b := buf.buf[offset : offset+length]
	return string(b), nil
}

// ReadBytes 从当前位置读取一个byte数组
// length 小于0 时 长度为 len(value)
func (buf *ByteBuffer) ReadBytes(value []byte, length int) error {
	if length < 0 {
		length = len(value)
	} else if length > len(value) {
		// 错误的长度
		return ErrorReadOverflow
	}
	offset, ok := buf.applyRead(length)
	if !ok {
		return ErrorReadOverflow
	}
	source := buf.buf[offset : offset+length]
	copy(value, source)
	return nil
}

// ReadInt16 从当前位置读取一个int16类型数值
func (buf *ByteBuffer) ReadInt16() (int16, error) {
	offset, ok := buf.applyRead(2)
	if !ok {
		return 0, ErrorReadOverflow
	}
	value := buf.order.Uint16(buf.buf[offset:])
	return int16(value), nil
}

// ReadUInt16 从当前位置读取一个uint16类型数值
func (buf *ByteBuffer) ReadUInt16() (uint16, error) {
	offset, ok := buf.applyRead(2)
	if !ok {
		return 0, ErrorReadOverflow
	}
	value := buf.order.Uint16(buf.buf[offset:])
	return value, nil
}

// ReadInt32 从当前位置读取一个int32类型数值
func (buf *ByteBuffer) ReadInt32() (int32, error) {
	offset, ok := buf.applyRead(4)
	if !ok {
		return 0, ErrorReadOverflow
	}
	value := buf.order.Uint32(buf.buf[offset:])
	return int32(value), nil
}

// ReadUInt32 从当前位置读取一个uint32类型数值
func (buf *ByteBuffer) ReadUInt32() (uint32, error) {
	offset, ok := buf.applyRead(4)
	if !ok {
		return 0, ErrorReadOverflow
	}
	value := buf.order.Uint32(buf.buf[offset:])
	return value, nil
}

// ReadInt64 从当前位置读取一个int64类型数值
func (buf *ByteBuffer) ReadInt64() (int64, error) {
	offset, ok := buf.applyRead(8)
	if !ok {
		return 0, ErrorReadOverflow
	}
	value := buf.order.Uint64(buf.buf[offset:])
	return int64(value), nil
}

// ReadUInt64 从当前位置读取一个uint64类型数值
func (buf *ByteBuffer) ReadUInt64() (uint64, error) {
	offset, ok := buf.applyRead(8)
	if !ok {
		return 0, ErrorReadOverflow
	}
	value := buf.order.Uint64(buf.buf[offset:])
	return value, nil
}

// ReadFloat32 从当前位置读取一个float32类型数值
func (buf *ByteBuffer) ReadFloat32() (float32, error) {
	offset, ok := buf.applyRead(4)
	if !ok {
		return 0, ErrorReadOverflow
	}
	value := buf.order.Uint32(buf.buf[offset:])
	return math.Float32frombits(value), nil
}

// ReadFloat64 从当前位置读取一个float64类型数值
func (buf *ByteBuffer) ReadFloat64() (float64, error) {
	offset, ok := buf.applyRead(8)
	if !ok {
		return 0, ErrorReadOverflow
	}
	value := buf.order.Uint64(buf.buf[offset:])
	return math.Float64frombits(value), nil
}

// **********************************************************************
// 不改变读/写位置情况下 写入数据
// **********************************************************************

// PutUInt16 在指定位置offset处填入一个uint16类型数值
// 该操作不会改变读写索引的位置
func (buf *ByteBuffer) PutUInt16(offset int, value uint16) error {
	if offset < 0 || offset+2 > buf.length {
		return ErrorPositionOverflow
	}
	buf.order.PutUint16(buf.buf[offset:], value)
	return nil
}

// PutUInt32 在指定位置offset处填入一个uint32类型数值
// 该操作不会改变读写索引的位置
func (buf *ByteBuffer) PutUInt32(offset int, value uint32) error {
	if offset < 0 || (offset+4) > buf.length {
		return ErrorPositionOverflow
	}
	buf.order.PutUint32(buf.buf[offset:], value)
	return nil
}

// PutUInt64 在指定位置offset处填入一个uint64类型数值
// 该操作不会改变读写索引的位置
func (buf *ByteBuffer) PutUInt64(offset int, value uint64) error {
	if offset < 0 || (offset+8) > buf.length {
		return ErrorPositionOverflow
	}
	buf.order.PutUint64(buf.buf[offset:], value)
	return nil
}

// **********************************************************************
// 不改变读/写位置情况下 读取数据
// **********************************************************************

// GetUInt16 从指定位置offset处读取一个uint16类型数值
// 该操作不会改变读写索引的位置
func (buf *ByteBuffer) GetUInt16(offset int) (value uint16, err error) {
	if offset < 0 || (offset+2) > buf.length {
		err = ErrorPositionOverflow
		return
	}
	value = buf.order.Uint16(buf.buf[offset:])
	return
}

// GetUInt32 从指定位置offset处读取一个uint32类型数值
// 该操作不会改变读写索引的位置
func (buf *ByteBuffer) GetUInt32(offset int) (value uint32, err error) {
	if offset < 0 || (offset+4) > buf.length {
		err = ErrorPositionOverflow
		return
	}
	value = buf.order.Uint32(buf.buf[offset:])
	return
}

// GetUInt64 从指定位置offset处读取一个uint64类型数值
// 该操作不会改变读写索引的位置
func (buf *ByteBuffer) GetUInt64(offset int) (value uint64, err error) {
	if offset < 0 || (offset+8) > buf.length {
		err = ErrorPositionOverflow
		return
	}
	value = buf.order.Uint64(buf.buf[offset:])
	return
}
