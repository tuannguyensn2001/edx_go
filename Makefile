migrate:
	goose mysql "root:@tcp(127.0.0.1:3306)/edx?charset=utf8mb4&parseTime=True&loc=Local" up