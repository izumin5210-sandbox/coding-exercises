puts gets.split(" ").map(&:to_i).combination(2).map { |x, y| 2 * x * y }.inject(0, :+)
