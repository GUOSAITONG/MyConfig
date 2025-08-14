package config

type Appconf struct {
	SendSms
}
type SendSms struct {
	Account  string
	Password string
}
