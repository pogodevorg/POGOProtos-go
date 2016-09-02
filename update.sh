#!/bin/sh

# Configure script
protos_repo="git@github.com:AeonLucid/POGOProtos.git"

# Collect paths of current script
origin_dir=$(pwd)
script_path=$(stat -f "%N" "$0")
base_dir=$(dirname $script_path)

# Collect full path of the submodule and base directory
cd $base_dir
base_dir=$(pwd)
protos_dir="$base_dir/POGOProtos"

# Get the project
if ! [ -d "$protos_dir" ]; then
	git clone $protos_repo $protos_dir
else
	cd $protos_dir
	git pull
fi

# Compile protobuffers and clean submodule
cd $protos_dir
current_commit=$(git rev-parse HEAD)
declare -a current_tags=($(git tag -l --contains $(git rev-parse HEAD)))
python ./compile_single.py -l=go --out_path=$base_dir
rm -rf "$protos_dir/tmp"

# Return back to original directory
cd $origin_dir

git add -A
git commit -m "Update POGOProtos $current_commit"

for tag in "${current_tags[@]}"; do
   git tag $tag
done
