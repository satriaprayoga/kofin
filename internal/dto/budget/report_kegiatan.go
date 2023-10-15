package dto

type KegiatanObjectReport struct {
	ObjectName string  `json:"name"`
	Volume     int     `json:"volume"`
	VolumeName string  `json:"volumeName"`
	Satuan     int     `json:"satuan"`
	SatuanName string  `json:"satuan_name"`
	Total      float64 `json:"total"`
}

type KegiatanReport struct {
	KegiatanKode string
	KegiatanName string
	Objects      []KegiatanObjectReport
}
