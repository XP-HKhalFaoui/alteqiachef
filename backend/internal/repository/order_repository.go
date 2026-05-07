package repository

import (
	"context"
	"database/sql"
	"fmt"

	"pos-backend/internal/models"

	"github.com/google/uuid"
)

// OrderRepository defines behaviour for order persistence
type OrderRepository interface {
	ListOrders(ctx context.Context, status, orderType string, limit, offset int) ([]models.Order, int, error)
	GetOrderByID(ctx context.Context, id string) (*models.Order, error)
	CreateOrder(ctx context.Context, req models.CreateOrderRequest, userID string, orderNumber string) (string, error)
	UpdateOrderStatus(ctx context.Context, orderID string, newStatus string, changedBy string, notes *string) error
}

// SQLiteOrderRepository is an implementation of OrderRepository using *sql.DB (SQLite)
type SQLiteOrderRepository struct {
	db *sql.DB
}

func NewPostgresOrderRepository(db *sql.DB) *SQLiteOrderRepository {
	return &SQLiteOrderRepository{db: db}
}

func (r *SQLiteOrderRepository) ListOrders(ctx context.Context, status, orderType string, limit, offset int) ([]models.Order, int, error) {
	queryBuilder := `
        SELECT DISTINCT o.id, o.order_number, o.table_id, o.user_id, o.customer_name,
               o.order_type, o.status, o.subtotal, o.tax_amount, o.discount_amount,
               o.total_amount, o.notes, o.created_at, o.updated_at, o.served_at, o.completed_at,
               t.table_number, t.location,
               u.username, u.first_name, u.last_name
        FROM orders o
        LEFT JOIN dining_tables t ON o.table_id = t.id
        LEFT JOIN users u ON o.user_id = u.id
        WHERE 1=1
    `

	var args []interface{}
	if status != "" {
		queryBuilder += " AND o.status = ?"
		args = append(args, status)
	}
	if orderType != "" {
		queryBuilder += " AND o.order_type = ?"
		args = append(args, orderType)
	}

	countQuery := "SELECT COUNT(*) FROM (" + queryBuilder + ") as count_query"
	var total int
	if err := r.db.QueryRowContext(ctx, countQuery, args...).Scan(&total); err != nil {
		return nil, 0, err
	}

	queryBuilder += " ORDER BY o.created_at DESC LIMIT ? OFFSET ?"
	args = append(args, limit, offset)

	rows, err := r.db.QueryContext(ctx, queryBuilder, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var orders []models.Order
	for rows.Next() {
		var order models.Order
		var tableNumber, tableLocation sql.NullString
		var username, firstName, lastName sql.NullString

		if err := rows.Scan(
			&order.ID, &order.OrderNumber, &order.TableID, &order.UserID, &order.CustomerName,
			&order.OrderType, &order.Status, &order.Subtotal, &order.TaxAmount, &order.DiscountAmount,
			&order.TotalAmount, &order.Notes, &order.CreatedAt, &order.UpdatedAt, &order.ServedAt, &order.CompletedAt,
			&tableNumber, &tableLocation,
			&username, &firstName, &lastName,
		); err != nil {
			return nil, 0, err
		}

		if tableNumber.Valid {
			order.Table = &models.DiningTable{
				TableNumber: tableNumber.String,
				Location:    &tableLocation.String,
			}
		}
		if username.Valid {
			order.User = &models.User{
				Username:  username.String,
				FirstName: firstName.String,
				LastName:  lastName.String,
			}
		}

		if err := r.loadOrderItems(ctx, &order); err != nil {
			return nil, 0, err
		}
		if err := r.loadOrderPayments(ctx, &order); err != nil {
			return nil, 0, err
		}

		orders = append(orders, order)
	}

	return orders, total, nil
}

func (r *SQLiteOrderRepository) GetOrderByID(ctx context.Context, id string) (*models.Order, error) {
	var order models.Order
	var tableNumber, tableLocation sql.NullString
	var username, firstName, lastName sql.NullString

	query := `
        SELECT o.id, o.order_number, o.table_id, o.user_id, o.customer_name,
               o.order_type, o.status, o.subtotal, o.tax_amount, o.discount_amount,
               o.total_amount, o.notes, o.created_at, o.updated_at, o.served_at, o.completed_at,
               t.table_number, t.location,
               u.username, u.first_name, u.last_name
        FROM orders o
        LEFT JOIN dining_tables t ON o.table_id = t.id
        LEFT JOIN users u ON o.user_id = u.id
        WHERE o.id = ?
    `

	if err := r.db.QueryRowContext(ctx, query, id).Scan(
		&order.ID, &order.OrderNumber, &order.TableID, &order.UserID, &order.CustomerName,
		&order.OrderType, &order.Status, &order.Subtotal, &order.TaxAmount, &order.DiscountAmount,
		&order.TotalAmount, &order.Notes, &order.CreatedAt, &order.UpdatedAt, &order.ServedAt, &order.CompletedAt,
		&tableNumber, &tableLocation,
		&username, &firstName, &lastName,
	); err != nil {
		return nil, err
	}

	if tableNumber.Valid {
		order.Table = &models.DiningTable{
			TableNumber: tableNumber.String,
			Location:    &tableLocation.String,
		}
	}
	if username.Valid {
		order.User = &models.User{
			Username:  username.String,
			FirstName: firstName.String,
			LastName:  lastName.String,
		}
	}

	if err := r.loadOrderItems(ctx, &order); err != nil {
		return nil, err
	}
	if err := r.loadOrderPayments(ctx, &order); err != nil {
		return nil, err
	}

	return &order, nil
}

func (r *SQLiteOrderRepository) CreateOrder(ctx context.Context, req models.CreateOrderRequest, userID string, orderNumber string) (string, error) {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return "", err
	}
	defer tx.Rollback()

	orderID := uuid.New().String()

	var subtotal float64
	for _, item := range req.Items {
		var price float64
		if err := tx.QueryRowContext(ctx, "SELECT price FROM products WHERE id = ? AND is_available = 1", item.ProductID).Scan(&price); err != nil {
			if err == sql.ErrNoRows {
				return "", fmt.Errorf("product_not_found: %w", err)
			}
			return "", err
		}
		subtotal += price * float64(item.Quantity)
	}

	taxRate := 0.10
	taxAmount := subtotal * taxRate
	totalAmount := subtotal + taxAmount

	orderQuery := `
        INSERT INTO orders (id, order_number, table_id, user_id, customer_name, order_type, status,
                           subtotal, tax_amount, discount_amount, total_amount, notes)
        VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
    `

	if _, err := tx.ExecContext(ctx, orderQuery, orderID, orderNumber, req.TableID, userID, req.CustomerName,
		req.OrderType, "pending", subtotal, taxAmount, 0, totalAmount, req.Notes); err != nil {
		return "", err
	}

	for _, item := range req.Items {
		var price float64
		if err := tx.QueryRowContext(ctx, "SELECT price FROM products WHERE id = ?", item.ProductID).Scan(&price); err != nil {
			return "", err
		}
		totalPrice := price * float64(item.Quantity)
		itemID := uuid.New().String()
		itemQuery := `
            INSERT INTO order_items (id, order_id, product_id, quantity, unit_price, total_price, special_instructions)
            VALUES (?, ?, ?, ?, ?, ?, ?)
        `
		if _, err := tx.ExecContext(ctx, itemQuery, itemID, orderID, item.ProductID, item.Quantity, price, totalPrice, item.SpecialInstructions); err != nil {
			return "", err
		}
	}

	if req.OrderType == "dine_in" && req.TableID != nil {
		if _, err := tx.ExecContext(ctx, "UPDATE dining_tables SET is_occupied = 1 WHERE id = ?", *req.TableID); err != nil {
			return "", err
		}
	}

	if err := tx.Commit(); err != nil {
		return "", err
	}

	return orderID, nil
}

