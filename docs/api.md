Create new user

```bash
http post api.dappi.ir/v1/users name=mahdi family=imani nationality=iran nationalCode=0011445732 username=mimani password=1jshi3 sex=male email=imani.mahdi@gmail.com phone=09124184801 birthday=1990-01-04
```

Get user data

```bash
http api.dappi.ir/v1/me
```

Get single user data

```bash
http api.dappi.ir/console/users/cba2a781-2536-4543-a22e-bfe3a0e3fd8c
```


Update user

```bash
http patch api.dappi.ir/console/users/cba2a781-2536-4543-a22e-bfe3a0e3fd8c name=ali
```