package models

// func main() {
// g := gen.NewGenerator(gen.Config{
// 	OutPath: "./models",
// 	Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
// })

// // gormdb, _ := gorm.Open(mysql.Open("root:@(127.0.0.1:3306)/demo?charset=utf8mb4&parseTime=True&loc=Local"))
// g.UseDB(initializers.DB) // reuse your gorm db

// // Generate basic type-safe DAO API for struct `model.User` following conventions
// g.ApplyBasic(model.Book{}, model.User{}, model.BorrowingRecord{})

// g.ApplyBasic(
// 	// Generate struct `User` based on table `users`
// 	g.GenerateModel("borrowing_records"),
// )
// g.Execute()
// }
