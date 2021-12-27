package model

type Location struct {
	LOCATION_ID    int    `gorm:"type:int;column:LOCATION_ID"`
	STREET_ADDRESS string `gorm:"type:varchar(40);column:STREET_ADDRESS"`
	POSTAL_CODE    string `gorm:"type:varchar(12);column:POSTAL_CODE"`
	CITY           string `gorm:"type:varchar(30);column:CITY"`
	STATE_PROVINCE string `gorm:"type:varchar(25);column:STATE_PROVINCE"`
	COUNTRY_ID     string `gorm:"type:char(2);column:COUNTRY_ID"`
}
func NewLocation(LOCATION_ID int,STREET_ADDRESS string ,POSTAL_CODE string ,CITY string ,STATE_PROVINCE string ,COUNTRY_ID string )Location{
	return Location{
	LOCATION_ID : LOCATION_ID,  
	STREET_ADDRESS :STREET_ADDRESS,
	POSTAL_CODE   : POSTAL_CODE,
	CITY    :    CITY,   
	STATE_PROVINCE :STATE_PROVINCE,
	COUNTRY_ID :  COUNTRY_ID,
	}
}