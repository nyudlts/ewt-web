package lib

type IndexResponse struct {
	Message     string   `json:"message"`
	EwtHome     string   `json:"ewt_home"`
	Directories []string `json:"directories"`
}

type ShowProjectResponse struct {
	ProjectConfig string `json:"project_config"`
}

type ProjectConfig struct {
	SIPLoc           string `json:"sip-location"`
	SourceLoc        string `json:"source-location"`
	PartnerCode      string `json:"partner-code"`
	CollectionCode   string `json:"collection-code"`
	ProjectLoc       string `json:"project-location"`
	LogLoc           string `json:"log-location"`
	AIPLoc           string `json:"aip-location"`
	AMTransferSource string `json:"archivematica-transfer-source"`
	XferLoc          string `json:"xfer-location"`
	AIPStoreLoc      string `json:"aipstore-location"`
	WorkLoc          string `json:"work-location"`
}
