func longestMountain(arr []int) int {
    maxLength := 0
    arrayLength := len(arr)
    
    for index, _ := range arr {
        
        if index > 0 && index < arrayLength-1 {
            
            // If peak
            if arr[index-1] < arr[index] && arr[index]> arr[index+1] {
                
                lengthOfAscent := 2
                lengthOfDescent := 1
                
                j := index-1
                for j>0 && arr[j-1] < arr[j] {
                    lengthOfAscent++
                    j--
                }
                
                j = index+1
                for j<arrayLength-1 && arr[j+1] < arr[j] {
                    lengthOfDescent++
                    j++
                }
                
                currentMountainLength := lengthOfAscent+lengthOfDescent
                if currentMountainLength > maxLength {
                    maxLength = currentMountainLength
                }
            }
        }
        
    }
    
    return maxLength
}
