package models

//InfoSatellite contiene la información de distancia y el mensaje de un satelite
type InfoSatellite struct {
	Distance float32  `json:"Distance" example:"100"`
	Message  []string `json:"Message"`
}

//Satellite contiene la informacion completa de un satelite
type Satellite struct {
	Name     string   `json:"Name" example:"kenobi"`
	Distance float32  `json:"Distance" example:"100"`
	Message  []string `json:"Message"`
}

//Satellites es la lista con la información de todos los satelites
type Satellites struct {
	Satellites []Satellite `json:"Satellites"`
}