func (r *SQLiteOrderRepository) UpdateOrderStatus(ctx context.Context, orderID string, newStatus string, changedBy string, notes *string) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	var currentStatus string
	if err := tx.QueryRowContext(ctx, "SELECT status FROM orders WHERE id = ?", orderID).Scan(&currentStatus); err != nil {
		return err
	}

	updateQuery := "UPDATE orders SET status = ?, updated_at = datetime('now')"
	if newStatus == "served" {
		updateQuery += ", served_at = datetime('now')"
	} else if newStatus == "completed" {
		updateQuery += ", completed_at = datetime('now')"
	}
	updateQuery += " WHERE id = ?"

	if _, err := tx.ExecContext(ctx, updateQuery, newStatus, orderID); err != nil {
		return err
	}

	historyID := uuid.New().String()
	historyQuery := `
        INSERT INTO order_status_history (id, order_id, previous_status, new_status, changed_by, notes)
        VALUES (?, ?, ?, ?, ?, ?)
    `
	if _, err := tx.ExecContext(ctx, historyQuery, historyID, orderID, currentStatus, newStatus, changedBy, notes); err != nil {
		return err
	}

	if newStatus == "completed" || newStatus == "cancelled" {
		tx.ExecContext(ctx, `
            UPDATE dining_tables
            SET is_occupied = 0
            WHERE id IN (SELECT table_id FROM orders WHERE id = ? AND table_id IS NOT NULL)
        `, orderID)
	}

	return tx.Commit()
}

