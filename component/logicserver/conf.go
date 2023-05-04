package logicserver

type Conf struct {
	ServerID string `csv:"cfgId"`
	Port     int    `csv:"noHolidaysTime"`
	Dev      bool   `csv:"holidaysTime"`
}
