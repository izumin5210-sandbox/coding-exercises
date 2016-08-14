_ = gets.to_i
a = gets.split(' ').map(&:to_i)

puts a.min.upto(a.max).map { |i| a.inject(0) { |s, a_i| s + (a_i - i) ** 2 } }.min
