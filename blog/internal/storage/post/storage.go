package post

import (
	"context"

	"github.com/fk4peace/golang_services/blog/internal/entity"
	"github.com/jackc/pgx/v5"
)

type Storage struct {
	client *pgx.Conn
}

func New(client *pgx.Conn) *Storage {
	return &Storage{
		client,
	}
}

func (s *Storage) GetPosts(limit, page int64) ([]entity.Post, error) {
	if limit <= 0 {
		limit = 10
	}
	if page <= 0 {
		page = 1
	}

	offset := (page - 1) * limit

	query := `
        SELECT id, content, person_id
        FROM post
        LIMIT $1 OFFSET $2;
    `

	rows, err := s.client.Query(context.Background(), query, limit, offset)
	if err != nil {
		return nil, ErrorFrom(err)
	}
	defer rows.Close()

	posts := []entity.Post{}

	for rows.Next() {
		var p entity.Post
		if err := rows.Scan(&p.Id, &p.Content, &p.PersonId); err != nil {
			return nil, ErrorFrom(err)
		}
		posts = append(posts, p)
	}

	if err := rows.Err(); err != nil {
		return nil, ErrorFrom(err)
	}

	return posts, nil
}

func (s *Storage) GetPostsTotal() (*int64, error) {
	query := `SELECT COUNT(*) FROM post;`

	var count int64
	err := s.client.QueryRow(context.Background(), query).Scan(&count)
	if err != nil {
		return nil, ErrorFrom(err)
	}

	return &count, nil
}

func (s Storage) GetPostById(postId int64) (*entity.Post, error) {
	query := `
        SELECT id, content
        FROM post
		WHERE id = $1
	`

	var post entity.Post

	err := s.client.QueryRow(context.Background(), query, postId).Scan(
		&post.Id,
		&post.Content,
	)
	if err != nil {
		return nil, ErrorFrom(err)
	}

	return &post, nil
}

func (s Storage) CreatePost(personId int64, content string) (*entity.Post, error) {
	query := `
		INSERT INTO post (person_id, content)
		VALUES ($1, $2)
		RETURNING id, content
	`

	var result entity.Post
	err := s.client.QueryRow(context.Background(), query, personId, content).Scan(&result.Id, &result.Content)
	if err != nil {
		return nil, ErrorFrom(err)
	}

	return &result, nil
}
