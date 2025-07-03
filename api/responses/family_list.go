package responses

import (
	"github.com/achmadardian/test-booking-api/models"
)

type FamilyListResponse struct {
	FLID       int    `json:"fl_id"`
	CSTID      int    `json:"cst_id"`
	FLRelation string `json:"fl_relation"`
	FLName     string `json:"fl_name"`
	FLDOB      string `json:"fl_Dob"`
}

func (f *FamilyListResponse) Map(data []models.FamilyList) []FamilyListResponse {
	families := make([]FamilyListResponse, 0, len(data))

	for _, d := range data {
		fml := FamilyListResponse{
			FLID:       d.FLID,
			CSTID:      d.CSTID,
			FLRelation: d.FLRelation,
			FLName:     d.FLName,
			FLDOB:      d.FLDOB.Format("2006-01-02"),
		}

		families = append(families, fml)
	}

	return families
}

func (f *FamilyListResponse) MapRow(data *models.FamilyList) FamilyListResponse {
	family := FamilyListResponse{
		FLID:       data.FLID,
		CSTID:      data.CSTID,
		FLRelation: data.FLRelation,
		FLName:     data.FLName,
		FLDOB:      data.FLDOB.Format("2006-01-02"),
	}

	return family
}
