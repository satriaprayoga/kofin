package dto

type ProgramResult struct {
	ProgramID   int     `json:"program_id"`
	ProgramName string  `json:"program_name"`
	ProgramKode string  `json:"program_kode"`
	UnitKode    string  `json:"unit_kode"`
	UnitName    string  `json:"unit_name"`
	Rank        float64 `json:"rank"`
}

type KegiatanResult struct {
	KegiatanID   int     `json:"kegiatan_id"`
	KegiatanName string  `json:"kegiatan_name"`
	KegiatanKode string  `json:"kegiatan_kode"`
	ProgramID    int     `json:"program_id"`
	ProgramName  string  `json:"program_name"`
	ProgramKode  string  `json:"program_kode"`
	Rank         float64 `json:"rank"`
}
