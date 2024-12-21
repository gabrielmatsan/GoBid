-- Write your migrate up statements here

CREATE TABLE IF NOT EXISTS products (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

  seller_id UUID NOT NULL REFERENCES users (id),
  product_name VARCHAR(255) NOT NULL,
  description TEXT NOT NULL,

  price FLOAT NOT NULL,
  auction_end TIMESTAMPTZ NOT NULL,
  is_sold BOOLEAN NOT NULL DEFAULT FALSE,

  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
)

---- create above / drop below ----

-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.


DROP TABLE IF EXISTS products;