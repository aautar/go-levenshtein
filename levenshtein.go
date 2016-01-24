package levenshtein

// min returns the minimum of 2 integers
func min(x int, y int) (int) {
    if(x < y) {
        return x
    }
    return y
}

// Distance computes the Levenshtein distance between 2 strings
func Distance(stringA string, stringB string) (int) {

    lenA := len([]rune(stringA)) + 1
    lenB := len([]rune(stringB)) + 1

    mat := make([][]int, lenA);
    for i:=0; i<lenA; i++ {
        mat[i] = make([]int, lenB)
    }

    for i:=0; i<lenA; i++ {
        mat[i][0] = i
    }

    for j:=0; j<lenB; j++ {
        mat[0][j] = j
    }

    for i:=1; i<=lenA-1; i++ {
        for j:=1; j<=lenB-1; j++ {
            if stringA[i-1] == stringB[j-1] {
                mat[i][j] = mat[i-1][j-1]
            } else {
                mat[i][j] = min( mat[i-1][j-1]+1, min(mat[i-1][j] + 1, mat[i][j-1] + 1) )
            }
        }
    }

    result := mat[lenA-1][lenB-1]
    return result
}
