#!/bin/bash
# go get -u github.com/golang/protobuf/protoc-gen-go
# go get -u google.golang.org/protobuf/cmd/protoc-gen-go
# go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc
git pull
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
# protoc-gen-grpc-gateway
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest
GOPATH=$(go env GOPATH)
export PATH=$PATH:$GOPATH/bin

DIRS=("common" "slice" "user" "userfile" "offline" "rpcs" "wrtc" "share" "webdav" "sftp" "ticket")


rm -rf ./go_temp
mkdir ./go_temp

OUT_DIR=`pwd`


function makeFile(){
    file=$1
    echo "Generate: ${file}"
    protoc --go_out=./go_temp --go-grpc_out=./go_temp --grpc-gateway_out=./go_temp ${file}
    protoc -I . --openapiv2_out ./gen/openapiv2 --openapiv2_opt logtostderr=true,allow_merge=true,merge_file_name=api ${file}
}


# for name in ${DIRS}
for(( i=0;i<${#DIRS[@]};i++)) 
do
    name=${DIRS[i]}
    echo "Generate from: ${name}"
    for file in ./${name}/*.proto   
    do
        makeFile $file
    done
done


echo "Clean Go Files"
rm -rf ./go
cp -r ./go_temp/github.com/city404/v6-public-rpc-proto/go ./
rm -rf ./go_temp
cd ./go
# cp ../main.go ./
go mod init github.com/city404/v6-public-rpc-proto/go
# go get google.golang.org/grpc
go get
# go get github.com/golang/protobuf@master
go get golang.org/x/net
go mod tidy
cd ..

# echo "Copy JavaScript Files"
# rm -rf ${nodejs_path}
# mkdir ${nodejs_path}
# cp -r ./user ${nodejs_path}
# cp -r ./store ${nodejs_path}
# cp -r ./file ${nodejs_path}
# cp -r ./common ${nodejs_path}
# cp -r ./remotetask ${nodejs_path}
# cp -r ./offline ${nodejs_path}
# cp -r ./share ${nodejs_path}
# cp -r ./ext ${nodejs_path}
# cp -r ./report ${nodejs_path}

#rm -rf ./temp
#mkdir ./temp
#cd ./temp
#git clone https://github.com/zzzhr1990/ts-api-define.git
#cd ./ts-api-define
#cp -R ../../ts-api-define/* ./
#rm -rf ./node_modules
git add .
git commit -m "Auto commit"
git push
# cd ../..
echo "Done at:"`pwd`