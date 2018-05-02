package T1ConnectionManager

type RemoteDeviceIterator interface {
	HasNext() bool
	Next() RemoteDevice
	Remove()
}
