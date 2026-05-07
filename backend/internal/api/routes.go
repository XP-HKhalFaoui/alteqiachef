package api

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"
	"time"

	"pos-backend/internal/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// SetupRoutes configures all API routes (orchestrator)
func SetupRoutes(router *gin.RouterGroup, db *sql.DB, authMiddleware gin.HandlerFunc) {
	SetupAuthRoutes(router, db, authMiddleware)
	SetupProductRoutes(router, db, authMiddleware)
	SetupTableRoutes(router, db, authMiddleware)
	SetupOrderRoutes(router, db, authMiddleware)
	SetupPaymentRoutes(router, db, authMiddleware)
	SetupServerRoutes(router, db, authMiddleware)
	SetupCounterRoutes(router, db, authMiddleware)
	SetupAdminReportsRoutes(router, db, authMiddleware)
	SetupAdminRoutes(router, db, authMiddleware)
	SetupKitchenRoutes(router, db, authMiddleware)
}

func createCategory(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req struct {
			Name        string  `json:"name" binding:"required"`
			Description *string `json:"description"`
			Color       *string `json:"color"`
			SortOrder   int     `json:"sort_order"`
		}

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"success": false, "message": "Invalid request body", "error": err.Error()})
			return
		}

		categoryID := uuid.New().String()
		_, err := db.Exec(`
			INSERT INTO categories (id, name, description, color, sort_order)
			VALUES (?, ?, ?, ?, ?)
		`, categoryID, req.Name, req.Description, req.Color, req.SortOrder)

		if err != nil {
			c.JSON(500, gin.H{"success": false, "message": "Failed to create category", "error": err.Error()})
			return
		}

		c.JSON(201, gin.H{"success": true, "message": "Category created successfully", "data": map[string]interface{}{"id": categoryID}})
	}
}

func updateCategory(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		categoryID := c.Param("id")

		var req struct {
			Name        *string `json:"name"`
			Description *string `json:"description"`
			Color       *string `json:"color"`
			SortOrder   *int    `json:"sort_order"`
			IsActive    *bool   `json:"is_active"`
		}

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"success": false, "message": "Invalid request body", "error": err.Error()})
			return
		}

		updates := []string{}
		args := []interface{}{}

		if req.Name != nil {
			updates = append(updates, "name = ?")
			args = append(args, *req.Name)
		}
		if req.Description != nil {
			updates = append(updates, "description = ?")
			args = append(args, req.Description)
		}
		if req.Color != nil {
			updates = append(updates, "color = ?")
			args = append(args, req.Color)
		}
		if req.SortOrder != nil {
			updates = append(updates, "sort_order = ?")
			args = append(args, *req.SortOrder)
		}
		if req.IsActive != nil {
			updates = append(updates, "is_active = ?")
			args = append(args, *req.IsActive)
		}

		if len(updates) == 0 {
			c.JSON(400, gin.H{"success": false, "message": "No fields to update"})
			return
		}

		updates = append(updates, "updated_at = datetime('now')")
		args = append(args, categoryID)

		query := fmt.Sprintf("UPDATE categories SET %s WHERE id = ?", strings.Join(updates, ", "))

		result, err := db.Exec(query, args...)
		if err != nil {
			c.JSON(500, gin.H{"success": false, "message": "Failed to update category", "error": err.Error()})
			return
		}

		rowsAffected, _ := result.RowsAffected()
		if rowsAffected == 0 {
			c.JSON(404, gin.H{"success": false, "message": "Category not found"})
			return
		}

		c.JSON(200, gin.H{"success": true, "message": "Category updated successfully"})
	}
}

func deleteCategory(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		categoryID := c.Param("id")

		var productCount int
		db.QueryRow("SELECT COUNT(*) FROM products WHERE category_id = ?", categoryID).Scan(&productCount)

		if productCount > 0 {
			c.JSON(400, gin.H{"success": false, "message": "Cannot delete category with existing products", "error": "category_has_products"})
			return
		}

		result, err := db.Exec("DELETE FROM categories WHERE id = ?", categoryID)
		if err != nil {
			c.JSON(500, gin.H{"success": false, "message": "Failed to delete category", "error": err.Error()})
			return
		}

		rowsAffected, _ := result.RowsAffected()
		if rowsAffected == 0 {
			c.JSON(404, gin.H{"success": false, "message": "Category not found"})
			return
		}

		c.JSON(200, gin.H{"success": true, "message": "Category deleted successfully"})
	}
}

