-- name: CreateExchange :exec
INSERT INTO exchanges (
	amount,
	rate,
	from_currency,
	to_currency,
	converted_value,
	currency_symbol
) VALUES (
	?, ?, ?, ?, ?, ?
);
