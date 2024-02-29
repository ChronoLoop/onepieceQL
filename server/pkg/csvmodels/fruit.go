package csvmodels

type DevilFruitCSV struct {
	JapaneseName  string `db:"japanese_name" csv:"japanese_name" json:"japaneseName"`
	EnglishName   string `db:"english_name" csv:"english_name" json:"englishName"`
	Meaning       string `db:"meaning" csv:"meaning" json:"meaning"`
	UsageDebut    string `db:"usage_debut" csv:"usage_debut" json:"usageDebut"`
	Type          string `db:"type" csv:"type" json:"type"`
	PreviousOwner string `db:"previous_owner" csv:"previous_owner" json:"previousOwner"`
	CurrentOwner  string `db:"current_owner" csv:"current_owner" json:"currentOwner"`
	Description   string `db:"description" csv:"description" json:"description"`
	AvatarSrc     string `db:"avatar_src" csv:"avatar_src" json:"avatarSrc"`
	URL           string `csv:"url" json:"url"`
}
