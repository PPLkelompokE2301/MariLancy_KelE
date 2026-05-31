	// Author: Fadhil
	// PBI: KF-12
	// Sprint: Sprint 2
	go func() {
		for {
			time.Sleep(1 * time.Hour)

			batasWaktu := time.Now().Add(-48 * time.Hour)

			var txs []models.Transaction
			err := config.DB.Where("status = ? AND created_at <= ?", "pending", batasWaktu).Find(&txs).Error
			if err == nil && len(txs) > 0 {
				for _, tx := range txs {
					config.DB.Model(&tx).Update("status", "success")

					config.DB.Model(&models.Project{}).Where("id = ?", tx.ProjectID).Update("payment_status", "paid")

					fmt.Printf("[AUTO-APPROVE] Transaksi ID %d otomatis sukses karena sudah melewati 2 hari tanpa rejeksi freelancer.\n", tx.ID)
				}
			}
		}
	}()

	r := gin.Default()