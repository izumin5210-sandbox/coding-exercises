n, x = gets.split(' ').map(&:to_i)

puts [n - x, x - 1].min
