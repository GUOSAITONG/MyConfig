package config

type Appconf struct {
	Mysql
	Redis
	Es
	MongoDB
	SendSms
	Upload
}
type Mysql struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
}
type Redis struct {
	Host     string
	Port     int
	Password string
}
type Es struct {
	Host string
	Port int
}
type MongoDB struct {
	Host     string
	Port     int
	User     string
	Password string
}
type SendSms struct {
	Account  string
	Password string
}
type Upload struct {
	AccessKeyID     string
	AccessKeySecret string
	Region          string
	Bucket          string
}
