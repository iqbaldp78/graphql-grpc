<!-- protoc -I. --micro_out=. --go_out=paths=source_relative:. proto/pb/user/users.proto

protoc -I. --micro_out=. --go_out=paths=source_relative:. proto/pb/rfqs/rfqs.proto

protoc -I. --micro_out=. --go_out=. proto/pb/rfqs/rfqs.proto

protoc -I. --micro_out=. --go_out=plugins=grpc,import_path=github.com/mbizmarket/dmp/radiance:. proto/pb/user/users.proto

protoc -I. --micro_out=. --go_out=. proto/pb/user/users.proto

protoc -I proto/pb/rfqs/rfqs.proto --micro_out=. --go_out=mod_root=github.com/mbizmarket/dmp/radiance:. proto/pb/rfqs/rfqs.proto

go list -f '{{ .Dir }}' -m github.com/golang/protobuf


protoc -I proto/pb/rfqs --micro_out=. --go_out=. proto/pb/rfqs/rfqs.proto

protoc -I proto/pb/user --micro_out=. --go_out=. proto/pb/user

protoc -I chaochaogege.com/filecatcher/common/ chaogege.com/filecatcher/common/ChunkInfo.proto --go_out=chaochaogege.com/filecatcher/common/

protoc -I proto/pb/user/ proto/pb/user/users.proto --go_out=proto/pb/user/

protoc --go_out=./ proto/pb/user/users.proto

go list -m -f "{{.Dir}}" github.com/mbizmarket/dmp/radiance

protoc -I. I`go list -m -f "{{.Dir}}" github.com/mbizmarket/dmp/radiance` --micro_out=. --go_out=. proto/pb/rfqs/rfqs.proto

protoc -I. I`go list -m -f "{{.Dir}}" /github.com/mbizmarket/dmp/radiance` --go_out=. proto/pb/rfqs/rfqs.proto

go_out=mod_root=github.com/user/myproject:.


protoc -I. --micro_out=. --go_out=mod_root=github.com/mbizmarket/dmp/radiance:. proto/pb/rfqs/rfqs.proto

protoc -I. --micro_out=. --go_out=mod_root=github.com/mbizmarket/dmp/radiance:. proto/pb/user/users.proto
 -->

protoc -I. --micro_out=. --go_out=. proto/pb/rfqs/rfqs.proto
protoc -I. --micro_out=. --go_out=. proto/pb/user/users.proto


note :
this command will create folder with option go_package = "github.com/mbizmarket/dmp/radiance/proto/pb/rfqs"; 
then copy pb file to user folder