func createProduct(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req struct {
			CategoryID      *string `json:"category_id"`
			Name            string  `json:"name" binding:"required"`
			Description     *string `json:"description"`
			Price           float64 `json:"price" binding:"required"`
			ImageURL        *string `json:"image_url"`
			Barcode         *string `json:"barcode"`
			SKU             *string `json:"sku"`
			PreparationTime int     `json:"preparation_time"`
			SortOrder       int     `json:"sort_order"`
		}

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"success": false, "message": "Invalid request body", "error": err.Error()})
			return
		}

		productID := uuid.New().String()
		_, err := db.Exec(`
			INSERT INTO products (id, category_id, name, description, price, image_url, barcode, sku, preparation_time, sort_order)
			VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
		`, productID, req.CategoryID, req.Name, req.Description, req.Price, req.ImageURL, req.Barcode, req.SKU, req.PreparationTime, req.SortOrder)

		if err != nil {
			c.JSON(500, gin.H{"success": false, "message": "Failed to create product", "error": err.Error()})
			return
		}

		c.JSON(201, gin.H{"success": true, "message": "Product created successfully", "data": map[string]interface{}{"id": productID}})
	}
}

func updateProduct(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		productID := c.Param("id")

		var req struct {
			CategoryID      *string  `json:"category_id"`
			Name            *string  `json:"name"`
			Description     *string  `json:"description"`
			Price           *float64 `json:"price"`
			ImageURL        *string  `json:"image_url"`
			Barcode         *string  `json:"barcode"`
			SKU             *string  `json:"sku"`
			IsAvailable     *bool    `json:"is_available"`
			PreparationTime *int     `json:"preparation_time"`
			SortOrder       *int     `json:"sort_order"`
		}

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"success": false, "message": "Invalid request body", "error": err.Error()})
			return
		}

		updates := []string{}
		args := []interface{}{}

		if req.CategoryID != nil {
			updates = append(updates, "category_id = ?")
			args = append(args, req.CategoryID)
		}
		if req.Name != nil {
			updates = append(updates, "name = ?")
			args = append(args, *req.Name)
		}
		if req.Description != nil {
			updates = append(updates, "description = ?")
			args = append(args, req.Description)
		}
		if req.Price != nil {
			updates = append(updates, "price = ?")
			args = append(args, *req.Price)
		}
		if req.ImageURL != nil {
			updates = append(updates, "image_url = ?")
			args = append(args, req.ImageURL)
		}
		if req.Barcode != nil {
			updates = append(updates, "barcode = ?")
			args = append(args, req.Barcode)
		}
		if req.SKU != nil {
			updates = append(updates, "sku = ?")
			args = append(args, req.SKU)
		}
		if req.IsAvailable != nil {
			updates = append(updates, "is_available = ?")
			args = append(args, *req.IsAvailable)
		}
		if req.PreparationTime != nil {
			updates = append(updates, "preparation_time = ?")
			args = append(args, *req.PreparationTime)
		}
		if req.SortOrder != nil {
			updates = append(updates, "sort_order = ?")
			args = append(args, *req.SortOrder)
		}

		if len(updates) == 0 {
			c.JSON(400, gin.H{"success": false, "message": "No fields to update"})
			return
		}

		updates = append(updates, "updated_at = datetime('now')")
		args = append(args, productID)

		query := fmt.Sprintf("UPDATE products SET %s WHERE id = ?", strings.Join(updates, ", "))

		result, err := db.Exec(query, args...)
		if err != nil {
			c.JSON(500, gin.H{"success": false, "message": "Failed to update product", "error": err.Error()})
			return
		}

		rowsAffected, _ := result.RowsAffected()
		if rowsAffected == 0 {
			c.JSON(404, gin.H{"success": false, "message": "Product not found"})
			return
		}

		c.JSON(200, gin.H{"success": true, "message": "Product updated successfully"})
	}
}

func deleteProduct(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		productID := c.Param("id")

		var orderCount int
		db.QueryRow(`
			SELECT COUNT(*)
			FROM order_items oi
			JOIN orders o ON oi.order_id = o.id
			WHERE oi.product_id = ? AND o.status NOT IN ('completed', 'cancelled')
		`, productID).Scan(&orderCount)

		if orderCount > 0 {
			c.JSON(400, gin.H{"success": false, "message": "Cannot delete product with active orders", "error": "product_has_active_orders"})
			return
		}

		result, err := db.Exec("DELETE FROM products WHERE id = ?", productID)
		if err != nil {
			c.JSON(500, gin.H{"success": false, "message": "Failed to delete product", "error": err.Error()})
			return
		}

		rowsAffected, _ := result.RowsAffected()
		if rowsAffected == 0 {
			c.JSON(404, gin.H{"success": false, "message": "Product not found"})
			return
		}

		c.JSON(200, gin.H{"success": true, "message": "Product deleted successfully"})
	}
}

