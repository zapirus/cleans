package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v4"
)

type Configs struct {
	Postgres struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		DBName   string `yaml:"dbname"`
		SSLMode  string `yaml:"sslmode"`
	}
}

type Postgres struct {
	cfg  *Configs
	conn *pgx.Conn
}

//func (p *Postgres) GetUserRepo(ctx context.Context, name string) (string, error) {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (p *Postgres) SendUserRepo(ctx context.Context, user string) error {
//	//TODO implement me
//	panic("implement me")
//}

func NewPostgres(cfg *Configs) *Postgres {
	p := &Postgres{
		cfg: cfg,
	}
	return p
}

func (p *Postgres) Start(ctx context.Context) error {
	fmt.Println("start postgres")

	connDB := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s",
		p.cfg.Postgres.Username, p.cfg.Postgres.Password, p.cfg.Postgres.Host, p.cfg.Postgres.Port, p.cfg.Postgres.DBName, p.cfg.Postgres.SSLMode)
	conn, err := pgx.Connect(ctx, connDB)
	if err != nil {
		return err
	}
	p.conn = conn
	return nil
}

func (p *Postgres) Shutdown(ctx context.Context) error {
	err := p.conn.Close(ctx)
	if err != nil {
		return err
	}
	fmt.Println("stop postgres")
	return nil
}

func (p *Postgres) GetUserRepo(ctx context.Context, name string) (string, error) {
	return "verify", nil
}
func (p *Postgres) SendUserRepo(ctx context.Context, user string) error {
	err := p.conn.QueryRow(ctx, "INSERT INTO users (name) VALUES ($1)", user)
	if err != nil {
		return errors.New("")
	}
	return nil
}
