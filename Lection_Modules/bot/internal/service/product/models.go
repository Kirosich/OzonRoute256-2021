package product

var allProducts = []Product{
	{Title: "GTA 5",
		Description: "Grand Theft Auto 5\n" +
			"That game about gangsters\n" +
			"Price: 100$"},

	{Title: "The Elder Scorlls",
		Description: "The Elder Scorlls\n" +
			"That game about dragons\n" +
			"Price: 70$"},

	{Title: "MeatBoy",
		Description: "MeatBoy\n" +
			"Really bloody adventure\n" +
			"Price: 25$"},

	{Title: "Counter-Strike",
		Description: "Counter-Strike\n" +
			"Default shooter\n" +
			"Price: 40$"},

	{Title: "Dota 2",
		Description: "Defend of the antions\n" +
			"That game about toxic players\n" +
			"Price: 10$"},

	{Title: "World Of Warcraft",
		Description: "We need more money\n" +
			"Fantasy game\n" +
			"Price: 40$"},

	{Title: "BattleField 2",
		Description: "Cool cool cool shooter\n" +
			"Tanks, Soldiers, Helicopters \n" +
			"Price: 25$"},

	{Title: "Golden ring",
		Description: "Qwa Qwa\n" +
			"A lot of rings, but golden\n" +
			"Price: 40$"},

	{Title: "GTA 6",
		Description: "Grand Theft Auto 5\n" +
			"That game about gangsters\n" +
			"Price: 100$"},
	{Title: "The Elder Scorlls 5",
		Description: "The Elder Scorlls\n" +
			"That game about dragons\n" +
			"Price: 70$"},
	{Title: "MeatBoy 2",
		Description: "MeatBoy\n" +
			"Really bloody adventure\n" +
			"Price: 25$"},
	{Title: "Counter-Strike 7",
		Description: "Counter-Strike\n" +
			"Default shooter\n" +
			"Price: 40$"},
}

type Product struct {
	Title       string
	Description string
}
