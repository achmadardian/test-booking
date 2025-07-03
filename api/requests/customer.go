package requests

type CreateCustomerRequest struct {
	NationalityID int                   `json:"nationality_id"`
	CstName       string                `json:"cst_name"`
	CstDOB        string                `json:"cst_dob"`
	CstPhoneNum   string                `json:"cst_phone_num"`
	CstEmail      string                `json:"cst_email"`
	Families      []CreateFamilyRequest `json:"families"`
}

type UpdateCustomerRequest struct {
	NationalityID int    `json:"nationality_id,omitempty"`
	CstName       string `json:"cst_name,omitempty"`
	CstDOB        string `json:"cst_dob,omitempty"`
	CstPhoneNum   string `json:"cst_phone_num,omitempty"`
	CstEmail      string `json:"cst_email,omitempty"`
}
