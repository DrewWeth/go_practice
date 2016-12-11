def merge_sort(arr)
  return arr if arr.length <= 1
  mid = (arr.length / 2).to_i
  left = merge_sort(arr[0..mid-1])
  right = merge_sort(arr[mid..arr.length-1])
  return merge(left, right)
end

def merge(left, right)
  sorted = []
  while left.length > 0 || right.length > 0
    if left.length > 0 && right.length > 0
      if left[0] <= right[0]
        sorted << left.slice!(0)
      else
        sorted << right.slice!(0)
      end
    elsif left.length > 0
      sorted.concat left.slice!(0..left.length-1)
    elsif right.length > 0
      sorted.concat right.slice!(0..right.length-1)
    end
  end
  puts "Sorted: #{sorted.inspect}"
  return sorted
end

arr = [3,2,8,7,0,-2]
puts merge_sort(arr)
