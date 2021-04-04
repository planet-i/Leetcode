/**
 * Note: The returned array must be malloced, assume caller calls free().
 */
#include <stdio.h>
#include <stdlib.h>
int* twoSum(int* nums, int numsSize, int target,int* returnSize){
    int *result = (int*)malloc(sizeof(int)*2);
    for(int i=0;i<numsSize-1;i++){
        for(int j=i+1;j<numsSize;j++){
            if(nums[j] == target - nums[i])
            {
                result[0]=i;
                result[1]=j;
                *returnSize = 2;
                return result;
            }    
        }
    }
    *returnSize = 0;
    return result;
}
int main(){
    int nums[4] = {2, 7, 11, 15};
    int target = 9;
    int* p=NULL;
    int* size = NULL;
    p = twoSum(nums,sizeof(nums),target,&size);
    for(int i=0;i<size;i++)
    {
        printf("%d",*p[i]);
    }
    return 0;
}