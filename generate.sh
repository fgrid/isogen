echo -n "update fgrid/iso20022 ... "
go get -u github.com/fgrid/iso20022/...
echo "done"

echo -n "download iso20022 e-repository ... "
VER=20160321
rm -f *.go 2>/dev/null
curl -s -O http://www.iso20022.org/documents/eRepositories/Metamodel/${VER}_ISO20022_eRepository.zip
unzip ${VER}_ISO20022_eRepository.zip
echo "done"

if [ -z $1 ]
then
	PACKAGE="github.com/fgrid/iso20022"
else
	PACKAGE=$1
fi
echo "generate code in $PACKAGE ... "
cat ${VER}_ISO20022_2013_eRepository.iso20022 | sed -e 's/xsi:type/xsitype/g' | isogen -package="$PACKAGE"

echo -n "format code in $PACKAGE ... "
gofmt -s -w *.go
echo "done"

export WD=`pwd`
for area in `ls -d ????`
do
	echo -n "format code in $PACKAGE/$area ... "
	cd $area && gofmt -s -w *.go 
	echo "done"
	cd $WD
done

echo -n "build ... "
go build
echo "done"

echo -n "cleanup ... "
rm -f ${VER}*.iso20022 ${VER}*.zip
echo "done"
