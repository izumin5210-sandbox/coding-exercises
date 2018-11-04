k = gets.to_i

n = 2

loop do
  break if n * (n - 1) + k <= 9 * n
  n += 1
end

i = n - 1 + k

buf = []
n.times do
  j = [9, i].min
  buf << j
  i -= j
end

puts n
puts buf.join(" ")
