# cmd-design

### 列出命令说明

```
agenda help
agenda help [command]
```

### 用户注册

```
agenda register -uUserName --password pass --email=a@xxx.com --phone=123456
```

### 用户登录

```
agenda login -uUserName --password pass
```

### 用户登出

```
agenda logout
```

### 用户查询

```
agenda query
```

### 用户删除

```
agenda delUser
```

### 创建会议

```
agenda cm --title=meeting --part=participator --start=start_time --end=end_time
```

### 增删会议参与者

```
agenda addPart --title=meeting --part=participator
agenda delPart --title=meeting --part=participator
```

### 查询会议

```
agenda qm --start=start_time  --end=end_time
```

### 取消会议

```
agenda dm --title=meeting
```

### 退出会议

```
agenda em --title=meeting
```

### 清空会议

```
agenda clear
```





