package setup_env

import "time"

const (
	CAcert_path          = "main/certificate/CAs/ca-cert.pem"
	Cert_path            = "main/certificate/client/client.cert"
	Key_path             = "main/certificate/client/client.key"
	MongoDB_addr         = "0.0.0.0:8080"
	MongoDB_service_addr = "0.0.0.0:8081"
	Redis_addr           = "0.0.0.0:6379"
	Redis_service_addr   = "0.0.0.0:8083"
	Create_user_addr     = "0.0.0.0:8084"
	Create_user_gw       = "0.0.0.0:8085"
	Login_user_addr      = "0.0.0.0:8086"
	Login_user_gw        = "0.0.0.0:8087"
	TokenDuration        = time.Minute * time.Duration(15)
	RefreshTokenDuration = time.Hour * time.Duration(24)
)
