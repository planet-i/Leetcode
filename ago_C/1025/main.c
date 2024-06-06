#include<stdio.h>
int main(){
    int N;   // 1 <= N <= 100 输入N
    int x;   //假设两位玩家以最佳状态参加游戏，分析x在最佳情况下的自动取值。根据奇偶次变换来判断谁赢了。
    int count=0;
    scanf("%d",&N);
    for(x=1;x<N;){      //控制x的变化
        if (0<x && x<N && N%x == 0 )
        {
            N -= x; 
            count++;
        }else{
            break;
        }
    }
    if (count%2 == 0)
    {
        printf("false\n");
    }else{
        printf("true\n");
    }
    return 0;             //自己的错误思路
}

bool divisorGame(int N){
    if(N==1)
        return false;
    if(N==2)
        return true;
    int dp[N + 5];
    memset(f, 0, sizeof(dp));
    //vector<int> f(N + 5, false);           C++
    //boolean[] f = new boolean[N + 5];      java      //+5   为了不处理下标偏移
    dp[1] = false;
    dp[2] = true;
    for(int i=3;i<=N;i++){
        dp[i] = false;
        for(int j=1;j<=i/2;j++){
            if ((i%j==0) && !dp[i-j])
            {
                dp[i]=true;
                break;
            }      
        }
    }
    return dp[N];
}
// func divisorGame(N int) bool {
//     f := make([]bool, N + 5)
//     f[1], f[2] = false, true
//     for i := 3; i <= N; i++ {
//         for j := 1; j < i; j++ {
//             if i % j == 0 && !f[i - j] {
//                 f[i] = true
//                 break
//             }
//         }
//     }
//     return f[N]
// }      go