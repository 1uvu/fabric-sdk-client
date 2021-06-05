#!/bin/bash

rm -rf keystore/
rm -rf wallet/

go mod vendor

ENV_DAL=`echo $DISCOVERY_AS_LOCALHOST`

echo "ENV_DAL:"$DISCOVERY_AS_LOCALHOST

if [ "$ENV_DAL" != "true" ]
then
	export DISCOVERY_AS_LOCALHOST=true
fi

export TEST_IN_SHELL=true

echo "DISCOVERY_AS_LOCALHOST="$DISCOVERY_AS_LOCALHOST
echo "TEST_IN_SHELL="$TEST_IN_SHELL
echo "run sdk app and admin client test..."

go test
