int searchInsert(int* nums, int numsSize, int target){
    int i=0;
    for(;i<numsSize;i++){
        if(nums[i] >= target){
            break;
        }
    }
    return i;
}