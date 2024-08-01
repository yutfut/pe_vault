package vault

type Password struct {
	Password string `json:"password"`
}

type LoginResponse struct {
	RequestId     string      `json:"request_id"`
	LeaseId       string      `json:"lease_id"`
	Renewable     bool        `json:"renewable"`
	LeaseDuration int         `json:"lease_duration"`
	Data          interface{} `json:"data"`
	WrapInfo      interface{} `json:"wrap_info"`
	Warnings      interface{} `json:"warnings"`
	Auth          struct {
		ClientToken   string   `json:"client_token"`
		Accessor      string   `json:"accessor"`
		Policies      []string `json:"policies"`
		TokenPolicies []string `json:"token_policies"`
		Metadata      struct {
			Username string `json:"username"`
		} `json:"metadata"`
		LeaseDuration  int         `json:"lease_duration"`
		Renewable      bool        `json:"renewable"`
		EntityId       string      `json:"entity_id"`
		TokenType      string      `json:"token_type"`
		Orphan         bool        `json:"orphan"`
		MfaRequirement interface{} `json:"mfa_requirement"`
		NumUses        int         `json:"num_uses"`
	} `json:"auth"`
	MountType string `json:"mount_type"`
}

type Secret struct {
	Data struct {
		Postgres struct {
			Host     string `json:"host"`
			Port     string `json:"port"`
			Username string `json:"username"`
			Password string `json:"password"`
			Database string `json:"database"`
		} `json:"postgres"`
		Redis struct {
			Host     string `json:"Host"`
			Port     int    `json:"Port"`
			Password string `json:"Password"`
		} `json:"redis"`
		Minio struct {
			Host    string `json:"host"`
			Port    string `json:"port"`
			ID      string `json:"id"`
			Secret  string `json:"secret"`
			Token   string `json:"token"`
			Bucket  string `json:"bucket"`
			Secure  bool   `json:"secure"`
			Expires int64  `json:"expires"`
		} `json:"minio"`
		CDN struct {
			Host string `json:"host"`
			Port string `json:"port"`
			Path string `json:"path"`
		} `json:"cdn"`
		Jaeger struct {
			ServiceName string  `json:"service_name"`
			Type        string  `json:"type"`
			Param       float64 `json:"param"`
			LogSpans    bool    `json:"log_spans"`
			Host        string  `json:"host"`
			Port        string  `json:"port"`
		} `json:"jaeger"`
		HTTPServer struct {
			Host string `json:"host"`
			Port string `json:"port"`
		} `json:"httpserver"`
	} `json:"data"`
}
