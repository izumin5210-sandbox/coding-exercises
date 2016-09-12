require 'securerandom'

class Card
  include Comparable

  TYPES = %w(3 4 5 6 7 8 9 10 J Q K A 2)

  attr_reader :id, :number

  def initialize(number)
    @id = SecureRandom.uuid
    @number = number
  end

  def <=>(other)
    TYPES.index(self.number) <=> TYPES.index(other.number)
  end
end

class Player
  include Comparable

  attr_reader :order, :cards

  def initialize(order, cards = [])
    @order = order
    @cards = cards
  end

  def <=>(other)
    self.order <=> other.order
  end

  def put?(card)
    card.nil? || @cards.any? { |c| c > card }
  end

  def put(card)
    puttable = nil
    if card.nil?
      puttable = @cards.min
    else
      puttable = @cards.select { |c| c > card }.min
    end
    @cards.delete_if { |c| c.id == puttable.id }
    puttable
  end

  def done?
    @cards.empty?
  end
end

players = gets.chomp.split(" ").map.with_index { |c, i| Player.new(i, [Card.new(c)]) }

ranks = {}
i = 0
rank = 0
field = nil
last_passed_at = nil

loop do
  player = players[i]

  unless player.done?
    if player.put?(field) || (!last_passed_at.nil? && (last_passed_at == i ))
      unless last_passed_at.nil?
        last_passed_at = nil
        field = nil
      end
      field = player.put(field)
      ranks[player.order] = (rank += 1) if player.done?
    elsif last_passed_at.nil?
      last_passed_at = i
    end

  end

  break if players.all?(&:done?)

  i += 1
  i = 0 if i == players.size
end

puts ranks.sort.map(&:last)
