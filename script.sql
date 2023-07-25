select
    p.id,
    p.name,
    p.barcode,
    p.price,
    p.category_id

    JSON_BUILD_OBJECT (
        'id', c.id,
        'name', c.name,
        'parent_id', c.parent_id,
        'created_at', c.created_at,
        'updated_at', c.updated_at
    ) AS category
    FROM product AS p
    LEFT JOIN category AS c ON p.category_id = c.id
    WHERE p.id = '01e6e80c-a2e4-4323-bb34-bc33dcd7ae67'



INSERT INTO pr (id,name)
VALUES(1,'a') 
ON CONFLICT (id,name) 
DO 
   UPDATE SET name='b';


INSERT INTO remainder(filial_id,category_id,name,barcode,amount,price,total_price,updated_at)
VALUES ($1, $2, $3, $4, $5, $6,$7, NOW())
ON CONFLICT (filial_id,barcode)
DO
    UPDATE SET
            filial_id = $1,
			category_id = $2,
			name = $3,
			barcode = %4,
			amount = $5,
			price = :$6,
			total_price =: $7,
			updated_at = NOW()
	WHERE id = $8



SELECT
			COUNT(*) OVER(),
			id,
			coming_id,
			category_id,
			name,
			barcode,
			amount,
			price,
			total_price,
			created_at,
			updated_at
		FROM coming_products

        where coming_id='P-0001';




                SELECT
                        COUNT(*) OVER(),
                        id,
                        coming_id,
                        category_id,
                        name,
                        barcode,
                        amount,
                        price,
                        total_price,
                        created_at,
                        updated_at
                FROM coming_products
         WHERE TRUE AND coming_id ILIKE '%' || 'a1567ea7-7f78-4c8b-a1d6-0b60b17f1bb8' || '%' OFFSET 0 LIMIT 10