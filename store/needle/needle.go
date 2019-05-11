package needle

import "store/types"

const HEAD_MAGIC uint32 = 0x902
const FOOT_MAGIC uint32 = 0x1503

type Header struct {
	MagicNumber uint64 ``
}

/*
* Needle
 */
/*
   uint8  : 0 to 255
   uint16 : 0 to 65535
   uint32 : 0 to 4294967295
   uint64 : 0 to 18446744073709551615
   int8   : -128 to 127
   int16  : -32768 to 32767
   int32  : -2147483648 to 2147483647
   int64  : -9223372036854775808 to 9223372036854775807
*/

type Needle struct {
	Header       uint64       ``
	Cookie       types.Cookie ``
	Size         uint32       ``
	DataSize     uint32       ``
	Data         []byte       ``
	Flags        byte         ``
	NameSize     uint8        ``
	Name         []byte       ``
	MimeSize     uint8        ``
	Mime         []byte       ``
	MetaDataSize uint16       ``
	MetaData     []byte       ``
	Footer       uint32       ``
	DataCheckSum uint64       ``
	Padding      []byte       ``
}
