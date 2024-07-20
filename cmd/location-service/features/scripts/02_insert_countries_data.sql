-- Insert test data into the countries table
INSERT INTO countries (name, code) VALUES
('United States', 'US'),
('Canada', 'CA'),
('United Kingdom', 'GB'),
('Australia', 'AU'),
('New Zealand', 'NZ'),
('France', 'FR')
ON CONFLICT (code) DO NOTHING;