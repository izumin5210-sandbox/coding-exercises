n = gets.to_i
p gets.split(" ").map(&:to_i).sort[n / 2]
