# buffer
支持内存大小端读写的固定数据缓冲区，实现了数据读写API
Support memory size end of the read and write fixed data buffer, the implementation of the data read and write API


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
buf.Data()
buf.Reset()


```

# API

##### NewByteBuffer(cap int, order binary.ByteOrder) *ByteBuffer
创建一个固定大小的内存缓冲区
Create a fixed-size memory buffer

参数cap指定缓冲区的最大容量
The cap parameter specifies the maximum capacity of the buffer

参数order指定缓冲区读写操作以哪种模式  `binary.LittleEndian` or `binary.BigEndian`
The order parameter specifies the mode in which buffer read and write operations are performed `binary.LittleEndian` or `binary.BigEndian`

--- 
## Base Method
##### ByteBuffer::Cap() int
获取缓冲区容量
Gets the capacity of the buffer
##### ByteBuffer::Length() int
获取缓冲区数据长度
Gets the data length of the buffer
##### ByteBuffer::Buffer() []byte
获取缓冲区byte数组对象
Gets a buffer byte array object
##### ByteBuffer::Data() []byte
获取缓冲区有效数据
Get buffer valid data
##### ByteBuffer::Reset()
重置缓冲区内数据并清除读/写位置
Reset get buffer data and clear read/write index
##### ByteBuffer::SetOrder()
设置缓冲区的读写模式
Set the buffer read and write mode

使用 `binary.LittleEndian` 或 `binary.BigEndian`
Used `binary.LittleEndian` or `binary.BigEndian`
##### ByteBuffer::SetPosition(value int) error
设置缓冲区当前读写位置
Sets the current buffer read/write location
##### ByteBuffer::Position() int
获取缓冲区当前的读写位置
Gets the current read/write location of the buffer

---
## Write

##### ByteBuffer::PutUInt16(offset int, value uint16) error
将Uint16写入缓冲区的指定位置，但不会改变缓冲区的Position
Writes Uint16 to the specified Position of the buffer, but does not change the position of the buffer
##### ByteBuffer::PutUInt32(offset int, value uint32) error
将Uint32写入缓冲区的指定位置，但不会改变缓冲区的Position
Writes Uint32 to the specified Position of the buffer, but does not change the position of the buffer
##### ByteBuffer::PutUInt64(offset int, value uint64) error
将Uint64写入缓冲区的指定位置，但不会改变缓冲区的Position
Writes Uint64 to the specified Position of the buffer, but does not change the position of the buffer

---

## Read
##### ByteBuffer::GetUInt16(offset int) (value uint16, err error)
从缓冲区的指定位置拿出一个UInt16数值，但不会改变缓冲区的Position
Takes a UInt16 value from the specified location of the buffer, but does not change the Position of the buffer
##### ByteBuffer::GetUInt32(offset int) (value uint32, err error)
从缓冲区的指定位置拿出一个UInt32数值，但不会改变缓冲区的Position
Takes a UInt32 value from the specified location of the buffer, but does not change the Position of the buffer
##### ByteBuffer::GetUInt64(offset int) (value uint64, err error)
从缓冲区的指定位置拿出一个UInt64数值，但不会改变缓冲区的Position
Takes a UInt64 value from the specified location of the buffer, but does not change the Position of the buffer





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

