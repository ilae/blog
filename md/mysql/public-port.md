# Public Port

Step 1:
```
sudo vim /etc/mysql/my.cnf
#bind-address  = 127.0.0.1
```

Step 2:
```
sudo /etc/init.d/mysql restart
```

Step 3:
```
mysql -uroot -p

GRANT ALL ON *.* to root@'%' IDENTIFIED BY 'password_you_set' WITH GRANT OPTION;
FLUSH PRIVILEGES;
```