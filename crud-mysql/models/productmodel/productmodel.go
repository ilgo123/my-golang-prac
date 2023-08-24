package productmodel

import (
	"web-native/config"
	"web-native/entities"
)

func GetAll(search string) []entities.Product {
	var where string
	if search != "" {
		where = `WHERE
				products.id LIKE "%` + search +`%"
			OR
				products.name LIKE "%` + search +`%"
			OR
				categories.name LIKE "%` + search +`%"
			OR
				products.stock LIKE "%` + search +`%"
			OR
				products.description LIKE "%` + search +`%"`
	}

	rows, err := config.DB.Query(`
	SELECT 
		products.id,
		products.name,
		categories.name as category_name,
		products.stock,
		products.description,
		products.created_at,
		products.updated_at
	FROM products 
	JOIN categories 
		ON products.category_id = categories.id
	` + where)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var products []entities.Product

	for rows.Next() {
		var product entities.Product
		err := rows.Scan(
			&product.Id,
			&product.Name,
			&product.Category.Name,
			&product.Stock,
			&product.Description,
			&product.CreatedAt,
			&product.UpdatedAt,
		)

		if err != nil {
			panic(err)
		}

		products = append(products, product)
	}

	return products
}

func Create(products entities.Product) bool {
	result, err := config.DB.Exec(`
		INSERT INTO products(
			name, 
			category_id,
			stock,
			description,
			created_at,
			updated_at
		) VALUES (?,?,?,?,?,?)`,
		products.Name,
		products.Category.Id,
		products.Stock,
		products.Description,
		products.CreatedAt,
		products.UpdatedAt,
	)

	if err != nil {
		panic(err)
	}

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	return lastInsertId > 0
}

func Detail(id int) entities.Product {
	row := config.DB.QueryRow(`
		SELECT 
			products.id,
			products.name,
			products.category_id,
			categories.name as category_name,
			products.stock,
			products.description,
			products.created_at,
			products.updated_at
		FROM products 
		JOIN categories 
			ON products.category_id = categories.id
		WHERE products.id = ?
	`, id)

	var product entities.Product
	err := row.Scan(
		&product.Id,
		&product.Name,
		&product.Category.Id,
		&product.Category.Name,
		&product.Stock,
		&product.Description,
		&product.CreatedAt,
		&product.UpdatedAt,
	)

	if err != nil {
		panic(err)
	}

	return product
}

func Update(id int, product entities.Product) bool {
	query, err := config.DB.Exec(`
		UPDATE products 
		SET 
			name = ?,
			category_id = ?,
			stock = ?,
			description = ?,
			updated_at = ?
		WHERE 
			id = ?
	`,
		product.Name,
		product.Category.Id,
		product.Stock,
		product.Description,
		product.UpdatedAt,
		id,
	)

	if err != nil {
		panic(err)
	}

	res, err := query.RowsAffected()
	if err != nil {
		panic(err)
	}

	return res > 0
}

func Delete(id int) error {
	_, err := config.DB.Exec(`DELETE FROM products WHERE id = ?`, id)
	return err
}
