# Syslog Max Length

rsyslog
```
sudo vim /etc/rsyslogd.conf

# add
$MaxMessageSize 64K

sudo /etc/init.d/rsyslog stop
sudo /etc/init.d/rsyslog start
```
