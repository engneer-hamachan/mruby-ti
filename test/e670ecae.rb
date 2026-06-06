def test
  x = '1'.downcase!

  return '1' if x.is_a?(NilClass)

  p x
end

