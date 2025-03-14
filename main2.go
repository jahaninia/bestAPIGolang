func main2() {
    repo, err := NewRepository()
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }

    ctx := context.Background()

    // Create
    itemID, err := repo.Create(ctx, Item{Name: "Item1", Price: 10.0})
    if err != nil {
        log.Println("Error creating item:", err)
        return
    }

    // Read
    item, err := repo.Read(ctx, itemID)
    if err != nil {
        log.Println("Error reading item:", err)
        return
    }
    fmt.Printf("Read Item: %+v\n", item)

    // List all items
    items, err := repo.List(ctx)
    if err != nil {
        log.Println("Error listing items:", err)
        return
    }
    fmt.Println("All items:", items)

    // Update
    item.Price = 15.0
    err = repo.Update(ctx, item)
    if err != nil {
        log.Println("Error updating item:", err)
        return
    }
    fmt.Printf("Updated Item: %+v\n", item)

    // Delete
    err = repo.Delete(ctx, itemID)
    if err != nil {
        log.Println("Error deleting item:", err)
        return
    }
    fmt.Println("Deleted Item with ID:", itemID)

    // List all items after deletion
    items, err = repo.List(ctx)
    if err != nil {
        log.Println("Error listing items:", err)
        return
    }
    fmt.Println("All items after deletion:", items)
}