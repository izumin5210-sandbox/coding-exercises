n = gets.to_i
a = gets.split(" ").map(&:to_i)

cnt = 0

loop do
  ok = true
  n.times do |i|
    if a[i].odd?
      ok = false
      break
    end
    a[i] /= 2
  end
  break if !ok
  cnt += 1
end

puts cnt
