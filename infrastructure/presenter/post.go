package presenter

import (
	"time"

	"github.com/venturabruno/bloghexagonal-go/domain"
)

type Post struct {
	ID          domain.EntityID `json:"id"`
	Title       string          `json:"title"`
	Subtitle    string          `json:"subtitle"`
	Context     string          `json:"context"`
	Status      domain.Status   `json:"status"`
	CreatedAt   time.Time       `json:"cretead_at"`
	PublishedAt domain.NullTime `json:"published_at"`
}
