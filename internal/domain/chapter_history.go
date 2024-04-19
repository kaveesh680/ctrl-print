package domain

type ChapterHistory struct {
	ChapterName    string `json:"chapter_name"`
	ProjectName    string `json:"project_name"`
	CompanyName    string `json:"company_name"`
	VersionHistory []struct {
		ChapterVersionID   int    `json:"chapter_version_id"`
		CreatedDate        string `json:"created_date"`
		VersionNumber      int    `json:"version_number"`
		Username           string `json:"username"`
		ApplicationVersion string `json:"application_version"`
	} `json:"version_history"`
}
