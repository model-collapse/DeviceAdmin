package usage

const (
	UnknownDevice = iota
	RaspberryPI_3B
	RaspberryPI_4B
	RaspberryPI_Zero
	JetsonNano
	ESP32CAM
	SparkfunBLE
)

const (
	UnknownRole = iota
	EdgeRole
	SensorRole
)

func StrToDevType(s string) (ret int) {
	switch s {
	case "pi3":
		ret = RaspberryPI_3B
	case "pi4":
		ret = RaspberryPI_4B
	case "pi0":
		ret = RaspberryPI_Zero
	case "jetson_nano", "jetson":
		ret = JetsonNano
	case "esp32cam":
		ret = ESP32CAM
	case "sparkfun_ble":
		ret = SparkfunBLE
	default:
		ret = UnknownDevice
	}

	return ret
}

func StrToRole(s string) (ret int) {
	switch s {
	case "edge":
		ret = EdgeRole
	case "sensor":
		ret = SensorRole
	}

	return
}
