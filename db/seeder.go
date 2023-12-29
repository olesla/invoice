package db

func Seed() {
	DB.AutoMigrate(&Customer{}, &Invoice{}, &Product{})
	DB.Create(&Customer{
		Name:    "Test bedrift",
		Number:  "000 00 000",
		Address: "Testadresse 1, 0000 Test",
		Email:   "test@test.no",
		Org:     "000 000 000",
	})
	DB.Create(&Customer{
		Name:    "Eksempel bedrift",
		Number:  "000 00 000",
		Address: "Eksempeladresse 1, 0000 Eksempel",
		Email:   "eksempel@eksempel.no",
		Org:     "000 000 000",
	})
	DB.Create(&Product{
		Name:        "Webtjenester",
		Description: "Beskrivelse av webtjenester",
		Code:        "P-001",
		Price:       500,
	})
	DB.Create(&Product{
		Name:        "Drift og vedlikehold",
		Description: "Bla bla vedlikehold osv",
		Code:        "P-002",
		Price:       999,
	})
}
