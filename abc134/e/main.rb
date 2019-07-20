last = -1
cnt = 1

gets.to_i.times do
  n = gets.to_i

  if last >= n
    cnt += 1
  end

  last = n
end

puts cnt
