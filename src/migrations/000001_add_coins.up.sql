CREATE TABLE IF NOT EXISTS crypto.coins
(
    id   UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    coinID         int,
    coinName       text,
    symbol         text,
    numMarketPairs int,
    dateAdded      timestamp,
    maxSupply      int4,
    price          float4,
    marketCap      float4,
    lastUpdated    timestamp
);