package dal

import "github.com/jackc/pgx/v5"

type (
	Dao struct {
		Item Item
	}
)

func Db(conn *pgx.Conn) Dao {
	return Dao{
		Item: Item{
			Conn: conn,
		},
	}
}
