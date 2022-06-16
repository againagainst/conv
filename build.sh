echo "rm bin/conv"
rm bin/conv

echo "rm bin/plugins/ip.so"
rm bin/plugins/ip.so

echo "cd cmd/conv"
cd cmd/conv

echo "go build -o conv main.go"
go build -o conv main.go

echo "mv conv ../../bin"
mv conv ../../bin

echo "cd ../../plugins/ip"
cd ../../plugins/ip

echo "go build -buildmode=plugin -o ip.so ip.go"
go build -buildmode=plugin -o ip.so ip.go

echo "mv ip.so ../../bin/plugins/ip.so"
mv ip.so ../../bin/plugins/ip.so

echo "cd ../../plugins/bits"
cd ../../plugins/bits

echo "go build -buildmode=plugin -o bits.so bits.go"
go build -buildmode=plugin -o bits.so bits.go

echo "mv bits.so ../../bin/plugins/bits.so"
mv bits.so ../../bin/plugins/bits.so