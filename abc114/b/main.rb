puts gets.chomp.chars.each_cons(3).map(&:join).map(&:to_i).map { |x| (753 - x).abs }.min
