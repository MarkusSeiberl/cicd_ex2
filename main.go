package main

func main() {
	a := App{}
	/*a.Initialize(
	os.Getenv("APP_DB_USERNAME"),
	os.Getenv("APP_DB_PASSWORD"),
	os.Getenv("APP_DB_NAME"))*/
	a.Initialize(
		"postgres",
		"postgres",
		"products")

	a.Run(":8010")
}
