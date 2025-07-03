package responses

import "github.com/achmadardian/test-booking-api/models"

type NationalityResponse struct {
	NationalityID   int    `json:"nationality_id"`
	NationalityName string `json:"nationality_name"`
	NationalityCode string `json:"nationality_code"`
}

func (n *NationalityResponse) Map(data []models.Nationality) []NationalityResponse {
	nationalities := make([]NationalityResponse, 0, len(data))

	for _, d := range data {
		nationality := NationalityResponse{
			NationalityID:   d.NationalityID,
			NationalityName: d.NationalityName,
			NationalityCode: d.NationalityCode,
		}

		nationalities = append(nationalities, nationality)
	}

	return nationalities
}
