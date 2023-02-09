mkdir -p gen/proto
#protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative propertyapi/property/v1/property.proto 
protoc --go_out=gen/proto/ --go_opt=paths=source_relative --go-grpc_out=gen/proto/ --go-grpc_opt=paths=source_relative propertyapi/property/v1/property.proto 
#protoc --go_out=gen/proto/property/ --go-grpc_out=gen/proto/property/ propertyapi/property/v1/property.proto 
