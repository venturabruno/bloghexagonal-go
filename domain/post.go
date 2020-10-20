package domain

import (
	"fmt"
	"strings"
	"time"
)

const (
	TitleMinLength    = 3
	TitleMaxLength    = 120
	SubtitleMinLength = 3
	SubtitleMaxLength = 255
	ContentMaxLength  = 65535
)

type Post struct {
	ID          EntityID
	Title       string
	Subtitle    string
	Content     string
	Status      Status
	CreatedAt   time.Time
	PublishedAt NullTime
}

type PostRepository interface {
	Create(post *Post) (EntityID, error)
	FindID(id EntityID) (*Post, error)
	All() ([]*Post, error)
	Update(post *Post) error
}

func NewPost(title string, subtitle string, content string) (*Post, error) {
	post := Post{
		ID:        NewID(),
		Title:     title,
		Subtitle:  subtitle,
		Status:    StatusDraft(),
		Content:   content,
		CreatedAt: time.Now(),
	}

	err := post.isValid()
	if err != nil {
		return nil, err
	}

	return &post, nil
}

func (post *Post) Publish() {
	post.PublishedAt = NullTime{time.Now(), true}
	post.Status = StatusPublished()
}

func (post *Post) isValid() error {
	if strings.TrimSpace(post.Title) == "" {
		return fmt.Errorf("Título não pode ser vazio")
	}

	if len(post.Title) < TitleMinLength {
		return fmt.Errorf("Título não pode ter menos que %d caracteres", TitleMinLength)
	}

	if len(post.Title) > TitleMaxLength {
		return fmt.Errorf("Título não pode ter mais que %d caracteres", TitleMaxLength)
	}

	if strings.TrimSpace(post.Subtitle) == "" {
		return fmt.Errorf("Subtitulo não pode ser vazio")
	}

	if len(post.Subtitle) < SubtitleMinLength {
		return fmt.Errorf("Subtitulo não pode ter menos que %d caracteres", SubtitleMinLength)
	}

	if len(post.Subtitle) > SubtitleMaxLength {
		return fmt.Errorf("Subtitulo não pode ter mais que %d caracteres", SubtitleMaxLength)
	}

	if strings.TrimSpace(post.Content) == "" {
		return fmt.Errorf("Conteúdo não pode ser vazio")
	}

	if len(post.Title) > ContentMaxLength {
		return fmt.Errorf("Conteúdo não pode ter mais que %d caracteres", ContentMaxLength)
	}

	return nil
}
