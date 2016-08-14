N, K = gets.split(" ").map(&:to_i)
D = gets.split(" ")

n = N

while n.to_s.chars.any? { |i| D.include?(i) }
  n += 1
end

puts n

