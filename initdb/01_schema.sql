\c loandb;

CREATE TABLE IF NOT EXISTS pmt_log (
    id SERIAL PRIMARY KEY,
    loan_amount BIGINT NOT NULL,
    interest_rate BIGINT NOT NULL,
    number_of_payments INT NOT NULL,
    monthly_repayment BIGINT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_pmt_log_created_at_desc ON pmt_log (created_at DESC);