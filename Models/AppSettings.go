package models

// SatellitesLocation contiene la ubicación por cada satelite
type SatellitesLocation struct {
	Name string
	X    float32
	Y    float32
}

//ApplicationInfo contiene informacion de la aplicacion
type ApplicationInfo struct {
	Port string
}

//Settings contiene los ajustes de la aplicación
type Settings struct {
	ApplicationInfo    ApplicationInfo
	SatellitesLocation []SatellitesLocation
}
