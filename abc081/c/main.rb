k = gets.split(' ')[1].to_i
puts gets.split(' ').group_by(&:itself).values.map(&:size).sort.reverse.drop(k).inject(0, :+)
