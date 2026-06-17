package estructuras

type Content struct {
	BName  [12]byte
	BInodo int32
}

type FolderBlock struct {
	BContent [4]Content
}