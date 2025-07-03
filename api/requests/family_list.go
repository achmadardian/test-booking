package requests

type CreateFamilyRequest struct {
	CstID      int    `json:"cst_id"`
	FLRelation string `json:"fl_relation"`
	FLName     string `json:"fl_name"`
	FLDOB      string `json:"fl_Dob"`
}

type UpdateFamilyRequest struct {
	CstID      int    `json:"cst_id,omitempty"`
	FLRelation string `json:"fl_relation,omitempty"`
	FLName     string `json:"fl_name,omitempty"`
	FLDOB      string `json:"fl_Dob,omitempty"`
}
