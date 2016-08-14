S = gets.chomp.chars

completed = false

2.upto(3) do |i|
  break if completed
  S.each_cons(i).with_index do |s, j|
    (puts "#{j + 1} #{i + j}"; break) if completed = (s.uniq.size != i)
  end
end

puts "-1 -1" unless completed
