package volume

import (
	"os"
	"store/needle"
	"sync"
)

const DEFAULT_VOLUME_SIZE = 4096

/*
1, volume用来管理实际数据文件 dat 和 index, 每一个volume对应一个dat 文件和一个index文件,
2, 卷的主要操作,新增volume,追加needle,读取index,
*/
type Volume struct {
	Id               uint64
	FilePath         string
	dataFile         *os.File
	indexFile        *os.File
	size             uint64
	freeSize         uint64
	fileAccessLocker sync.Mutex
}

func NewVolume(volumeId uint64) (v *Volume, err error) {
	return
}

func LoadVolume() (v *Volume, err error) {
	return
}

/*
static size default 4GB
*/
func (v *Volume) Size() (size uint64) {
	return
}

/*
volume free space
*/
func (v *Volume) FreeSize() (size uint64) {
	return
}

func (v *Volume) DataFile() (file *os.File) {
	return
}

func (v *Volume) Count() (total uint64) {
	return
}

func (v *Volume) AddNeedle(m *needle.Needle) (offset uint64, size uint32, err error) {
	return
}

func (v *Volume) GetNeedle(offset uint64, size uint32) (m *needle.Needle, err error) {
	return
}

func (v *Volume) DeleteNeedle(offset uint64, size uint32) (m *needle.Needle, err error) {
	return
}
