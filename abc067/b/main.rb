n, k = gets.split(' ').map(&:to_i)
puts gets.split(' ').map(&:to_i).sort.reverse[0, k].inject(0, :+)
