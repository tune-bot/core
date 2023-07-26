#!/bin/bash

sed -i "/#\$nrconf{restart} = 'i';/s/.*/\$nrconf{restart} = 'a';/" /etc/needrestart/needrestart.conf

apt update
apt upgrade -y
apt install -y curl mysql-server python-is-python3 ffmpeg

mkdir -p library
mkdir -p bin

curl -L https://github.com/yt-dlp/yt-dlp/releases/latest/download/yt-dlp -o bin/download
chmod a+rx bin/download

bin/download -U

source vars/database.env
sed -i "s|DB_USER|$DB_USER|g" database/install/create.sql
sed -i "s|DB_PASS|$DB_PASS|g" database/install/create.sql
sed -i "s|DB_HOST|$DB_HOST|g" database/install/create.sql
sed -i "s|DB_USER|$DB_USER|g" database/install/delete.sql
sed -i "s|DB_PASS|$DB_PASS|g" database/install/delete.sql
sed -i "s|DB_HOST|$DB_HOST|g" database/install/delete.sql

service mysql start
mysql --defaults-extra-file=/etc/mysql/debian.cnf < database/install/create.sql
service mysql stop

echo "#!/bin/bash" > bin/database
echo "source vars/database.env" >> bin/database
echo "service mysql start" >> bin/database
chmod a+rx bin/database