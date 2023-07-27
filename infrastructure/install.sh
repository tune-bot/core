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

source infrastructure/database.env
sed -i "s|DB_USER|$DB_USER|g" database/infrastructure/create.sql
sed -i "s|DB_PASS|$DB_PASS|g" database/infrastructure/create.sql
sed -i "s|DB_HOST|$DB_HOST|g" database/infrastructure/create.sql
sed -i "s|DB_USER|$DB_USER|g" database/infrastructure/delete.sql
sed -i "s|DB_PASS|$DB_PASS|g" database/infrastructure/delete.sql
sed -i "s|DB_HOST|$DB_HOST|g" database/infrastructurel/delete.sql

service mysql start
mysql --defaults-extra-file=/etc/mysql/debian.cnf < database/infrastructure/create.sql
service mysql stop

echo "#!/bin/bash" > bin/core
echo "source vars/database.env" >> bin/core
echo "service mysql start" >> bin/core
chmod a+rx bin/core