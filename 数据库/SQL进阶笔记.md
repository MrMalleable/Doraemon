# 第一章 神奇的SQL
## 1-1 CASE 表达式

1. 写CASE的时候一定记得加上ELSE,不加ELSE的话，走到ELSE的时候执行结果为NULL,可能会影响最终的结果
2. 在发现为真的when子句时，case表达式的真假值判断就会终止，而剩余的when子句会被忽略，所以为了避免不必要的混乱，when子句一定要注意排他性
3. 统一各分支返回的数据类型
4. 不要忘记写END

示例SQL:

- 将已有编号方式转换为新的方式并统计
```sql
-- 把县编号转换成地区编号
SELECT CASE pref_name 
        WHEN '德岛' THEN '四国'
        WHEN '香川' THEN '四国'
        WHEN '爱媛' THEN '四国'
        WHEN '高知' THEN '四国'
        WHEN '福冈' THEN '九州'
        WHEN '佐贺' THEN '九州'
        WHEN '长崎' THEN '九州'
    ELSE '其他' END AS district
FROM PopTbl
GROUP BY CASE pref_name
    WHEN '德岛' THEN '四国'
            WHEN '香川' THEN '四国'
            WHEN '爱媛' THEN '四国'
            WHEN '高知' THEN '四国'
            WHEN '福冈' THEN '九州'
            WHEN '佐贺' THEN '九州'
            WHEN '长崎' THEN '九州'
        ELSE '其他' END;

```
- 用一条SQL语句进行不同条件的统计
```sql
-- 男性人口
SELECT pref_name, sum(population) 
    from PopTbl2
    where sex = '1'
    GROUP BY pref_name;

-- 女性人口
SELECT pref_name,sum(population)
    from PopTbl2
    where sex = '2'
    GROUP BY pref_name;


-- 优化：
select pref_name,
    -- 男性人口
    sum(case when sex ='1' then population else 0 END) as cnt_m,
    -- 女性人口
    sum(case when sex ='2' then population else 0 END) as cnt_f
    From PopTbl2
    group by pref_name;
```