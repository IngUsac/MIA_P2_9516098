package api

import (
	"encoding/json"
	"net/http"
	"os"
	"path/filepath"
)

type DiskInfo struct {
	Name string `json:"name"`
	Path string `json:"path"`
	Size int64  `json:"size"`
}

func GetDisksHandler(w http.ResponseWriter, r *http.Request) {

	var disks []DiskInfo

	root := "./SALIDAS/discos"

	_ = filepath.Walk(root, func(path string, info os.FileInfo, err error) error {

		if err != nil {
			return nil
		}

		if info.IsDir() {
			return nil
		}

		if filepath.Ext(info.Name()) != ".dsk" {
			return nil
		}

		disks = append(disks, DiskInfo{
			Name: info.Name(),
			Path: path,
			Size: info.Size(),
		})

		return nil
	})

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(disks)
}