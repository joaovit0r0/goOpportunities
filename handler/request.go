package handler

import "fmt"

type CreateOpeningRequest struct {
	Role     string `json: "role"`
	Company  string `json: "company"`
	Location string `json: "location"`
	Remote   *bool  `json: "remote"` // it is a point cause boolean is a default value (zero value)
	Link     string `json: "link"`
	Salary   int64  `json: "salary"`
}

func errParamIsRequired(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}

func (req *CreateOpeningRequest) Validate() error {
	if req.Role == "" && req.Company == "" && req.Location == "" && req.Remote == nil && req.Link == "" && req.Salary <= 0 {
		return fmt.Errorf("send a valid JSON")
	}
	if req.Role == "" {
		return errParamIsRequired("role", "string")
	}
	if req.Company == "" {
		return errParamIsRequired("company", "string")
	}
	if req.Location == "" {
		return errParamIsRequired("location", "string")
	}
	if req.Link == "" {
		return errParamIsRequired("link", "string")
	}
	if req.Remote == nil {
		return errParamIsRequired("remote", "bool")
	}
	if req.Salary <= 0 {
		return errParamIsRequired("salary", "int64")
	}

	return nil
}
