#！/bin/bash

create_folder () {
	if [ -d "$1" ]; then
		echo "$1 exist"
	else
		mkdir "$1"
		echo "create $1 success"
	fi
}

root_path=/home
profile_path=/etc/profile

logs_path=$root_path/logs
create_folder "$logs_path"

tools_path=$root_path/tools
create_folder "$tools_path"

go_workspace_path=$root_path/go_workspace
create_folder "$go_workspace_path"
create_folder "$go_workspace_path/bin"
create_folder "$go_workspace_path/pkg"
create_folder "$go_workspace_path/src"

# install git
if [ `command -v git` ]; then
	echo "Git has been installed"
else
	echo "install git start--------------------"
	yum install -y git
	git --version
	echo "install git end--------------------"
fi


# install go
if [ `command -v go` ]; then
	echo "Go has been installed"
else
	echo "install go start--------------------"
	cd "$tools_path"
	wget https://studygolang.com/dl/golang/go1.20.1.linux-amd64.tar.gz
	tar -xvf go1.20.1.linux-amd64.tar.gz
	mv go /usr/local/
	echo "Modify $profile_path"
	echo ''>>"$profile_path"
	echo 'export GOPROXY=https://goproxy.cn,direct'>>"$profile_path"
	echo 'export GOROOT=/usr/local/go'>>"$profile_path"
	echo 'export GOPATH=/home/go_workspace'>>"$profile_path"
	echo 'export GOBIN=$GOPATH/bin'>>"$profile_path"
	echo 'export PATH=$PATH:$GOROOT/bin'>>"$profile_path"
	echo ''>>"$profile_path"
	source "$profile_path"
	echo "install go end -------------------"
fi


# install beego
if [ `command -v bee`  ]; then
	echo "Beego has been installed"
else
	echo "install beego start--------------------"
	cd "$go_workspace_path"
	git clone https://github.com/beego/bee.git
	cd "$go_workspace_path/bee"
	go build -o bee
	echo "Modify $profile_path"
	sed -i '/export PATH=$PATH:$GOROOT/d' "$profile_path"
	echo 'export BEEPATH=/home/go_workspace/bee'>>"$profile_path"
	echo 'export PATH=$PATH:$GOROOT/bin:$BEEPATH'>>"$profile_path"
	source "$profile_path"
	echo "install beego end--------------------"
fi


if [ `command -v ipfs` ]; then
        echo "IPFS has been installed"
else
	echo "install ipfs start--------------------"
	cd "$tools_path"
	wget https://dist.ipfs.tech/kubo/v0.18.1/kubo_v0.18.1_linux-amd64.tar.gz --no-check-certificate
	tar -xvzf kubo_v0.18.1_linux-amd64.tar.gz
	cd kubo
	bash install.sh
	ipfs --version
	echo "install ipfs end--------------------"

	echo "init ipfs start--------------------"
	ipfs init
	# todo : config ipfs
	echo "init ipfs end--------------------"
fi

# install mysql
if [ `command -v mysql` ]; then
        echo "MySQL has been installed"
else
	echo "install mysql start--------------------"
	dnf -y install @mysql
	echo "install mysql end--------------------"
fi

# mtv-server
mtv_server_path=$go_workspace_path/mtv-server
if [ -d "$mtv_server_path" ]; then
	echo "Update mtv-server"
	cd "$mtv_server_path"
	git pull
else
	echo "Download mtv-server"
	cd "$go_workspace_path"
	git clone -b develop https://gitee.com/tinyverse-space/mtv-server.git
fi
cd "$mtv_server_path"
nohup bee run mtv > ../../logs/mtv.log 2>&1 &

# qasks
qasks_path=$go_workspace_path/qasks
if [ -d "$qasks_path" ]; then
	echo "Update qasks"
	cd "$qasks_path"
	git pull
else
	cd "$go_workspace_path"
	git clone -b develop https://gitee.com/tinyverse-space/qasks.git
fi
cd "$qasks_path"
nohup bee run qasks  > ../../logs/qasks.log 2>&1 &
