def test
  x = '1'.downcase!

  return '1' if x.is_a?(NilClass)

  dbtp x
end

