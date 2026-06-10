package estructuras

type MBR struct {
	MbrTamano        int32
	MbrFechaCreacion [20]byte
	MbrDiskSignature int32
	DskFit           byte
	MbrPartitions    [4]Partition
}