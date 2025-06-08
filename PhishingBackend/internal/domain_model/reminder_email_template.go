package domain_model

type ReminderEmailTemplate struct {
	Id       int `gorm:"primary_key;"`
	Template string
	Subject  string
}
