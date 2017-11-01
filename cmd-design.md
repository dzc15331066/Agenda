# cmd-design

### 列出命令说明

```
Agenda help
Agenda help [command]
```

### 用户注册

```
Agenda register -uUserName -p pass --email=a@xxx.com --contact=123456
```

### 用户登录

```
Agenda login -uUserName --password pass
```

### 用户登出

```
Agenda logout
```

### 用户查询

```
Agenda query
```

### 用户删除

```
Agenda delUser
```

### 创建会议

```
Agenda cm --title=meeting --part=participator --start=start_time --end=end_time
```

### 增删会议参与者

```
Agenda addPart --title=meeting --part=participator
Agenda delPart --title=meeting --part=participator
```

### 查询会议

```
Agenda qm --start=start_time  --end=end_time
```

### 取消会议

```
Agenda dm --title=meeting
```

### 退出会议

```
Agenda em --title=meeting
```

### 清空会议

```
Agenda clear
```





