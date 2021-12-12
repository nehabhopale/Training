package geo
import "errors"
type Coordinates struct {
	latitude  float64
	longitude float64
}
type Landmark struct{
	name string
	Coordinates
}
func(l*Landmark)Name()string{
	return l.name
}
func(l*Landmark)SetName(name string)error{
	if name==""{
		return errors.New("invalid")
	}
	l.name=name
	return nil
}
func (c *Coordinates) SetLatitude(latitude float64) {
	c.latitude = latitude
}
func (c *Coordinates) SetLongitude(longitude float64) {
	c.longitude = longitude
}
func (c *Coordinates) Latitude() float64 {
	return c.latitude
}
func (c *Coordinates) Longitude() float64 {
	return c.longitude
}
func(c *Coordinates)SetLatitude(latitude float64) error{
	if latitude<-90 || latitude>90{
		return errors.New("invalid")
	}
	c.latitude=latitude
	return nil
}
func(c *Coordinates)SetLongitude(longitude float64) error{
	if latitude<-180 || latitude>180{
		return errors.New("invalid long")
	}
	c.longitude=longitude
	return nil
}

