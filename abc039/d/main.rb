require "set"

H, W = gets.split(' ').map(&:to_i)
S = H.times.map { gets.chomp.chars }

black_cells = Set.new
shrinked_cells = Set.new

-1.upto(H).each_cons(3) do |h|
  -1.upto(W).each_cons(3) do |w|
    if h.all? { |i| w.all? { |j| !(0...H).cover?(i) || !(0...W).cover?(j) || S[i][j] == ?# } }
      black_cells.add([h[1], w[1]])
      if (0...H).cover?(h[0])
        shrinked_cells.add([h[0], w[0]]) if (0...W).cover?(w[0])
        shrinked_cells.add([h[0], w[1]])
        shrinked_cells.add([h[0], w[2]]) if (0...W).cover?(w[2])
      end
      shrinked_cells.add([h[1], w[0]]) if (0...W).cover?(w[0])
      shrinked_cells.add([h[1], w[2]]) if (0...W).cover?(w[2])
      if (0...H).cover?(h[2])
        shrinked_cells.add([h[2], w[0]]) if (0...W).cover?(w[0])
        shrinked_cells.add([h[2], w[1]])
        shrinked_cells.add([h[2], w[2]]) if (0...W).cover?(w[2])
      end
    end
  end
end

if H.times.all? { |i| W.times.all? { |j| S[i][j] == ?. || shrinked_cells.include?([i, j]) || black_cells.include?([i, j])} }
  puts "possible"
  H.times do |i|
    W.times do |j|
      if black_cells.include?([i, j])
        print ?#
      elsif shrinked_cells.include?([i, j])
        print ?.
        else
          print S[i][j]
        end
    end
    puts ""
  end
else
  puts "impossible"
end

