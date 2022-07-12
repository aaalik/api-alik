package item

const (
	fieldItem = `
		id,
		name,
		COALESCE(price, 0) price,
		COALESCE(description, '') description,
		created_at,
		COALESCE(updated_at, '') updated_at
	`
	getItems = `
		SELECT
			` + fieldItem + `
		FROM
			item
	`

	findItemById = `
		SELECT
		` + fieldItem + `
		FROM
			item
		WHERE
			id = ?
	`

	insertItem = `
		INSERT INTO
			item(
				name,
				price,
				description
			)
			VALUES(
				:name,
				:price,
				:description
			)
	`

	updateItem = `
		UPDATE
			item
		SET
			name = :name,
			price = :price,
			description = :description
		WHERE
			id = :id
	`

	deleteItem = `
			DELETE
			FROM
				item
			WHERE
				id = :id
	`
)
