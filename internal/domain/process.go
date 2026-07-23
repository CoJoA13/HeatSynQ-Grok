package domain

type ProcessMaster struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	ProcessID   string `json:"process_id" gorm:"unique"`
	Name        string `json:"name"`
	Revision    string `json:"revision"`
	Status      string `json:"status"`
}
