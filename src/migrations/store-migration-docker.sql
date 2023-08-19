CREATE TABLE IF NOT EXISTS stores (
    id SERIAL PRIMARY KEY,
    tradingname VARCHAR(255),
    ownername VARCHAR(125),
    document VARCHAR(50),
    coveragearea JSONB,
    addrs JSONB
);

INSERT INTO stores (tradingName, ownerName, document, coverageArea, addrs)
VALUES (
    'Adega da Cerveja - Pinheiros',
    'Zé da Silva',
    '1432132123891/0001',
    '{
        "type": "MultiPolygon",
        "coordinates": [
            [[[30, 20], [45, 40], [10, 40], [30, 20]]],
            [[[15, 5], [40, 10], [10, 20], [5, 10], [15, 5]]]
        ]
    }',
    '{
        "type": "Point",
        "coordinates": [-46.57421, -21.785741]
    }'
);