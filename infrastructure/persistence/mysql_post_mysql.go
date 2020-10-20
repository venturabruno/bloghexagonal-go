package persistence

import (
	"database/sql"
	"errors"

	"github.com/venturabruno/bloghexagonal-go/domain"
)

type mySQLPostRepository struct {
	db *sql.DB
}

func NewMySQLPostRepository(db *sql.DB) *mySQLPostRepository {
	return &mySQLPostRepository{
		db: db,
	}
}

func (repository *mySQLPostRepository) Create(post *domain.Post) (domain.EntityID, error) {
	stmt, err := repository.db.Prepare(`
	insert into post (id, title, subtitle, content, status, created_at)
	values (?,?,?,?,?,?)`)
	if err != nil {
		return post.ID, err
	}

	_, err = stmt.Exec(
		post.ID,
		post.Title,
		post.Subtitle,
		post.Content,
		post.Status,
		post.CreatedAt,
	)
	if err != nil {
		return post.ID, err
	}

	err = stmt.Close()
	if err != nil {
		return post.ID, err
	}

	return post.ID, nil
}

func (repository *mySQLPostRepository) FindID(id domain.EntityID) (*domain.Post, error) {
	stmt, err := repository.db.Prepare(`select id, title, subtitle, content, status, created_at, published_at from post where id = ? limit 1`)
	if err != nil {
		return nil, err
	}

	var post domain.Post
	err = stmt.QueryRow(id).Scan(&post.ID, &post.Title, &post.Subtitle, &post.Content, &post.Status, &post.CreatedAt, &post.PublishedAt)

	switch {
	case err == sql.ErrNoRows:
		return nil, nil
	case err != nil:
		return nil, err
	default:
		return &post, nil
	}
}

func (repository *mySQLPostRepository) All() ([]*domain.Post, error) {
	stmt, err := repository.db.Prepare(`select id, title, subtitle, content, status, created_at, published_at from post`)
	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query()
	defer stmt.Close()
	if err != nil {
		return nil, err
	}

	var posts []*domain.Post
	for rows.Next() {
		var post domain.Post
		err = rows.Scan(&post.ID, &post.Title, &post.Subtitle, &post.Content, &post.Status, &post.CreatedAt, &post.PublishedAt)
		if err != nil {
			return nil, err
		}

		posts = append(posts, &post)
	}

	if len(posts) == 0 {
		return nil, errors.New("Not Found")
	}

	return posts, nil
}

func (repository *mySQLPostRepository) Update(post *domain.Post) error {
	stmt, err := repository.db.Prepare(`update post set title = ?, subtitle = ?, content = ?, status = ?, created_at = ?, published_at = ? where id = ?`)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(
		post.Title,
		post.Subtitle,
		post.Content,
		post.Status,
		post.CreatedAt,
		post.PublishedAt,
		post.ID,
	)
	if err != nil {
		return err
	}

	return nil
}
