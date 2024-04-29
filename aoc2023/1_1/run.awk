BEGIN {
	FS="" 
} 

{
	j=0;
	for (i=0; i<=NF; i=i+1) { 
		if (length($i) == 1 && $i+0 != 0) {
			nums[2*NR+j] = $i;
			nums[2*NR+j+1] = $i;
			print (":"NR" "$i+0" "j);
			if (j==0) {
				j=j+1;
			}
		} else {
			# print ("!"$i)
		}
	}
} 

END { 
	print "Solution:"
	sum=0;
	for( i=0; i<=NR; i=i+1 ){ 
		num = nums[2*i] nums[2*i+1]
		print i ": " num
		sum = sum + num
	}
	print "Total: " sum
}
