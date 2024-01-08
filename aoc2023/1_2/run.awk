BEGIN {
	FS=""
} 

# get first digit 
{
	curr = $0
	gsub("twone", "2one", curr); #
	gsub("one", "1", curr);
	gsub("eightwo", "8two", curr); #
	gsub("two", "2", curr);
	gsub("eighthree", "8three", curr); #
	gsub("three", "3", curr);
	gsub("four", "4", curr);
	gsub("five", "5", curr);
	gsub("six", "6", curr);
	gsub("seven", "7", curr);
	gsub("nineight", "9eight", curr); #
	gsub("eight", "8", curr);
	gsub("nine", "9", curr);

	split(curr, arr, "")
	for (i=1; i<=NF; i=i+1) { 
		if (arr[i]+0 != 0) {
			nums[NR]=(arr[i]+0)*10
			break
		}
	}
}

# get last digit
{
	curr = ""
	for (i=NF; i>0; i-=1) {
		curr = sprintf("%s%s", curr, $i);
	}

	gsub("thgieno", "8eno", curr);
	gsub("eno", "1", curr);
	gsub("owt", "2", curr);
	gsub("thgieerht", "8eerht", curr);
	gsub("eerht", "3", curr);
	gsub("ruof", "4", curr);
	gsub("thgievif", "8evif", curr);
	gsub("evif", "5", curr);
	gsub("xis", "6", curr);
	gsub("thgienineves", "8enineves", curr);
	gsub("enineves", "9neves", curr);
	gsub("neves", "7", curr);
	gsub("thgie", "8", curr);
	gsub("enin", "9", curr);
	
	split(curr, arr, "")
	for (i=1; i<=NF; i=i+1) { 
		if (arr[i]+0 != 0) {
			nums[NR]=nums[NR]+arr[i]+0
			break
		}
	}

}

END { 
	# print "Solution:"
	sum=0;
	for( i=1; i<=NR; i=i+1 ){ 
		num = nums[i]
		# print i ": " num
		sum = sum + num
	}
	print "Total: " sum
}
