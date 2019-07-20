N = gets.to_i
arr = N.times.map { gets.to_i }

sorted = arr.sort
max_1 = sorted[-1]
max_2 = sorted[-2]

N.times do |i|
  if arr[i] == max_1
    puts max_2
  else
    puts max_1
  end
end
