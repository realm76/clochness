package infrastructureA

import (
	"github.com/realm76/clochness/db/clochness/public/model"
	"github.com/realm76/clochness/internal/entity"
)

func EntriesToEntry(entries model.Entries) entity.Entry {
	return entity.Entry{
		ID:          entries.ID,
		UserID:      entries.UserID,
		ProjectID:   entries.ProjectID,
		Description: entries.Description,
		StartDate:   entries.StartDate,
		EndDate:     entries.EndDate,
		CreatedAt:   entries.CreatedAt,
		UpdatedAt:   entries.UpdatedAt,
	}
}
