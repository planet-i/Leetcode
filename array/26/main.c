int removeDuplicates(int* nums, int numsSize){
    int pre=0,cur = 0;
    if( numsSize == 0)
        return 0;
    while(cur < numsSize ){
        if(nums[pre] == nums[cur]){
            cur++;
        }else{
            pre++;
            nums[pre] = nums[cur];
            cur ++;
        }
    }
    return pre+1;
}