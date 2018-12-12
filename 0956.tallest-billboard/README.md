i still cant fully understand how can u write out such beauty && nice code,how??
first, should understand the meaning of the dp[diff]=a, which represent the max sum (a+a+diff) \
the pair(a,a+diff) can get. its very hard. when u calculate dp[diff]=a, actually, dp[rods]=(sum-rod)/2 \
or dp(diff+-rod) = "sum".
second, u should understand the dp[]'s meaning during the calculation.
for rod in rods
    for d,y in dp
        dp[d+x] = max(dp[d+x],y)
        dp[abs(d-x)] = max(dp[abs(d-x),y+min(rod,d))

this two dp procession in the circlution include all the possible add or sub every rods
but whats the no use rod?? and where the?? on no, i forget what to say
i just understand why return dp[0], for diff  = 0.
i think if list all the dp[] and try to explan them and think the calculation, may understand more.

if some rods are not used, the diff = sum(no use rods).
if dp[0] is updated, then some rods'sum is equal,
the diff is kind of cumulative result, so the dp[diff] is the max of some kind of cumulative result
