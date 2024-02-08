package challenges

//  You are going to build a stone wall. The wall should be straight and N meters long,
//  and its thickness should be constant; however, it should have different heights
//  in different places. The height of the wall is specified by an array H of N positive
//  integers. H[I] is the height of the wall from I to I+1 meters to the right of its
//  left end. In particular, H[0] is the height of the wall's left end and H[N−1] is
//  the height of the wall's right end.

//  The wall should be built of cuboid stone blocks (that is, all sides of such blocks
//  are rectangular). Your task is to compute the minimum number of blocks needed to build the wall.

//  Write a function:

//  def solution(H)

//  that, given an array H of N positive integers specifying the height of the wall,
//  returns the minimum number of blocks needed to build it.

//  For example, given array H containing N = 9 integers:

//    H[0] = 8    H[1] = 8    H[2] = 5
//    H[3] = 7    H[4] = 9    H[5] = 8
//    H[6] = 7    H[7] = 4    H[8] = 8
//  the function should return 7. The figure shows one possible arrangement of seven blocks.

//  Write an efficient algorithm for the following assumptions:

//  N is an integer within the range [1..100,000];
//  each element of array H is an integer within the range [1..1,000,000,000].

func SolutionStoneWall(A []int) int {
	intArr := A

	count := 1
	for i := 1; i < len(intArr); i++ {
		prevHeight := intArr[i-1]
		newHeight := intArr[i]
		if newHeight > prevHeight {
			count += 1
		} else if newHeight < prevHeight {
			counted := false
			for j := 0; j < i; j++ { // check if the new height is already counted in a previous block
				jHeight := intArr[j]
				if newHeight > jHeight {
					break
				} else if newHeight == jHeight {
					counted = true
					break
				}
			}
			if !counted {
				count += 1
			}
		}
	}
	return count
}
