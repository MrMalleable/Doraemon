第 83 场周赛

据说是中国赛首赛

[830. 较大分组的位置](https://leetcode-cn.com/problems/positions-of-large-groups/)

思路： 这题的话比较常规，直接计算长度大于 3 的连续字符串就行

```java
class Solution {
    public List<List<Integer>> largeGroupPositions(String s) {
        List<List<Integer>> res = new ArrayList();
        //计算目前已经有多少个字符是相同的
        int cnt = 1;
        //记录前一个字符是什么字符
        char prev = s.charAt(0);
        //计算前一个相同字符的索引
        int prevIndex = 0;
        for(int i = 1; i < s.length(); i++){
            char cur = s.charAt(i);
            if(cur==prev){
                //如果当前字符和前一个字符是相同的
                cnt++;
            }else{
                //如果长度大于等于3则记录前后的索引位置
                if(cnt>= 3){
                    List<Integer> d = new ArrayList();
                    d.add(prevIndex);
                    d.add(i-1);
                    res.add(d);
                }
                //重置统计信息
                prev = cur;
                prevIndex = i;
                cnt = 1;
            }
        }
        //可能在结束遍历时相同字符的数目也达到了3个以上
        if(cnt >=3){
            List<Integer> d = new ArrayList();
            d.add(prevIndex);
            d.add(s.length()-1);
            res.add(d);
        }
        return res;
    }
}
```

[831.隐藏个人信息](https://leetcode-cn.com/problems/masking-personal-information/)

思路：单纯是一道阅读理解题，死办法，慢慢处理就好

```java
class Solution {
    public String maskPII(String S) {
        if(S.contains("@")){
            //处理邮箱
            String[] arr = S.split("@");
            String newStr = arr[0].toLowerCase();
            int len = newStr.length();
            return newStr.charAt(0)+"*****"+newStr.charAt(len-1)+"@"+arr[1].toLowerCase();
        }else{
            //处理电话号码
            StringBuilder sb = new StringBuilder();
            //只要是数字就加入
            for(int i = 0; i < S.length(); i++){
                char cur = S.charAt(i);
                if(Character.isDigit(cur)){
                    sb.append(cur);
                }
            }
            int len = sb.length();
            StringBuilder res = new StringBuilder();
            //判断超过10位就在前面加+
            if(len - 10 > 0){
                res.append("+");
            }
            //超过10位的多少位用*表示
            for(int i = 0; i < len - 10; i++){
                res.append("*");
            }
            if(len -10 > 0){
                res.append("-");
            }
            //后面十位本地号码表示比较简单
            res.append("***-***-"+sb.substring(len-4));
            return res.toString();

        }
    }
}
```

[829. 连续整数求和](https://leetcode-cn.com/problems/consecutive-numbers-sum/)

思路：看了题解才知道这是一道数学题，咋自己就想不出来呢

```java
class Solution {
    // n(2x+n+1)=2*N x必为整数
    public int consecutiveNumbersSum(int N) {
        int res = 0;
        int upper = (int)Math.sqrt(2*N);
        for (int n = 1; n <= upper; ++n)
            if (2 * N % n == 0) {
                int y = 2 * N / n - n - 1;
                if (y % 2 == 0 && y >= 0)
                    res++;
            }
        return res;
    }
}
```

[828. 统计子串中的唯一字符](https://leetcode-cn.com/problems/count-unique-characters-of-all-substrings-of-a-given-string/)

思路：我是这样理解的，当前字符位置为 i,往前遍历到相同的字符位置为 j 往后遍历到的相同的字符位置为 k

j.............i................k

我们知道从 i 往左边包括位置 i 的这个字符符合条件的字符串就是 i - j

我们知道从 i 往右边包括位置 i 的这个字符符合条件的字符串就是 k - i

然后以位置 i 的字符为中心，左边可以多加上的字符为（i - j - 1）个， 右边可以多加上的字符为（k - i - 1）

因为不能左边都是 0 个字符，这种只包含单边的已经包含在上述两种情况里去了

那么现在以位置 i 的字符为中心，左右必须都有字符，可选择的方案有（i - j - 1）\*（k - i - 1）种

加上上面所有的方案， 总数为（i- j）\*（k - i）

```java
class Solution {
    /**
     * 对于每一个字符，我们以当前字符为起始点，向前、向后去寻找最近的和该字符相等的字符
        例如：当前字符位置记为i,前面的和它相等的字符位置记为j，后面的和它相等的字符位置记为k,那么对于区间[j,...i,...k],
        i位置的字符对 [j+1, k-1]区间所做的贡献是(i-j)*(k-i).
        为什么呢？很显然j、k已经是截止点了。那么包含i的唯一字符串的个数：
        (i-j+1)(k-i+1)+(k-j-1) = (i-j)*(k-i).
     */
    public int uniqueLetterString(String s) {
        long re=0;
        for (int i=0; i<s.length(); i++) {
             int j=i-1, k=i+1;
             while (j>=0) {
                 if (s.charAt(j) == s.charAt(i)) {
                     break;
                 }
                 j--;
             }
             while (k<s.length()) {
                 if (s.charAt(k) == s.charAt(i)) {
                     break;
                 }
                 k++;
             }
             int cur = (i-j)*(k-i);
             re += cur;
        }
        return (int)re%1000000007;
    }
}
```

done
