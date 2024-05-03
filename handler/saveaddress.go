package handler

import "fmt"

func SaveStaticAddress(address string, latency, downloadspeed, uploadspeed float64, isp string) error {

	db, err := getDBConnection()
	config := GetConfig()

	if err != nil {
		return err
	}

	// Insert data into the database
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	iD := 0
	query := fmt.Sprintf("INSERT INTO public.%s (address, latency, downloadspeed, uploadspeed, isp ) VALUES ($1, $2, $3, $4, $5)RETURNING id", config.TB_NAME[0])
	err = tx.QueryRow(query, address, latency, downloadspeed, uploadspeed, isp).Scan(&iD)

	if err != nil {
		fmt.Printf("Failed insert %s, %f, %f, %f, %s into %s with error: %s\n", address, latency, downloadspeed, uploadspeed, isp, config.TB_NAME[0], err)
		return err
	}

	// Commit the transaction
	err = tx.Commit()
	if err != nil {
		return err
	}

	CloseDB(db)

	return err
}
