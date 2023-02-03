enum SearchProvider {
  leetx('1337x'),
  tpb('ThePirateBay'),
  // rarbg('RARBG'),
  yts('YTS'),
  kat('Kickass'),
  lime('LimeTorrents');

  const SearchProvider(this.name);
  final String name;
}
