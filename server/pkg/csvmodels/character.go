package csvmodels

type CharacterCSV struct {
	JapaneseName   string `db:"japanese_name" csv:"japanese_name" json:"japaneseName"`
	EnglishName    string `db:"english_name" csv:"english_name" json:"englishName"`
	Debut          string `db:"debut" csv:"debut" json:"debut"`
	Affiliations   string `db:"affiliations" csv:"affiliations" json:"affiliations"`
	Origin         string `db:"origin" csv:"origin" json:"origin"`
	Age            string `db:"age" csv:"age" json:"age"`
	Birthday       string `db:"birthday" csv:"birthday" json:"birthday"`
	BloodType      string `db:"blood_type" csv:"blood_type" json:"bloodType"`
	Bounty         string `db:"bounty" csv:"bounty" json:"bounty"`
	Description    string `db:"description" csv:"description" json:"description"`
	DevilFruitName string `db:"devil_fruit_name" csv:"devil_fruit_name" json:"devilFruitName"`
	AvatarSrc      string `db:"avatar_src" csv:"avatar_src" json:"avatarSrc"`
}
