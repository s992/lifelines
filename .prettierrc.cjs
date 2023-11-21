module.exports = {
  semi: true,
  trailingComma: 'all',
  singleQuote: true,
  printWidth: 80,
  importOrder: ['<THIRD_PARTY_MODULES>', '^[./]'],
  importOrderSeparation: true,
  importOrderSortSpecifiers: true,
  xmlWhitespaceSensitivity: 'ignore',
  plugins: ['@trivago/prettier-plugin-sort-imports', '@prettier/plugin-xml'],
};
