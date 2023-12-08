package repository

import "context"

func (p *Postgres) checkLogin(ctx context.Context, login string) bool {
	var count int
	query := "SELECT COUNT(*) FROM users WHERE login = $1"
	err := p.conn.QueryRow(ctx, query, login).Scan(&count)
	if err != nil {
		return false
	}
	return count > 0

}
