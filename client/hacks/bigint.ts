type SerializableBigInt = bigint & { toJSON(): string };

(BigInt.prototype as SerializableBigInt).toJSON = function () {
  return this.toString();
};
