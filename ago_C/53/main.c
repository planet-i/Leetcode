int maxSubArray(int* nums, int numsSize){
    int i,j;
    int max = nums[0];
    for(int i=0;i<numsSize;i++){
        int temp = 0;
        for (j=i;j<numsSize;j++){
            temp += nums[j];
            max = max >temp? max:temp;
        }
    }
    return max;
}