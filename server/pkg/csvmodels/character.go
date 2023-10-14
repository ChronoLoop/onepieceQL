package csvmodels

type CharacterCSV struct {
	JapaneseName   string `db:"japanese_name" csv:"japanese_name"`
	EnglishName    string `db:"english_name" csv:"english_name"`
	Debut          string `db:"debut" csv:"debut"`
	Affiliations   string `db:"affiliations" csv:"affiliations"`
	Origin         string `db:"origin" csv:"origin"`
	Age            string `db:"age" csv:"age"`
	Birthday       string `db:"birthday" csv:"birthday"`
	BloodType      string `db:"blood_type" csv:"blood_type"`
	Bounty         string `db:"bounty" csv:"bounty"`
	Description    string `db:"description" csv:"description"`
	DevilFruitName string `db:"devil_fruit_name" csv:"devil_fruit_name"`
	AvatarSrc      string `db:"avatar_src" csv:"avatar_src"`
}
