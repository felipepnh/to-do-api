CREATE TABLE IF NOT EXISTS todo(
    id uuid DEFAULT gen_random_uuid() PRIMARY KEY,
    title text,
    descriprion text
);