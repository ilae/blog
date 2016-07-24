# Register Hostname To DHCP


```
sudo vim /etc/dhcp/dhclient.conf

send host-name 'name-you-want'
#or
send host-name = gethostname();

sudo reboot now
```
