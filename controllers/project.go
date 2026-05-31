// Author: Fadhil
// PBI: KF-12
// Sprint: Sprint 2
func ConfirmPayment(c *gin.Context) {
	projectID := c.Param("id")

	var project models.Project
	if err := config.DB.Preload("Job").Preload("Transactions").Where("id = ?", projectID).First(&project).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Project tidak ditemukan"})
		return
	}

	if len(project.Transactions) > 0 {
		lastTx := project.Transactions[len(project.Transactions)-1]

		if lastTx.Status == "success" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Project ini sudah lunas, tidak perlu membayar lagi."})
			return
		}
		if lastTx.Status == "pending" || lastTx.Status == "process" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Pembayaran sebelumnya sedang diproses. Harap tunggu konfirmasi."})
			return
		}
	}

	nominalStr := c.PostForm("nominal")
	nominal, errParse := strconv.ParseFloat(nominalStr, 64)
	if errParse != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Nominal pembayaran tidak valid!"})
		return
	}

	budgetStr := project.Job.Budget
	parts := strings.Split(budgetStr, "-")
	minBudgetStr := strings.TrimSpace(parts[0])

	cleanBudgetStr := strings.Map(func(r rune) rune {
		if r >= '0' && r <= '9' {
			return r
		}
		return -1
	}, minBudgetStr)

	minBudget, _ := strconv.ParseFloat(cleanBudgetStr, 64)

	if nominal < minBudget {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("Pembayaran minimal adalah Rp %.0f", minBudget),
		})
		return
	}

	file, errFile := c.FormFile("bukti_transfer")
	if errFile != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Harap unggah bukti transfer!",
		})
		return
	}

	fileName := fmt.Sprintf("pay_%d_%s", time.Now().Unix(), filepath.Base(file.Filename))
	savePath := filepath.Join("uploads", fileName)

	if err := c.SaveUploadedFile(file, savePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Gagal simpan file",
		})
		return
	}

	fileURL := "/" + filepath.ToSlash(savePath)

	newTransaction := models.Transaction{
		ProjectID:     project.ID,
		ClientID:      project.ClientID,
		FreelancerID:  project.FreelancerID,
		Nominal:       nominal,
		BuktiTransfer: fileURL,
		Status:        "pending",
	}

	if err := config.DB.Create(&newTransaction).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Gagal simpan transaksi",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Pembayaran berhasil dikirim ulang dan sedang diproses.",
	})
}

// Author: Fadhil
// PBI: KF-12
// Sprint: Sprint 2
func GetClientTransactions(c *gin.Context) {
	userID, _ := getUserID(c)
	var txs []models.Transaction
	config.DB.Preload("Project.Job").Preload("Project.Freelancer").Where("client_id = ?", userID).Find(&txs)
	c.JSON(200, txs)
}

// Author: Fadhil
// PBI: KF-12
// Sprint: Sprint 2
func UpdateTransactionStatus(c *gin.Context) {
	id := c.Param("id")
	var input struct {
		Status string `json:"status"`
	}
	c.ShouldBindJSON(&input)
	config.DB.Model(&models.Transaction{}).Where("id = ?", id).Update("status", input.Status)
	c.JSON(200, gin.H{"message": "Status updated"})
}