package estructuras

type Partition struct {
	PartStatus      byte
	PartType        byte
	PartFit         byte
	PartStart       int32
	PartSize        int32
	PartName        [16]byte
	PartCorrelative int32
	PartID          [4]byte
}