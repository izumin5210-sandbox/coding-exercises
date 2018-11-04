_ = gets.to_i
t, a = gets.split(" ").map(&:to_i)

puts gets.split(" ").map(&:to_i).each_with_index.min { |(h1, _), (h2, _)| (a - (t - (h1 * 0.006))).abs <=> (a - (t - (h2 * 0.006))).abs }[1] + 1
