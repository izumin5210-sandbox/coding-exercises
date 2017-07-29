N = gets.to_i
A = gets.split(' ').map(&:to_i)

def calc(i)
  (A[0...i].inject(0, :+) - A[i...N].inject(0, :+)).abs
end

sub = (1...(N - 1)).to_a

loop do
  n = sub.size
  break if n < 2
  sub1, sub2 = sub[0...(n / 2 + (n.odd? ? 1 : 0))], sub[(n / 2)...n]
  sub1.push(sub2[0]) if sub1.size.even?
  sub2.push(sub1[-1]) if sub2.size.even?
  puts "#{sub1} #{sub2}"
  puts "#{sub1[sub1.size / 2]}: #{calc(sub1[sub1.size / 2])}, #{sub2[sub2.size / 2]}: #{calc(sub2[sub2.size / 2])}"
  sub = (calc(sub1[sub1.size / 2]) < calc(sub2[sub2.size / 2]) ? sub1 : sub2).uniq
end

puts calc(sub[0] || 1)
