_ = gets.to_i
gets.split(" ").map(&:to_i).map.with_index { |n, i| [n, i] }.sort_by { |a| -a[0] }.each do |i|
  puts i[1] + 1
end


