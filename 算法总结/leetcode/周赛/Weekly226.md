第226场周赛

这周题感觉不是很难，虽然我只做出来了第一题（因为我是菜鸡），但是思路上还是比较好理解的。下面就来具体看下吧。

[1742. 盒子中小球的最大数量](https://leetcode-cn.com/problems/maximum-number-of-balls-in-a-box/)

思路：使用一个Map来保存每个和出现的次数，每次都将最新的max更新，最后返回即可。

这里计算每个数字的各位之和，是通过计算每一个字符对应的数字来算的，比较简单！

```java
class Solution {
    public int countBalls(int low, int high) {
        Map<Integer,Integer> map = new HashMap();
        int max = 0;
        for(int i = low; i <= high; i++){
            String cur = i+"";
            // 计算各位之和
            int total = 0;
            for(char c: cur.toCharArray()){
                total += c-'0';
            }
            int res = map.getOrDefault(total,0) + 1;
            //更新最新的结果
            max = Math.max(max, res);
            map.put(total, res);
        }
        return max;
    }
}
```

[1743. 从相邻元素对还原数组](https://leetcode-cn.com/problems/restore-the-array-from-adjacent-pairs/)

思路：Map<Integer,List<Integer>> key：存放每个数 List<Integer>存放key对应的邻居

分析得知，我们知道整个最后的数组 第一个元素和最后一个元素的邻居只有一个，由于题目不要求顺序，所以开头和结尾的数肯定是只有一个邻居的数

所以：

1. 计算每个数的邻居

2. 找到邻居只有一个数的那个数作为开始

3. 取当前数的两个邻居，因为一个邻居已经在这个数前面放好了，那我们只需要在当前数的后面加上这个邻居

4. 以此循环即可

```java
class Solution {
    public int[] restoreArray(int[][] a) {
        //计算所有数的邻居
        Map<Integer, List<Integer>> map = new HashMap();
        int n = a.length + 1;
        for(int i = 0; i < n - 1; i++){
            int first = a[i][0];
            int second = a[i][1];
            if(map.containsKey(first)){
                map.get(first).add(second);
            }else{
                map.put(first, new ArrayList());
                map.get(first).add(second);
            }

            if(map.containsKey(second)){
                map.get(second).add(first);
            }else{
                map.put(second, new ArrayList());
                map.get(second).add(first);
            }
        }
        
        int start = 0;
        
        LinkedList<Integer> res = new LinkedList();
        
        //找到第一个只有一个邻居的那个数
        for(Map.Entry<Integer,List<Integer>> e: map.entrySet()){
            List<Integer> val = e.getValue();
            if(val.size()==1){
                start = e.getKey();
                break;
            }
        }
        
        //在数组的开始添加这个只有一个邻居的数
        res.add(start);
        //将这个数的邻居放到它的后面
        res.add(map.get(start).get(0));
        
        //当res的长度为n时停止循环
        while(res.size() < n){
            int cur = res.get(res.size() - 1);
            int prev = res.get(res.size() - 2);
            for(int i: map.get(cur)){
                if(i!=prev){
                    res.add(i);
                }
            }
        }
        
        
        int[] r = new int[n];
        int index = 0;
        for(int i: res){
            r[index++] = i;
        }
        
        return r;
        
    }
}
```

[1744. 你能在你最喜欢的那天吃到你最喜欢的糖果吗？](https://leetcode-cn.com/problems/can-you-eat-your-favorite-candy-on-your-favorite-day/)

思路：
这道题看题解是真简单，但是自己做的时候怎么就想不到呢我去

吃到favoriteDay的时候，能吃到的糖果的范围为： (favoriteDayi+1,(favoriteDayi+1)×dailyCapi)

对应favoriteType的糖果范围：(sum[favoriteTypei−1]+1,sum[favoriteTypei])

只要求这两个范围是否有交集即可。

```java
class Solution {
    // candiesCount[i] 表示你拥有的第 i 类糖果的数目
    // queries[i] = [favoriteTypei, favoriteDayi, dailyCapi]是一个查询, 表示为: 从第0天开始吃糖果, 每次可以吃1 ~ dailyCapi个, 其中只有吃完了0 ~ i - 1类型的糖果才能吃i类型糖果, 判断是否可在favoriteDayi天中吃到favoriteTypei类型的糖果
    
    // 思路分析: 
    // 1. 给所有类型的糖果进行编号; => 例如candiesCount = [3,2] => 编号 => [1,2,3,4,5]
    //                                                   0 1             0 0 0 1 1
    // 2. 从第0天开始吃糖果, 那么到favoriteDayi天时, 其一定最少能吃到 favoriteDayi + 1个糖果, 因为每天都必须吃一个糖果; 其每天最多能吃 (favoriteDayi + 1) * dailyCapi 个糖果, 因为每天最多可以吃dailyCapi个糖果. 所以到favoriteDayi天, 能吃到的糖果范围为[favoriteDayi + 1, favoriteDayi + 1) * dailyCapi] 个糖果
    // 3. 经过第一步的编号, 每种类型的糖果编号范围都确定了, 例如第一步类型0的糖果范围为[1,2]. 那么要判断在favoriteDayi是否能吃到类型x的糖果, 类型x的让过范围为[x, y], 直接判断[favoriteDayi + 1, favoriteDayi + 1) * dailyCapi] 与 [x, y]是否存在交集即可.
    
    public boolean[] canEat(int[] candiesCount, int[][] queries) {
        // sum[i]: 存储了下标i之前的元素之和
        int n = candiesCount.length;
        long[] sum = new long[n + 1];
        sum[0] = 0L;
        // 前缀和计算
        for (int i = 1; i <= n; i++) {
            sum[i] = sum[i - 1] + candiesCount[i - 1];
        }
        
        // 保存答案
        boolean[] answer = new boolean[queries.length];
        int idx = 0;
        for (int i = 0; i < queries.length; i++) {
            
            int type = queries[i][0];
            int day = queries[i][1];
            int num = queries[i][2];
            
            // 计算到favoriteDayi天能吃到的糖果范围 [favoriteDayi + 1, (favoriteDayi + 1) * dailyCapi]
            long a = day + 1;
            long b = a * num;
            // 计算type类型的编号范围
            long c = sum[type] + 1;
            long d = sum[type + 1];
          
            answer[i] = checkIntersection (a, b, c, d);   
        
        }
        
        return answer;
        
    }
    
    // 判断是否存在交集
    public boolean checkIntersection (long a, long b, long c, long d) {
        if (b < c || d < a)
            return false;
        return true;
    }

}
```

[1745. 回文串分割 IV](https://leetcode-cn.com/problems/palindrome-partitioning-iv/)

思路：

先用动态规划求出从i到j的位置的字符串是否为回文 d[i][j]=1代表就是回文

拷贝代码

```cpp
const int maxn = 2010;
bool dp[maxn][maxn];

class Solution {
public:
    bool checkPartitioning(string s) {
        memset(dp, 0, sizeof(dp));
        int n = s.length();
        for(int i = 0; i < n; i++){
            dp[i][i] = 1;
            if(i < n - 1){
                if(s[i]==s[i+1])dp[i][i+1] = 1;
            }
        }
        for(int L = 3; L <= n; L++){
            for(int i = 0; i + L - 1 < n; i++){
                int j = i + L - 1;
                if(s[i]==s[j] && dp[i+1][j-1]){
                    dp[i][j] = 1;
                }
            }
        }
        for(int s = 1; s <= n - 2; s++)
            for(int e = s; e <= n - 2; e++)
                if(dp[0][s-1]&&dp[s][e]&&dp[e+1][n-1])
                    return true;
        return false;
    }
};
```