package main

func createDatabase(products *Products) error {
	// Exec return not value
	_ , err := db.Exec(
		"INSERT INTO public.products(name, price) VALUES ($1, $2);",
		products.Name,
		products.Price,
	)

	return err
}

func getProduct(id int) (Products,error) {
	var p Products
	// QueryRow return value
	row := db.QueryRow(
		"SELECT * FROM products WHERE id = $1;",
		id,
 	)
	// add value from database to value p (products)
	err := row.Scan(&p.ID, &p.Name, &p.Price)
	
	if err != nil {
		return Products{}, err
	}
	return p, nil
}

func updateProducts(id int, products *Products) error {
	_ , err := db.Exec(
		"UPDATE public.products SET name=$1, price=$2 WHERE id=$3;",
		products.Name,
		products.Price,
		id,
	)

	return err
}

func delProducts(id int) error {
	// Exec return not value
	_ , err := db.Exec(
		"DELETE FROM products WHERE id = $1;",
		id,
	)

	return err
}

func getProducts() ([]Products, error) {
	// Query is a command used to retrieve data from multiple rows.
	rows, err := db.Query("SELECT id, name, price FROM products")
	if err != nil {
	  return nil, err
	}
	defer rows.Close()
  
	var products []Products

	// command Next() It will keep looking at the value until it runs out.
	for rows.Next() {
	  var p Products
	  err := rows.Scan(&p.ID, &p.Name, &p.Price)
	  if err != nil {
		return nil, err
	  }
	  products = append(products, p)
	}
  
	// Check for errors from iterating over rows
	if err = rows.Err(); err != nil {
	  return nil, err
	}
  
	return products, nil
  }
