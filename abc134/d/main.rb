require 'set'

N = gets.to_i
arr = gets.split(' ').map(&:to_i)
arr.unshift nil

invalid = []

1.upto(n) do |i|
  cur = i
  cnt = 0

  loop do
    break if cur >= arr.size

    cnt += arr[cur]
    cur += i
  end

  if cnt % 2 != arr[i]
    invalid << i
  end
end
