# 第一章 神奇的 SQL

## 1-1 CASE 表达式

1. 写 CASE 的时候一定记得加上 ELSE,不加 ELSE 的话，走到 ELSE 的时候执行结果为 NULL,可能会影响最终的结果
2. 在发现为真的 when 子句时，case 表达式的真假值判断就会终止，而剩余的 when 子句会被忽略，所以为了避免不必要的混乱，when 子句一定要注意排他性
3. 统一各分支返回的数据类型
4. 不要忘记写 END

示例 SQL:

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

- 用一条 SQL 语句进行不同条件的统计

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

- 用 CHECK 约束定义多个列的条件关系
- 在 UPDATE 语句里进行条件分支
- 表之间的数据匹配
- 在 CASE 表达式中使用聚合函数

## 1-2 自连接的用法

- 可重排列、排列、组合
- 删除重复行
- 查找局部不一致的列
- 排序

## 1-3 三值逻辑和 NULL

大多数编程语言都是基于二值逻辑的，即逻辑真值只有真和假两个。而 SQL 语言则采用一种特别的逻辑体系--三值逻辑，即逻辑真值除了真和假，还有第三个之“不确定”。

## 1-4 HAVING 子句的力量

- 寻找缺失的编号
- 用 HAVING 子句进行子查询：求众数
- 用 HAVING 子句进行自连接：求中位数
- 查询不包含 NULL 的集合
- 用关系除法运算进行购物篮分析

## 1-5 外连接的用法

- 用外连接进行行列转换(1)(行-列):制作交叉表
- 用外连接进行行列转换(2)(行-列):汇总重复项于一列
- 在交叉表里制作嵌套式表侧栏
- 作为乘法运算的连接
- 全外连接
- 用外连接进行集合运算
- 用外连接求差集：A-B
- 用全外连接求异或集

## 1-6 用关联子查询比较行与行

- 增长、减少、维持现状
- 用列表展示与上一年的比较结果
- 时间轴有间断时：和过去最邻近的时间进行比较
- 移动累计值和移动平均值
- 查询重叠的时间区间

## 1-7 用 SQL 进行集合运算

## 1-8 EXISTS 谓词的用法

## 1-9 用 SQL 处理数列

## 1-10 HAVING 子句又回来了

## 1-11 让 SQL 飞起来

- 使用高效的查询

  - 参数是子查询时，使用 EXISTS 代替 IN
  - 参数是子查询时，使用连接代替 IN

  - 避免排序

    - GROUP BY 子句
    - ORDER BY 子句
    - 聚合函数（SUM\COUNT\AVG\MAX\MIN）
    - DISTINCT
    - 集合运算符（UNION、INTERSECT、EXCEPT）
    - 窗口函数（RANK\ROW_NUMBER 等）

  - 灵活使用集合运算符的 ALL 可选项
  - 使用 EXISTS 代替 DISTINCT
  - 在极值函数中使用索引（MAX/MIN）
  - 能写在 where 子句的条件不要写在 HAVING 子句里
  - 在 group by 子句和 order by 子句中使用索引

- 真的用到索引了吗？

## 1-12 SQL 编程方法
