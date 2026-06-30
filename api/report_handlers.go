package api

import (
	"encoding/json"
	"net/http"
	"os"
	"path/filepath"
)

type ReportInfo struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

func GetReportsHandler(
	w http.ResponseWriter,
	r *http.Request,
) {

	var reports []ReportInfo

	root := "./SALIDAS/reportes"

	_ = os.MkdirAll(root, 0755)

	_ = filepath.Walk(root, func(path string, info os.FileInfo, err error) error {

		if err != nil || info == nil {
			return nil
		}

		if info.IsDir() {
			return nil
		}

		ext := filepath.Ext(info.Name())

		switch ext {

		case ".png", ".jpg", ".jpeg", ".pdf":

			reports = append(reports, ReportInfo{
				Name: info.Name(),
				Path: path,
			})

		}

/*		reports = append(reports, ReportInfo{
			Name: info.Name(),
			Path: path,
		})
*/
		return nil
	})

	if reports == nil {
		reports = []ReportInfo{}
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(reports)
}