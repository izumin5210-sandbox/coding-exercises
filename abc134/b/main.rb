n, d = gets.split(' ').map(&:to_i)
puts (n / (2 * d + 1).to_f).ceil
