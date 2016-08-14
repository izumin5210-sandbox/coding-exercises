a, b = gets.split(" ").map(&:to_i)

if a * b > 0
  if b < 0 && (b.abs - a.abs).even?
    puts "Negative"
  else
    puts "Positive"
  end
else
  puts "Zero"
end
