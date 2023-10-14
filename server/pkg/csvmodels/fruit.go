package csvmodels

type DevilFruitCSV struct {
	JapaneseName string `db:"japanese_name" csv:"japanese_name"`
	EnglishName  string `db:"english_name" csv:"english_name"`
	Meaning      string `db:"meaning" csv:"meaning"`
	UsageDebut   string `db:"usage_debut" csv:"usage_debut"`
	Type         string `db:"type" csv:"type"`
	PreviousUser string `db:"previous_user" csv:"previous_user"`
	CurrentOwner string `db:"current_owner" csv:"current_owner"`
	Description  string `db:"description" csv:"description"`
	AvatarSrc    string `db:"avatar_src" csv:"avatar_src"`
}
