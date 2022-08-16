package docker_reg_models

type V2TagsList struct {
	Name string   `json:"name"`
	Tags []string `json:"tags"`
}

type layerStruct struct {
	MediaType string `json:"mediaType"`
	Size      int    `json:"size"`
	Digest    string `json:"digest"`
	Platform  struct {
		Architecture string   `json:"architecture,omitempty"`
		OS           string   `json:"os,omitempty"`
		OSVersion    string   `json:"os.version,omitempty"`
		OSFeatures   []string `json:"os.features,omitempty"`
		Variant      string   `json:"variant,omitempty"`
		Features     []string `json:"features,omitempty"`
	} `json:"platform,omitempty"`
}

type V2Manifest struct {
	SchemaVersion int           `json:"schemaVersion"`
	MediaType     string        `json:"mediaType"`
	Config        layerStruct   `json:"config"`
	Layers        []layerStruct `json:"layers"`
}

type V2ManifestHead struct {
	ContentLength                int    `json:"Content-Length,omitempty"`
	ContentType                  string `json:"Content-Type,omitempty"`
	DockerContentDigest          string `json:"Docker-Content-Digest,omitempty"`
	DockerDistributionApiVersion string `json:"Docker-Distribution-Api-Version,omitempty"`
	Etag                         string `json:"Etag,omitempty"`
	Date                         string `json:"Date,omitempty"`
}
