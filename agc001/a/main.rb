N = gets.to_i
L = gets.split(" ").map(&:to_i)

puts L.sort.each_slice(2).map(&:min).inject(0, :+)
