#/bin/bash

dirList=$(ls -l ./ |awk '/^d/ {print $NF}')

for dir in $dirList
do
  cd $dir
  testResult=$(go test -failfast -coverprofile=coverage.out)
  echo "test $testResult"
  if [[ "$testResult" =~ "FAIL" ]]
  then
    echo "Failed"
    break
    exit 1
  else
    echo "Successfully"
    cd ..
  fi
done

