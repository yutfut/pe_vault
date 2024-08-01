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
