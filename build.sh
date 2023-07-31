#!/bin/zsh

export GOOS=linux
export GOARCH=arm64
export CGO_ENABLED=0

home=$PWD
root="$home/cmd"

cd $root

for i in $(ls -d **/)
do
    if [[ -f "$i/main.go" ]]
    then
        cd $i

        newFileName=${${i%/}//\//-}

        zipFileName="$home/bin/$newFileName.zip" 
        outFileName="$home/bin/bootstrap" 

        rm -f $zipFileName $outFileName

        go build -o $outFileName
        echo "Creating zip: $zipFileName"
        zip -jD  $zipFileName $outFileName
        
        cd $root
    fi
done

cd "$home"