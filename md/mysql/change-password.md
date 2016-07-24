# Change Password

Method 1:
```
sudo -i
mysqladmin -u root -p password "newpassword"
```

If forget the password
Method 2:
```
sudo vim /etc/mysql/my.cnf

# add
skip-grant-tables
skip-networking

sudo /etc/init.d/mysqld restart

mysql -uroot

UPDATE mysql.user SET Password=PASSWORD('MyNewPass')WHERE User='root';
FLUSH PRIVILEGES;

sudo vim /etc/mysql/my.cnf

#skip-grant-tables
#skip-networking

sudo /etc/init.d/mysqld restart
```

