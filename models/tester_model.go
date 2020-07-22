package models

type Tester struct {
	Id    int    `json:"id" orm:"column(tester_id)"`
	Key   string `json:"key" orm:"column(tester_key)"`
	Value string `json:"value" orm:"column(tester_value)"`
	// MetaId   int    `json:"metaId" orm:"column(tester_meta_id)"`
	MetaType string `json:"metaType" orm:"column(tester_meta_type)"`
	User     *User  `json:"user" orm:"rel(fk);column(tester_meta_id)":` // ;column(tester_meta_id);null
}

func (t *Tester) TableName() string {
	return "m_tester"
}