func (r *SQLiteOrderRepository) loadOrderItems(ctx context.Context, order *models.Order) error {
	query := `
        SELECT oi.id, oi.product_id, oi.quantity, oi.unit_price, oi.total_price,
               oi.special_instructions, oi.status, oi.created_at, oi.updated_at,
               p.name, p.description, p.price, p.preparation_time
        FROM order_items oi
        JOIN products p ON oi.product_id = p.id
        WHERE oi.order_id = ?
        ORDER BY oi.created_at
    `

	rows, err := r.db.QueryContext(ctx, query, order.ID)
	if err != nil {
		return err
	}
	defer rows.Close()

	var items []models.OrderItem
	for rows.Next() {
		var item models.OrderItem
		var productName, productDescription sql.NullString
		var productPrice sql.NullFloat64
		var preparationTime sql.NullInt32

		if err := rows.Scan(
			&item.ID, &item.ProductID, &item.Quantity, &item.UnitPrice, &item.TotalPrice,
			&item.SpecialInstructions, &item.Status, &item.CreatedAt, &item.UpdatedAt,
			&productName, &productDescription, &productPrice, &preparationTime,
		); err != nil {
			return err
		}

		item.OrderID = order.ID
		var desc *string
		if productDescription.Valid {
			desc = &productDescription.String
		}
		var price float64
		if productPrice.Valid {
			price = productPrice.Float64
		}
		prep := 0
		if preparationTime.Valid {
			prep = int(preparationTime.Int32)
		}

		item.Product = &models.Product{
			ID:              item.ProductID,
			Name:            productName.String,
			Description:     desc,
			Price:           price,
			PreparationTime: prep,
		}

		items = append(items, item)
	}

	order.Items = items
	return nil
}

func (r *SQLiteOrderRepository) loadOrderPayments(ctx context.Context, order *models.Order) error {
	query := `
        SELECT p.id, p.payment_method, p.amount, p.reference_number, p.status,
               p.processed_by, p.processed_at, p.created_at,
               u.username, u.first_name, u.last_name
        FROM payments p
        LEFT JOIN users u ON p.processed_by = u.id
        WHERE p.order_id = ?
        ORDER BY p.created_at
    `

	rows, err := r.db.QueryContext(ctx, query, order.ID)
	if err != nil {
		return err
	}
	defer rows.Close()

	var payments []models.Payment
	for rows.Next() {
		var payment models.Payment
		var username, firstName, lastName sql.NullString

		if err := rows.Scan(
			&payment.ID, &payment.PaymentMethod, &payment.Amount, &payment.ReferenceNumber,
			&payment.Status, &payment.ProcessedBy, &payment.ProcessedAt, &payment.CreatedAt,
			&username, &firstName, &lastName,
		); err != nil {
			return err
		}

		payment.OrderID = order.ID
		if username.Valid {
			payment.ProcessedByUser = &models.User{
				Username:  username.String,
				FirstName: firstName.String,
				LastName:  lastName.String,
			}
		}

		payments = append(payments, payment)
	}

	order.Payments = payments
	return nil
}
