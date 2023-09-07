package model

type Store struct {
    ID        uint    `gorm:"primaryKey"`
    Name      string  `json:"name" gorm:"not null"`
    Address   string  `json:"address"`
    Longitude float64 `json:"longitude"`
    Latitude  float64 `json:"latitude"`
    Rating    int     `json:"rating"`
    Weather   string  `json:"weather"`
}