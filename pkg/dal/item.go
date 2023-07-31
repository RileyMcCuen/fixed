package dal

import (
	"context"
	"fmt"
	"log"
	"poc/pkg/model"

	"github.com/jackc/pgx/v5"
)

type (
	Item struct {
		Conn *pgx.Conn
	}
)

func (i *Item) Insert(ctx context.Context, item model.Item) error {
	_, err := i.Conn.Exec(ctx, `insert into item (id, "itemName", "itemImageUrl") values ($1, $2, $3);`, item.Id, item.ItemName, item.ItemImageUrl)
	if err != nil {
		log.Printf("insert item failed: %s", err.Error())
		return fmt.Errorf("%w; Dao.Item.Insert", err)
	}

	return nil
}

func (i *Item) GetAll(ctx context.Context) ([]model.Item, error) {
	items := []model.Item{}

	rows, err := i.Conn.Query(context.Background(), `select id, "itemName", "itemImageUrl" from item;`)
	if err != nil {
		return items, fmt.Errorf("%w; Dao.Item.GetAll", err)
	}
	defer rows.Close()

	for rows.Next() {
		item := model.Item{}

		if err := rows.Scan(&item.Id, &item.ItemName, &item.ItemImageUrl); err != nil {
			return items, fmt.Errorf("%w; Dao.Item.GetAll", err)
		}

		items = append(items, item)
	}

	return items, nil
}

func (i *Item) GetById(ctx context.Context, id string) (model.Item, error) {
	item := model.Item{}

	row := i.Conn.QueryRow(context.Background(), `select id, "itemName", "itemImageUrl" from item where id = $1;`, id)
	if err := row.Scan(&item.Id, &item.ItemName, &item.ItemImageUrl); err != nil {
		return item, fmt.Errorf("%w; Dao.Item.GetById", err)
	}

	return item, nil
}

func (i *Item) Update(ctx context.Context, item model.Item) error {
	_, err := i.Conn.Exec(context.Background(), `update item set "itemName" = $1, "itemImageUrl" = $2 where id = $3;`, item.ItemName, item.ItemImageUrl, item.Id)
	if err != nil {
		return fmt.Errorf("%w; Dao.Item.Update", err)
	}

	return nil
}
