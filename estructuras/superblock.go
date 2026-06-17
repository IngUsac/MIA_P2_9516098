package estructuras

type SuperBlock struct {

	// Cantidades
	SFilesystemType   int32
	SInodesCount      int32
	SBlocksCount      int32
	SFreeBlocksCount  int32
	SFreeInodesCount  int32

	// Fechas
	SMtime [20]byte
	SUmtime [20]byte

	// Montajes
	SMntCount int32

	// Magic EXT2
	SMagic int32

	// Tamaños
	SInodeSize int32
	SBlockSize int32

	// Primeros libres
	SFirstIno int32
	SFirstBlo int32

	// Inicios
	SBmInodeStart int32
	SBmBlockStart int32
	SInodeStart int32
	SBlockStart int32
}