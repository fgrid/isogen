echo -n "update fgrid/iso20022 ..."
go get -u github.com/fgrid/iso20022/...
echo "done"

echo -n "download iso20022 e-repository ... "
VER=20160321
rm *.go 2>/dev/null
curl -s -O http://www.iso20022.org/documents/eRepositories/Metamodel/${VER}_ISO20022_eRepository.zip
unzip ${VER}_ISO20022_eRepository.zip
echo "done"

echo -n "generate code ... "
cat ${VER}_ISO20022_2013_eRepository.iso20022 | sed -e 's/xsi:type/xsitype/g' | isogen
echo "done"

echo -n "format code ... "
gofmt -s -w *.go
echo "done"

echo -n "build ... "
go build
echo "done"

echo -n "cleanup ... "
rm -f ${VER}*.iso20022 ${VER}*.zip
echo "done"
