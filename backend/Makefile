.PHONY: proto server client ca

certificate_path = main/certificate

proto:
	protoc --proto_path=proto --go_out=pb --go-grpc_out=pb \
		--grpc-gateway_out=pb --grpc-gateway_opt=paths=source_relative \
		--go_opt=paths=source_relative --go-grpc_opt=paths=source_relative proto/*.proto

server:
	openssl req -new -nodes -newkey rsa:4096 -keyout ${certificate_path}/server/server.key -out ${certificate_path}/server/server.csr -subj "/C=VN/L=Ho Chi Minh/CN=localhost.local"
	openssl x509 -req -in ${certificate_path}/server/server.csr -CA ${certificate_path}/CAs/ca-cert.pem -CAkey ${certificate_path}/CAs/ca-key.pem -CAcreateserial -out ${certificate_path}/server/server.cert -extfile ${certificate_path}/server/server-ext.cnf

client:
	openssl req -new -nodes -newkey rsa:4096 -keyout ${certificate_path}/client/client.key -out ${certificate_path}/client/client.csr -subj "/C=VN/L=Ho Chi Minh/CN=localhost.local"
	openssl x509 -req -in ${certificate_path}/client/client.csr -CA ${certificate_path}/CAs/ca-cert.pem -CAkey ${certificate_path}/CAs/ca-key.pem -CAcreateserial -out ${certificate_path}/client/client.cert -extfile ${certificate_path}/client/client-ext.cnf

ca:
	openssl req -nodes -x509 -newkey rsa:4096 -keyout ${certificate_path}/CAs/ca-key.pem -out ${certificate_path}/CAs/ca-cert.pem -subj "/C=VN/CN=makerzone.net/emailAddress=nguyendt456@gmail.com"
