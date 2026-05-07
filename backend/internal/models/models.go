package models

// User represents a system user/staff member
type User struct {
	ID           string  `json:"id"`
	Username     string  `json:"username"`
	Email        string  `json:"email"`
	PasswordHash string  `json:"-"`
	FirstName    string  `json:"first_name"`
	LastName     string  `json:"last_name"`
	Role         string  `json:"role"`
	IsActive     bool    `json:"is_active"`
	CreatedAt    string  `json:"created_at"`
	UpdatedAt    string  `json:"updated_at"`
}

// Category represents a product category
type Category struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
	Color       *string `json:"color"`
	SortOrder   int     `json:"sort_order"`
	IsActive    bool    `json:"is_active"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
}

// Product represents a menu item
type Product struct {
	ID              string    `json:"id"`
	CategoryID      *string   `json:"category_id"`
	Name            string    `json:"name"`
	Description     *string   `json:"description"`
	Price           float64   `json:"price"`
	ImageURL        *string   `json:"image_url"`
	Barcode         *string   `json:"barcode"`
	SKU             *string   `json:"sku"`
	IsAvailable     bool      `json:"is_available"`
	PreparationTime int       `json:"preparation_time"`
	SortOrder       int       `json:"sort_order"`
	CreatedAt       string    `json:"created_at"`
	UpdatedAt       string    `json:"updated_at"`
	Category        *Category `json:"category,omitempty"`
}

// DiningTable represents a table or dining area
type DiningTable struct {
	ID              string  `json:"id"`
	TableNumber     string  `json:"table_number"`
	SeatingCapacity int     `json:"seating_capacity"`
	Location        *string `json:"location"`
	IsOccupied      bool    `json:"is_occupied"`
	CreatedAt       string  `json:"created_at"`
	UpdatedAt       string  `json:"updated_at"`
}

// Order represents a customer order
type Order struct {
	ID             string       `json:"id"`
	OrderNumber    string       `json:"order_number"`
	TableID        *string      `json:"table_id"`
	UserID         *string      `json:"user_id"`
	CustomerName   *string      `json:"customer_name"`
	OrderType      string       `json:"order_type"`
	Status         string       `json:"status"`
	Subtotal       float64      `json:"subtotal"`
	TaxAmount      float64      `json:"tax_amount"`
	DiscountAmount float64      `json:"discount_amount"`
	TotalAmount    float64      `json:"total_amount"`
	Notes          *string      `json:"notes"`
	CreatedAt      string       `json:"created_at"`
	UpdatedAt      string       `json:"updated_at"`
	ServedAt       *string      `json:"served_at"`
	CompletedAt    *string      `json:"completed_at"`
	Table          *DiningTable `json:"table,omitempty"`
	User           *User        `json:"user,omitempty"`
	Items          []OrderItem  `json:"items,omitempty"`
	Payments       []Payment    `json:"payments,omitempty"`
}

// OrderItem represents an item within an order
type OrderItem struct {
	ID                  string   `json:"id"`
	OrderID             string   `json:"order_id"`
	ProductID           string   `json:"product_id"`
	Quantity            int      `json:"quantity"`
	UnitPrice           float64  `json:"unit_price"`
	TotalPrice          float64  `json:"total_price"`
	SpecialInstructions *string  `json:"special_instructions"`
	Status              string   `json:"status"`
	CreatedAt           string   `json:"created_at"`
	UpdatedAt           string   `json:"updated_at"`
	Product             *Product `json:"product,omitempty"`
}

// Payment represents a payment transaction
type Payment struct {
	ID              string  `json:"id"`
	OrderID         string  `json:"order_id"`
	PaymentMethod   string  `json:"payment_method"`
	Amount          float64 `json:"amount"`
	ReferenceNumber *string `json:"reference_number"`
	Status          string  `json:"status"`
	ProcessedBy     *string `json:"processed_by"`
	ProcessedAt     *string `json:"processed_at"`
	CreatedAt       string  `json:"created_at"`
	ProcessedByUser *User   `json:"processed_by_user,omitempty"`
}

// Inventory represents product inventory
type Inventory struct {
	ID              string   `json:"id"`
	ProductID       string   `json:"product_id"`
	CurrentStock    int      `json:"current_stock"`
	MinimumStock    int      `json:"minimum_stock"`
	MaximumStock    int      `json:"maximum_stock"`
	UnitCost        *float64 `json:"unit_cost"`
	LastRestockedAt *string  `json:"last_restocked_at"`
	CreatedAt       string   `json:"created_at"`
	UpdatedAt       string   `json:"updated_at"`
	Product         *Product `json:"product,omitempty"`
}

// OrderStatusHistory tracks order status changes
type OrderStatusHistory struct {
	ID             string  `json:"id"`
	OrderID        string  `json:"order_id"`
	PreviousStatus *string `json:"previous_status"`
	NewStatus      string  `json:"new_status"`
	ChangedBy      *string `json:"changed_by"`
	Notes          *string `json:"notes"`
	CreatedAt      string  `json:"created_at"`
	ChangedByUser  *User   `json:"changed_by_user,omitempty"`
}

// CreateOrderRequest represents the request to create a new order
type CreateOrderRequest struct {
	TableID      *string           `json:"table_id"`
	CustomerName *string           `json:"customer_name"`
	OrderType    string            `json:"order_type"`
	Items        []CreateOrderItem `json:"items"`
	Notes        *string           `json:"notes"`
}

// CreateOrderItem represents an item in the order creation request
type CreateOrderItem struct {
	ProductID           string  `json:"product_id"`
	Quantity            int     `json:"quantity"`
	SpecialInstructions *string `json:"special_instructions"`
}

// UpdateOrderStatusRequest represents the request to update order status
type UpdateOrderStatusRequest struct {
	Status string  `json:"status"`
	Notes  *string `json:"notes"`
}

// ProcessPaymentRequest represents the request to process a payment
type ProcessPaymentRequest struct {
	PaymentMethod   string  `json:"payment_method"`
	Amount          float64 `json:"amount"`
	ReferenceNumber *string `json:"reference_number"`
}

// LoginRequest represents the login request
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// LoginResponse represents the login response
type LoginResponse struct {
	Token string `json:"token"`
	User  User   `json:"user"`
}

// APIResponse represents a generic API response
type APIResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   *string     `json:"error,omitempty"`
}

// PaginatedResponse represents a paginated API response
type PaginatedResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Meta    MetaData    `json:"meta"`
}

// MetaData represents pagination metadata
type MetaData struct {
	CurrentPage int `json:"current_page"`
	PerPage     int `json:"per_page"`
	Total       int `json:"total"`
	TotalPages  int `json:"total_pages"`
}
