package model

type VulnerabilityReport struct {
	Kind     string `json:"kind"`
	Metadata struct {
		Labels struct {
			ResourceKind      string `json:"trivy-operator.resource.kind"`
			ResourceName      string `json:"trivy-operator.resource.name"`
			ResourceNameSpace string `json:"trivy-operator.resource.namespace"`
		} `json:"labels"`
	} `json:"metadata"`
	Report struct {
		Os struct {
			Family string `json:"family"`
			Name   string `json:"name"`
		} `json:"os"`
		Summary struct {
			CriticalCount int `json:"criticalCount"`
			HighCount     int `json:"highCount"`
			MediumCount   int `json:"mediumCount"`
			LowCount      int `json:"lowCount"`
			UnknownCount  int `json:"unknownCount"`
			NoneCount     int `json:"noneCount"`
		} `json:"summary"`
		Vulnerabilities []struct {
			VulnerabilityID  string  `json:"vulnerabilityID"`
			Resource         string  `json:"resource"`
			InstalledVersion string  `json:"installedVersion"`
			FixedVersion     string  `json:"fixedVersion"`
			PublishedDate    string  `json:"publishedDate"`
			LastModifiedDate string  `json:"lastModifiedDate"`
			Severity         string  `json:"severity"`
			Title            string  `json:"title"`
			PrimaryLink      string  `json:"primaryLink"`
			Score            float64 `json:"score"`
			Target           string  `json:"target"`
		} `json:"vulnerabilities"`
	} `json:"report"`
}
