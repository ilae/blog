# SSH Keys

PublicKey like lock

PrivateKey like key

A is client

B is server

A generate PrivateKey & PublicKey, then send PublicKey to B. A can login B's ssh server.

A:

```
ssh-keygen
```
will create /root/.ssh/id_rsa (PrivateKey) and /root/.ssh/id_rsa.pub (PublicKey)

B:
```
cat id_rsa.pub >> /root/.ssh/authorized_keys
```

A login:
```
ssh root@B_host
```