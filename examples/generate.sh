echo -n "update fgrid/isogen ... "
go get -u github.com/fgrid/isogen
echo "done"

VER=20160603
echo -n "download iso20022 e-repository-${VER} ... "
if [ -r ${VER}_ISO20022_eRepository.zip ]
then
	echo "already here"
else
	curl -s -O http://www.iso20022.org/documents/eRepositories/Metamodel/${VER}_ISO20022_eRepository.zip
	echo done
fi

echo -n "unpack iso20022 e-repository ... "
if [ -r ${VER}_ISO20022_2013_eRepository.iso20022 ]
then
	echo "alread done"
else
	unzip ${VER}_ISO20022_eRepository.zip >/dev/null
	echo "done"
fi

if [ -z $1 ]
then
	PACKAGE="github.com/fgrid/iso20022"
else
	PACKAGE=$1
fi

if [ -z $2 ]
then
	echo "no message filter set"
else
	MESSAGE_OPTS="-message=$2"
	echo "MESSAGE_OPTS=$MESSAGE_OPTS"
fi
echo "generate code in $PACKAGE ... "
cat ${VER}_ISO20022_2013_eRepository.iso20022 | sed -e 's/xsi:type/xsitype/g' | isogen -package="$PACKAGE" $MESSAGE_OPTS

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

#echo -n "cleanup ... "
#rm -f ${VER}*.iso20022 ${VER}*.zip
#echo "done"