func createTable(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req struct {
			TableNumber     string  `json:"table_number" binding:"required"`
			SeatingCapacity int     `json:"seating_capacity"`
			Location        *string `json:"location"`
		}

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"success": false, "message": "Invalid request body", "error": err.Error()})
			return
		}

		tableID := uuid.New().String()
		_, err := db.Exec(`
			INSERT INTO dining_tables (id, table_number, seating_capacity, location)
			VALUES (?, ?, ?, ?)
		`, tableID, req.TableNumber, req.SeatingCapacity, req.Location)

		if err != nil {
			c.JSON(500, gin.H{"success": false, "message": "Failed to create table", "error": err.Error()})
			return
		}

		c.JSON(201, gin.H{"success": true, "message": "Table created successfully", "data": map[string]interface{}{"id": tableID}})
	}
}

func updateTable(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		tableID := c.Param("id")

		var req struct {
			TableNumber     *string `json:"table_number"`
			SeatingCapacity *int    `json:"seating_capacity"`
			Location        *string `json:"location"`
			IsOccupied      *bool   `json:"is_occupied"`
		}

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"success": false, "message": "Invalid request body", "error": err.Error()})
			return
		}

		updates := []string{}
		args := []interface{}{}

		if req.TableNumber != nil {
			updates = append(updates, "table_number = ?")
			args = append(args, *req.TableNumber)
		}
		if req.SeatingCapacity != nil {
			updates = append(updates, "seating_capacity = ?")
			args = append(args, *req.SeatingCapacity)
		}
		if req.Location != nil {
			updates = append(updates, "location = ?")
			args = append(args, req.Location)
		}
		if req.IsOccupied != nil {
			updates = append(updates, "is_occupied = ?")
			args = append(args, *req.IsOccupied)
		}

		if len(updates) == 0 {
			c.JSON(400, gin.H{"success": false, "message": "No fields to update"})
			return
		}

		updates = append(updates, "updated_at = datetime('now')")
		args = append(args, tableID)

		query := fmt.Sprintf("UPDATE dining_tables SET %s WHERE id = ?", strings.Join(updates, ", "))

		result, err := db.Exec(query, args...)
		if err != nil {
			c.JSON(500, gin.H{"success": false, "message": "Failed to update table", "error": err.Error()})
			return
		}

		rowsAffected, _ := result.RowsAffected()
		if rowsAffected == 0 {
			c.JSON(404, gin.H{"success": false, "message": "Table not found"})
			return
		}

		c.JSON(200, gin.H{"success": true, "message": "Table updated successfully"})
	}
}

func deleteTable(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		tableID := c.Param("id")

		var orderCount int
		db.QueryRow(`
			SELECT COUNT(*)
			FROM orders
			WHERE table_id = ? AND status NOT IN ('completed', 'cancelled')
		`, tableID).Scan(&orderCount)

		if orderCount > 0 {
			c.JSON(400, gin.H{"success": false, "message": "Cannot delete table with active orders", "error": "table_has_active_orders"})
			return
		}

		result, err := db.Exec("DELETE FROM dining_tables WHERE id = ?", tableID)
		if err != nil {
			c.JSON(500, gin.H{"success": false, "message": "Failed to delete table", "error": err.Error()})
			return
		}

		rowsAffected, _ := result.RowsAffected()
		if rowsAffected == 0 {
			c.JSON(404, gin.H{"success": false, "message": "Table not found"})
			return
		}

		c.JSON(200, gin.H{"success": true, "message": "Table deleted successfully"})
	}
}

