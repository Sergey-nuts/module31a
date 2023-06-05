package postgres

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"

	"GoNews/pkg/storage"
)

type Postgres struct {
	db *pgxpool.Pool
}

func New(conf string) (*Postgres, error) {
	db, err := pgxpool.Connect(context.Background(), conf)
	if err != nil {
		return nil, err
	}
	return &Postgres{db}, err
}

// получение всех публикаций
func (p *Postgres) Posts() ([]storage.Post, error) {
	rows, err := p.db.Query(context.Background(), `
		SELECT 
			p.id,
			p.title,
			p.content,
			p.author_id,
			a.name,
			p.created_at
		FROM posts p
		LEFT JOIN authors a ON a.id = p.author_id
		ORDER BY p.id;
	`)
	if err != nil {
		return nil, err
	}
	var posts []storage.Post
	var post storage.Post
	for rows.Next() {
		err := rows.Scan(&post.ID, &post.Title, &post.Content, &post.AuthorID, &post.AuthorName, &post.CreatedAt)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, rows.Err()
}

// создание новой публикации
func (p *Postgres) AddPost(post storage.Post) error {
	_, err := p.db.Query(context.Background(), `
		INSERT INTO posts(author_id, title, content)
		VALUES ($1, $2, $3);
	`, post.AuthorID, post.Title, post.Content)
	if err != nil {
		return err
	}
	return nil
}

// обновление публикации
func (p *Postgres) UpdatePost(post storage.Post) error {
	_, err := p.db.Query(context.Background(), `
		UPDATE posts
		SET
			title = $2,
			content = $3,
			author_id = $4
		WHERE id = $1;
	`, post.ID, post.Title, post.Content, post.AuthorID)
	if err != nil {
		return err
	}
	return nil
}

// удаление публикации по ID
func (p *Postgres) DeletePost(post storage.Post) error {
	_, err := p.db.Query(context.Background(), `
		DELETE FROM posts
		WHERE id = $1;
	`, post.ID)
	if err != nil {
		return err
	}
	return nil
}
