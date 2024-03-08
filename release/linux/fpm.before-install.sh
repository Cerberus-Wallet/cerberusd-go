getent group cerberusd >/dev/null || groupadd -r cerberusd
getent group plugdev >/dev/null || groupadd -r plugdev
getent passwd cerberusd >/dev/null || useradd -r -g cerberusd -d /var -s /bin/false -c "Cerberus Bridge" cerberusd
usermod -a -G plugdev cerberusd
