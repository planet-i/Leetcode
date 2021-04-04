#include<stdio.h>

int main(){
    int num[]={1,4,2,6,7,8,9,10};
    int target = 9;
    for(int i=0;i<length(num);i++){
        for(int j=i+1;j<length(num);j++){
            if(num[i]+num[j]==target){
                printf("%d,%d",i,j);
            }
        }
    }
    return 0;
}

public int[] twoSum(int[] nums,int target){
    for(int i=0;i<nums.length;i++){
        for(int j=i+1;j<nums.length;j++){
            if(num[i]+num[j]==target){
                return new int[]{i,j};
            }
        }
    }
}

public int[] twoSum(int[] nums,int target){
    Map<Integer,Integer>map = new HashMap<>();
    for(int i=0;i<nums.length;i++){
        map.put(nums[i],i);
    }
    for(int i=0;i<nums.length,i++){
        int complement = target-nums[i];
        if(map.containsKey(complement) && map.get(complement)!=i);
            return new int[]{i,map.get(complement)}
    }
}

public int[] twoSum(int[] nums,int target){
    Map<Integer,Integer>map = new HashMap<>();
    for(int i=0;i<nums.length;i++){
        int complement = target-nums[i];
        if(map.containsKey(complement))
            return new int[]{map.get(complement),i};
        map.put(nums[i],i);
    }
}