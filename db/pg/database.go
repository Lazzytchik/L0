package db

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"lazzytchik/L0/models"
	"log"
)

type Postgres struct {
	Conn   *pgx.Conn
	Logger *log.Logger
}

func New(config DataBaseConfig, logger *log.Logger) (Postgres, error) {
	pg := Postgres{Logger: logger}
	err := pg.Connect(context.Background(), config.GetConnString())

	if err != nil {
		return pg, err
	}

	return pg, nil
}

func (pg *Postgres) Connect(ctx context.Context, connString string) (err error) {
	pg.Conn, err = pgx.Connect(ctx, connString)
	return
}

func (pg *Postgres) Close(ctx context.Context) error {
	return pg.Conn.Close(ctx)
}

func (pg *Postgres) Insert(obj Storable) (int, error) {
	sql := obj.Insert()
	id := 0

	err := pg.Conn.QueryRow(context.Background(), sql).Scan(&id)

	if err != nil {
		pg.Logger.Println("Insert problem: ", err)
		return 0, err
	}

	pg.Logger.Println("Successful Insert, id =", id)

	return id, err
}

func (pg *Postgres) InsertOrder(order models.Order) (int, error) {
	tx, txErr := pg.Conn.BeginTx(context.Background(), pgx.TxOptions{})

	if txErr != nil {
		return 0, txErr
	}

	deliveryId, delErr := pg.Insert(order.Delivery)

	if delErr != nil {
		pg.Logger.Println("Aborting transaction... Delivery inserting error:", delErr)
		tx.Rollback(context.Background())
		return 0, delErr
	}

	paymentId, payErr := pg.Insert(order.Payment)

	if payErr != nil {
		pg.Logger.Println("Aborting transaction... Payment inserting error:", payErr)
		tx.Rollback(context.Background())
		return 0, payErr
	}

	itemsIds := make([]int, len(order.Items))

	for i, item := range order.Items {
		var itemErr error
		itemsIds[i], itemErr = pg.Insert(item)
		if itemErr != nil {
			pg.Logger.Println("Aborting transaction... Item inserting error:", itemErr)
			tx.Rollback(context.Background())
			return 0, itemErr
		}
	}

	sql := fmt.Sprintf(
		"INSERT INTO %s (order_uid, track_number, entry, delivery_id, payment_id, locale, internal_signature, customer_id, delivery_service, shardkey, sm_id, date_created, oof_shard) VALUES ('%s', '%s', '%s', '%d', '%d', '%s', '%s', '%s', '%s', '%s', '%d', to_timestamp('%d'), '%s') RETURNING id",
		order.TableName(),
		order.OrderUid,
		order.TrackNumber,
		order.Entry,
		deliveryId,
		paymentId,
		order.Locale,
		order.InternalSignature,
		order.CustomerID,
		order.DeliveryService,
		order.ShardKey,
		order.SMID,
		order.DateCreated,
		order.OofShard,
	)

	id := 0

	err := pg.Conn.QueryRow(context.Background(), sql).Scan(&id)

	if err != nil {
		pg.Logger.Println("Aborting transaction... Order inserting error:", err)
		tx.Rollback(context.Background())
		return 0, err
	}

	for _, item := range itemsIds {
		itemId := 0
		sql = fmt.Sprintf("INSERT INTO order_items (order_id, item_id) VALUES ('%d', '%d') RETURNING id", id, item)
		err := pg.Conn.QueryRow(context.Background(), sql).Scan(&itemId)

		if err != nil {
			pg.Logger.Println("Aborting transaction... Order items inserting error:", err)
			tx.Rollback(context.Background())
			return 0, err
		}
	}

	tx.Commit(context.Background())

	return id, err

}

//func (pg *Postgres) GetOrders() map[int]models.Order {
//
//	sql := "SELECT * FROM orders LIMIT 10000"
//	rows, err := pg.Conn.Query(context.Background(), sql)
//
//	if err != nil {
//		pg.Logger.Println("Insert problem: ", err)
//	}
//
//	pg.Logger.Println("Successful Insert")
//}

func (pg *Postgres) GetOrderById(id int) (models.Order, error) {
	var order models.Order

	sql := fmt.Sprintf("SELECT order_uid, track_number, entry, locale, internal_signature, customer_id, delivery_service, shardkey, sm_id, extract(epoch from date_created)::bigint, oof_shard FROM orders WHERE %s = %d", order.PrimaryColumn(), id)
	err := pg.Conn.QueryRow(context.Background(), sql).Scan(
		&order.OrderUid,
		&order.TrackNumber,
		&order.Entry,
		&order.Locale,
		&order.InternalSignature,
		&order.CustomerID,
		&order.DeliveryService,
		&order.ShardKey,
		&order.SMID,
		&order.DateCreated,
		&order.OofShard,
	)
	if err != nil {
		pg.Logger.Println("Querying Orders issue:", err)
		return order, err
	}

	sql = fmt.Sprintf("SELECT name, phone, zip, city, address, region, email FROM deliveries WHERE %s = %d", order.Delivery.PrimaryColumn(), id)
	delErr := pg.Conn.QueryRow(context.Background(), sql).Scan(
		&order.Delivery.Name,
		&order.Delivery.Phone,
		&order.Delivery.Zip,
		&order.Delivery.City,
		&order.Delivery.Address,
		&order.Delivery.Region,
		&order.Delivery.Email,
	)
	if delErr != nil {
		pg.Logger.Println("Querying Deliveries issue:", delErr)
		return order, delErr
	}

	sql = fmt.Sprintf("SELECT transaction, request_id, currency, provider, amount, payment_dt, bank, delivery_cost, goods_total, custom_fee FROM payments WHERE %s = %d", order.Payment.PrimaryColumn(), id)
	payErr := pg.Conn.QueryRow(context.Background(), sql).Scan(
		&order.Payment.Transaction,
		&order.Payment.RequestID,
		&order.Payment.Currency,
		&order.Payment.Provider,
		&order.Payment.Amount,
		&order.Payment.PaymentDT,
		&order.Payment.Bank,
		&order.Payment.DeliveryCost,
		&order.Payment.GoodsTotal,
		&order.Payment.CustomFee,
	)
	if payErr != nil {
		pg.Logger.Println("Querying Payments issue:", payErr)
		return order, payErr
	}

	sql = fmt.Sprintf(
		"SELECT chrt_id, track_number, price, rid, name, sale, size, total_price, nm_id, brand, status FROM items JOIN order_items oi on items.id = oi.item_id WHERE order_id = %d",
		id,
	)
	rows, itemErr := pg.Conn.Query(context.Background(), sql)
	if itemErr != nil {
		pg.Logger.Println("Querying Items issue:", itemErr)
		return order, itemErr
	}
	defer rows.Close()

	for rows.Next() {
		var item models.Item

		if err := rows.Scan(
			&item.ChrtId,
			&item.TrackNumber,
			&item.Price,
			&item.RID,
			&item.Name,
			&item.Sale,
			&item.Size,
			&item.TotalPrice,
			&item.NMID,
			&item.Brand,
			&item.Status,
		); err != nil {
			pg.Logger.Println("Parsing items issue:", err)
			return order, err
		}

		order.Items = append(order.Items, item)
	}

	return order, err
}
