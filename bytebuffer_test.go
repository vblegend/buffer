package buffer

import (
	"encoding/binary"
	"testing"
)

var buf = NewByteBuffer(1024, binary.LittleEndian)
var wBytes = []byte{1, 2, 3, 4, 5}
var rBytes = []byte{0, 0, 0, 0, 0}

func BenchmarkByteBuffer(b *testing.B) {

	Assert := func(val bool, ss error) {
		if !val {
			b.Errorf("xxxxx %v", ss)
		}
	}

	b.Logf("start:")
	b.ResetTimer()

	str := "123哈哈的@##￥@！abc"

	for i := 0; i < b.N; i++ {
		buf.Reset()
		buf.WriteInt8(100)
		buf.WriteInt16(32767)
		buf.WriteInt32(25)
		buf.WriteString(str)
		buf.WriteBoolean(true)
		buf.WriteFloat64(3.1415926)
		buf.WriteFloat32(3.1415926)
		buf.WriteBytes(wBytes)

		// ===============
		originPos := buf.Position()
		buf.SetPosition(0)
		// ===============

		v, e := buf.ReadInt8()
		Assert(v == 100, e)
		v2, e := buf.ReadInt16()
		Assert(v2 == 32767, e)

		v3, e := buf.ReadInt32()
		Assert(v3 == 25, e)

		v4, e := buf.ReadString(25)
		Assert(v4 == str, e)

		v5, e := buf.ReadBoolean()
		Assert(v5 == true, e)

		v6, e := buf.ReadFloat64()
		Assert(v6 == 3.1415926, e)

		v7, e := buf.ReadFloat32()
		Assert(v7 == 3.1415926, e)

		e = buf.ReadBytes(rBytes, -1)

		// Assert(rBytes[0] == wBytes[0] && e == nil, e)
		// Assert(rBytes[1] == wBytes[1] && e == nil, e)
		// Assert(rBytes[2] == wBytes[2] && e == nil, e)
		// Assert(rBytes[3] == wBytes[3] && e == nil, e)
		// Assert(rBytes[4] == wBytes[4] && e == nil, e)

		currentPos := buf.Position()
		Assert(currentPos == originPos, e) // errors.New("Position")

	}

	b.StopTimer()
	b.Logf("end:")
}
