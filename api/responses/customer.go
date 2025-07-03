package responses

import "github.com/achmadardian/test-booking-api/models"

type CustomerResponse struct {
	CstID       int                 `json:"cst_id"`
	CstName     string              `json:"cst_name"`
	CstDOB      string              `json:"cst_dob"`
	CstPhoneNum string              `json:"cst_phone_num"`
	CstEmail    string              `json:"cst_email"`
	Nationality NationalityResponse `json:"nationality"`
}

func (c *CustomerResponse) Map(data []models.Customer) []CustomerResponse {
	customers := make([]CustomerResponse, 0, len(data))

	for _, d := range data {
		cst := CustomerResponse{
			CstID:       d.CstID,
			CstName:     d.CstName,
			CstDOB:      d.CstDOB.Format("2006-01-02"),
			CstPhoneNum: d.CstPhoneNum,
			CstEmail:    d.CstEmail,
			Nationality: NationalityResponse{
				NationalityID:   d.Nationality.NationalityID,
				NationalityName: d.Nationality.NationalityName,
				NationalityCode: d.Nationality.NationalityCode,
			},
		}

		customers = append(customers, cst)
	}

	return customers
}

func (c *CustomerResponse) MapRow(data *models.Customer) CustomerResponse {
	customers := CustomerResponse{
		CstID:       data.CstID,
		CstName:     data.CstName,
		CstDOB:      data.CstDOB.Format("2006-01-02"),
		CstPhoneNum: data.CstPhoneNum,
		CstEmail:    data.CstEmail,
		Nationality: NationalityResponse{
			NationalityID:   data.Nationality.NationalityID,
			NationalityName: data.Nationality.NationalityName,
			NationalityCode: data.Nationality.NationalityCode,
		},
	}

	return customers
}