func createUser(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req struct {
			Username  string `json:"username" binding:"required"`
			Email     string `json:"email" binding:"required"`
			Password  string `json:"password" binding:"required"`
			FirstName string `json:"first_name" binding:"required"`
			LastName  string `json:"last_name" binding:"required"`
			Role      string `json:"role" binding:"required"`
		}

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"success": false, "message": "Invalid request body", "error": err.Error()})
			return
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(500, gin.H{"success": false, "message": "Failed to hash password", "error": err.Error()})
			return
		}

		userID := uuid.New().String()
		err = func() error {
			_, err := db.Exec(`
				INSERT INTO users (id, username, email, password_hash, first_name, last_name, role)
				VALUES (?, ?, ?, ?, ?, ?, ?)
			`, userID, req.Username, req.Email, string(hashedPassword), req.FirstName, req.LastName, req.Role)
			return err
		}()

		if err != nil {
			c.JSON(500, gin.H{"success": false, "message": "Failed to create user", "error": err.Error()})
			return
		}

		c.JSON(201, gin.H{"success": true, "message": "User created successfully", "data": map[string]interface{}{"id": userID}})
	}
}

func updateUser(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.Param("id")

		var req struct {
			Username  *string `json:"username"`
			Email     *string `json:"email"`
			Password  *string `json:"password"`
			FirstName *string `json:"first_name"`
			LastName  *string `json:"last_name"`
			Role      *string `json:"role"`
			IsActive  *bool   `json:"is_active"`
		}

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"success": false, "message": "Invalid request body", "error": err.Error()})
			return
		}

		updates := []string{}
		args := []interface{}{}

		if req.Username != nil {
			updates = append(updates, "username = ?")
			args = append(args, *req.Username)
		}
		if req.Email != nil {
			updates = append(updates, "email = ?")
			args = append(args, *req.Email)
		}
		if req.Password != nil {
			hashedPassword, err := bcrypt.GenerateFromPassword([]byte(*req.Password), bcrypt.DefaultCost)
			if err != nil {
				c.JSON(500, gin.H{"success": false, "message": "Failed to hash password", "error": err.Error()})
				return
			}
			updates = append(updates, "password_hash = ?")
			args = append(args, string(hashedPassword))
		}
		if req.FirstName != nil {
			updates = append(updates, "first_name = ?")
			args = append(args, *req.FirstName)
		}
		if req.LastName != nil {
			updates = append(updates, "last_name = ?")
			args = append(args, *req.LastName)
		}
		if req.Role != nil {
			updates = append(updates, "role = ?")
			args = append(args, *req.Role)
		}
		if req.IsActive != nil {
			updates = append(updates, "is_active = ?")
			args = append(args, *req.IsActive)
		}

		if len(updates) == 0 {
			c.JSON(400, gin.H{"success": false, "message": "No fields to update"})
			return
		}

		updates = append(updates, "updated_at = datetime('now')")
		args = append(args, userID)

		query := fmt.Sprintf("UPDATE users SET %s WHERE id = ?", strings.Join(updates, ", "))

		result, err := db.Exec(query, args...)
		if err != nil {
			c.JSON(500, gin.H{"success": false, "message": "Failed to update user", "error": err.Error()})
			return
		}

		rowsAffected, _ := result.RowsAffected()
		if rowsAffected == 0 {
			c.JSON(404, gin.H{"success": false, "message": "User not found"})
			return
		}

		c.JSON(200, gin.H{"success": true, "message": "User updated successfully"})
	}
}

func deleteUser(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.Param("id")

		var orderCount int
		db.QueryRow("SELECT COUNT(*) FROM orders WHERE user_id = ?", userID).Scan(&orderCount)

		if orderCount > 0 {
			c.JSON(400, gin.H{"success": false, "message": "Cannot delete user with existing orders", "error": "user_has_orders"})
			return
		}

		result, err := db.Exec("DELETE FROM users WHERE id = ?", userID)
		if err != nil {
			c.JSON(500, gin.H{"success": false, "message": "Failed to delete user", "error": err.Error()})
			return
		}

		rowsAffected, _ := result.RowsAffected()
		if rowsAffected == 0 {
			c.JSON(404, gin.H{"success": false, "message": "User not found"})
			return
		}

		c.JSON(200, gin.H{"success": true, "message": "User deleted successfully"})
	}
}

