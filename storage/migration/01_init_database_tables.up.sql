
CREATE TABLE IF NOT EXISTS whois (
	"method_name" VARCHAR(255) UNIQUE NOT NULL,
	"whoclicked" VARCHAR(255) NOT NULL,
	"created_at" TIMESTAMP DEFAULT now(),
 );
