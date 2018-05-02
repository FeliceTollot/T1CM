package T1ConnectionManager

type T1Client interface {
	Connect()
	Discovery() *RemoteDeviceIterator
}
