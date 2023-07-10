package interfaces

type IVehicle interface {
	IsRunning() bool
	Start() (bool, error)
	Stop() (bool, error)
}
