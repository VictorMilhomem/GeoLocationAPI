CREATE TABLE IF NOT EXISTS store (
    id SERIAL PRIMARY KEY,
    tradingName TEXT,
    ownerName TEXT,
    document TEXT,
    coverageArea GEOMETRY(MULTIPOLYGON),
    addrs GEOMETRY(POINT)
);

INSERT INTO store_data (tradingName, ownerName, document, coverageArea, addrs)
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
    }'::geometry,
    '{
        "type": "Point",
        "coordinates": [-46.57421, -21.785741]
    }'::geometry
);