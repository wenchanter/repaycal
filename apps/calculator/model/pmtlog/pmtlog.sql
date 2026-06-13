CREATE TABLE pmt_log (
    id SERIAL PRIMARY KEY,
    loan_amount BIGINT NOT NULL,
    interest_rate BIGINT NOT NULL,
    number_of_payments INT NOT NULL,
    monthly_repayment BIGINT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE INDEX idx_pmt_log_created_at ON pmt_log(created_at);