func getAdminUsers(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		page := 1
		perPage := 20
		role := c.Query("role")
		isActive := c.Query("active")
		search := c.Query("search")

		if pageStr := c.Query("page"); pageStr != "" {
			if p, err := strconv.Atoi(pageStr); err == nil && p > 0 {
				page = p
			}
		}
		if perPageStr := c.Query("per_page"); perPageStr != "" {
			if pp, err := strconv.Atoi(perPageStr); err == nil && pp > 0 && pp <= 100 {
				perPage = pp
			}
		}

		offset := (page - 1) * perPage

		queryBuilder := "SELECT id, username, email, first_name, last_name, role, is_active, created_at FROM users WHERE 1=1"
		args := []interface{}{}

		if role != "" {
			queryBuilder += " AND role = ?"
			args = append(args, role)
		}
		if isActive != "" {
			queryBuilder += " AND is_active = ?"
			args = append(args, isActive == "true")
		}
		if search != "" {
			// SQLite: each ? needs its own arg
			queryBuilder += " AND (first_name LIKE ? OR last_name LIKE ? OR username LIKE ? OR email LIKE ?)"
			s := "%" + search + "%"
			args = append(args, s, s, s, s)
		}

		countQuery := "SELECT COUNT(*) FROM (" + queryBuilder + ") as count_query"
		var total int
		if err := db.QueryRow(countQuery, args...).Scan(&total); err != nil {
			c.JSON(500, gin.H{"success": false, "message": "Failed to count users", "error": err.Error()})
			return
		}

		queryBuilder += " ORDER BY created_at DESC LIMIT ? OFFSET ?"
		args = append(args, perPage, offset)

		rows, err := db.Query(queryBuilder, args...)
		if err != nil {
			c.JSON(500, gin.H{"success": false, "message": "Failed to fetch users", "error": err.Error()})
			return
		}
		defer rows.Close()

		var users []map[string]interface{}
		for rows.Next() {
			var id, username, email, firstName, lastName, userRole string
			var isActiveVal bool
			var createdAt time.Time

			err := rows.Scan(&id, &username, &email, &firstName, &lastName, &userRole, &isActiveVal, &createdAt)
			if err != nil {
				c.JSON(500, gin.H{"success": false, "message": "Failed to scan user data", "error": err.Error()})
				return
			}

			users = append(users, map[string]interface{}{
				"id":         id,
				"username":   username,
				"email":      email,
				"first_name": firstName,
				"last_name":  lastName,
				"role":       userRole,
				"is_active":  isActiveVal,
				"created_at": createdAt,
			})
		}

		totalPages := (total + perPage - 1) / perPage
		c.JSON(200, gin.H{
			"success": true,
			"message": "Users retrieved successfully",
			"data":    users,
			"meta": gin.H{
				"current_page": page,
				"per_page":     perPage,
				"total":        total,
				"total_pages":  totalPages,
			},
		})
	}
}

func getAdminCategories(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		page := 1
		perPage := 20
		activeOnly := c.Query("active_only") == "true"
		search := c.Query("search")

		if pageStr := c.Query("page"); pageStr != "" {
			if p, err := strconv.Atoi(pageStr); err == nil && p > 0 {
				page = p
			}
		}
		if perPageStr := c.Query("per_page"); perPageStr != "" {
			if pp, err := strconv.Atoi(perPageStr); err == nil && pp > 0 && pp <= 100 {
				perPage = pp
			}
		}

		offset := (page - 1) * perPage

		queryBuilder := "SELECT id, name, description, color, sort_order, is_active, created_at, updated_at FROM categories WHERE 1=1"
		args := []interface{}{}

		if activeOnly {
			queryBuilder += " AND is_active = 1"
		}
		if search != "" {
			queryBuilder += " AND (name LIKE ? OR description LIKE ?)"
			s := "%" + search + "%"
			args = append(args, s, s)
		}

		countQuery := "SELECT COUNT(*) FROM (" + queryBuilder + ") as count_query"
		var total int
		if err := db.QueryRow(countQuery, args...).Scan(&total); err != nil {
			c.JSON(500, gin.H{"success": false, "message": "Failed to count categories", "error": err.Error()})
			return
		}

		queryBuilder += " ORDER BY sort_order ASC, name ASC LIMIT ? OFFSET ?"
		args = append(args, perPage, offset)

		rows, err := db.Query(queryBuilder, args...)
		if err != nil {
			c.JSON(500, gin.H{"success": false, "message": "Failed to fetch categories", "error": err.Error()})
			return
		}
		defer rows.Close()

		var categories []models.Category
		for rows.Next() {
			var category models.Category
			err := rows.Scan(
				&category.ID, &category.Name, &category.Description, &category.Color,
				&category.SortOrder, &category.IsActive, &category.CreatedAt, &category.UpdatedAt,
			)
			if err != nil {
				c.JSON(500, gin.H{"success": false, "message": "Failed to scan category", "error": err.Error()})
				return
			}
			categories = append(categories, category)
		}

		totalPages := (total + perPage - 1) / perPage
		c.JSON(200, gin.H{
			"success": true,
			"message": "Categories retrieved successfully",
			"data":    categories,
			"meta": gin.H{
				"current_page": page,
				"per_page":     perPage,
				"total":        total,
				"total_pages":  totalPages,
			},
		})
	}
}

