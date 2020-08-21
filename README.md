# 【golang】gorm出现incorrect datetime value '0000-0-0 00:00:00' for column问题

```sql
set global sql_mode='ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION';
```

# docker启动redis
```bash
docker run -p 6379:6379 -v $PWD/data:/data --name redis_1 -d redis redis-server --appendonly yes
```