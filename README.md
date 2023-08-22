# buffer
支持内存大小端读写的固定数据缓冲区
A fixed data buffer that supports reads and writes at the memory size end

用于高并发下的[]Byte缓存池复用
Used for []Byte cache pool multiplexing under high concurrency

# Example

``` golang
var buf = NewByteBuffer(1024, binary.LittleEndian)

buf.WriteString("123456!@#$%^asdfgh")
buf.WriteInt32(0)
buf.WriteFloat32(3.1415926)
buf.WriteInt32(12345678)

buf.SetPosition(0)

buf.ReadString(-1)
buf.ReadInt32()
buf.ReadFloat32()
buf.ReadInt32()

buf.Reset()
buf.Reset()

```

# API

##### NewByteBuffer(int, binary.ByteOrder) *ByteBuffer

--- 
## Base Method
##### ByteBuffer::Cap() int
##### ByteBuffer::Length() int
##### ByteBuffer::Buffer() []byte
##### ByteBuffer::Data() []byte
##### ByteBuffer::Reset()
##### ByteBuffer::SetOrder()
##### ByteBuffer::SetPosition(value int) error
### ByteBuffer::Position() int

---
## Write

##### ByteBuffer::PutUInt16(offset int, value uint16) error
##### ByteBuffer::PutUInt32(offset int, value uint32) error
##### ByteBuffer::PutUInt64(offset int, value uint64) error

---

## Read
##### ByteBuffer::GetUInt16(offset int) (value uint16, err error)
##### ByteBuffer::GetUInt32(offset int) (value uint32, err error)
##### ByteBuffer::GetUInt64(offset int) (value uint64, err error)





--- 
## Writer API
##### ByteBuffer::WriteInt8(value byte) error
##### ByteBuffer::WriteBoolean(value bool) error
##### ByteBuffer::WriteString(value string) (wlen int, err error)
##### ByteBuffer::WriteBytes(value []byte) (int, error)
##### ByteBuffer::WriteInt16(value int16) error
##### ByteBuffer::WriteUInt16(value uint16) error
##### ByteBuffer::WriteInt32(value int32) error
##### ByteBuffer::WriteUInt32(value uint32) error
##### ByteBuffer::WriteInt64(value int64) error
##### ByteBuffer::WriteUInt64(value uint64) error
##### ByteBuffer::WriteFloat32(value float32) error
##### ByteBuffer::WriteFloat64(value float64) error

--- 
## Reader API

##### ByteBuffer::ReadInt8() (byte, error)
##### ByteBuffer::ReadBoolean() (bool, error)
##### ByteBuffer::ReadString(length int) (string, error) 
##### ByteBuffer::ReadBytes(value []byte, length int) (int, error)
##### ByteBuffer::ReadInt16() (int16, error)
##### ByteBuffer::ReadUInt16() (uint16, error)
##### ByteBuffer::ReadInt32() (int32, error)
##### ByteBuffer::ReadUInt32() (uint32, error)
##### ByteBuffer::ReadInt64() (int64, error)
##### ByteBuffer::ReadUInt64() (uint64, error)
##### ByteBuffer::ReadFloat32() (float32, error)
##### ByteBuffer::ReadFloat64() (float64, error)

