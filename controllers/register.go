package controllers

// func CreateOrders(c *gin.Context) {
// 	var order Order
// 	if err := c.ShouldBindJSON(&order); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 			return
// 	}

// 	// Retrieve the ticket from the database
// 	var ticket Ticket
// 	if err := db.Where("ticket_id = ?", order.TicketID).First(&ticket).Error; err != nil {
// 			c.JSON(http.StatusNotFound, gin.H{"error": "Ticket not found"})
// 			return
// 	}

// 	// Check if there is enough stock
// 	if ticket.CurrentQuantity < 1 {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": "Insufficient stock"})
// 			return
// 	}

// 	// Decrement the ticket's current quantity
// 	ticket.CurrentQuantity -= 1

// 	// Save the order and the updated ticket
// 	if err := db.Transaction(func(tx *gorm.DB) error {
// 			if err := tx.Create(&order).Error; err != nil {
// 					return err
// 			}
// 			if err := tx.Save(&ticket).Error; err != nil {
// 					return err
// 			}
// 			return nil
// 	}); err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 			return
// 	}

// 	c.JSON(http.StatusCreated, order)
// }

// // GetOrders retrieves a list of all orders.
// func GetOrders(c *gin.Context) {
// 	var orders []Order
// 	if err := db.Find(&orders).Error; err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 			return
// 	}
// 	c.JSON(http.StatusOK, orders)
// }

// // GetOrder retrieves a single order by ID.
// func GetOrder(c *gin.Context) {
// 	id, err := strconv.Atoi(c.Param("id"))
// 	if err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
// 			return
// 	}
// 	var order Order
// 	if err := db.Where("order_id = ?", id).First(&order).Error; err != nil {
// 			c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
// 			return
// 	}
// 	c.JSON(http.StatusOK, order)
// }