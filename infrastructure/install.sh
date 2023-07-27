#!/bin/bash

sed -i "/#\$nrconf{restart} = 'i';/s/.*/\$nrconf{restart} = 'a';/" /etc/needrestart/needrestart.conf

apt update
apt upgrade -y
apt install -y curl mysql-server python-is-python3 ffmpeg

mkdir -p library
mkdir -p bin

curl -L https://github.com/yt-dlp/yt-dlp/releases/latest/download/yt-dlp -o bin/download
chmod a+rx bin/download

source infrastructure/database.env
sed -i "s|DB_USER|$DB_USER|g" core/infrastructure/create.sql
sed -i "s|DB_PASS|$DB_PASS|g" core/infrastructure/create.sql
sed -i "s|DB_HOST|$DB_HOST|g" core/infrastructure/create.sql
sed -i "s|DB_USER|$DB_USER|g" core/infrastructure/delete.sql
sed -i "s|DB_PASS|$DB_PASS|g" core/infrastructure/delete.sql
sed -i "s|DB_HOST|$DB_HOST|g" core/infrastructure/delete.sql

service mysql start
mysql --defaults-extra-file=/etc/mysql/debian.cnf < core/infrastructure/create.sql
service mysql stop

echo "#!/bin/bash" > bin/database
echo "source vars/database.env" >> bin/database
echo "service mysql start" >> bin/database
chmod a+rx bin/database