func getAdminTables(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		page := 1
		perPage := 20
		location := c.Query("location")
		status := c.Query("status")
		search := c.Query("search")

		if pageStr := c.Query("page"); pageStr != "" {
			if p, err := strconv.Atoi(pageStr); err == nil && p > 0 {
				page = p
			}
		}
		if perPageStr := c.Query("per_page"); perPageStr != "" {
			if pp, err := strconv.Atoi(perPageStr); err == nil && pp > 0 && pp <= 100 {
				perPage = pp
			}
		}

		offset := (page - 1) * perPage

		queryBuilder := `
			SELECT t.id, t.table_number, t.seating_capacity, t.location, t.is_occupied,
			       t.created_at, t.updated_at,
			       o.id as order_id, o.order_number, o.customer_name, o.status as order_status,
			       o.created_at as order_created_at, o.total_amount
			FROM dining_tables t
			LEFT JOIN orders o ON t.id = o.table_id AND o.status NOT IN ('completed', 'cancelled')
			WHERE 1=1
		`

		args := []interface{}{}

		if location != "" {
			queryBuilder += " AND t.location LIKE ?"
			args = append(args, "%"+location+"%")
		}
		if status == "occupied" {
			queryBuilder += " AND t.is_occupied = 1"
		} else if status == "available" {
			queryBuilder += " AND t.is_occupied = 0"
		}
		if search != "" {
			queryBuilder += " AND (t.table_number LIKE ? OR t.location LIKE ?)"
			s := "%" + search + "%"
			args = append(args, s, s)
		}

		countQuery := "SELECT COUNT(*) FROM (" + queryBuilder + ") as count_query"
		var total int
		if err := db.QueryRow(countQuery, args...).Scan(&total); err != nil {
			c.JSON(500, gin.H{"success": false, "message": "Failed to count tables", "error": err.Error()})
			return
		}

		queryBuilder += " ORDER BY t.table_number ASC LIMIT ? OFFSET ?"
		args = append(args, perPage, offset)

		rows, err := db.Query(queryBuilder, args...)
		if err != nil {
			c.JSON(500, gin.H{"success": false, "message": "Failed to fetch tables", "error": err.Error()})
			return
		}
		defer rows.Close()

		var tables []map[string]interface{}
		for rows.Next() {
			var table models.DiningTable
			var orderID, orderNumber, customerName, orderStatus sql.NullString
			var orderCreatedAt sql.NullString
			var totalAmount sql.NullFloat64

			err := rows.Scan(
				&table.ID, &table.TableNumber, &table.SeatingCapacity, &table.Location, &table.IsOccupied,
				&table.CreatedAt, &table.UpdatedAt,
				&orderID, &orderNumber, &customerName, &orderStatus, &orderCreatedAt, &totalAmount,
			)
			if err != nil {
				c.JSON(500, gin.H{"success": false, "message": "Failed to scan table", "error": err.Error()})
				return
			}

			tableData := map[string]interface{}{
				"id":               table.ID,
				"table_number":     table.TableNumber,
				"seating_capacity": table.SeatingCapacity,
				"location":         table.Location,
				"is_occupied":      table.IsOccupied,
				"created_at":       table.CreatedAt,
				"updated_at":       table.UpdatedAt,
				"current_order":    nil,
			}

			if orderID.Valid {
				tableData["current_order"] = map[string]interface{}{
					"id":            orderID.String,
					"order_number":  orderNumber.String,
					"customer_name": customerName.String,
					"status":        orderStatus.String,
					"created_at":    orderCreatedAt.String,
					"total_amount":  totalAmount.Float64,
				}
			}

			tables = append(tables, tableData)
		}

		totalPages := (total + perPage - 1) / perPage
		c.JSON(200, gin.H{
			"success": true,
			"message": "Tables retrieved successfully",
			"data":    tables,
			"meta": gin.H{
				"current_page": page,
				"per_page":     perPage,
				"total":        total,
				"total_pages":  totalPages,
			},
		})
	}
}

func stringPtr(s string) *string {
	return &s
}
