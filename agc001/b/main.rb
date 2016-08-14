N, X = gets.split(' ').map(&:to_i)

def dev(n, x)
  m = n % x
  2 * (n / x) * x + ((m > 0) ? dev(x, m) : -x)
end

puts N + dev(N - X, X)
