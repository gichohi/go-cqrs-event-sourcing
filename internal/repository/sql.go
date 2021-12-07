package repository

const (
	CreateProductQuery = `INSERT INTO products (product_id, title, description, price, created_at, updated_at) 
	VALUES ($1, $2, $3, $4, now(), now()) RETURNING user_id, name, email, created_at, updated_at`

	UpdateUserQuery = `UPDATE users u SET 
                      name=COALESCE(NULLIF($1, ''), name), 
                      email=COALESCE(NULLIF($2, ''), email), 
                      password=COALESCE(NULLIF($3, 0), password),
                      updated_at = now()
                      WHERE user_id=$4
                      RETURNING user_id, name, email, created_at, updated_at`

	GetUserByIdQuery = `SELECT u.user_id, u.name, u.email, u.created_at, u.updated_at 
	FROM users u WHERE u.user_id = $1`
)